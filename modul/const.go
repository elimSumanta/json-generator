package modul

var Template = `
{
    "user_id": "73709207",
    "shop_id": %s,
    "master_layout_id": 15,
    "layout_id": 0,
    "merchant_tier_id": 1,
    "header": "{\"bgColor\":[%s],\"bgImages\":[\"%s\"],\"textColor\":\"%s\",\"isDark\":%t,\"widgetHeaderStyle\":\"round\",\"widgetHeaderBgColors\":[\"%s\"]}",
    "widgets": [%s],
    "status": 1
}`

var (
	MapWidgetMasterID = map[string]int{
		"slider_banner":         1,
		"slider_square":         2,
		"display_single_column": 3,
		"display_double_column": 4,
		"display_triple_column": 5,
		"video":                 6,
	}

	WidgetOrder = []string{
		"banner_timer",
		"product_highlight",
		"voucher",
		"banner_tambahan",
		"banner_tambahan",
		"banner_tambahan",
		"banner_tambahan",
		"banner_tambahan",
		"product_highlight",
		"product_highlight",
		"play",
		"slider_banner_highlight",
	}

	BannerWidget = `
  {
    "widget_id": 0,
    "widget_master_id": 39,
    "layout_order": %d,
    "name": "banner_timer",
    "type": "display",
    "header": "{\"isActive\": 1}",
    "data": "[]"
  }`

	ProductHighlightWidget = `
  {
    "widget_id": 0,
    "widget_master_id": 41,
    "layout_order": %d,
    "name": "product_highlight",
    "type": "campaign",
    "header": "{\"title\": \"%s\", \"isActive\": 1}",
    "data": "[{\"linkID\": %d, \"linkType\": \"group_order\"}]"
  }`

	Voucher = `
  {
    "widget_id": 0,
    "widget_master_id": 42,
    "layout_order": %d,
    "name": "voucher",
    "type": "voucher_slider",
    "header": "{\"isActive\": 1}",
    "data": "[{\"linkID\": 0, \"linkType\": \"external\"}]"
  }`

	BannerTambahan = `
  {
    "widget_id": 0,
    "widget_master_id": %d,
    "layout_order": %d,
    "name": "%s",
    "type": "display",
    "header": "{\"cover\": \"\", \"ratio\": \"%s\", \"title\": \"%s\", \"isATC\": 0, \"ctaLink\": \"\", \"ctaText\": \"\", \"isActive\": %d,\"sizeOption\": \"small\", \"isShowEtalaseName\": 0, \"isBundleAutoActive\": 0}",
    "data": "[%s]"
  }`

	BannerTambahanData = `{\"linkType\":\"%s\",\"linkID\":%d,\"videoURL\":\"%s\",\"linkName\":\"\",\"webLink\":\"\",\"imageID\": \"%s\", \"imageURL\": \"%s\", \"desktopImageUrl\": \"%s\"}`

	InactiveBannerTambahan = `
 
  {
    "widget_id": 0,
    "widget_master_id": 2,
    "layout_order": %d,
    "name": "slider_square",
    "type": "display",
    "header": "{\"cover\": \"\", \"ratio\": \"1:1\", \"title\": \"\", \"isATC\": 0, \"ctaLink\": \"\", \"ctaText\": \"\", \"isActive\": 0,\"sizeOption\": \"small\", \"isShowEtalaseName\": 0, \"isBundleAutoActive\": 0}",
    "data": "[]"
  }`

	SliderBanner = `
  {
    "widget_id": 0,
    "widget_master_id": 40,
    "layout_order": %d,
    "name": "slider_banner_highlight",
    "type": "display",
    "header": "{\"ratio\":\"1:1\",\"title\":\"Masih banyak produk di toko ini, lho!\",\"isATC\": 0,\"ctaText\":\"Jelajahi Toko\",\"isActive\": 1}",
    "data": ""
  }`

	Play = `
  {
    "widget_id": 0,
    "widget_master_id": 24,
    "layout_order": %d,
    "name": "play",
    "type": "dynamic",
    "header": "{\"title\": \"join The Epic Live\", \"isActive\": 1}",
    "data": ""
  }`
)
