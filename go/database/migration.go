package database

type Listbata struct {
	IdList      int    `gorm:"primarykey;autoIncrement" json:"id_list"`
	Namabrg     string `gorm:"column:nama_brg" json:"nama_brg"`
	Foto        []byte `gorm:"column:foto" json:"foto"`
	EncodedFoto string `json:"-"`
	Detailbrg   string `gorm:"column:detail_brg" json:"detail_brg"`
	Jumlah      int    `gorm:"column:jumlah" json:"jumlah"`
	Harga       string `gorm:"column:harga_brg" json:"harga_brg"`
}

type Bata struct {
	IdBata      int    `gorm:"primarykey;autoIncrement" json:"id_bata"`
	Namabrg     string `gorm:"column:nama_brg" json:"nama_brg"`
	Foto        []byte `gorm:"column:foto" json:"foto"`
	EncodedFoto string `json:"-"`
	Detailbrg   string `gorm:"column:detail_brg" json:"detail_brg"`
	Jumlah      int    `gorm:"column:jumlah" json:"jumlah"`
	Harga       string `gorm:"column:harga_brg" json:"harga_brg"`
}

type Roster struct {
	IdRoster    int    `gorm:"primarykey;autoIncrement" json:"id_roster"`
	Namabrg     string `gorm:"column:nama_brg" json:"nama_brg"`
	Foto        []byte `gorm:"column:foto" json:"foto"`
	EncodedFoto string `json:"-"`
	Detailbrg   string `gorm:"column:detail_brg" json:"detail_brg"`
	Jumlah      int    `gorm:"column:jumlah" json:"jumlah"`
	Harga       string `gorm:"column:harga_brg" json:"harga_brg"`
}

type User struct {
	IdUser   int    `gorm:"primaryKey;autoIncrement" json:"id_user"`
	Username string `gorm:"column:username" json:"username"`
	Nama     string `gorm:"column:nama_user" json:"nama_user"`
	Email    string `gorm:"column:email" json:"email"`
	NoTelp   string `gorm:"column:no_telp" json:"no_telp"`
	Alamat   string `gorm:"column:alamat" json:"alamat"`
	Password string `gorm:"column:password" json:"-"`
}

type Riwayat struct {
	IdPesanan  int    `gorm:"primaryKey;autoIncrement" json:"id_pesanan"`
	idUser     int    `gorm:"column:id_user" json:"id_user"`
	Namabrg    string `gorm:"column:nama_brg" json:"nama_brg"`
	Jumlah     string `gorm:"column:jumlah" json:"jumlah"`
	TotalHarga string `gorm:"column:total_harga" json:"total_harga"`
}

type Keranjang struct {
	IdBrg    int    `gorm:"primaryKey;autoIncrement" json:"id_brg"`
	IdUser   int    `gorm:"column:id_user" json:"id_user"`
	Namabrg  string `gorm:"column:nama_brg" json:"nama_brg"`
	Jumlah   string `gorm:"column:jumlah" json:"jumlah"`
	Hargabrg string `gorm:"column:harga_brg" json:"harga_brg"`
}

func (Listbata) TableName() string {
	return "listbata"
}

func (Roster) TableName() string {
	return "roster"
}

func (Bata) TableName() string {
	return "bata"
}

func (User) TableName() string {
	return "user"
}

func (Riwayat) TableName() string {
	return "riwayat"
}

func (Keranjang) TableName() string {
	return "keranjang"
}
