package utils

// [Clien 클리앙], clien
// https://www.clien.net/service/board/jirum?&od=T31&po=0
const ClienBaseUrl = "https://www.clien.net"
const ClienResourcePath = "/service/board/jirum?&od=T31&po="
const ClienTrSelector = "div.list_item.symph_row"
const ClienSubSelector = "span.list_subject > a:nth-of-type(1)"
const ClienTitleSelector = "span.list_subject > a:nth-of-type(1)"
const ClienCategorySelector = "a.icon_keyword"
const ClienDateSelector = "#div_content > div.post_view > div.post_author > span:nth-child(1)"
const ClienContentSelector = "div.post_content > article"

// [CoolenJoy 쿨앤조이], cool
const CoolBaseUrl = "https://coolenjoy.net"
const CoolResourcePath = "/bbs/jirum/p"
const CoolTrSelector = "#fboardlist > div > table > tbody > tr:not(.bo_notice)"
const CoolSubSelector = "td a"
const CoolTitleSelector = "#bo_v_title"
const CoolDateSelector = "#bo_v_info > strong:nth-child(4)"
const CoolContentSelector = "#bo_v_con"

// [eomisae 어미새], eomi
// https://eomisae.co.kr/index.php?mid=fs&sort_index=regdate&order_type=desc&page=1
const EomiBaseUrl = "https://eomisae.co.kr"
const EomiResourcePath = "/index.php?mid=fs&sort_index=regdate&order_type=desc&page="
const EomiTrSelector = "div.card_el"
const EomiSubSelector = "h3 > a"
const EomiTitleSelector = "#D_ > div._wrapper > div._hd.clear > div._section > h2 > a"
const EomiDateSelector = "span:nth-of-type(2)"
const EomiContentSelector = "article > div:nth-of-type(1)"

// [ppomppu 뽐뿌국내], ppomin
// http://www.ppomppu.co.kr/zboard/zboard.php?id=ppomppu&divpage=63&page=1
const PpominBaseUrl = "https://www.ppomppu.co.kr"
const PpominResourcePath = "/zboard/zboard.php?id=ppomppu&divpage=63&page="
const PpominTrSelector = "tr.list0, tr.list1"
const PpominSubSelector = "td[valign=\"middle\"] > a"
const PpominTitleSelector = "td[valign=\"middle\"] > a"
const PpominCategorySelector = "td.han4.list_vspace > nobr"
const PpominDateSelector = "div.sub-top-contents-box"
const PpominContentSelector = "td.board-contents[align=\"left\"]"

// [ppomppu 뽑뿌해외], ppomex
// http://www.ppomppu.co.kr/zboard/zboard.php?id=ppomppu4&divpage=21&page=1
const PpomexBaseUrl = "https://www.ppomppu.co.kr"
const PpomexResourcePath = "/zboard/zboard.php?id=ppomppu4&divpage=21&page="
const PpomexSubLinkPath = "/zboard/"
const PpomexTrSelector = "tr.list0, tr.list1"
const PpomexSubSelector = "td[valign=\"middle\"] > a"
const PpomexTitleSelector = "td[valign=\"middle\"] > a"
const PpomexCategorySelector = "td.han4.list_vspace > nobr"
const PpomexDateSelector = "div.sub-top-contents-box"
const PpomexContentSelector = "td.board-contents[align=\"left\"]"

// [quasarzon 퀘이사존], quas
// https://quasarzone.com/bbs/qb_saleinfo?page=1
const QuasBaseUrl = "https://quasarzone.com"
const QuasResourcePath = "/bbs/qb_saleinfo?page="
const QuasTrSelector = "div.market-type-list > table > tbody > tr"
const QuasSubSelector = "p.tit > a"
const QuasTitleSelector = "p.tit > a"
const QuasCategorySelector = "span.category"
const QuasUserSelector = "span.user-nick-wrap"
const QuasDateSelector = "#content > div.sub-content-wrap > div.left-con-wrap > div.common-view-wrap.market-info-view-wrap > div > dl > dt > div.util-area > p > span"
const QuasContentSelector = "div.view-content > div.note-editor"
var QuasSellers = []string{
	"alphascan", "LG대장", "Micronics", "AntecKorea", "3RSYS", "이엠텍홍보팀", "루선배",
	"제이웍스코리아", "LG온라인샵", "제닉스홀릭", "LGEMONITOR", "삼성디스플레이", "HyperX",
	"waycos", "인터픽셀", "컴디씨", "기가바이트", "스카이디지탈", "컴대문", "컴스빌", "터틀비치",
	"COUGAR게이밍", "intech", "제이씨현시스템", "MAXTILL", "에이수스코리아", "잘만테크", "BRAVOTEC",
	"서린씨앤아이", "한성컴퓨터"}





