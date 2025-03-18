package products

type Product struct {
	Name string
	Url  string
	//Selector       string
	PriceThreshold float64
}

var Products = []Product{
	{
		"Пена NIVEA MEN",
		"https://www.ozon.ru/product/pena-dlya-britya-uspokaivayushchaya-nivea-men-dlya-chuvstvitelnoy-kozhi-bez-spirta-200-ml-142922708/",
		//".lz4_29.l4z_29.zl8_29",
		230,
	},
	{
		"Сменные Кассеты Gillette Fusion5",
		"https://www.ozon.ru/product/smennye-kassety-gillette-fusion5-dlya-muzhskoy-britvy-12-sht-s-5-lezviyami-c-tochnym-trimmerom-dlya-305543242/",
		//".y9l_29.ly8_29",
		3000,
	},
	{
		"Успокаивающий бальзам после бритья NIVEA MEN",
		"https://www.ozon.ru/product/uspokaivayushchiy-balzam-posle-britya-nivea-men-dlya-chuvstvitelnoy-kozhi-bez-spirta-100-ml-1603872693/",
		//".lz4_29.l4z_29.zl8_29",
		350,
	},
	{
		"Gillette Fusion5 ProGlide Power мужская бритва",
		"https://www.ozon.ru/product/gillette-fusion5-proglide-power-muzhskaya-britva-1-kasseta-s-5-lezviyami-s-tehnologiey-flexball-s-1609771128/",
		//".lz4_29.l4z_29.zl8_29",
		1000,
	},
	{
		"Gillette Mach3 Turbo мужская бритва",
		"https://www.ozon.ru/product/nabor-gillette-mach3-turbo-muzhskaya-britva-4-kassety-s-3-lezviyami-prochnee-chem-stal-dlya-tochnogo-239779171/?from_sku=239779171&oos_search=false",
		//".lz4_29.l4z_29.zl8_29",
		1300,
	},
	{
		"Gillette Fusion5 ProGlide Мужская Бритва, 5 кассет",
		"https://www.ozon.ru/product/gillette-fusion5-proglide-muzhskaya-britva-5-kasset-s-5-lezviyami-s-uglerodnym-pokrytiem-s-1657595919/?from_sku=239779181&oos_search=false",
		//".lz4_29.l4z_29.zl8_29",
		1500,
	},
	{
		"OLLIN PROFESSIONAL Шампунь против перхоти",
		"https://www.ozon.ru/product/ollin-professional-shampun-protiv-perhoti-zhenskiy-care-anti-dandruff-1000-ml-160411881/",
		//".lz4_29.l4z_29.zl8_29",
		800,
	},
	{
		"ЕЛИЗАР, кислородный пятновыводитель",
		"https://www.ozon.ru/product/elizar-kislorodnyy-pyatnovyvoditel-otbelivatel-ochistitel-kontsentrat-1-kg-dlya-tsvetnogo-i-belogo-212587655/",
		//".lz4_29.l4z_29.zl8_29",
		270,
	},
	{
		"Himalaya Крем",
		"https://www.ozon.ru/product/himalaya-krem-dlya-litsa-i-tela-pitatelnyy-uvlazhnyayushchiy-smyagchayushchiy-i-uspokaivayushchiy-ot-154953298/",
		//".lz4_29.l4z_29.zl8_29",
		300,
	},
	{
		"OLD SPICE",
		"https://www.ozon.ru/product/old-spice-muzhskoy-gel-dlya-dusha-shampun-3v1-whitewater-1-l-1086926377/?from_sku=1086926377&oos_search=false",
		//".lz4_29.l4z_29.zl8_29",
		600,
	},
}
