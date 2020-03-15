package entity

type KetQuaDangKy struct {
	ID           int    `json:"id"`
	MailQuanLy   string `json:"mail_quan_ly"`
	TenQuanLy    string `json:"ten_quan_ly"`
	TenNhanVien  string `json:"ten_nhan_vien"`
	MailNhanVien string `json:"mail_nhan_vien"`
	TenCaLamViec string `json:"ten_ca_lam_viec"`
	NgayPhep     string `json:"ngay_phep"`
}
