package mod

import (
	"github.com/AkuLaper/akulaperbe/mod/daftar"
	"github.com/AkuLaper/akulaperbe/mod/helpdesk"
	"github.com/AkuLaper/akulaperbe/mod/idgrup"
	"github.com/AkuLaper/akulaperbe/mod/kyc"
	"github.com/AkuLaper/akulaperbe/mod/lms"
	"github.com/AkuLaper/akulaperbe/mod/lmsdesa"
	"github.com/AkuLaper/akulaperbe/mod/posint"
	"github.com/AkuLaper/akulaperbe/mod/presensi"
	"github.com/AkuLaper/akulaperbe/mod/siakad"
	"github.com/AkuLaper/akulaperbe/mod/tasklist"
	"github.com/AkuLaper/akulaperbe/mod/unsubscribe"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

func Caller(Profile itmodel.Profile, Modulename string, Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
	switch Modulename {
	case "unsubscribe":
		reply = unsubscribe.Unsubscribe(Pesan, db)
	case "idgrup":
		reply = idgrup.IDGroup(Pesan)
	case "feedbackhelpdesk":
		reply = helpdesk.FeedbackHelpdesk(Profile, Pesan, db)
	case "endhelpdesk":
		reply = helpdesk.EndHelpdesk(Profile, Pesan, db)
	case "helpdesk":
		reply = helpdesk.HelpdeskPDLMS(Profile, Pesan, db)
	case "helpdeskpusat":
		reply = helpdesk.HelpdeskPusat(Profile, Pesan, db)
	case "adminopenusertiket":
		reply = helpdesk.AdminOpenSessionCurrentUserTiket(Profile, Pesan, db)
	case "presensi-masuk":
		reply = presensi.PresensiMasuk(Pesan, db)
	case "presensi-pulang":
		reply = presensi.PresensiPulang(Pesan, db)
	case "upload-lmsdesa-file":
		reply = lmsdesa.ArsipFile(Pesan, db)
	case "upload-lmsdesa-gambar":
		reply = lmsdesa.ArsipGambar(Pesan, db)
	case "lms":
		reply = lms.ReplyRekapUsers(Profile, Pesan, db)
	case "cek-ktp":
		reply = kyc.CekKTP(Profile, Pesan, db)
	case "selfie-masuk":
		reply = presensi.CekSelfieMasuk(Profile, Pesan, db)
	case "selfie-pulang":
		reply = presensi.CekSelfiePulang(Profile, Pesan, db)
	case "tasklist-append":
		reply = tasklist.TaskListAppend(Pesan, db)
	case "tasklist-reset":
		reply = tasklist.TaskListReset(Pesan, db)
	case "tasklist-save":
		reply = tasklist.TaskListSave(Pesan, db)
	case "domyikado-user":
		reply = daftar.DaftarDomyikado(Pesan, db)
	case "login-siakad":
		reply = siakad.LoginSiakad(Pesan, db)
	case "minta-bap":
		reply = siakad.MintaBAP(Pesan, db)
	case "approve-bimbingan":
		reply = siakad.ApproveBimbingan(Pesan, db)
	case "approve-bimbinganbypoin":
		reply = siakad.ApproveBimbinganbyPoin(Pesan, db)
	case "prohibited-items":
		reply = posint.GetProhibitedItems(Pesan, db)
	}

	return
}
