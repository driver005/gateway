package utils

type Country struct {
	Alpha2  string
	Name    string
	Alpha3  string
	Numeric string
}

var Countries = []Country{
	{Alpha2: "AF", Name: "Afghanistan", Alpha3: "AFG", Numeric: "004"},
	{Alpha2: "AL", Name: "Albania", Alpha3: "ALB", Numeric: "008"},
	{Alpha2: "DZ", Name: "Algeria", Alpha3: "DZA", Numeric: "012"},
	{Alpha2: "AS", Name: "American Samoa", Alpha3: "ASM", Numeric: "016"},
	{Alpha2: "AD", Name: "Andorra", Alpha3: "AND", Numeric: "020"},
	{Alpha2: "AO", Name: "Angola", Alpha3: "AGO", Numeric: "024"},
	{Alpha2: "AI", Name: "Anguilla", Alpha3: "AIA", Numeric: "660"},
	{Alpha2: "AQ", Name: "Antarctica", Alpha3: "ATA", Numeric: "010"},
	{Alpha2: "AG", Name: "Antigua and Barbuda", Alpha3: "ATG", Numeric: "028"},
	{Alpha2: "AR", Name: "Argentina", Alpha3: "ARG", Numeric: "032"},
	{Alpha2: "AM", Name: "Armenia", Alpha3: "ARM", Numeric: "051"},
	{Alpha2: "AW", Name: "Aruba", Alpha3: "ABW", Numeric: "533"},
	{Alpha2: "AU", Name: "Australia", Alpha3: "AUS", Numeric: "036"},
	{Alpha2: "AT", Name: "Austria", Alpha3: "AUT", Numeric: "040"},
	{Alpha2: "AZ", Name: "Azerbaijan", Alpha3: "AZE", Numeric: "031"},
	{Alpha2: "BS", Name: "Bahamas", Alpha3: "BHS", Numeric: "044"},
	{Alpha2: "BH", Name: "Bahrain", Alpha3: "BHR", Numeric: "048"},
	{Alpha2: "BD", Name: "Bangladesh", Alpha3: "BGD", Numeric: "050"},
	{Alpha2: "BB", Name: "Barbados", Alpha3: "BRB", Numeric: "052"},
	{Alpha2: "BY", Name: "Belarus", Alpha3: "BLR", Numeric: "112"},
	{Alpha2: "BE", Name: "Belgium", Alpha3: "BEL", Numeric: "056"},
	{Alpha2: "BZ", Name: "Belize", Alpha3: "BLZ", Numeric: "084"},
	{Alpha2: "BJ", Name: "Benin", Alpha3: "BEN", Numeric: "204"},
	{Alpha2: "BM", Name: "Bermuda", Alpha3: "BMU", Numeric: "060"},
	{Alpha2: "BT", Name: "Bhutan", Alpha3: "BTN", Numeric: "064"},
	{Alpha2: "BO", Name: "Bolivia", Alpha3: "BOL", Numeric: "068"},
	{
		Alpha2:  "BQ",
		Name:    "Bonaire, Sint Eustatius and Saba",
		Alpha3:  "BES",
		Numeric: "535",
	},
	{
		Alpha2:  "BA",
		Name:    "Bosnia and Herzegovina",
		Alpha3:  "BIH",
		Numeric: "070",
	},
	{Alpha2: "BW", Name: "Botswana", Alpha3: "BWA", Numeric: "072"},
	{Alpha2: "BV", Name: "Bouvet Island", Alpha3: "BVD", Numeric: "074"},
	{Alpha2: "BR", Name: "Brazil", Alpha3: "BRA", Numeric: "076"},
	{
		Alpha2:  "IO",
		Name:    "British Indian Ocean Territory",
		Alpha3:  "IOT",
		Numeric: "086",
	},
	{Alpha2: "BN", Name: "Brunei Darussalam", Alpha3: "BRN", Numeric: "096"},
	{Alpha2: "BG", Name: "Bulgaria", Alpha3: "BGR", Numeric: "100"},
	{Alpha2: "BF", Name: "Burkina Faso", Alpha3: "BFA", Numeric: "854"},
	{Alpha2: "BI", Name: "Burundi", Alpha3: "BDI", Numeric: "108"},
	{Alpha2: "KH", Name: "Cambodia", Alpha3: "KHM", Numeric: "116"},
	{Alpha2: "CM", Name: "Cameroon", Alpha3: "CMR", Numeric: "120"},
	{Alpha2: "CA", Name: "Canada", Alpha3: "CAN", Numeric: "124"},
	{Alpha2: "CV", Name: "Cape Verde", Alpha3: "CPV", Numeric: "132"},
	{Alpha2: "KY", Name: "Cayman Islands", Alpha3: "CYM", Numeric: "136"},
	{
		Alpha2:  "CF",
		Name:    "Central African Republic",
		Alpha3:  "CAF",
		Numeric: "140",
	},
	{Alpha2: "TD", Name: "Chad", Alpha3: "TCD", Numeric: "148"},
	{Alpha2: "CL", Name: "Chile", Alpha3: "CHL", Numeric: "152"},
	{Alpha2: "CN", Name: "China", Alpha3: "CHN", Numeric: "156"},
	{Alpha2: "CX", Name: "Christmas Island", Alpha3: "CXR", Numeric: "162"},
	{
		Alpha2:  "CC",
		Name:    "Cocos (Keeling) Islands",
		Alpha3:  "CCK",
		Numeric: "166",
	},
	{Alpha2: "CO", Name: "Colombia", Alpha3: "COL", Numeric: "170"},
	{Alpha2: "KM", Name: "Comoros", Alpha3: "COM", Numeric: "174"},
	{Alpha2: "CG", Name: "Congo", Alpha3: "COG", Numeric: "178"},
	{
		Alpha2:  "CD",
		Name:    "Congo, the Democratic Republic of the",
		Alpha3:  "COD",
		Numeric: "180",
	},
	{Alpha2: "CK", Name: "Cook Islands", Alpha3: "COK", Numeric: "184"},
	{Alpha2: "CR", Name: "Costa Rica", Alpha3: "CRI", Numeric: "188"},
	{Alpha2: "CI", Name: "Cote D'Ivoire", Alpha3: "CIV", Numeric: "384"},
	{Alpha2: "HR", Name: "Croatia", Alpha3: "HRV", Numeric: "191"},
	{Alpha2: "CU", Name: "Cuba", Alpha3: "CUB", Numeric: "192"},
	{Alpha2: "CW", Name: "Curaçao", Alpha3: "CUW", Numeric: "531"},
	{Alpha2: "CY", Name: "Cyprus", Alpha3: "CYP", Numeric: "196"},
	{Alpha2: "CZ", Name: "Czech Republic", Alpha3: "CZE", Numeric: "203"},
	{Alpha2: "DK", Name: "Denmark", Alpha3: "DNK", Numeric: "208"},
	{Alpha2: "DJ", Name: "Djibouti", Alpha3: "DJI", Numeric: "262"},
	{Alpha2: "DM", Name: "Dominica", Alpha3: "DMA", Numeric: "212"},
	{Alpha2: "DO", Name: "Dominican Republic", Alpha3: "DOM", Numeric: "214"},
	{Alpha2: "EC", Name: "Ecuador", Alpha3: "ECU", Numeric: "218"},
	{Alpha2: "EG", Name: "Egypt", Alpha3: "EGY", Numeric: "818"},
	{Alpha2: "SV", Name: "El Salvador", Alpha3: "SLV", Numeric: "222"},
	{Alpha2: "GQ", Name: "Equatorial Guinea", Alpha3: "GNQ", Numeric: "226"},
	{Alpha2: "ER", Name: "Eritrea", Alpha3: "ERI", Numeric: "232"},
	{Alpha2: "EE", Name: "Estonia", Alpha3: "EST", Numeric: "233"},
	{Alpha2: "ET", Name: "Ethiopia", Alpha3: "ETH", Numeric: "231"},
	{
		Alpha2:  "FK",
		Name:    "Falkland Islands (Malvinas)",
		Alpha3:  "FLK",
		Numeric: "238",
	},
	{Alpha2: "FO", Name: "Faroe Islands", Alpha3: "FRO", Numeric: "234"},
	{Alpha2: "FJ", Name: "Fiji", Alpha3: "FJI", Numeric: "242"},
	{Alpha2: "FI", Name: "Finland", Alpha3: "FIN", Numeric: "246"},
	{Alpha2: "FR", Name: "France", Alpha3: "FRA", Numeric: "250"},
	{Alpha2: "GF", Name: "French Guiana", Alpha3: "GUF", Numeric: "254"},
	{Alpha2: "PF", Name: "French Polynesia", Alpha3: "PYF", Numeric: "258"},
	{
		Alpha2:  "TF",
		Name:    "French Southern Territories",
		Alpha3:  "ATF",
		Numeric: "260",
	},
	{Alpha2: "GA", Name: "Gabon", Alpha3: "GAB", Numeric: "266"},
	{Alpha2: "GM", Name: "Gambia", Alpha3: "GMB", Numeric: "270"},
	{Alpha2: "GE", Name: "Georgia", Alpha3: "GEO", Numeric: "268"},
	{Alpha2: "DE", Name: "Germany", Alpha3: "DEU", Numeric: "276"},
	{Alpha2: "GH", Name: "Ghana", Alpha3: "GHA", Numeric: "288"},
	{Alpha2: "GI", Name: "Gibraltar", Alpha3: "GIB", Numeric: "292"},
	{Alpha2: "GR", Name: "Greece", Alpha3: "GRC", Numeric: "300"},
	{Alpha2: "GL", Name: "Greenland", Alpha3: "GRL", Numeric: "304"},
	{Alpha2: "GD", Name: "Grenada", Alpha3: "GRD", Numeric: "308"},
	{Alpha2: "GP", Name: "Guadeloupe", Alpha3: "GLP", Numeric: "312"},
	{Alpha2: "GU", Name: "Guam", Alpha3: "GUM", Numeric: "316"},
	{Alpha2: "GT", Name: "Guatemala", Alpha3: "GTM", Numeric: "320"},
	{Alpha2: "GG", Name: "Guernsey", Alpha3: "GGY", Numeric: "831"},
	{Alpha2: "GN", Name: "Guinea", Alpha3: "GIN", Numeric: "324"},
	{Alpha2: "GW", Name: "Guinea-Bissau", Alpha3: "GNB", Numeric: "624"},
	{Alpha2: "GY", Name: "Guyana", Alpha3: "GUY", Numeric: "328"},
	{Alpha2: "HT", Name: "Haiti", Alpha3: "HTI", Numeric: "332"},
	{
		Alpha2:  "HM",
		Name:    "Heard Island And Mcdonald Islands",
		Alpha3:  "HMD",
		Numeric: "334",
	},
	{
		Alpha2:  "VA",
		Name:    "Holy See (Vatican City State)",
		Alpha3:  "VAT",
		Numeric: "336",
	},
	{Alpha2: "HN", Name: "Honduras", Alpha3: "HND", Numeric: "340"},
	{Alpha2: "HK", Name: "Hong Kong", Alpha3: "HKG", Numeric: "344"},
	{Alpha2: "HU", Name: "Hungary", Alpha3: "HUN", Numeric: "348"},
	{Alpha2: "IS", Name: "Iceland", Alpha3: "ISL", Numeric: "352"},
	{Alpha2: "IN", Name: "India", Alpha3: "IND", Numeric: "356"},
	{Alpha2: "ID", Name: "Indonesia", Alpha3: "IDN", Numeric: "360"},
	{
		Alpha2:  "IR",
		Name:    "Iran, Islamic Republic of",
		Alpha3:  "IRN",
		Numeric: "364",
	},
	{Alpha2: "IQ", Name: "Iraq", Alpha3: "IRQ", Numeric: "368"},
	{Alpha2: "IE", Name: "Ireland", Alpha3: "IRL", Numeric: "372"},
	{Alpha2: "IM", Name: "Isle Of Man", Alpha3: "IMN", Numeric: "833"},
	{Alpha2: "IL", Name: "Israel", Alpha3: "ISR", Numeric: "376"},
	{Alpha2: "IT", Name: "Italy", Alpha3: "ITA", Numeric: "380"},
	{Alpha2: "JM", Name: "Jamaica", Alpha3: "JAM", Numeric: "388"},
	{Alpha2: "JP", Name: "Japan", Alpha3: "JPN", Numeric: "392"},
	{Alpha2: "JE", Name: "Jersey", Alpha3: "JEY", Numeric: "832"},
	{Alpha2: "JO", Name: "Jordan", Alpha3: "JOR", Numeric: "400"},
	{Alpha2: "KZ", Name: "Kazakhstan", Alpha3: "KAZ", Numeric: "398"},
	{Alpha2: "KE", Name: "Kenya", Alpha3: "KEN", Numeric: "404"},
	{Alpha2: "KI", Name: "Kiribati", Alpha3: "KIR", Numeric: "296"},
	{
		Alpha2:  "KP",
		Name:    "Korea, Democratic People's Republic of",
		Alpha3:  "PRK",
		Numeric: "408",
	},
	{Alpha2: "KR", Name: "Korea, Republic of", Alpha3: "KOR", Numeric: "410"},
	{Alpha2: "XK", Name: "Kosovo", Alpha3: "XKX", Numeric: "900"},
	{Alpha2: "KW", Name: "Kuwait", Alpha3: "KWT", Numeric: "414"},
	{Alpha2: "KG", Name: "Kyrgyzstan", Alpha3: "KGZ", Numeric: "417"},
	{
		Alpha2:  "LA",
		Name:    "Lao People's Democratic Republic",
		Alpha3:  "LAO",
		Numeric: "418",
	},
	{Alpha2: "LV", Name: "Latvia", Alpha3: "LVA", Numeric: "428"},
	{Alpha2: "LB", Name: "Lebanon", Alpha3: "LBN", Numeric: "422"},
	{Alpha2: "LS", Name: "Lesotho", Alpha3: "LSO", Numeric: "426"},
	{Alpha2: "LR", Name: "Liberia", Alpha3: "LBR", Numeric: "430"},
	{
		Alpha2:  "LY",
		Name:    "Libyan Arab Jamahiriya",
		Alpha3:  "LBY",
		Numeric: "434",
	},
	{Alpha2: "LI", Name: "Liechtenstein", Alpha3: "LIE", Numeric: "438"},
	{Alpha2: "LT", Name: "Lithuania", Alpha3: "LTU", Numeric: "440"},
	{Alpha2: "LU", Name: "Luxembourg", Alpha3: "LUX", Numeric: "442"},
	{Alpha2: "MO", Name: "Macao", Alpha3: "MAC", Numeric: "446"},
	{
		Alpha2:  "MK",
		Name:    "Macedonia, the Former Yugoslav Republic of",
		Alpha3:  "MKD",
		Numeric: "807",
	},
	{Alpha2: "MG", Name: "Madagascar", Alpha3: "MDG", Numeric: "450"},
	{Alpha2: "MW", Name: "Malawi", Alpha3: "MWI", Numeric: "454"},
	{Alpha2: "MY", Name: "Malaysia", Alpha3: "MYS", Numeric: "458"},
	{Alpha2: "MV", Name: "Maldives", Alpha3: "MDV", Numeric: "462"},
	{Alpha2: "ML", Name: "Mali", Alpha3: "MLI", Numeric: "466"},
	{Alpha2: "MT", Name: "Malta", Alpha3: "MLT", Numeric: "470"},
	{Alpha2: "MH", Name: "Marshall Islands", Alpha3: "MHL", Numeric: "584"},
	{Alpha2: "MQ", Name: "Martinique", Alpha3: "MTQ", Numeric: "474"},
	{Alpha2: "MR", Name: "Mauritania", Alpha3: "MRT", Numeric: "478"},
	{Alpha2: "MU", Name: "Mauritius", Alpha3: "MUS", Numeric: "480"},
	{Alpha2: "YT", Name: "Mayotte", Alpha3: "MYT", Numeric: "175"},
	{Alpha2: "MX", Name: "Mexico", Alpha3: "MEX", Numeric: "484"},
	{
		Alpha2:  "FM",
		Name:    "Micronesia, Federated States of",
		Alpha3:  "FSM",
		Numeric: "583",
	},
	{Alpha2: "MD", Name: "Moldova, Republic of", Alpha3: "MDA", Numeric: "498"},
	{Alpha2: "MC", Name: "Monaco", Alpha3: "MCO", Numeric: "492"},
	{Alpha2: "MN", Name: "Mongolia", Alpha3: "MNG", Numeric: "496"},
	{Alpha2: "ME", Name: "Montenegro", Alpha3: "MNE", Numeric: "499"},
	{Alpha2: "MS", Name: "Montserrat", Alpha3: "MSR", Numeric: "500"},
	{Alpha2: "MA", Name: "Morocco", Alpha3: "MAR", Numeric: "504"},
	{Alpha2: "MZ", Name: "Mozambique", Alpha3: "MOZ", Numeric: "508"},
	{Alpha2: "MM", Name: "Myanmar", Alpha3: "MMR", Numeric: "104"},
	{Alpha2: "NA", Name: "Namibia", Alpha3: "NAM", Numeric: "516"},
	{Alpha2: "NR", Name: "Nauru", Alpha3: "NRU", Numeric: "520"},
	{Alpha2: "NP", Name: "Nepal", Alpha3: "NPL", Numeric: "524"},
	{Alpha2: "NL", Name: "Netherlands", Alpha3: "NLD", Numeric: "528"},
	{Alpha2: "NC", Name: "New Caledonia", Alpha3: "NCL", Numeric: "540"},
	{Alpha2: "NZ", Name: "New Zealand", Alpha3: "NZL", Numeric: "554"},
	{Alpha2: "NI", Name: "Nicaragua", Alpha3: "NIC", Numeric: "558"},
	{Alpha2: "NE", Name: "Niger", Alpha3: "NER", Numeric: "562"},
	{Alpha2: "NG", Name: "Nigeria", Alpha3: "NGA", Numeric: "566"},
	{Alpha2: "NU", Name: "Niue", Alpha3: "NIU", Numeric: "570"},
	{Alpha2: "NF", Name: "Norfolk Island", Alpha3: "NFK", Numeric: "574"},
	{
		Alpha2:  "MP",
		Name:    "Northern Mariana Islands",
		Alpha3:  "MNP",
		Numeric: "580",
	},
	{Alpha2: "NO", Name: "Norway", Alpha3: "NOR", Numeric: "578"},
	{Alpha2: "OM", Name: "Oman", Alpha3: "OMN", Numeric: "512"},
	{Alpha2: "PK", Name: "Pakistan", Alpha3: "PAK", Numeric: "586"},
	{Alpha2: "PW", Name: "Palau", Alpha3: "PLW", Numeric: "585"},
	{
		Alpha2:  "PS",
		Name:    "Palestinian Territory, Occupied",
		Alpha3:  "PSE",
		Numeric: "275",
	},
	{Alpha2: "PA", Name: "Panama", Alpha3: "PAN", Numeric: "591"},
	{Alpha2: "PG", Name: "Papua New Guinea", Alpha3: "PNG", Numeric: "598"},
	{Alpha2: "PY", Name: "Paraguay", Alpha3: "PRY", Numeric: "600"},
	{Alpha2: "PE", Name: "Peru", Alpha3: "PER", Numeric: "604"},
	{Alpha2: "PH", Name: "Philippines", Alpha3: "PHL", Numeric: "608"},
	{Alpha2: "PN", Name: "Pitcairn", Alpha3: "PCN", Numeric: "612"},
	{Alpha2: "PL", Name: "Poland", Alpha3: "POL", Numeric: "616"},
	{Alpha2: "PT", Name: "Portugal", Alpha3: "PRT", Numeric: "620"},
	{Alpha2: "PR", Name: "Puerto Rico", Alpha3: "PRI", Numeric: "630"},
	{Alpha2: "QA", Name: "Qatar", Alpha3: "QAT", Numeric: "634"},
	{Alpha2: "RE", Name: "Reunion", Alpha3: "REU", Numeric: "638"},
	{Alpha2: "RO", Name: "Romania", Alpha3: "ROM", Numeric: "642"},
	{Alpha2: "RU", Name: "Russian Federation", Alpha3: "RUS", Numeric: "643"},
	{Alpha2: "RW", Name: "Rwanda", Alpha3: "RWA", Numeric: "646"},
	{Alpha2: "BL", Name: "Saint Barthélemy", Alpha3: "BLM", Numeric: "652"},
	{Alpha2: "SH", Name: "Saint Helena", Alpha3: "SHN", Numeric: "654"},
	{
		Alpha2:  "KN",
		Name:    "Saint Kitts and Nevis",
		Alpha3:  "KNA",
		Numeric: "659",
	},
	{Alpha2: "LC", Name: "Saint Lucia", Alpha3: "LCA", Numeric: "662"},
	{
		Alpha2:  "MF",
		Name:    "Saint Martin (French part)",
		Alpha3:  "MAF",
		Numeric: "663",
	},
	{
		Alpha2:  "PM",
		Name:    "Saint Pierre and Miquelon",
		Alpha3:  "SPM",
		Numeric: "666",
	},
	{
		Alpha2:  "VC",
		Name:    "Saint Vincent and the Grenadines",
		Alpha3:  "VCT",
		Numeric: "670",
	},
	{Alpha2: "WS", Name: "Samoa", Alpha3: "WSM", Numeric: "882"},
	{Alpha2: "SM", Name: "San Marino", Alpha3: "SMR", Numeric: "674"},
	{
		Alpha2:  "ST",
		Name:    "Sao Tome and Principe",
		Alpha3:  "STP",
		Numeric: "678",
	},
	{Alpha2: "SA", Name: "Saudi Arabia", Alpha3: "SAU", Numeric: "682"},
	{Alpha2: "SN", Name: "Senegal", Alpha3: "SEN", Numeric: "686"},
	{Alpha2: "RS", Name: "Serbia", Alpha3: "SRB", Numeric: "688"},
	{Alpha2: "SC", Name: "Seychelles", Alpha3: "SYC", Numeric: "690"},
	{Alpha2: "SL", Name: "Sierra Leone", Alpha3: "SLE", Numeric: "694"},
	{Alpha2: "SG", Name: "Singapore", Alpha3: "SGP", Numeric: "702"},
	{Alpha2: "SX", Name: "Sint Maarten", Alpha3: "SXM", Numeric: "534"},
	{Alpha2: "SK", Name: "Slovakia", Alpha3: "SVK", Numeric: "703"},
	{Alpha2: "SI", Name: "Slovenia", Alpha3: "SVN", Numeric: "705"},
	{Alpha2: "SB", Name: "Solomon Islands", Alpha3: "SLB", Numeric: "090"},
	{Alpha2: "SO", Name: "Somalia", Alpha3: "SOM", Numeric: "706"},
	{Alpha2: "ZA", Name: "South Africa", Alpha3: "ZAF", Numeric: "710"},
	{
		Alpha2:  "GS",
		Name:    "South Georgia and the South Sandwich Islands",
		Alpha3:  "SGS",
		Numeric: "239",
	},
	{Alpha2: "SS", Name: "South Sudan", Alpha3: "SSD", Numeric: "728"},
	{Alpha2: "ES", Name: "Spain", Alpha3: "ESP", Numeric: "724"},
	{Alpha2: "LK", Name: "Sri Lanka", Alpha3: "LKA", Numeric: "144"},
	{Alpha2: "SD", Name: "Sudan", Alpha3: "SDN", Numeric: "729"},
	{Alpha2: "SR", Name: "Suriname", Alpha3: "SUR", Numeric: "740"},
	{
		Alpha2:  "SJ",
		Name:    "Svalbard and Jan Mayen",
		Alpha3:  "SJM",
		Numeric: "744",
	},
	{Alpha2: "SZ", Name: "Swaziland", Alpha3: "SWZ", Numeric: "748"},
	{Alpha2: "SE", Name: "Sweden", Alpha3: "SWE", Numeric: "752"},
	{Alpha2: "CH", Name: "Switzerland", Alpha3: "CHE", Numeric: "756"},
	{Alpha2: "SY", Name: "Syrian Arab Republic", Alpha3: "SYR", Numeric: "760"},
	{
		Alpha2:  "TW",
		Name:    "Taiwan, Province of China",
		Alpha3:  "TWN",
		Numeric: "158",
	},
	{Alpha2: "TJ", Name: "Tajikistan", Alpha3: "TJK", Numeric: "762"},
	{
		Alpha2:  "TZ",
		Name:    "Tanzania, United Republic of",
		Alpha3:  "TZA",
		Numeric: "834",
	},
	{Alpha2: "TH", Name: "Thailand", Alpha3: "THA", Numeric: "764"},
	{Alpha2: "TL", Name: "Timor Leste", Alpha3: "TLS", Numeric: "626"},
	{Alpha2: "TG", Name: "Togo", Alpha3: "TGO", Numeric: "768"},
	{Alpha2: "TK", Name: "Tokelau", Alpha3: "TKL", Numeric: "772"},
	{Alpha2: "TO", Name: "Tonga", Alpha3: "TON", Numeric: "776"},
	{Alpha2: "TT", Name: "Trinidad and Tobago", Alpha3: "TTO", Numeric: "780"},
	{Alpha2: "TN", Name: "Tunisia", Alpha3: "TUN", Numeric: "788"},
	{Alpha2: "TR", Name: "Turkey", Alpha3: "TUR", Numeric: "792"},
	{Alpha2: "TM", Name: "Turkmenistan", Alpha3: "TKM", Numeric: "795"},
	{Alpha2: "TC", Name: "Turks and Caicos Islands", Alpha3: "TCA", Numeric: "796"},
	{Alpha2: "TV", Name: "Tuvalu", Alpha3: "TUV", Numeric: "798"},
	{Alpha2: "UG", Name: "Uganda", Alpha3: "UGA", Numeric: "800"},
	{Alpha2: "UA", Name: "Ukraine", Alpha3: "UKR", Numeric: "804"},
	{Alpha2: "AE", Name: "United Arab Emirates", Alpha3: "ARE", Numeric: "784"},
	{Alpha2: "GB", Name: "United Kingdom", Alpha3: "GBR", Numeric: "826"},
	{Alpha2: "US", Name: "United States", Alpha3: "USA", Numeric: "840"},
	{Alpha2: "UM", Name: "United States Minor Outlying Islands", Alpha3: "UMI", Numeric: "581"},
	{Alpha2: "UY", Name: "Uruguay", Alpha3: "URY", Numeric: "858"},
	{Alpha2: "UZ", Name: "Uzbekistan", Alpha3: "UZB", Numeric: "860"},
	{Alpha2: "VU", Name: "Vanuatu", Alpha3: "VUT", Numeric: "548"},
	{Alpha2: "VE", Name: "Venezuela", Alpha3: "VEN", Numeric: "862"},
	{Alpha2: "VN", Name: "Viet Nam", Alpha3: "VNM", Numeric: "704"},
	{Alpha2: "VG", Name: "Virgin Islands, British", Alpha3: "VGB", Numeric: "092"},
	{Alpha2: "VI", Name: "Virgin Islands, U.S.", Alpha3: "VIR", Numeric: "850"},
	{Alpha2: "WF", Name: "Wallis and Futuna", Alpha3: "WLF", Numeric: "876"},
	{Alpha2: "EH", Name: "Western Sahara", Alpha3: "ESH", Numeric: "732"},
	{Alpha2: "YE", Name: "Yemen", Alpha3: "YEM", Numeric: "887"},
	{Alpha2: "ZM", Name: "Zambia", Alpha3: "ZMB", Numeric: "894"},
	{Alpha2: "ZW", Name: "Zimbabwe", Alpha3: "ZWE", Numeric: "716"},
	{Alpha2: "AX", Name: "Åland Islands", Alpha3: "ALA", Numeric: "248"},
}
