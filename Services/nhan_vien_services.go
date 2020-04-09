package Services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmdl-user-service/pb"
	"hmdl-user-service/repository"
)

type NhanVienServicePro struct {
	RepoNhanVien repository.NhanVienRepository
}

func (u *NhanVienServicePro) GetNhanVien(context.Context, *pb.ReadRequest) (*pb.ReadResponse, error) {

	dsNhanVien, err := u.RepoNhanVien.GetAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro server: %v", err)
	}

	var dataResponse []*pb.NhanVien
	for _, item := range dsNhanVien {
		nv := pb.NhanVien{
			DmNhanVienId: uint32(item.DM_NhanVienId),
			MaNhanVien:   item.MaNhanVien,
			TenNhanVien:  item.TenNhanVien,
			FullName:     item.FullName,
			LastName:     item.LastName,
			PhongId:      uint32(item.PhongID),
			DmChucDanhId: uint32(item.DM_ChucDanhID),
			DmChucVuId:   uint32(item.DM_ChucVuID),
			SoDienThoai:  item.SoDienThoai,
		}
		dataResponse = append(dataResponse, &nv)
	}

	res := &pb.ReadResponse{
		Api:      "v1",
		NhanVien: dataResponse,
	}
	return res, nil
}
