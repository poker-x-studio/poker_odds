/*
功能：下注类型map
说明：
*/
package BETTING_TYPE

type Item struct {
	t       TYPE
	txt_eng string
	txt_chs string
}

var (
	type_map = make(map[TYPE]Item, 0)
)

func init() {
	items := []Item{
		{ALL_IN, "all_in", "全下"},
		{FIXED_AMOUNT, "fixed_amount", "固定额度"},
		{MARTEGAL, "martegal", "马丁格尔策略:输了加倍"},
		{FIBONACCI, "fibonacci", "斐波那契策略:累加"},
		{KELLY, "kelly", "凯利策略:每次投注总筹码的某一百分比"},
	}
	for k, v := range items {
		type_map[v.t] = items[k]
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
