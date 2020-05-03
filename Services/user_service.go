package Services

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmdl-user-service/auth"
	"hmdl-user-service/pb"
	"hmdl-user-service/repository"
	"log"
	"time"
)

type UserService struct {
	RepoNhanVien      repository.NhanVienRepository
	RepoThemSoHeThong repository.DmThamSoHeThongRepo
}

func (u *UserService) GetThamSosByCode(ctx context.Context, request *pb.ThamSoHeThongRequest) (*pb.ThamSoHeThongsResponse, error) {
	fmt.Println("Call : GetThamSoByCode")
	data, err := u.RepoThemSoHeThong.GetThamSoByCode(ctx, request.Code)

	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Erro server: %v", err)
	}

	var dataResponse []*pb.ThamSoHeThong
	for _, item := range data {
		nv := pb.ThamSoHeThong{
			Id:        int64(item.Id),
			MaThamSo:  item.MaThamSo,
			TenThamSo: item.TenThamSo,
			GiaTri:    item.GiaTri,
			DienGiai:  item.DienGiai,
		}
		dataResponse = append(dataResponse, &nv)
	}

	res := &pb.ThamSoHeThongsResponse{
		Thamso: dataResponse,
	}
	return res, nil
}

func (u *UserService) GetThamSoByCode(ctx context.Context, request *pb.ThamSoHeThongRequest) (*pb.ThamSoHeThongResponse, error) {
	fmt.Println("Call : GetThamSoByCode")
	data, err := u.RepoThemSoHeThong.GetThamSoValueByCode(ctx, request.Code)

	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Erro server: %v", err)
	}

	res := pb.ThamSoHeThongResponse{
		Thamso: &pb.ThamSoHeThong{
			Id:        int64(data.Id),
			MaThamSo:  data.MaThamSo,
			TenThamSo: data.TenThamSo,
			GiaTri:    data.GiaTri,
			DienGiai:  data.DienGiai,
		},
	}

	return &res, nil
}

func (u *UserService) GetToken(ctx context.Context, request *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	fmt.Println("Call : GetToken")
	data, _, err := auth.GenTokenNhanVienId(int(request.NhanVienId), 168*time.Hour)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Erro server: %v", err)
	}
	res := &pb.GetTokenResponse{
		Token: data,
	}
	return res, nil
}

func (u *UserService) GetDanhSachNhanVien(ctx context.Context, request *pb.ReadRequest) (*pb.DanhSachNhanVienResponse, error) {
	fmt.Println("Call : GetDanhSachNhanVien")
	dsNhanVien, err := u.RepoNhanVien.GetAll()

	if err != nil {
		log.Println(err)
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

	res := &pb.DanhSachNhanVienResponse{
		NhanVien: dataResponse,
	}
	return res, nil
}

func (u *UserService) GetPhongBanNhanVien(ctx context.Context, request *pb.PhongBanNhanVienRequest) (*pb.DanhSachNhanVienResponse, error) {
	fmt.Println("Call : GetPhongBanNhanVien")

	dsNhanVien, err := u.RepoNhanVien.GetNhanVienByPhongBanId(int(request.PhongBanId))
	if err != nil {
		fmt.Println(err)
		return nil, err
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
	res := &pb.DanhSachNhanVienResponse{
		NhanVien: dataResponse,
	}

	return res, nil
}

func (u *UserService) Ping(context.Context, *pb.ReadRequest) (*pb.PingResponse, error) {
	dataResponse := pb.PingResponse{
		Message: "Hello",
	}
	return &dataResponse, nil
}

func (u *UserService) GetNhanVienById(ctx context.Context, request *pb.ReadRequest) (*pb.NhanVienResponse, error) {
	item, err := u.RepoNhanVien.GetNhanVienByNhanVienId(int(request.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro server: %v", err)
	}

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

	res := &pb.NhanVienResponse{
		NhanVien: &nv,
	}
	return res, nil
}
