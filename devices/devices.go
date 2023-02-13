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

// ============================================================================Котёл 2А============================================================================

var k2a_1 = []string{"Q_per_par_3a", "Q_pit_voda_3a"}

var k2a_2 = []string{"P_para_2a", "T_para_nitk_A_2a", "T_para_nitk_B_2a"}

var k2a_3 = []string{"T_pitvod_kotl_2a", "T_mazut_2a", "T_uhodgaz_ventur1_2a", "T_uhodgaz_ventur2_2a", "T_uhodgaz_ventur3_2a", "T_uhodgaz_ventur4_2a"}

var k2a_4 = []string{"T_vozdh_pered_rvp_2a", "T_vozdh_pered_tvpslev_2a", "T_vozdh_pered_tvpsprav_2a"}

var k2a_5 = []string{"T_vozdh_posl_rvp_2a", "T_vozdh_posl_tvpslev_2a", "T_vozdh_posl_tvpsprav_2a", "T_perd_ekonomiz_2a",
	"T_dym_zarvp_2a", "T_dym_zatvp_slev_2a", "T_dym_zatvp_sprav_2a"}

var k2a_6 = []string{"O2_slev_2a"}

var k2a_7 = []string{"O2_sprav_2a"}

var k2a_8 = []string{"Q_gaz_2a"}

// ============================================================================Котёл 2Б============================================================================

var k2b_1 = []string{"Q_per_par_2b", "Q_pit_voda_2b"}

var k2b_2 = []string{"P_para_2b", "T_para_nitk_A_2b", "T_para_nitk_B_2b"}

var k2b_3 = []string{"T_pitvod_kotl_2b", "T_mazut_2b", "T_uhodgaz_ventur1_2b", "T_uhodgaz_ventur2_2b", "T_uhodgaz_ventur3_2b", "T_uhodgaz_ventur4_2b"}

var k2b_4 = []string{"T_vozdh_pered_rvp_2b", "T_vozdh_pered_tvp_2b"}

var k2b_5 = []string{"T_vozdh_posl_rvp_2b", "T_vozdh_posl_tvpslev_2b", "T_vozdh_posl_tvpsprav_2b", "T_perd_ekonomiz_2b",
	"T_dym_zarvp_2b", "T_dym_zatvp_slev_2b", "T_dym_zatvp_sprav_2b"}

var k2b_6 = []string{"O2_slev_2b"}

var k2b_7 = []string{"O2_sprav_2b"}

var k2b_8 = []string{"Q_gaz_2b"}

// ============================================================================Турбина 2============================================================================

var t2_1 = []string{"T_ostrpar_pered_2GPZ2_nitka_A"}

var t2_2 = []string{"T_ostrpar_pered_2GPZ2_nitka_B"}

var t2_3 = []string{"P_para_do_stopor_klapan_nitka_A"}

var t2_4 = []string{"P_para_do_stopor_klapan_nitka_B"}

var t2_5 = []string{"Q_para_turbina_nitka_A"}

var t2_6 = []string{"Q_para_turbina_nitka_B"}

var t2_7 = []string{"vakum"}

var t2_8 = []string{"T_cirk_voda_k_kondesator_N1", "T_cirk_voda_k_kondesator_N2", "T_cirk_voda_mezh_kondesator_N1", "T_cirk_voda_mezh_kondesator_N2",
	"T_cirk_voda_sliv_N1", "T_cirk_voda_sliv_N2", "T_maslo_do_MO", "T_ohl_voda_do_MO"}

var t2_9 = []string{"T_kondensat_posle_PND4", "T_set_voda_do_PSG1_0m", "T_set_voda_do_PSG1_3m", "T_set_voda_za_PSG1", "T_set_voda_za_PSG2", "T_cirk_voda_vstroenih_puchkov"}

var t2_10 = []string{"P_para_psg1"}

var t2_11 = []string{"P_para_psg2"}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
// ============================================================================ИРТ5920Н============================================================================

//var irt_1 = new_device(1, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x37, 0x36, 0x32, 0x37, 0x0D}, k2a_6) // панель:3-2N  котёл:2A

//var irt_2 = new_device(2, "irt5920n", 1, []byte{0xFF, 0x3A, 0x32, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x31, 0x31, 0x39, 0x37, 0x39, 0x0D}, k2a_7) // панель:3-2N  котёл:2A

//var irt_3 = new_device(3, "irt5920n", 1, []byte{0xFF, 0x3A, 0x33, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x36, 0x35, 0x34, 0x38, 0x32, 0x0D}, k2a_8) // панель:3-2N  котёл:2A

var irt_4 = new_device(4, "irt5920n", 1, []byte{0xFF, 0x3A, 0x34, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x31, 0x38, 0x36, 0x33, 0x35, 0x0D}, k2b_6) // панель:3-3N  котёл:2Б

//var irt_5 = new_device(5, "irt5920n", 1, []byte{0xFF, 0x3A, 0x35, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x33, 0x39, 0x33, 0x37, 0x30, 0x0D}, k2b_7) // панель:3-3N  котёл:2Б

var irt_6 = new_device(6, "irt5920n", 1, []byte{0xFF, 0x3A, 0x36, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x34, 0x33, 0x37, 0x32, 0x32, 0x0D}, k2b_8) // панель:3-3N  котёл:2Б

var irt_7 = new_device(7, "irt5920n", 1, []byte{0xFF, 0x3A, 0x37, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x33, 0x31, 0x36, 0x39, 0x31, 0x0D}, t2_10) // панель:2S-2  турбина:2

//var irt_8 = new_device(8, "irt5920n", 1, []byte{0xFF, 0x3A, 0x38, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x33, 0x33, 0x39, 0x39, 0x35, 0x0D}, t2_11) // панель:2S-2  турбина:2

var irt_9 = new_device(9, "irt5920n", 1, []byte{0xFF, 0x3A, 0x39, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x32, 0x31, 0x39, 0x36, 0x32, 0x0D}, t2_1) // панель:2S-2  турбина:2

var irt_10 = new_device(10, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x30, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x35, 0x33, 0x36, 0x31, 0x0D}, t2_2) // панель:2S-2  турбина:2

var irt_11 = new_device(11, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x31, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x35, 0x30, 0x36, 0x37, 0x32, 0x0D}, t2_3) // панель:2S1  турбина:2

var irt_12 = new_device(12, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x32, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x36, 0x33, 0x32, 0x31, 0x36, 0x0D}, t2_4) // панель:2S1  турбина:2

var irt_13 = new_device(13, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x33, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x31, 0x30, 0x32, 0x32, 0x35, 0x0D}, t2_5) // панель:2S1  турбина:2

var irt_14 = new_device(14, "irt5920n", 1, []byte{0xFF, 0x3A, 0x31, 0x34, 0x3B, 0x31, 0x3B, 0x30, 0x3B, 0x33, 0x37, 0x31, 0x30, 0x34, 0x0D}, t2_6) // панель:2S1  турбина:2

// ============================================================================Ф1772============================================================================

//var F_1 = new_device(15, "f1772", 2, []byte{0x0f, 0x04, 0x02, 0x00, 0x00, 0x04, 0xf1, 0x5f}, k2a_1) // панель:2-2N  котёл:2А

//var F_2 = new_device(16, "f1772", 3, []byte{0x10, 0x04, 0x02, 0x00, 0x00, 0x06, 0x72, 0xf1}, k2a_2) // панель:1-2N  котёл:2А

var F_3 = new_device(17, "f1772", 2, []byte{0x11, 0x04, 0x02, 0x00, 0x00, 0x04, 0xf2, 0xe1}, k2b_1) // панель:2-3N  котёл:2Б

var F_4 = new_device(18, "f1772", 3, []byte{0x12, 0x04, 0x02, 0x00, 0x00, 0x06, 0x73, 0x13}, k2b_2) // панель:1-3N  котёл:2Б

var F_5 = new_device(19, "f1772", 1, []byte{0x13, 0x04, 0x02, 0x00, 0x00, 0x02, 0x73, 0x01}, t2_7) // панель:2U  турбина:2

// ============================================================================ТМ5104Д============================================================================

//var tm_1 = new_device(20, "tm5104d", 7, []byte{0x14, 0x03, 0x05, 0x00, 0x00, 0x0e, 0xc6, 0x07}, k2a_5) // панель:3a-2N  котёл:2A

var tm_2 = new_device(21, "tm5104d", 7, []byte{0x15, 0x03, 0x05, 0x00, 0x00, 0x0e, 0xc7, 0xd6}, k2b_5) // панель:3a-3N  котёл:2Б

// ============================================================================ТРМ138============================================================================

//var trm_1 = new_device(24, "trm138", 6, []byte{0x18, 0x04, 0x00, 0x00, 0x00, 0x1E, 0x72, 0x0b}, k2a_3) // панель:1-2N  котёл:2A

//var trm_2 = new_device(32, "trm138", 6, []byte{0x20, 0x04, 0x00, 0x00, 0x00, 0x1E, 0x76, 0xb3}, k2b_3) // панель:1-3N  котёл:2Б

//var trm_3 = new_device(40, "trm138", 3, []byte{0x28, 0x04, 0x00, 0x00, 0x00, 0x0f, 0xb7, 0xf7}, k2a_4) // панель:3a-2N  котёл:2A

//var trm_4 = new_device(48, "trm138", 2, []byte{0x30, 0x04, 0x00, 0x00, 0x00, 0x0a, 0x74, 0x2c}, k2b_4) // панель:3a-3N  котёл:2Б

//var trm_5 = new_device(56, "trm138", 8, []byte{0x38, 0x04, 0x00, 0x00, 0x00, 0x28, 0xf5, 0x7d}, t2_8) // панель:2S-2a  турбина:2

//var trm_6 = new_device(64, "trm138", 6, []byte{0x40, 0x04, 0x00, 0x00, 0x00, 0x1e, 0x7f, 0x13}, t2_9) // панель:2S-4  турбина:2

// ========================================================================================================================================================

var Devices = []Device{irt_4}
