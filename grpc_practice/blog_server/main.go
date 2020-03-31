package main

//https://itnext.io/learning-go-mongodb-crud-with-grpc-98e425aeaae6 참고
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	blogpb "first_golang/grpc_practice/blog"
)

//TODO : COBRA로 CLI만들어서 블로그 테스트

type BlogServiceServer struct {
	blogpb.UnimplementedBlogServiceServer
}

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

const port = ":8080"

//전역변수
var (
	mongoLink string
	db        *mongo.Client
	mongoCtx  context.Context
	blogdb    *mongo.Collection
)

func main() {

	getMongoLink()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Printf("%s번 포트에서 서버를 실행합니다.\n", port)

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("해당 포트에 서버를 연결할 수 없습니다. : %v", err)
	}

	//set options
	opts := []grpc.ServerOption{}
	//options으로 gprc 생성
	s := grpc.NewServer(opts...)
	//BlogService 타입 생성
	srv := &BlogServiceServer{}
	//서버와 서비스를 등록
	blogpb.RegisterBlogServiceServer(s, srv)

	//고루틴에서 서버 시작
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("서버 실행 실패 : %v", err)
		}
	}()

	fmt.Printf("서버가 성공적으로 %s 포트에서 실행되고 있습니다.\n", port)

	initMongoInstance()

	//프로그램 강제 종료시 시그널을 받아서 서버도 같이 종료시킴
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	//시그널을 받기 전까지 메인 고루틴을 막음
	<-c

	//ctrl + c 누를 시 종료
	fmt.Println("서버를 종료하는 중입니다...")
	s.Stop()
	listener.Close()
	fmt.Println("MongoDB 연결을 중지합니다.")
	db.Disconnect(mongoCtx)
	fmt.Println("모두 성공적으로 종료되었습니다.")
}

func getMongoLink() {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		mongoLink = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//몽고디비 설정
func initMongoInstance() {
	fmt.Println("MongoDB에 연결을 시도합니다.")
	mongoCtx = context.Background()
	db, _ = mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoLink))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Check whether the connection was succesful by pinging the MongoDB server
	_err := db.Ping(mongoCtx, nil)
	if _err != nil {
		log.Fatalf("MongoDB에 연결할 수 없습니다.: %v\n", _err)
	} else {
		fmt.Println("MongoDB에 연결되었습니다.")
	}

	//컬렉션 추가
	blogdb = db.Database("mydb").Collection("blog")
}

//create
func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {

	blog := req.GetBlog()

	data := BlogItem{
		//id는 없어도 자동생성될 수 있게 테이블 생성할 당시 id에 auth_increment 적용했음.
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	//데이터베이스에 추가
	result, err := blogdb.InsertOne(mongoCtx, data)
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID) //추가된 아이디

	blog.Id = oid.Hex() //string값으로 변환

	return &blogpb.CreateBlogRes{Blog: blog}, nil
}

//read
func (s *BlogServiceServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogReq) (*blogpb.ReadBlogRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	result := blogdb.FindOne(ctx, bson.M{"_id": oid})

	data := BlogItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s: %v", req.GetId(), err))
	}

	response := &blogpb.ReadBlogRes{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}

	return response, nil

}

//delete
func (s *BlogServiceServer) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogReq) (*blogpb.DeleteBlogRes, error) {

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	_, err = blogdb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", req.GetId(), err))
	}

	return &blogpb.DeleteBlogRes{
		Success: true,
	}, nil
}

//update
func (s *BlogServiceServer) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogReq) (*blogpb.UpdateBlogRes, error) {
	blog := req.GetBlog()

	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err),
		)
	}

	update := bson.M{
		"author_id": blog.GetAuthorId(),
		"title":     blog.GetTitle(),
		"content":   blog.GetContent(),
	}

	filter := bson.M{"_id": oid}

	result := blogdb.FindOneAndUpdate(ctx, filter, bson.M{"&set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	decoded := BlogItem{}

	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find blog with supplied ID: %v", err),
		)
	}

	return &blogpb.UpdateBlogRes{
		Blog: &blogpb.Blog{
			Id:       decoded.ID.Hex(),
			AuthorId: decoded.AuthorID,
			Title:    decoded.Title,
			Content:  decoded.Content,
		},
	}, nil

}

//listing
func (s *BlogServiceServer) ListBlogs(req *blogpb.ListBlogsReq, stream blogpb.BlogService_ListBlogsServer) error {

	data := &BlogItem{}

	cursor, err := blogdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		stream.Send(&blogpb.ListBlogsRes{
			Blog: &blogpb.Blog{
				Id:       data.ID.Hex(),
				AuthorId: data.AuthorID,
				Content:  data.Content,
				Title:    data.Title,
			},
		})
	}
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	return nil
}
