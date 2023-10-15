package modelproc

import "github.com/sakurtek/goserver/bookingremyconcept/pkg/model"

func GetOnlineNews() []model.Newsline {

	var headline []model.Newsline
	var new1, new2, new3, new4 model.Newsline

	new1.SetConstructor(123, "Go is the best", "Go merupakan bahasa yang sangat baik yang diluncurkan oleh Google...", "begin - ghajd fjah jhasd fhja gjhaj ghajs dfhj ghaj gjahs djhajdhf ahj gjag ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a  ga dhf ajhgajhsdjf hajghajshd fjhaj hasjkgh ajsdhf jah gjkashd jfha jghajskdhfjahgjahsd jhgajh gjkahsd jkfhasjdgh jkashg jash jahs djfhajsdhg jkah jkadshfjah gjhasdjkf hajghjaks hgjhasjdf jakh gjkashdjkhasjkdfh ajkhg jkashdjk ahs gkjhasdjkf hajkgh jaksdh jkf ajhg ajksg. \n\n<strong>BARIS BARU</strong> fjah jhasd fhja gjhaj ghajs dfhj ghaj gjahs djhajdhf ahj gjag ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a  ga dhf ajhgajhsdjf hajghajshd fjhaj hasjkgh ajsdhf jah gjkashd jfha jghajskdhfjahgjahsd jhgajh gjkahsd jkfhasjdgh jkashg jash jahs djfhajsdhg jkah jkadshfjah gjhasdjkf hajghjaks hgjhasjdf jakh gjkashdjkhasjkdfh ajkhg jkashdjk ahs gkjhasdjkf hajkgh jaksdh jkf ajhg ajksg--end", "stendy", "koresponden tahuna")

	new2.SetConstructor(234, "Why using GOlang for Web Developer??", "Mengapa harus menggunakan Golang? kita sudah memiliki PHP, Django Rust dan lainnya...", "begin - ghajd fjah jhasd fhja gjhaj ghajs dfhj ghaj gjahs djhajdhf ahj gjag ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg aajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a--- end", "stendy", "koresponden tahuna")

	new3.SetConstructor(345, "Golang For make own API REST.", "Kemampuan Golang bukan hanya untuk membuat webdev namun sebagai Backend untuk membangunan RESTFUL API...", "begin - ghajd fjah jhasd fhja gjhaj ghajs dfhj ghaj gjahs djhajdhf ahj gjag ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg aajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg aajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a--- end", "stendy", "koresponden tahuna")

	new4.SetConstructor(456, "Golang The programming Language is Easy, Kompatibel and Good Performance", "Golang dibuat untuk memudahkan programmer dalam membuat program untuk crossplatform...", "begin - ghajd fjah jhasd fhja gjhaj ghajs dfhj ghaj gjahs djhajdhf ahj gjag ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a ajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg aajsd hjfha jghas jkdhfjah gjha sjdhd jhf ajhg ajs gjashd jhgjsdh gajksh gjahsd gjha jgha sjgha sjdhf ja gjah gjhasd jfajdhg ajsg ajsdhg jash gjas gjhas jdghajskg hajsdh jkahsgasdasg a--- end", "stendy", "koresponden tahuna")

	headline = append(headline, new1)
	headline = append(headline, new2)
	headline = append(headline, new3)
	headline = append(headline, new4)

	/*
		onlineNews := []model.Newsline{
			{
				MyTitle: "title 1",
				Summery: "summersy 1",
				Writer: model.Personal{
					Nama:   "stendy",
					Status: "koresponden tahun",
				},
			},
			{
				MyTitle: "titel 2",
				Summery: "Summery 2",
				Writer: model.Personal{
					Nama:   "kahfi",
					Status: "koresponedn jakarta",
				},
			},
		}

	*/
	return headline
}

func GetDataKaryawan() []model.Karyawan {

	var karyawan []model.Karyawan
	var kahfi, almeera, winny, stendy model.Karyawan

	kahfi.SetConstructor("Muhammad Isa Al kahfi sakur", "kahfi.sakur@gmail.com", 8, 1000000)
	almeera.SetConstructor("Almeera Qiana Azzahra Sakur", "almeera.sakur@gmail.com", 2, 200000)
	winny.SetConstructor("Winnya ayu perwitasari", "winny.ayu@gmail.com", 30, 300000)
	stendy.SetConstructor("Stendy Sakur", "sakur.stendy@gmail.com", 40, 200000)

	karyawan = append(karyawan, kahfi)
	karyawan = append(karyawan, almeera)
	karyawan = append(karyawan, winny)
	karyawan = append(karyawan, stendy)

	/*
		karyawan := []model.Karyawan{
			{
				NamaKaryawan:  "almeera qiana azzahra sakur",
				UmurKaryawan:  2,
				EmailKaryawan: "almeera.sakur@gmail.com",
				GajiBulanan:   1200000,
			},
			{
				NamaKaryawan:  "muhammad isa alkahfi sakur",
				UmurKaryawan:  8,
				EmailKaryawan: "kahfi.sakur@gmail.com",
				GajiBulanan:   2700000,
			},
			{
				NamaKaryawan:  "winny ayu perwitasari",
				UmurKaryawan:  30,
				EmailKaryawan: "winny.ayu@gmail.com",
				GajiBulanan:   3000000,
			},
		}
	*/

	return karyawan
}
