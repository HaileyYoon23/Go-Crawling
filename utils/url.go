package utils

// [Clien 클리앙], clien
// https://www.clien.net/service/board/jirum?&od=T31&po=0
const clienBaseUrl = "https://www.clien.net"
const clienResourcePath = "/service/board/jirum?&od=T31&po="
const clienTrSelector = "div.list_item.symph_row"
const clienSubSelector = "span.list_subject > a:nth-of-type(1)"
const clienTitleSelector = "span.list_subject > a:nth-of-type(1)"
const clienCategorySelector = "a.icon_keyword"
const clienDateSelector = "#div_content > div.post_view > div.post_author > span:nth-child(1)"
const clienContentSelector = "div.post_content > article"

// [CoolenJoy 쿨앤조이], cool
const coolBaseUrl = "https://coolenjoy.net"
const coolResourcePath = "/bbs/jirum/p"
const coolTrSelector = "#fboardlist > div > table > tbody > tr:not(.bo_notice)"
const coolSubSelector = "td a"
const coolTitleSelector = "#bo_v_title"
const coolDateSelector = "#bo_v_info > strong:nth-child(4)"
const coolContentSelector = "#bo_v_con"

// [eomisae 어미새], eomi
// https://eomisae.co.kr/index.php?mid=fs&sort_index=regdate&order_type=desc&page=1
const eomiBaseUrl = "https://eomisae.co.kr"
const eomiResourcePath = "/index.php?mid=fs&sort_index=regdate&order_type=desc&page="
const eomiTrSelector = "div.card_el"
const eomiSubSelector = "h3 > a"
const veomitTitleSelector = "#D_ > div._wrapper > div._hd.clear > div._section > h2 > a"
const eomiDateSelector = "span:nth-of-type(2)"
const eomiContentSelector = "article > div:nth-of-type(1)"

// [ppomppu 뽐뿌국내], ppomin
// http://www.ppomppu.co.kr/zboard/zboard.php?id=ppomppu&divpage=63&page=1
const ppominBaseUrl = "http://www.ppomppu.co.kr"
const ppominResourcePath = "/zboard/zboard.php?id=ppomppu&divpage=63&page="
const ppominTrSelector = "tr.list0, tr.list1"
const ppominSubSelector = "td[valign=\"middle\"] > a"
const ppominTitleSelector = "td[valign=\"middle\"] > a"
const ppominCategorySelector = "td.han4.list_vspace > nobr"
const ppominDateSelector = "div.sub-top-contents-box"
const ppominContentSelector = "td.board-contents[align=\"left\"]"

// [ppomppu 뽑뿌해외], ppomex
// http://www.ppomppu.co.kr/zboard/zboard.php?id=ppomppu4&divpage=21&page=1
const ppomexBaseUrl = "http://www.ppomppu.co.kr"
const ppomexResourcePath = "/zboard/zboard.php?id=ppomppu4&divpage=21&page="
const ppomexTrSelector = "tr.list0, tr.list1"
const ppomexSubSelector = "td[valign=\"middle\"] > a"
const ppomexTitleSelector = "td[valign=\"middle\"] > a"
const ppomexCategorySelector = "td.han4.list_vspace > nobr"
const ppomexDateSelector = "div.sub-top-contents-box"
const ppomexContentSelector = "td.board-contents[align=\"left\"]"

// [quasarzon 퀘이사존], quas
// https://quasarzone.com/bbs/qb_saleinfo?page=1
const quasBaseUrl = "https://quasarzone.com"
const quasResourcePath = "/bbs/qb_saleinfo?page="
const quasTrSelector = "div.market-type-list > table > tbody > tr"
const quasSubSelector = "p.tit > a"
const quasTitleSelector = "p.tit > a"
const quasCategorySelector = "span.category"
const quasUserSelector = "span.user-nick-wrap"
const quasDateSelector = "#content > div.sub-content-wrap > div.left-con-wrap > div.common-view-wrap.market-info-view-wrap > div > dl > dt > div.util-area > p > span"
const quasContentSelector = "div.view-content > div.note-editor"
var quasSellers = []string{
"alphascan", "LG대장", "Micronics", "AntecKorea", "3RSYS", "이엠텍홍보팀", "루선배",
"제이웍스코리아", "LG온라인샵", "제닉스홀릭", "LGEMONITOR", "삼성디스플레이", "HyperX",
"waycos", "인터픽셀", "컴디씨", "기가바이트", "스카이디지탈", "컴대문", "컴스빌", "터틀비치",
"COUGAR게이밍", "intech", "제이씨현시스템", "MAXTILL", "에이수스코리아", "잘만테크", "BRAVOTEC",
"서린씨앤아이", "한성컴퓨터"}





