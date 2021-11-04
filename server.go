package main
import (
	"context"
	"log"
	//"math/rand"
	"net"
	//"time"

	pb "github.com/Escariuz/SD/proto_archivos"

	"google.golang.org/grpc"
)

//Puerto definido
const (
	port = ":50051"
)

//Implementación del método propuesto por proto, ya que está sin implementar
type ComunicacionJLServer struct {
	pb.UnimplementedComunicacionJLServer
}

//Implementación del método definido en proto
func (s *ComunicacionJLServer) GetRespuesta(ctx context.Context, in *pb.SolicitudParticipacion) (*pb.RespuestaLider, error) {
	log.Printf("Recibida solicitud de juego")

	arreglo := in.GetParticipacion()

	log.Printf(arreglo[0])



	
	
	//Lider contesta a jugador un 0 o un 1
	//Si contesta 0, el jugador no puede unirse al juego
	//Si contesta 1, el jugador puede unirse al juego
	//var numero int32 = int32(rand.Intn(max-min+1) + min)

	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	return &pb.RespuestaLider{Respuesta: arr}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fallo al escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterComunicacionJLServer(s, &ComunicacionJLServer{})
	log.Printf("Servidor escuchando en %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}

}
