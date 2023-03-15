package devices

type Device struct {
	Id          int
	Name        string
	Qan_chanels int32
	Request     []byte
	Chanels     []string
}

func new_device(adress int, name string, qantity int32, req []byte, ch []string) Device {
	return Device{Id: adress, Name: name, Qan_chanels: qantity, Request: req, Chanels: ch}
}

// ============================================================================Котёл 1============================================================================

var k1_1 = []string{"o2_slev"}

var k1_2 = []string{"o2_sprav"}

var k1_3 = []string{"q_gaz"}

var k1_4 = []string{"t_para_nitk_a", "t_para_nitk_b", "p_para"}

var k1_5 = []string{"q_per_par", "q_pit_voda"}

var k1_6 = []string{"t_vozdh_posl_rvpa", "t_vozdh_posl_rvpb", "t_vozdh_posl_tvp_slev", "t_vozdh_posl_tvp_sprav", "t_posle_ekonomiz_slev",
	"t_posle_ekonomiz_sprav", "t_dym_zarvp_slev", "t_dym_zarvp_sprav", "t_dym_zatvp_slev", "t_dym_zatvp_sprav"}

var k1_7 = []string{"t_pitvod_kotl", "t_uhodgaz_ventur1", "t_uhodgaz_ventur2", "t_mazut", "t_uhodgaz_ventur3", "t_uhodgaz_ventur4"}

var k1_8 = []string{"t_kontensat_sbor", "t_kontensat_vihkollect", "t_vod_kollect_slev", "t_vod_kollect_sprav", "t_gaz_dsa", "t_gaz_dsb"}

var k1_9 = []string{"t_vozdh_pered_rvpa", "t_vozdh_pered_rvpb", "t_vozdh_pered_tvpslev", "t_vozdh_pered_tvpsprav"}

// ============================================================================Турбина 1============================================================================

var t1_1 = []string{"p_par_do_sk"}

var t1_2 = []string{"q_par"}

var t1_3 = []string{"t_par_posle_sk"}

var t1_4 = []string{"vakum"}

var t1_5 = []string{"t_par_psg1_prov1", "t_par_psg1_prov2", "t_par_psg2", "t_vod_do_psg1", "t_vod_posle_psg1", "t_vod_vihod_psg12"}

var t1_6 = []string{"t_par_posle_ou1", "t_par_posle_ou2", "t_kondesat_do kn", "t_cirk_vod_posle_kondensat", "t_cirk_vod_do_kondensat",}

var t1_7 = []string{"t_par_cnd_sprav", "t_par_cnd_slev", "t_par_uplotn_kollect", "t_kondest_posle_pnd4", "t_vod_obvod_pvd7", 
    "t_par_pered_pn130", "t_vod_posle_pvd7"}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
// ============================================================================ИРТ5920Н============================================================================

var irt_1 = new_device(1, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x37, 0x36, 0x32, 0x37, 0x0D}, k1_1) // панель:17-1N  котёл:1

var irt_2 = new_device(2, "irt5920n", 1, []byte{0xFF, 0x3A, 0x32, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x31, 0x31, 0x39, 0x37, 0x39, 0x0D}, k1_2) // панель:17-1N  котёл:1

var irt_3 = new_device(3, "irt5920n", 1, []byte{0xFF, 0x3A, 0x33, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x36, 0x35, 0x34, 0x38, 0x32, 0x0D}, k1_3) // панель:1-1N  котёл:1

var irt_4 = new_device(3, "irt5920n", 1, []byte{0xFF, 0x3A, 0x33, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x36, 0x35, 0x34, 0x38, 0x32, 0x0D}, t1_1) // панель:3-1S  турбина:1

var irt_5 = new_device(4, "irt5920n", 1, []byte{0xFF, 0x3A, 0x34, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x31, 0x38, 0x36, 0x33, 0x35, 0x0D}, t1_2) // панель:3-1S  турбина:1

var irt_6 = new_device(5, "irt5920n", 1, []byte{0xFF, 0x3A, 0x35, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x33, 0x39, 0x33, 0x37, 0x30, 0x0D}, t1_3) // панель:3-1S  турбина:1

// ============================================================================Ф1772============================================================================

var f_1 = new_device(5, "f1772", 3, []byte{0x05, 0x04, 0x02, 0x00, 0x00, 0x06, 0x70, 0x34}, k1_4) // панель:16-1N  котёл:1

var f_2 = new_device(6, "f1772", 2, []byte{0x06, 0x04, 0x02, 0x00, 0x00, 0x04, 0xf1, 0xc6}, k1_5) // панель:1-1N  котёл:1

var f_3 = new_device(6, "f1772", 1, []byte{0x06, 0x04, 0x02, 0x00, 0x00, 0x02, 0x71, 0xc4}, t1_4) // панель:1-1S  турбина:1

// ============================================================================ТМ5104Д============================================================================

var tm_1 = new_device(7, "tm5104d", 10, []byte{0x07, 0x03, 0x05, 0x00, 0x00, 0x14, 0x45, 0x6f}, k1_6) // панель:3a-1N  котёл:1

// ============================================================================ТРМ138============================================================================

var trm_1 = new_device(8, "trm138", 6, []byte{0x08, 0x04, 0x00, 0x00, 0x00, 0x1e, 0x70, 0x9b}, k1_7) // панель:17-1N  котёл:1

var trm_2 = new_device(16, "trm138", 6, []byte{0x10, 0x04, 0x00, 0x00, 0x00, 0x1e, 0x73, 0x43}, k1_8) // панель:2A-1N  котёл:1

var trm_3 = new_device(24, "trm138", 4, []byte{0x18, 0x04, 0x00, 0x00, 0x00, 0x14, 0xf2, 0x0c}, k1_9) // панель:3A-1N  котёл:1

var trm_4 = new_device(8, "trm138", 6, []byte{0x08, 0x04, 0x00, 0x00, 0x00, 0x1e, 0x70, 0x9b}, t1_5) // панель:1-1S  турбина:1

var trm_5 = new_device(16, "trm138", 5, []byte{0x10, 0x04, 0x00, 0x00, 0x00, 0x19, 0x32, 0x81}, t1_6) // панель:1-1S  турбина:1

var trm_6 = new_device(24, "trm138", 7, []byte{0x18, 0x04, 0x00, 0x00, 0x00, 0x23, 0xb3, 0xda}, t1_7) // панель:3A-1S  турбина:1

// ========================================================================================================================================================

var Devices_kettle = []Device{irt_1, irt_2, irt_3, f_1, f_2, tm_1, trm_1, trm_2, trm_3}

var Devices_turbo = []Device{irt_4, irt_5, irt_6, f_3, trm_4, trm_5, trm_6}
