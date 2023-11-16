package measurement

import (
	"context"

	// session "github.com/MikaelHans/catea/session-management/api"
	// "github.com/MikaelHans/catea/user-service/pkg/structs"
	pb "github.com/MikaelHans/catea/product-management-service/api/measurement"
	repo "github.com/MikaelHans/catea/product-management-service/internal/measurement/repository"
)

type Server struct {
	pb.UnimplementedMeasurementServiceServer
}

func (s *Server) GetAllMeasurements(ctx context.Context, none*pb.None) (*pb.Measurements, error) {
	rows, err := repo.Get_Measurements()
	if err != nil{
		return nil, err
	}
	var measurements[]*pb.Measurement
	for rows.Next() {
		var measurement_data *pb.Measurement
		rows.Scan(
			&measurement_data.Name,
		)
		measurements = append(measurements, measurement_data)
	}
	return &pb.Measurements{Measurements: measurements}, err
}
