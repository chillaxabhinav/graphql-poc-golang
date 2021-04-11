package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
)

func (r *queryResolver) FlightSrp(ctx context.Context) (*string, error) {
	const json2 = `{
		"providers": [
		  "1044"
		],
		"results": {
		  "1044": {
			"outR": {
			  "flights": {
				"DEL-AMD-6E161": "161|6E|1210|1340|0|130||0|||127|rl1h1q4hzqhdzdhtptdpkpwhqb06",
				"VNS-BOM-6E5377": "5377|6E|1330|1545|0|215||0|||127|30lhi6xhzqhvswwhtptdpkpwhqb06",
				"DEL-BHO-AI437": "437|AI|2000|2120|0|120||0|||127|3c0h1q4hr9hkswhtptdpkpwhqb06",
				"DEL-RAJ-SG2643": "2643|SG|0615|0855|0|240||0|||127|oryh1q4hxfhtzkshtptdpkpwhqb06",
				"DEL-HYD-UK829": "829|UK|0705|0925|0|220||0|||127|c71h1q4hjehutnhtptdpkpwhqb06",
				"DEL-BOM-6E171": "171|6E|0230|0435|0|205||0|||127|30lh1q4hzqhdwdhtptdpkpwhqb06",
				"DEL-IXC-UK637": "637|UK|1710|1815|0|105||0|||127|9~bh1q4hjehzswhtptdpkpwhqb06",
				"DEL-STV-SG8483": "8483|SG|1800|1950|0|150||0|||127|x8ih1q4hxfhukushtptdpkpwhqb06",
				"HYD-BOM-UK874": "874|UK|0820|0955|0|135||0|||127|30lhc71hjehuwkhtptdpkpwhqb06",
				"BLR-BOM-UK846": "846|UK|0830|1015|0|145||0|||127|30lh34ohjehukzhtptdpkpwhqb06",
				"DEL-GOI-G8286": "286|G8|1040|1320|0|240||0|||127|f09h1q4hfuhtuzhtptdpkpwhqb06",
				"HYD-BOM-UK876": "876|UK|2135|2300|0|125||0|||127|30lhc71hjehuwzhtptdpkpwhqb06",
				"DEL-HYD-6E6205": "6205|6E|1355|1610|0|215||0|||127|c71h1q4hzqhztpvhtptdpkpwhqb06",
				"DEL-JDH-AI475": "475|AI|1255|1415|0|120||0|||127|y1ch1q4hr9hkwvhtptdpkpwhqb06",
				"DEL-BOM-AI887": "887|AI|0700|0905|0|205||0|||127|30lh1q4hr9huuwhtptdpkpwhqb06",
				"DEL-BOM-AI665": "665|AI|0800|1010|0|210||0|||127|30lh1q4hr9hzzvhtptdpkpwhqb06",
				"DEL-BOM-6E6197": "6197|6E|0845|1100|0|215||0|||127|30lh1q4hzqhzdnwhtptdpkpwhqb06",
				"LKO-BOM-G8307": "307|G8|1335|1555|0|220||0|||127|30lh4e0hfuhspwhtptdpkpwhqb06",
				"DEL-AMD-6E6373": "6373|6E|0715|0855|0|140||0|||127|rl1h1q4hzqhzswshtptdpkpwhqb06",
				"DEL-GOI-AI883": "883|AI|2200|0035|1|235||0|||127|f09h1q4hr9huushtptdpkpwhqb06",
				"DEL-AMD-6E153": "153|6E|1815|1945|0|130||0|||127|rl1h1q4hzqhdvshtptdpkpwhqb06",
				"DED-BOM-SG290": "290|SG|1250|1505|0|215||0|||127|30lh1q1hxfhtnphtptdpkpwhqb06",
				"CCU-BOM-I5316": "316|I5|2210|0115|1|305||0|||127|30lhbbjh9vhsdzhtptdpkpwhqb06",
				"DEL-BOM-6E181": "181|6E|0930|1145|0|215||0|||127|30lh1q4hzqhdudhtptdpkpwhqb06",
				"DEL-IDR-6E5051": "5051|6E|0700|0830|0|130||0|||127|91oh1q4hzqhvpvdhtptdpkpwhqb06",
				"GOI-BOM-6E5306": "5306|6E|1630|1750|0|120||0|||127|30lhf09hzqhvspzhtptdpkpwhqb06",
				"RPR-VTZ-AI651": "651|AI|1230|1335|0|105||0|||127|i8mhogohr9hzvdhtptdpkpwhqb06",
				"DEL-STV-SG8473": "8473|SG|0540|0740|0|200||0|||127|x8ih1q4hxfhukwshtptdpkpwhqb06",
				"NAG-BOM-6E297": "297|6E|1050|1220|0|130||0|||127|30lh6rfhzqhtnwhtptdpkpwhqb06",
				"GOI-BOM-I5472": "472|I5|0050|0205|0|115||0|||127|30lhf09h9vhkwthtptdpkpwhqb06",
				"DEL-VNS-G8404": "404|G8|1020|1140|0|120||0|||127|i6xh1q4hfuhkpkhtptdpkpwhqb06",
				"DEL-MAA-I5761": "761|I5|0510|0805|0|255||0|||127|lrrh1q4h9vhwzdhtptdpkpwhqb06",
				"DEL-LKO-AI411": "411|AI|0705|0815|0|110||0|||127|4e0h1q4hr9hkddhtptdpkpwhqb06",
				"DEL-BOM-6E6186": "6186|6E|1745|1955|0|210||0|||127|30lh1q4hzqhzduzhtptdpkpwhqb06",
				"HYD-BOM-AI620": "620|AI|0530|0710|0|140||0|||127|30lhc71hr9hztphtptdpkpwhqb06",
				"DEL-MAA-AI439": "439|AI|0600|0855|0|255||0|||127|lrrh1q4hr9hksnhtptdpkpwhqb06",
				"DEL-AMD-6E5072": "5072|6E|1345|1520|0|135||0|||127|rl1h1q4hzqhvpwthtptdpkpwhqb06",
				"BLR-BOM-UK864": "864|UK|1130|1315|0|145||0|||127|30lh34ohjehuzkhtptdpkpwhqb06",
				"DEL-HYD-6E2187": "2187|6E|1135|1350|0|215||0|||127|c71h1q4hzqhtduwhtptdpkpwhqb06",
				"DEL-GOI-6E367": "367|6E|1210|1450|0|240||0|||127|f09h1q4hzqhszwhtptdpkpwhqb06",
				"AMD-BOM-G82505": "2505|G8|1510|1625|0|115||0|||127|30lhrl1hfuhtvpvhtptdpkpwhqb06",
				"BLR-BOM-UK866": "866|UK|2140|2325|0|145||0|||127|30lh34ohjehuzzhtptdpkpwhqb06",
				"DEL-BOM-AI865": "865|AI|1040|1250|0|210||0|||127|30lh1q4hr9huzvhtptdpkpwhqb06",
				"DEL-BOM-AI624": "624|AI|1900|2115|0|215||0|||127|30lh1q4hr9hztkhtptdpkpwhqb06",
				"BLR-BOM-UK850": "850|UK|1955|2150|0|155||0|||127|30lh34ohjehuvphtptdpkpwhqb06",
				"MAA-BOM-UK828": "828|UK|0700|0900|0|200||0|||127|30lhlrrhjehutuhtptdpkpwhqb06",
				"MAA-BOM-UK826": "826|UK|1230|1430|0|200||0|||127|30lhlrrhjehutzhtptdpkpwhqb06",
				"BLR-BOM-UK854": "854|UK|1825|2015|0|150||0|||127|30lh34ohjehuvkhtptdpkpwhqb06",
				"MAA-BOM-I5803": "803|I5|1225|1425|0|200||0|||127|30lhlrrh9vhupshtptdpkpwhqb06",
				"DEL-MAA-AI429": "429|AI|0945|1240|0|255||0|||127|lrrh1q4hr9hktnhtptdpkpwhqb06",
				"BHO-BOM-AI634": "634|AI|0805|0940|0|135||0|||127|30lh3c0hr9hzskhtptdpkpwhqb06",
				"LKO-BOM-6E891": "891|6E|1850|2105|0|215||0|||127|30lh4e0hzqhundhtptdpkpwhqb06",
				"BLR-BOM-UK852": "852|UK|0930|1125|0|155||0|||127|30lh34ohjehuvthtptdpkpwhqb06",
				"BHO-BOM-AI632": "632|AI|2150|2325|0|135||0|||127|30lh3c0hr9hzsthtptdpkpwhqb06",
				"HYD-BOM-AI616": "616|AI|0825|0940|0|115||0|||127|30lhc71hr9hzdzhtptdpkpwhqb06",
				"BLR-BOM-AI610": "610|AI|1910|2105|0|155||0|||127|30lh34ohr9hzdphtptdpkpwhqb06",
				"MAA-BOM-UK824": "824|UK|2030|2235|0|205||0|||127|30lhlrrhjehutkhtptdpkpwhqb06",
				"MAA-BOM-UK822": "822|UK|0945|1140|0|155||0|||127|30lhlrrhjehutthtptdpkpwhqb06",
				"BLR-BOM-AI604": "604|AI|0835|1035|0|200||0|||127|30lh34ohr9hzpkhtptdpkpwhqb06",
				"DEL-VNS-UK671": "671|UK|1615|1740|0|125||0|||127|i6xh1q4hjehzwdhtptdpkpwhqb06",
				"DEL-MAA-AI540": "540|AI|2040|2335|0|255||0|||127|lrrh1q4hr9hvkphtptdpkpwhqb06",
				"DEL-GOI-I5775": "775|I5|1155|1435|0|240||0|||127|f09h1q4h9vhwwvhtptdpkpwhqb06",
				"DEL-BOM-6E843": "843|6E|1830|2035|0|205||0|||127|30lh1q4hzqhukshtptdpkpwhqb06",
				"BLR-BOM-UK858": "858|UK|0640|0815|0|135||0|||127|30lh34ohjehuvuhtptdpkpwhqb06",
				"DEL-GOI-I5773": "773|I5|1945|2225|0|240||0|||127|f09h1q4h9vhwwshtptdpkpwhqb06",
				"DEL-IXC-G8108": "108|G8|0900|1000|0|100||0|||127|9~bh1q4hfuhdpuhtptdpkpwhqb06",
				"DEL-AMD-SG8913": "8913|SG|0545|0725|0|140||0|||127|rl1h1q4hxfhundshtptdpkpwhqb06",
				"DEL-BOM-I5716": "716|I5|1155|1415|0|220||0|||127|30lh1q4h9vhwdzhtptdpkpwhqb06",
				"DEL-HYD-AI560": "560|AI|0820|1030|0|210||0|||127|c71h1q4hr9hvzphtptdpkpwhqb06",
				"DEL-BOM-6E5115": "5115|6E|1640|1850|0|210||0|||127|30lh1q4hzqhvddvhtptdpkpwhqb06",
				"RAJ-BOM-AI656": "656|AI|1805|1915|0|110||0|||127|30lhoryhr9hzvzhtptdpkpwhqb06",
				"HYD-BOM-6E5312": "5312|6E|1800|1940|0|140||0|||127|30lhc71hzqhvsdthtptdpkpwhqb06",
				"DEL-NAG-AI469": "469|AI|0600|0740|0|140||0|||127|6rfh1q4hr9hkznhtptdpkpwhqb06",
				"DEL-BOM-6E6202": "6202|6E|1200|1405|0|205||0|||127|30lh1q4hzqhztpthtptdpkpwhqb06",
				"AMD-BOM-I5359": "359|I5|0850|0955|0|105||0|||127|30lhrl1h9vhsvnhtptdpkpwhqb06",
				"DEL-CCU-UK705": "705|UK|0730|0935|0|205||0|||127|bbjh1q4hjehwpvhtptdpkpwhqb06",
				"DEL-CCU-UK707": "707|UK|1730|1935|0|205||0|||127|bbjh1q4hjehwpwhtptdpkpwhqb06",
				"DEL-BLR-AI504": "504|AI|2115|2355|0|240||0|||127|34oh1q4hr9hvpkhtptdpkpwhqb06",
				"DEL-BLR-AI502": "502|AI|1330|1615|0|245||0|||127|34oh1q4hr9hvpthtptdpkpwhqb06",
				"RAJ-BOM-SG224": "224|SG|1715|1845|0|130||0|||127|30lhoryhxfhttkhtptdpkpwhqb06",
				"DEL-BLR-AI506": "506|AI|0945|1230|0|245||0|||127|34oh1q4hr9hvpzhtptdpkpwhqb06",
				"VNS-BOM-AI696": "696|AI|1325|1540|0|215||0|||127|30lhi6xhr9hznzhtptdpkpwhqb06",
				"DEL-BHO-AI481": "481|AI|0840|1015|0|135||0|||127|3c0h1q4hr9hkudhtptdpkpwhqb06",
				"DEL-AMD-6E5006": "5006|6E|2000|2135|0|135||0|||127|rl1h1q4hzqhvppzhtptdpkpwhqb06",
				"DEL-SXR-G83011": "3011|G8|0900|1030|0|130||0|||127|x~oh1q4hfuhspddhtptdpkpwhqb06",
				"DEL-UDR-6E2274": "2274|6E|1320|1440|0|120||0|||127|j1oh1q4hzqhttwkhtptdpkpwhqb06",
				"DEL-BOM-6E5347": "5347|6E|2030|2235|0|205||0|||127|30lh1q4hzqhvskwhtptdpkpwhqb06",
				"IXC-BOM-G8382": "382|G8|1130|1355|0|225||0|||127|30lh9~bhfuhsuthtptdpkpwhqb06",
				"BBI-BOM-I5630": "630|I5|1845|2120|0|235||0|||127|30lh339h9vhzsphtptdpkpwhqb06",
				"HYD-BOM-6E5325": "5325|6E|1600|1735|0|135||0|||127|30lhc71hzqhvstvhtptdpkpwhqb06",
				"IDR-BOM-6E957": "957|6E|1615|1735|0|120||0|||127|30lh91ohzqhnvwhtptdpkpwhqb06",
				"DEL-DED-SG2033": "2033|SG|1830|1920|0|050||0|||127|1q1h1q4hxfhtpsshtptdpkpwhqb06",
				"DEL-SXR-G83009": "3009|G8|0500|0620|0|120||0|||127|x~oh1q4hfuhsppnhtptdpkpwhqb06",
				"DEL-VNS-UK812": "812|UK|1320|1445|0|125||0|||127|i6xh1q4hjehudthtptdpkpwhqb06",
				"AMD-BOM-6E6811": "6811|6E|1815|1935|0|120||0|||127|30lhrl1hzqhzuddhtptdpkpwhqb06",
				"AMD-BOM-6E6813": "6813|6E|2240|2355|0|115||0|||127|30lhrl1hzqhzudshtptdpkpwhqb06",
				"DEL-JAI-AI491": "491|AI|1200|1255|0|055||0|||127|yr9h1q4hr9hkndhtptdpkpwhqb06",
				"DEL-GOI-6E2167": "2167|6E|1855|2130|0|235||0|||127|f09h1q4hzqhtdzwhtptdpkpwhqb06",
				"DEL-HYD-6E816": "816|6E|1000|1220|0|220||0|||127|c71h1q4hzqhudzhtptdpkpwhqb06",
				"DEL-UDR-UK627": "627|UK|1320|1445|0|125||0|||127|j1oh1q4hjehztwhtptdpkpwhqb06",
				"JDH-BOM-AI646": "646|AI|1155|1335|0|140||0|||127|30lhy1chr9hzkzhtptdpkpwhqb06",
				"BLR-BOM-I5306": "306|I5|2240|0030|1|150||0|||127|30lh34oh9vhspzhtptdpkpwhqb06",
				"GOI-BOM-UK844": "844|UK|1440|1600|0|120||0|||127|30lhf09hjehukkhtptdpkpwhqb06",
				"DEL-BOM-6E5328": "5328|6E|1030|1240|0|210||0|||127|30lh1q4hzqhvstuhtptdpkpwhqb06",
				"DEL-RAJ-AI495": "495|AI|1655|1845|0|150||0|||127|oryh1q4hr9hknvhtptdpkpwhqb06",
				"DEL-DED-6E966": "966|6E|1425|1510|0|045||0|||127|1q1h1q4hzqhnzzhtptdpkpwhqb06",
				"DEL-AMD-G8717": "717|G8|1645|1820|0|135||0|||127|rl1h1q4hfuhwdwhtptdpkpwhqb06",
				"IXC-BOM-UK652": "652|UK|1905|2135|0|230||0|||127|30lh9~bhjehzvthtptdpkpwhqb06",
				"JAI-BOM-AI612": "612|AI|1400|1550|0|150||0|||127|30lhyr9hr9hzdthtptdpkpwhqb06",
				"DEL-AMD-G8719": "719|G8|0730|0900|0|130||0|||127|rl1h1q4hfuhwdnhtptdpkpwhqb06",
				"DEL-LKO-G8268": "268|G8|1200|1310|0|110||0|||127|4e0h1q4hfuhtzuhtptdpkpwhqb06",
				"DEL-AMD-G8713": "713|G8|2005|2145|0|140||0|||127|rl1h1q4hfuhwdshtptdpkpwhqb06",
				"GOI-BOM-6E418": "418|6E|2320|0045|1|125||0|||127|30lhf09hzqhkduhtptdpkpwhqb06",
				"DEL-HYD-6E5406": "5406|6E|1900|2120|0|220||0|||127|c71h1q4hzqhvkpzhtptdpkpwhqb06",
				"BDQ-BOM-6E628": "628|6E|1700|1815|0|115||0|||127|30lh312hzqhztuhtptdpkpwhqb06",
				"AMD-BOM-6E214": "214|6E|0935|1100|0|125||0|||127|30lhrl1hzqhtdkhtptdpkpwhqb06",
				"DEL-BDQ-6E6245": "6245|6E|1355|1535|0|140||0|||127|312h1q4hzqhztkvhtptdpkpwhqb06",
				"DEL-BOM-G8392": "392|G8|1545|1800|0|215||0|||127|30lh1q4hfuhsnthtptdpkpwhqb06",
				"HYD-BOM-6E903": "903|6E|1410|1550|0|140||0|||127|30lhc71hzqhnpshtptdpkpwhqb06",
				"RAJ-BOM-SG436": "436|SG|0830|0940|0|110||0|||127|30lhoryhxfhkszhtptdpkpwhqb06",
				"DEL-LKO-G82509": "2509|G8|0650|0755|0|105||0|||127|4e0h1q4hfuhtvpnhtptdpkpwhqb06",
				"HYD-BOM-6E235": "235|6E|2315|0055|1|140||0|||127|30lhc71hzqhtsvhtptdpkpwhqb06",
				"DEL-LKO-AI431": "431|AI|1220|1325|0|105||0|||127|4e0h1q4hr9hksdhtptdpkpwhqb06",
				"DEL-BOM-AI678": "678|AI|0900|1115|0|215||0|||127|30lh1q4hr9hzwuhtptdpkpwhqb06",
				"DEL-BOM-6E5031": "5031|6E|0715|0925|0|210||0|||127|30lh1q4hzqhvpsdhtptdpkpwhqb06",
				"MAA-BOM-AI672": "672|AI|1455|1705|0|210||0|||127|30lhlrrhr9hzwthtptdpkpwhqb06",
				"DEL-JLR-SG3029": "3029|SG|0545|0745|0|200||0|||127|y4oh1q4hxfhsptnhtptdpkpwhqb06",
				"LKO-BOM-6E362": "362|6E|0750|1005|0|215||0|||127|30lh4e0hzqhszthtptdpkpwhqb06",
				"DEL-LKO-6E6612": "6612|6E|0505|0615|0|110||0|||127|4e0h1q4hzqhzzdthtptdpkpwhqb06",
				"DEL-MAA-UK837": "837|UK|1725|2015|0|250||0|||127|lrrh1q4hjehuswhtptdpkpwhqb06",
				"AMD-BOM-G8413": "413|G8|0035|0200|0|125||0|||127|30lhrl1hfuhkdshtptdpkpwhqb06",
				"DEL-CCU-UK747": "747|UK|0615|0820|0|205||0|||127|bbjh1q4hjehwkwhtptdpkpwhqb06",
				"DEL-BOM-UK953": "953|UK|2040|2250|0|210||0|||127|30lh1q4hjehnvshtptdpkpwhqb06",
				"DEL-MAA-UK835": "835|UK|1955|2245|0|250||0|||127|lrrh1q4hjehusvhtptdpkpwhqb06",
				"DEL-BOM-UK955": "955|UK|1745|2005|0|220||0|||127|30lh1q4hjehnvvhtptdpkpwhqb06",
				"DEL-MAA-UK833": "833|UK|0720|1010|0|250||0|||127|lrrh1q4hjehusshtptdpkpwhqb06",
				"DEL-GOI-G83019": "3019|G8|1630|1920|0|250||0|||127|f09h1q4hfuhspdnhtptdpkpwhqb06",
				"AMD-BOM-G8537": "537|G8|1955|2110|0|115||0|||127|30lhrl1hfuhvswhtptdpkpwhqb06",
				"STV-BOM-SG2773": "2773|SG|0940|1050|0|110||0|||127|30lhx8ihxfhtwwshtptdpkpwhqb06",
				"DEL-BOM-I5482": "482|I5|2100|2310|0|210||0|||127|30lh1q4h9vhkuthtptdpkpwhqb06",
				"DEL-BBI-I5814": "814|I5|1355|1615|0|220||0|||127|339h1q4h9vhudkhtptdpkpwhqb06",
				"HYD-BOM-6E884": "884|6E|2200|2340|0|140||0|||127|30lhc71hzqhuukhtptdpkpwhqb06",
				"DEL-BOM-AI805": "805|AI|2000|2210|0|210||0|||127|30lh1q4hr9hupvhtptdpkpwhqb06",
				"DEL-JAI-AI9843": "9843|AI|0755|0855|0|100||0|||127|yr9h1q4hr9hnukshtptdpkpwhqb06",
				"DEL-BOM-G8346": "346|G8|1945|2200|0|215||0|||127|30lh1q4hfuhskzhtptdpkpwhqb06",
				"VNS-BOM-G8350": "350|G8|1315|1530|0|215||0|||127|30lhi6xhfuhsvphtptdpkpwhqb06",
				"DEL-LKO-AI811": "811|AI|1830|1940|0|110||0|||127|4e0h1q4hr9huddhtptdpkpwhqb06",
				"CCU-BOM-UK776": "776|UK|1735|2020|0|245||0|||127|30lhbbjhjehwwzhtptdpkpwhqb06",
				"DEL-BOM-UK941": "941|UK|1655|1905|0|210||0|||127|30lh1q4hjehnkdhtptdpkpwhqb06",
				"RAJ-BOM-AI602": "602|AI|0640|0750|0|110||0|||127|30lhoryhr9hzpthtptdpkpwhqb06",
				"DEL-BOM-6E5025": "5025|6E|1405|1610|0|205||0|||127|30lh1q4hzqhvptvhtptdpkpwhqb06",
				"CCU-BOM-UK772": "772|UK|1025|1300|0|235||0|||127|30lhbbjhjehwwthtptdpkpwhqb06",
				"DEL-BOM-UK943": "943|UK|0730|0945|0|215||0|||127|30lh1q4hjehnkshtptdpkpwhqb06",
				"DEL-BOM-UK945": "945|UK|1140|1400|0|220||0|||127|30lh1q4hjehnkvhtptdpkpwhqb06",
				"IXU-BOM-AI441": "441|AI|2020|2135|0|115||0|||127|30lh9~jhr9hkkdhtptdpkpwhqb06",
				"DEL-BOM-G8338": "338|G8|1055|1310|0|215||0|||127|30lh1q4hfuhssuhtptdpkpwhqb06",
				"GOI-BOM-G82606": "2606|G8|1725|1850|0|125||0|||127|30lhf09hfuhtzpzhtptdpkpwhqb06",
				"DEL-BOM-G8330": "330|G8|2050|2320|0|230||0|||127|30lh1q4hfuhssphtptdpkpwhqb06",
				"DEL-LKO-6E546": "546|6E|1630|1735|0|105||0|||127|4e0h1q4hzqhvkzhtptdpkpwhqb06",
				"DEL-BOM-G8336": "336|G8|1430|1640|0|210||0|||127|30lh1q4hfuhsszhtptdpkpwhqb06",
				"DEL-HYD-6E967": "967|6E|1705|1925|0|220||0|||127|c71h1q4hzqhnzwhtptdpkpwhqb06",
				"DEL-BOM-G8334": "334|G8|0800|1010|0|210||0|||127|30lh1q4hfuhsskhtptdpkpwhqb06",
				"DEL-HYD-AI542": "542|AI|0945|1155|0|210||0|||127|c71h1q4hr9hvkthtptdpkpwhqb06",
				"IDR-BOM-6E232": "232|6E|0940|1105|0|125||0|||127|30lh91ohzqhtsthtptdpkpwhqb06",
				"PAT-BOM-G8352": "352|G8|1420|1655|0|235||0|||127|30lhgr8hfuhsvthtptdpkpwhqb06",
				"DEL-BOM-6E5379": "5379|6E|0530|0740|0|210||0|||127|30lh1q4hzqhvswnhtptdpkpwhqb06",
				"UDR-BOM-AI644": "644|AI|1620|1750|0|130||0|||127|30lhj1ohr9hzkkhtptdpkpwhqb06",
				"DEL-BOM-6E5012": "5012|6E|0800|1010|0|210||0|||127|30lh1q4hzqhvpdthtptdpkpwhqb06",
				"DEL-BOM-6E5013": "5013|6E|2340|0145|1|205||0|||127|30lh1q4hzqhvpdshtptdpkpwhqb06",
				"DEL-AMD-SG8193": "8193|SG|2010|2150|0|140||0|||127|rl1h1q4hxfhudnshtptdpkpwhqb06",
				"DEL-CCU-UK727": "727|UK|2030|2245|0|215||0|||127|bbjh1q4hjehwtwhtptdpkpwhqb06",
				"DEL-BOM-UK933": "933|UK|1530|1740|0|210||0|||127|30lh1q4hjehnsshtptdpkpwhqb06",
				"DEL-NAG-AI641": "641|AI|2330|0120|1|150||0|||127|6rfh1q4hr9hzkdhtptdpkpwhqb06",
				"GOI-BOM-G83008": "3008|G8|2050|2210|0|120||0|||127|30lhf09hfuhsppuhtptdpkpwhqb06",
				"IXC-BOM-6E264": "264|6E|1250|1520|0|230||0|||127|30lh9~bhzqhtzkhtptdpkpwhqb06",
				"LKO-BOM-G8396": "396|G8|1900|2120|0|220||0|||127|30lh4e0hfuhsnzhtptdpkpwhqb06",
				"GOI-BOM-G83006": "3006|G8|1850|2010|0|120||0|||127|30lhf09hfuhsppzhtptdpkpwhqb06",
				"DEL-RPR-AI477": "477|AI|0525|0705|0|140||0|||127|ogoh1q4hr9hkwwhtptdpkpwhqb06",
				"DEL-JAI-6E736": "736|6E|2200|2255|0|055||0|||127|yr9h1q4hzqhwszhtptdpkpwhqb06",
				"DEL-SXR-G8190": "190|G8|0520|0645|0|125||0|||127|x~oh1q4hfuhdnphtptdpkpwhqb06",
				"IDR-BOM-AI636": "636|AI|1705|1815|0|110||0|||127|30lh91ohr9hzszhtptdpkpwhqb06",
				"DEL-SXR-G8191": "191|G8|0925|1055|0|130||0|||127|x~oh1q4hfuhdndhtptdpkpwhqb06",
				"GOI-BOM-G8380": "380|G8|2350|0110|1|120||0|||127|30lhf09hfuhsuphtptdpkpwhqb06",
				"GOI-BOM-AI664": "664|AI|1520|1635|0|115||0|||127|30lhf09hr9hzzkhtptdpkpwhqb06",
				"DEL-BOM-UK927": "927|UK|0930|1140|0|210||0|||127|30lh1q4hjehntwhtptdpkpwhqb06",
				"HYD-BOM-AI698": "698|AI|2200|2330|0|130||0|||127|30lhc71hr9hznuhtptdpkpwhqb06",
				"DEL-HYD-UK899": "899|UK|1445|1700|0|215||0|||127|c71h1q4hjehunnhtptdpkpwhqb06",
				"DEL-IDR-6E2278": "2278|6E|1345|1510|0|125||0|||127|91oh1q4hzqhttwuhtptdpkpwhqb06",
				"DEL-BLR-I52879": "2879|I5|1740|1945|0|205||0|||127|34oh1q4h9vhtuwnhtptdpkpwhqb06",
				"DEL-ATQ-6E322": "322|6E|1055|1210|0|115||0|||127|r82h1q4hzqhstthtptdpkpwhqb06",
				"DEL-DED-SG2770": "2770|SG|0700|0755|0|055||0|||127|1q1h1q4hxfhtwwphtptdpkpwhqb06",
				"SXR-BOM-G8287": "287|G8|1525|1825|0|300||0|||127|30lhx~ohfuhtuwhtptdpkpwhqb06",
				"DEL-BOM-G82501": "2501|G8|0240|0455|0|215||0|||127|30lh1q4hfuhtvpdhtptdpkpwhqb06",
				"AMD-BOM-6E167": "167|6E|1155|1310|0|115||0|||127|30lhrl1hzqhdzwhtptdpkpwhqb06",
				"DEL-BOM-6E5071": "5071|6E|1545|1750|0|205||0|||127|30lh1q4hzqhvpwdhtptdpkpwhqb06",
				"DEL-BLR-UK813": "813|UK|1735|2025|0|250||0|||127|34oh1q4hjehudshtptdpkpwhqb06",
				"DEL-BOM-UK995": "995|UK|1020|1235|0|215||0|||127|30lh1q4hjehnnvhtptdpkpwhqb06",
				"DEL-BLR-UK815": "815|UK|0805|1050|0|245||0|||127|34oh1q4hjehudvhtptdpkpwhqb06",
				"DEL-PAT-G82511": "2511|G8|1100|1240|0|140||0|||127|gr8h1q4hfuhtvddhtptdpkpwhqb06",
				"DEL-BLR-UK811": "811|UK|0615|0900|0|245||0|||127|34oh1q4hjehuddhtptdpkpwhqb06",
				"DEL-BLR-UK809": "809|UK|1950|2240|0|250||0|||127|34oh1q4hjehupnhtptdpkpwhqb06",
				"DEL-BOM-SG8701": "8701|SG|0720|0935|0|215||0|||127|30lh1q4hxfhuwpdhtptdpkpwhqb06",
				"ATQ-BOM-6E448": "448|6E|1405|1640|0|235||0|||127|30lhr82hzqhkkuhtptdpkpwhqb06",
				"DEL-BLR-UK807": "807|UK|2040|2320|0|240||0|||127|34oh1q4hjehupwhtptdpkpwhqb06",
				"DEL-AMD-I5795": "795|I5|2130|2305|0|135||0|||127|rl1h1q4h9vhwnvhtptdpkpwhqb06",
				"DEL-BOM-UK993": "993|UK|1250|1500|0|210||0|||127|30lh1q4hjehnnshtptdpkpwhqb06",
				"DEL-BOM-SG8709": "8709|SG|1845|2105|0|220||0|||127|30lh1q4hxfhuwpnhtptdpkpwhqb06",
				"GOI-BOM-AI684": "684|AI|0730|0845|0|115||0|||127|30lhf09hr9hzukhtptdpkpwhqb06",
				"DEL-UDR-AI471": "471|AI|1320|1435|0|115||0|||127|j1oh1q4hr9hkwdhtptdpkpwhqb06",
				"DEL-JAI-AI9643": "9643|AI|1930|2035|0|105||0|||127|yr9h1q4hr9hnzkshtptdpkpwhqb06",
				"JLR-BOM-SG2871": "2871|SG|1325|1550|0|225||0|||127|30lhy4ohxfhtuwdhtptdpkpwhqb06",
				"DEL-GOI-UK847": "847|UK|1110|1350|0|240||0|||127|f09h1q4hjehukwhtptdpkpwhqb06",
				"DEL-BLR-I5721": "721|I5|2045|2330|0|245||0|||127|34oh1q4h9vhwtdhtptdpkpwhqb06",
				"DEL-SXR-G8357": "357|G8|0530|0835|0|305||0|||127|x~oh1q4hfuhsvwhtptdpkpwhqb06",
				"DEL-HYD-UK871": "871|UK|2035|2255|0|220||0|||127|c71h1q4hjehuwdhtptdpkpwhqb06",
				"DEL-BOM-UK985": "985|UK|1950|2200|0|210||0|||127|30lh1q4hjehnuvhtptdpkpwhqb06",
				"VNS-BOM-UK622": "622|UK|1610|1820|0|210||0|||127|30lhi6xhjehztthtptdpkpwhqb06",
				"DEL-BOM-SG8169": "8169|SG|1945|2205|0|220||0|||127|30lh1q4hxfhudznhtptdpkpwhqb06",
				"DEL-HYD-UK879": "879|UK|1740|1950|0|210||0|||127|c71h1q4hjehuwnhtptdpkpwhqb06",
				"DEL-BOM-I5332": "332|I5|1540|1750|0|210||0|||127|30lh1q4h9vhssthtptdpkpwhqb06",
				"DEL-BOM-UK981": "981|UK|2140|2350|0|210||0|||127|30lh1q4hjehnudhtptdpkpwhqb06",
				"DEL-BLR-UK817": "817|UK|1605|1850|0|245||0|||127|34oh1q4hjehudwhtptdpkpwhqb06",
				"DEL-BOM-6E665": "665|6E|2155|0005|1|210||0|||127|30lh1q4hzqhzzvhtptdpkpwhqb06",
				"DEL-HYD-AI839": "839|AI|2110|2335|0|225||0|||127|c71h1q4hr9husnhtptdpkpwhqb06",
				"DEL-AMD-6E517": "517|6E|0520|0700|0|140||0|||127|rl1h1q4hzqhvdwhtptdpkpwhqb06",
				"DEL-BLR-UK819": "819|UK|1415|1705|0|250||0|||127|34oh1q4hjehudnhtptdpkpwhqb06",
				"DEL-GOI-I5796": "796|I5|1605|1845|0|240||0|||127|f09h1q4h9vhwnzhtptdpkpwhqb06",
				"DEL-BOM-G8530": "530|G8|0650|0855|0|205||0|||127|30lh1q4hfuhvsphtptdpkpwhqb06",
				"UDR-BOM-UK614": "614|UK|1420|1600|0|140||0|||127|30lhj1ohjehzdkhtptdpkpwhqb06",
				"DEL-CCU-I5578": "578|I5|1240|1455|0|215||0|||127|bbjh1q4h9vhvwuhtptdpkpwhqb06",
				"DEL-GOI-I5798": "798|I5|1240|1520|0|240||0|||127|f09h1q4h9vhwnuhtptdpkpwhqb06",
				"DEL-VNS-AI406": "406|AI|1015|1135|0|120||0|||127|i6xh1q4hr9hkpzhtptdpkpwhqb06",
				"LKO-BOM-AI626": "626|AI|1105|1325|0|220||0|||127|30lh4e0hr9hztzhtptdpkpwhqb06",
				"DEL-RAJ-SG3237": "3237|SG|1705|1935|0|230||0|||127|oryh1q4hxfhstswhtptdpkpwhqb06",
				"BLR-BOM-AI640": "640|AI|0640|0835|0|155||0|||127|30lh34ohr9hzkphtptdpkpwhqb06",
				"DEL-BOM-UK975": "975|UK|0600|0800|0|200||0|||127|30lh1q4hjehnwvhtptdpkpwhqb06",
				"DEL-BOM-UK977": "977|UK|1855|2115|0|220||0|||127|30lh1q4hjehnwwhtptdpkpwhqb06",
				"BLR-BOM-I5941": "941|I5|0445|0640|0|155||0|||127|30lh34oh9vhnkdhtptdpkpwhqb06",
				"DED-BOM-6E933": "933|6E|1655|1915|0|220||0|||127|30lh1q1hzqhnsshtptdpkpwhqb06",
				"UDR-BOM-6E753": "753|6E|1825|2000|0|135||0|||127|30lhj1ohzqhwvshtptdpkpwhqb06",
				"DEL-BLR-AI801": "801|AI|1925|2155|0|230||0|||127|34oh1q4hr9hupdhtptdpkpwhqb06",
				"DEL-BOM-SG8723": "8723|SG|0850|1100|0|210||0|||127|30lh1q4hxfhuwtshtptdpkpwhqb06",
				"DEL-BOM-6E993": "993|6E|0630|0840|0|210||0|||127|30lh1q4hzqhnnshtptdpkpwhqb06",
				"DEL-IDR-AI636": "636|AI|1505|1630|0|125||0|||127|91oh1q4hr9hzszhtptdpkpwhqb06",
				"AMD-BOM-SG636": "636|SG|0525|0635|0|110||0|||127|30lhrl1hxfhzszhtptdpkpwhqb06",
				"NAG-BOM-AI630": "630|AI|2110|2305|0|155||0|||127|30lh6rfhr9hzsphtptdpkpwhqb06",
				"MAA-BOM-AI569": "569|AI|0620|0805|0|145||0|||127|30lhlrrhr9hvznhtptdpkpwhqb06",
				"DEL-BOM-G8323": "323|G8|1820|2035|0|215||0|||127|30lh1q4hfuhstshtptdpkpwhqb06",
				"SXR-BOM-G83002": "3002|G8|1130|1430|0|300||0|||127|30lhx~ohfuhsppthtptdpkpwhqb06",
				"DEL-BOM-6E5162": "5162|6E|1940|2145|0|205||0|||127|30lh1q4hzqhvdzthtptdpkpwhqb06",
				"DEL-BLR-I5740": "740|I5|0755|1040|0|245||0|||127|34oh1q4h9vhwkphtptdpkpwhqb06",
				"DEL-IXC-6E636": "636|6E|1055|1145|0|050||0|||127|9~bh1q4hzqhzszhtptdpkpwhqb06",
				"JAI-BOM-6E6031": "6031|6E|0240|0425|0|145||0|||127|30lhyr9hzqhzpsdhtptdpkpwhqb06",
				"DEL-IXC-UK706": "706|UK|1340|1445|0|105||0|||127|9~bh1q4hjehwpzhtptdpkpwhqb06",
				"DEL-BOM-UK963": "963|UK|0850|1110|0|220||0|||127|30lh1q4hjehnzshtptdpkpwhqb06",
				"DEL-NAG-6E774": "774|6E|0605|0755|0|150||0|||127|6rfh1q4hzqhwwkhtptdpkpwhqb06",
				"DEL-LKO-AI9605": "9605|AI|0655|0830|0|135||0|||127|4e0h1q4hr9hnzpvhtptdpkpwhqb06",
				"DEL-BOM-SG2871": "2871|SG|1130|1550|0|420||0|||127|30lh1q4hxfhtuwdhtptdpkpwhqb06",
				"NAG-BOM-AI628": "628|AI|0755|0920|0|125||0|||127|30lh6rfhr9hztuhtptdpkpwhqb06",
				"DEL-VNS-6E6613": "6613|6E|1045|1210|0|125||0|||127|i6xh1q4hzqhzzdshtptdpkpwhqb06",
				"VTZ-BOM-AI651": "651|AI|1415|1625|0|210||0|||127|30lhi8mhr9hzvdhtptdpkpwhqb06",
				"AMD-BOM-6E5303": "5303|6E|1505|1645|0|140||0|||127|30lhrl1hzqhvspshtptdpkpwhqb06",
				"DEL-RAJ-AI403": "403|AI|1250|1435|0|145||0|||127|oryh1q4hr9hkpshtptdpkpwhqb06",
				"DEL-IXU-AI441": "441|AI|1750|1935|0|145||0|||127|9~jh1q4hr9hkkdhtptdpkpwhqb06",
				"DEL-BOM-I5314": "314|I5|0545|0755|0|210||0|||127|30lh1q4h9vhsdkhtptdpkpwhqb06"
			  },
			  "fares": {
				"DEL-RAJ-AI403*RAJ-BOM-AI602": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5348,
					"c": "INR",
					"bf": 4300,
					"t": 1048,
					"f": 0,
					"d": 0,
					"tf": 5348,
					"ttf": 5348,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.309,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI602": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI403": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "RAJ-BOM-AI602": {
						"at": "Terminal 2"
					  },
					  "DEL-RAJ-AI403": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5260.18,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "320"
					},
					"ind": 320,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYzFzAPQHSElHL6In4RsMih9KSW8uLo/DcVgyKbCNfYgwNW7vEtm2b9D7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI602": {
						"sr": 3,
						"ft": "S"
					  },
					  "AI403": {
						"sr": 3,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-UK815*BLR-BOM-UK854": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 20.264,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK854": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK815": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK854": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK815": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/uL4pG4roUzJsTFVyTuq7khnsBGurXwalCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK854": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK815": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BLR-UK815*BLR-BOM-UK850": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.298,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK815": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK850": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK850": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK815": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/uL4pG4roUzJsTFVyTuq7kwuDPPA3SCzRCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK815": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK850": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-MAA-UK837*MAA-BOM-UK822": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 12.59,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK822": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK837": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-MAA-UK837": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  },
					  "MAA-BOM-UK822": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEffXEFc37ROMFiqpJp4mySNz4KJ7+tYKJtCvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK822": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK837": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-NAG-AI469*NAG-BOM-AI630": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4301,
					"c": "INR",
					"bf": 3415,
					"t": 886,
					"f": 0,
					"d": 0,
					"tf": 4301,
					"ttf": 4301,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 21.514,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI469": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI630": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-NAG-AI469": {
						"dt": "Terminal 3"
					  },
					  "NAG-BOM-AI630": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 4231.01,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7dJ1pG5cHoUDairDDCNywWC8Zb7HUh2aehyketvUi0rpgyKbCNfYgwKnvBXMfui447Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI469": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI630": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-MAA-AI540*MAA-BOM-AI672": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4448,
					"c": "INR",
					"bf": 3555,
					"t": 893,
					"f": 0,
					"d": 0,
					"tf": 4448,
					"ttf": 4448,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.335,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI672": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI540": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-AI672": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-AI540": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 4375.27,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEfCmAjViqaIvViqpJp4mySNa1t6T5iCfBtgyKbCNfYgwKASgztLw8Uj7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI672": {
						"sr": 9,
						"ft": "T"
					  },
					  "AI540": {
						"sr": 9,
						"ft": "T"
					  }
					}
				  }
				],
				"DEL-MAA-AI439*MAA-BOM-AI672": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4028,
					"c": "INR",
					"bf": 3155,
					"t": 873,
					"f": 0,
					"d": 0,
					"tf": 4028,
					"ttf": 4028,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 18.95,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI439": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI672": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-AI672": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-AI439": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3963.07,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEcdA+v9Imoo11iqpJp4mySNa1t6T5iCfBtgyKbCNfYgwHP+b1KOy8lu7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI439": {
						"sr": 6,
						"ft": "S"
					  },
					  "AI672": {
						"sr": 6,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-6E171": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.651,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E171": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E171": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/Ckr0IqMzbGclW6QQXu2eRdRi7zbWSMU9th+ZEoUooTw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E171": {
						"sr": 84,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK815*BLR-BOM-UK858": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.602,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK858": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK815": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK858": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK815": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/uL4pG4roUzJsTFVyTuq7kVx8ZFA3pK3VgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK858": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK815": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-GOI-G83019*GOI-BOM-G8380": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 6126,
					"c": "INR",
					"bf": 4283,
					"t": 1843,
					"f": 0,
					"d": 0,
					"tf": 6126,
					"ttf": 6126,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.642,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83019": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8380": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 329,
					  "tdf": 329
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "GOI-BOM-G8380": {
						"at": "Terminal 1"
					  },
					  "DEL-GOI-G83019": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 5996.32,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "367"
					},
					"ind": 367,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/vzplVzoZHRFhzO8TfE8yKmY2k34IPDCCXLLuUiv5eMNnpj7tV/21oRu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83019": {
						"sr": 6,
						"ft": "GS"
					  },
					  "G8380": {
						"sr": 6,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-IXC-6E636*IXC-BOM-6E264": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.346,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E636": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E264": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-IXC-6E636": {
						"dt": "Terminal 1"
					  },
					  "IXC-BOM-6E264": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8nhUpJa+b/g8AcTBXd5yTzur2T6v3s9V/5gyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E636": {
						"sr": 35,
						"ft": "R"
					  },
					  "6E264": {
						"sr": 35,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-SXR-G8190*SXR-BOM-G83002": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 10056,
					"c": "INR",
					"bf": 8025,
					"t": 2031,
					"f": 0,
					"d": 0,
					"tf": 10056,
					"ttf": 10056,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.2,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83002": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8190": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 369,
					  "tdf": 369
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-SXR-G8190": {
						"dt": "Terminal 2"
					  },
					  "SXR-BOM-G83002": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 9833.57,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "603"
					},
					"ind": 603,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy4pini6vPTw/oy5NkBwJfYS0uWq3bpB7vvLLuUiv5eMNobRjEyj2A0zkymwHBX0SUo=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83002": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G8190": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-JAI-AI9843*JAI-BOM-AI612": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3718,
					"c": "INR",
					"bf": 2780,
					"t": 938,
					"f": 0,
					"d": 0,
					"tf": 3718,
					"ttf": 3718,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.897,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI612": {
						"hb": "5 Kg",
						"cb": "15 Kilograms"
					  },
					  "AI9843": {
						"hb": "5 Kg",
						"cb": "15 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "JAI-BOM-AI612": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  },
					  "DEL-JAI-AI9843": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3659.97,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz71gQeS9pQ7+hvfvLVYY+he5NfvMiApCzuXnDGeaIDiU7LLuUiv5eMNnQtP/2EdxaG4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI612": {
						"sr": 2,
						"ft": "S"
					  },
					  "AI9843": {
						"sr": 2,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-MAA-UK837*MAA-BOM-UK826": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 13.723,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK826": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK837": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK826": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK837": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEffXEFc37ROMFiqpJp4mySNFvQ+t4TcKT1CvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK826": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK837": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-AI887": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.504,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI887": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-AI887": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u92W3mcGVC6yclW6QQXu2eRNk9795mi/AFAxnl2QHB1lw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI887": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-VNS-6E6613*VNS-BOM-6E5377": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4314,
					"c": "INR",
					"bf": 3619,
					"t": 695,
					"f": 0,
					"d": 0,
					"tf": 4314,
					"ttf": 4314,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.088,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5377": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E6613": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-VNS-6E6613": {
						"dt": "Terminal 1"
					  },
					  "VNS-BOM-6E5377": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 4276.63,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7IvjHpZPJB54yM+9F74D3ZuO2vzyLaoTfbdDk/vlq4lvJVukEF7tnkSm9JjTBH9Lll76+7lczXkQ=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5377": {
						"sr": 12,
						"ft": "R"
					  },
					  "6E6613": {
						"sr": 12,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-MAA-UK837*MAA-BOM-UK828": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.523,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK837": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK828": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK828": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK837": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEffXEFc37ROMFiqpJp4mySNa9Js+usI3rRCvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK837": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK828": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-RAJ-AI495*RAJ-BOM-AI656": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5348,
					"c": "INR",
					"bf": 4300,
					"t": 1048,
					"f": 0,
					"d": 0,
					"tf": 5348,
					"ttf": 5348,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.242,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI656": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI495": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-RAJ-AI495": {},
					  "RAJ-BOM-AI656": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5260.18,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "320"
					},
					"ind": 320,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYyoisccFR7vGr6In4RsMih9D0ZWGmpYkiVgyKbCNfYgwNW7vEtm2b9D7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI656": {
						"sr": 3,
						"ft": "S"
					  },
					  "AI495": {
						"sr": 3,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-AI801*BLR-BOM-AI640": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5241,
					"c": "INR",
					"bf": 4310,
					"t": 931,
					"f": 0,
					"d": 0,
					"tf": 5241,
					"ttf": 5241,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.911,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI801": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI640": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-AI801": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-AI640": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5153.56,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "314"
					},
					"ind": 314,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+E/o6IZoYuhZsTFVyTuq7kzEVSengXhiNgyKbCNfYgwP8PjH7aMk617Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI801": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI640": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-AI665": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.537,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI665": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-AI665": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/yDFWANZmHhslW6QQXu2eRNk9795mi/AFAxnl2QHB1lw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI665": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-UK815*BLR-BOM-UK866": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.069,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK866": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK815": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK866": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK815": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/uL4pG4roUzJsTFVyTuq7km5uyOzaujphgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK866": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK815": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-6E6197": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2209,
					"t": 572,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.569,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6197": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E6197": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2756.91,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+eXl8ZyOYtgFBjMaVPIqa3HT+7NAjbv3J3B/cWMT8z7Q==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6197": {
						"sr": 70,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-AMD-G8719*AMD-BOM-G82505": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3622,
					"c": "INR",
					"bf": 2297,
					"t": 1325,
					"f": 0,
					"d": 0,
					"tf": 3622,
					"ttf": 3622,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.24,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8719": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G82505": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-G82505": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "DEL-AMD-G8719": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3547.6,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H9C7qMrp2mZs6opjOalXrMxFoQzcG48PanLLuUiv5eMNkpMfEtSQuTeSMMLUxrC9LM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8719": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G82505": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-UDR-6E2274*UDR-BOM-6E753": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.246,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E2274": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E753": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "UDR-BOM-6E753": {
						"at": "Terminal 1"
					  },
					  "DEL-UDR-6E2274": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7bC7vxRRnG9c2jUYQh9YLHGY0XipmcJcc9E5KIYuJCy3LLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E2274": {
						"sr": 49,
						"ft": "R"
					  },
					  "6E753": {
						"sr": 49,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-6E181": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.717,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E181": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E181": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9jGGnz5KOlM8lW6QQXu2eRdRi7zbWSMU9th+ZEoUooTw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E181": {
						"sr": 75,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-CCU-UK747*CCU-BOM-UK772": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 19.329,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK747": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK772": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK772": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK747": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQrdVlY3SiIfD/T4z+KegT3lz8sYXzW1wVNgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK747": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK772": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-GOI-6E367*GOI-BOM-6E5306": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4314,
					"c": "INR",
					"bf": 3619,
					"t": 695,
					"f": 0,
					"d": 0,
					"tf": 4314,
					"ttf": 4314,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.355,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5306": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E367": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "GOI-BOM-6E5306": {
						"at": "Terminal 2"
					  },
					  "DEL-GOI-6E367": {
						"dt": "Terminal 1"
					  }
					},
					"ottf": 4276.63,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/tuuGoaRtojXrzGcMImpgPo+82/ekzPBSbLLuUiv5eMNnX5sYdJtptR7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5306": {
						"sr": 8,
						"ft": "R"
					  },
					  "6E367": {
						"sr": 8,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-CCU-UK747*CCU-BOM-UK776": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 22.263,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK776": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK747": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK776": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK747": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQrdVlY3SiIfD/T4z+KegT3li9E9THdgGzJgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK776": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK747": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-CCU-UK705*CCU-BOM-UK776": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 21.763,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK776": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK705": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK776": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK705": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQqADlMROGbpufT4z+KegT3li9E9THdgGzJgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK776": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK705": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-LKO-AI9605*LKO-BOM-AI626": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5157,
					"c": "INR",
					"bf": 4150,
					"t": 1007,
					"f": 0,
					"d": 0,
					"tf": 5157,
					"ttf": 5157,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.194,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI9605": {
						"hb": "5 Kg",
						"cb": "15 Kilograms"
					  },
					  "AI626": {
						"hb": "5 Kg",
						"cb": "15 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "LKO-BOM-AI626": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  },
					  "DEL-LKO-AI9605": {}
					},
					"ottf": 5072.27,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "309"
					},
					"ind": 309,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTmlpA1IEVXDcYgJ5p0V6IqLYKpekuZDA6XLLuUiv5eMNriNdn797qcj4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI9605": {
						"sr": 4,
						"ft": "S"
					  },
					  "AI626": {
						"sr": 4,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-CCU-UK705*CCU-BOM-UK772": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 18.829,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK705": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK772": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK772": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK705": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQqADlMROGbpufT4z+KegT3lz8sYXzW1wVNgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK705": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK772": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-6E6186": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.684,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6186": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E6186": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/4to7PKfgUa2K4jkcXOKaheo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6186": {
						"sr": 79,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-IXC-G8108*IXC-BOM-G8382": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4182,
					"c": "INR",
					"bf": 2431,
					"t": 1751,
					"f": 0,
					"d": 0,
					"tf": 4182,
					"ttf": 4182,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.976,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8108": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8382": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "IXC-BOM-G8382": {
						"at": "Terminal 1"
					  },
					  "DEL-IXC-G8108": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 4098.23,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8nYM+9hAqgVWAcTBXd5yTzu9Kkb5lU7PS1gyKbCNfYgwBGIkIdO+sTA7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8108": {
						"sr": 2,
						"ft": "GS"
					  },
					  "G8382": {
						"sr": 2,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-MAA-I5761*MAA-BOM-I5803": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3879,
					"c": "INR",
					"bf": 3278,
					"t": 601,
					"f": 0,
					"d": 0,
					"tf": 3879,
					"ttf": 3879,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.527,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5803": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5761": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-I5803": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "DEL-MAA-I5761": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 3636.66,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEdWrF+2BPhfUliqpJp4mySNb4MnueTKdcdgyKbCNfYgwBxlVh3VIVln7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5803": {
						"sr": 119,
						"ft": "EP"
					  },
					  "I5761": {
						"sr": 119,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-RAJ-SG3237*RAJ-BOM-SG224": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5848,
					"c": "INR",
					"bf": 4332,
					"t": 1516,
					"f": 0,
					"d": 0,
					"tf": 5848,
					"ttf": 5848,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.789,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG224": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG3237": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "RAJ-BOM-SG224": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-RAJ-SG3237": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "350"
					},
					"ind": 350,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYyKq0HVl85SnEP/6XSAlyB0zvClE3FGdSTLLuUiv5eMNjpnL/ZNHyuifZk/4Aoz45k=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG224": {
						"sr": 30,
						"ft": "RS"
					  },
					  "SG3237": {
						"sr": 30,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-GOI-6E2167*GOI-BOM-6E418": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.913,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E418": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E2167": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-6E2167": {
						"dt": "Terminal 2"
					  },
					  "GOI-BOM-6E418": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/uOvk1Sr6loEddhpcW4Dc9PNhe/s1cr4GDLLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E418": {
						"sr": 55,
						"ft": "R"
					  },
					  "6E2167": {
						"sr": 55,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-HYD-AI839*HYD-BOM-AI616": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.115,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI839": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI616": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-AI839": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-AI616": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAb3yL5b0vnoqv6LuWTGDycsyi5/eohLjc1gyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI839": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI616": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-HYD-UK879*HYD-BOM-UK876": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5548,
					"c": "INR",
					"bf": 4682,
					"t": 866,
					"f": 0,
					"d": 0,
					"tf": 5548,
					"ttf": 5548,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 17.562,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK876": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK879": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-UK876": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-UK879": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAbNbPWF7SuP1v6LuWTGDycszz5kLbi0kqdgyKbCNfYgwMHZIKrQKuHy",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK876": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK879": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-HYD-AI542*HYD-BOM-AI616": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.682,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI616": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI542": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-AI542": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-AI616": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAY2oYSRm6wczf6LuWTGDycsyi5/eohLjc1gyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI616": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI542": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-STV-SG8473*STV-BOM-SG2773": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5887,
					"c": "INR",
					"bf": 5140,
					"t": 747,
					"f": 0,
					"d": 0,
					"tf": 5887,
					"ttf": 5887,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.099,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG2773": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG8473": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "STV-BOM-SG2773": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-STV-SG8473": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 5774.41,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "353"
					},
					"ind": 353,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy7w1vVdoOEB8xG/xTdFRCIwHq60WJyxi47JVukEF7tnkarHua5pQMMYN4bk1SCBReU=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG2773": {
						"sr": 5,
						"ft": "RS"
					  },
					  "SG8473": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-MAA-AI429*MAA-BOM-AI672": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4448,
					"c": "INR",
					"bf": 3555,
					"t": 893,
					"f": 0,
					"d": 0,
					"tf": 4448,
					"ttf": 4448,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 17.702,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI429": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI672": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-AI672": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-AI429": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4375.27,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEez4DGr6m+AHViqpJp4mySNa1t6T5iCfBtgyKbCNfYgwKASgztLw8Uj7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI429": {
						"sr": 9,
						"ft": "T"
					  },
					  "AI672": {
						"sr": 9,
						"ft": "T"
					  }
					}
				  }
				],
				"DEL-BLR-UK819*BLR-BOM-UK866": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 9.464,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK866": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK819": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK819": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK866": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/0IBC//3EOcZsTFVyTuq7km5uyOzaujphCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK866": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK819": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-UK879*HYD-BOM-UK874": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5548,
					"c": "INR",
					"bf": 4682,
					"t": 866,
					"f": 0,
					"d": 0,
					"tf": 5548,
					"ttf": 5548,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.328,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK879": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK874": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-UK874": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-UK879": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAbNbPWF7SuP1v6LuWTGDycsKdyrZm9jJP5gyKbCNfYgwMHZIKrQKuHy",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK879": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK874": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-AI865": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.537,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI865": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-AI865": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8y8NPn32ZlHclW6QQXu2eRNk9795mi/AFAxnl2QHB1lw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI865": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-AI624": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.57,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI624": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-AI624": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8P8XKaoGz418lW6QQXu2eRNk9795mi/AFAxnl2QHB1lw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI624": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-HYD-AI542*HYD-BOM-AI620": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.682,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI620": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI542": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-AI620": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-AI542": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAY2oYSRm6wczf6LuWTGDycseQmyYSyE3wNgyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI620": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI542": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-UK811*BLR-BOM-UK850": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.169,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK811": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK850": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK850": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK811": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u804i1680QA75sTFVyTuq7kwuDPPA3SCzRgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK811": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK850": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-CCU-I5578*CCU-BOM-I5316": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5517,
					"c": "INR",
					"bf": 4838,
					"t": 679,
					"f": 0,
					"d": 0,
					"tf": 5517,
					"ttf": 5517,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.844,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5316": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5578": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-I5316": {
						"at": "Terminal 1"
					  },
					  "DEL-CCU-I5578": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5463.9,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "331"
					},
					"ind": 331,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQoJKnRxq3mwFPT4z+KegT3l+/9ezOt3PxJgyKbCNfYgwCQDa0vi2pnY7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5316": {
						"sr": 20,
						"ft": "EP"
					  },
					  "I5578": {
						"sr": 20,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-LKO-AI431*LKO-BOM-AI626": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5199,
					"c": "INR",
					"bf": 4270,
					"t": 929,
					"f": 0,
					"d": 0,
					"tf": 5199,
					"ttf": 5199,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.653,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI626": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI431": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "LKO-BOM-AI626": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  },
					  "DEL-LKO-AI431": {}
					},
					"ottf": 5112.33,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "311"
					},
					"ind": 311,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTnc0+dVzZIAOePf4zDMLBU0B4n4/jZYjBlgyKbCNfYgwKaz0vDNDl4/7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI626": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI431": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-AMD-G8717*AMD-BOM-G8413": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3622,
					"c": "INR",
					"bf": 2297,
					"t": 1325,
					"f": 0,
					"d": 0,
					"tf": 3622,
					"ttf": 3622,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.373,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8413": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8717": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-AMD-G8717": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  },
					  "AMD-BOM-G8413": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3547.6,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H+LNJS8Es4eK6opjOalXrMx4DV5tA4UIadgyKbCNfYgwKpgjBeJ3PhG7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8413": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G8717": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-DED-SG2770*DED-BOM-SG290": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7430,
					"c": "INR",
					"bf": 6611,
					"t": 819,
					"f": 0,
					"d": 0,
					"tf": 7430,
					"ttf": 7430,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.191,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG2770": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG290": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-DED-SG2770": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  },
					  "DED-BOM-SG290": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 7322.73,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "445"
					},
					"ind": 445,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7oDfdgjLN6viDKOU7rSFTpxqV40j2JcE9RpN3/EBw8nnLLuUiv5eMNuefS/984CP64Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG2770": {
						"sr": 25,
						"ft": "RS"
					  },
					  "SG290": {
						"sr": 25,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-HYD-6E5406*HYD-BOM-6E235": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.946,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E235": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E5406": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-6E5406": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-6E235": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAZ5HvWCPvZgTRdIUGZiw6HcuO7Hv2tW4bjLLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E235": {
						"sr": 38,
						"ft": "R"
					  },
					  "6E5406": {
						"sr": 38,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-AI502*BLR-BOM-AI610": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5241,
					"c": "INR",
					"bf": 4310,
					"t": 931,
					"f": 0,
					"d": 0,
					"tf": 5241,
					"ttf": 5241,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.678,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI610": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI502": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-AI502": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-AI610": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5153.56,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "314"
					},
					"ind": 314,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8exXCEBgw/mpsTFVyTuq7kicuW7f84gbdgyKbCNfYgwP8PjH7aMk617Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI610": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI502": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-DED-SG2033*DED-BOM-SG290": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7430,
					"c": "INR",
					"bf": 6611,
					"t": 819,
					"f": 0,
					"d": 0,
					"tf": 7430,
					"ttf": 7430,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.191,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG2033": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG290": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-DED-SG2033": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  },
					  "DED-BOM-SG290": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 7322.73,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "445"
					},
					"ind": 445,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7oDfdgjLN6viH0qZbg1B4gEWVvwW/UL7tRpN3/EBw8nnLLuUiv5eMNuefS/984CP64Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG2033": {
						"sr": 24,
						"ft": "RS"
					  },
					  "SG290": {
						"sr": 24,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-AMD-G8717*AMD-BOM-G8537": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3622,
					"c": "INR",
					"bf": 2297,
					"t": 1325,
					"f": 0,
					"d": 0,
					"tf": 3622,
					"ttf": 3622,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.44,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8537": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8717": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-AMD-G8717": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  },
					  "AMD-BOM-G8537": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3547.6,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H+LNJS8Es4eK6opjOalXrMxF72Ta1Mo+HFgyKbCNfYgwKpgjBeJ3PhG7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8537": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G8717": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-MAA-AI429*MAA-BOM-AI569": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4448,
					"c": "INR",
					"bf": 3555,
					"t": 893,
					"f": 0,
					"d": 0,
					"tf": 4448,
					"ttf": 4448,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.102,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI569": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI429": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-MAA-AI429": {
						"dt": "Terminal 3"
					  },
					  "MAA-BOM-AI569": {
						"dt": "Terminal 4",
						"at": "Terminal 2"
					  }
					},
					"ottf": 4375.27,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEez4DGr6m+AHViqpJp4mySNg0tM8AbFpU9gyKbCNfYgwKASgztLw8Uj7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI569": {
						"sr": 9,
						"ft": "T"
					  },
					  "AI429": {
						"sr": 9,
						"ft": "T"
					  }
					}
				  }
				],
				"DEL-SXR-G83011*SXR-BOM-G8287": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 11137,
					"c": "INR",
					"bf": 9054,
					"t": 2083,
					"f": 0,
					"d": 0,
					"tf": 11137,
					"ttf": 11137,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.949,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8287": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G83011": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 369,
					  "tdf": 369
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "SXR-BOM-G8287": {
						"at": "Terminal 1"
					  },
					  "DEL-SXR-G83011": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 10889.07,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "668"
					},
					"ind": 668,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy5nl5xsfyDIPPBsPAddvgcMzVy42wIfM9TLLuUiv5eMNjyKgxn0ypf8Y2zK8M5MquM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8287": {
						"sr": 2,
						"ft": "GS"
					  },
					  "G83011": {
						"sr": 2,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BLR-UK819*BLR-BOM-UK850": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 8.831,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK819": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK850": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK850": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK819": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/0IBC//3EOcZsTFVyTuq7kwuDPPA3SCzRCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK819": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK850": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-6E843": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.651,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E843": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E843": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u92OU+e3NfVH8lW6QQXu2eRdRi7zbWSMU9th+ZEoUooTw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E843": {
						"sr": 103,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-GOI-I5798*GOI-BOM-I5472": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5753,
					"c": "INR",
					"bf": 5064,
					"t": 689,
					"f": 0,
					"d": 0,
					"tf": 5753,
					"ttf": 5753,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 11.318,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5798": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5472": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "GOI-BOM-I5472": {
						"at": "Terminal 1"
					  },
					  "DEL-GOI-I5798": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5697.63,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "345"
					},
					"ind": 345,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/vfVlgEsa18rLzGcMImpgPorGjs/TVXScxgyKbCNfYgwK6RoVcToHb67Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5798": {
						"sr": 2,
						"ft": "EP"
					  },
					  "I5472": {
						"sr": 2,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-UK819*BLR-BOM-UK854": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 17.798,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK854": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK819": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK854": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK819": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/0IBC//3EOcZsTFVyTuq7khnsBGurXwalCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK854": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK819": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-UDR-UK627*UDR-BOM-UK614": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5979,
					"c": "INR",
					"bf": 5093,
					"t": 886,
					"f": 0,
					"d": 0,
					"tf": 5979,
					"ttf": 5979,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.754,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK614": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK627": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-UDR-UK627": {
						"dt": "Terminal 3"
					  },
					  "UDR-BOM-UK614": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5693.72,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "358"
					},
					"ind": 358,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7bC7vxRRnG9c2uR5ymZnZmrxQJMDNGQIt79i3dkYKfrpgyKbCNfYgwL6KNFyEYnXq7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK614": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK627": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BLR-UK819*BLR-BOM-UK858": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.135,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK858": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK819": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK819": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK858": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/0IBC//3EOcZsTFVyTuq7kVx8ZFA3pK3VgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK858": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK819": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-DED-6E966*DED-BOM-6E933": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4621,
					"c": "INR",
					"bf": 3911,
					"t": 710,
					"f": 0,
					"d": 0,
					"tf": 4621,
					"ttf": 4621,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.206,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E933": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E966": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-DED-6E966": {
						"dt": "Terminal 1"
					  },
					  "DED-BOM-6E933": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 4580.97,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7oDfdgjLN6vijlXBFvlMYR+Qohn4JfUOW/onunqa+De5gyKbCNfYgwCKGUCei4SMI7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E933": {
						"sr": 9,
						"ft": "R"
					  },
					  "6E966": {
						"sr": 9,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-RAJ-AI495*RAJ-BOM-AI602": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5348,
					"c": "INR",
					"bf": 4300,
					"t": 1048,
					"f": 0,
					"d": 0,
					"tf": 5348,
					"ttf": 5348,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 11.675,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI602": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI495": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-RAJ-AI495": {},
					  "RAJ-BOM-AI602": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5260.18,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "320"
					},
					"ind": 320,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYyoisccFR7vGr6In4RsMih9KSW8uLo/DcVgyKbCNfYgwNW7vEtm2b9D7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI602": {
						"sr": 3,
						"ft": "S"
					  },
					  "AI495": {
						"sr": 3,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-I5716": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2232,
					"t": 549,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.602,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5716": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-I5716": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2754.23,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8jocT6hxFZUslW6QQXu2eRFpoT8PVfX8ls/MBZ2wJYrw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5716": {
						"sr": 97,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-UK811*BLR-BOM-UK864": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 18.198,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK864": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK811": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK864": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK811": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u804i1680QA75sTFVyTuq7kSqy2XPVrgwNCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK864": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK811": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-6E5115": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.684,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5115": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5115": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8+BpEscm+212l8UNbXcFHEeo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5115": {
						"sr": 109,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-6E6202": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2209,
					"t": 572,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.502,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6202": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E6202": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2756.91,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9cQ9wybZNNnLfl1H2cOBbVHT+7NAjbv3J3B/cWMT8z7Q==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6202": {
						"sr": 65,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK811*BLR-BOM-UK854": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 20.998,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK854": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK811": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK854": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK811": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u804i1680QA75sTFVyTuq7khnsBGurXwalCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK854": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK811": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-AMD-SG8193*AMD-BOM-SG636": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3777,
					"c": "INR",
					"bf": 3132,
					"t": 645,
					"f": 0,
					"d": 0,
					"tf": 3777,
					"ttf": 3777,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.933,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG8193": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG636": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-SG636": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-AMD-SG8193": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3722.47,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H9GHDry+7pt17JLfxynqjQWiXg7xJx3J3XLLuUiv5eMNsVBZRcd7tV+UKgAPjC/U6M=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG8193": {
						"sr": 5,
						"ft": "RS"
					  },
					  "SG636": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-RAJ-AI403*RAJ-BOM-AI656": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5348,
					"c": "INR",
					"bf": 4300,
					"t": 1048,
					"f": 0,
					"d": 0,
					"tf": 5348,
					"ttf": 5348,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 17.875,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI656": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI403": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "RAJ-BOM-AI656": {
						"at": "Terminal 2"
					  },
					  "DEL-RAJ-AI403": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5260.18,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "320"
					},
					"ind": 320,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYzFzAPQHSElHL6In4RsMih9D0ZWGmpYkiVgyKbCNfYgwNW7vEtm2b9D7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI656": {
						"sr": 4,
						"ft": "S"
					  },
					  "AI403": {
						"sr": 4,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-6E5347": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.651,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5347": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5347": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/4jTQT1rPmd1BjMaVPIqa3eo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5347": {
						"sr": 96,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-SXR-G83011*SXR-BOM-G83002": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 9561,
					"c": "INR",
					"bf": 7554,
					"t": 2007,
					"f": 0,
					"d": 0,
					"tf": 9561,
					"ttf": 9561,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.437,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83002": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G83011": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 359,
					  "tdf": 359
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-SXR-G83011": {
						"dt": "Terminal 2"
					  },
					  "SXR-BOM-G83002": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 9350.25,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "573"
					},
					"ind": 573,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy5nl5xsfyDIPPBsPAddvgcMlG2Ch1+BERbJVukEF7tnkZ72t3bKBDNgsqWo29Ar9sY=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83002": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G83011": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-IXC-UK637*IXC-BOM-UK652": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 8371,
					"c": "INR",
					"bf": 7371,
					"t": 1000,
					"f": 0,
					"d": 0,
					"tf": 8371,
					"ttf": 8371,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.289,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK637": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK652": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 349,
					  "tdf": 349
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "IXC-BOM-UK652": {
						"at": "Terminal 2"
					  },
					  "DEL-IXC-UK637": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 8052.84,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "502"
					},
					"ind": 502,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8nlAPI0wGFUGQcTBXd5yTzuNW2O4duYRPtgyKbCNfYgwPK8QzQZ14fI7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK637": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK652": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-AMD-6E517*AMD-BOM-6E214": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.846,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E517": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E214": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-6E214": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "DEL-AMD-6E517": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H9HBNt4AHysFaopjOalXrMxsYHE+k59/h9gyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E517": {
						"sr": 68,
						"ft": "R"
					  },
					  "6E214": {
						"sr": 68,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-GOI-I5796*GOI-BOM-I5472": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5525,
					"c": "INR",
					"bf": 4846,
					"t": 679,
					"f": 0,
					"d": 0,
					"tf": 5525,
					"ttf": 5525,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.815,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5796": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5472": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-I5796": {
						"dt": "Terminal 3"
					  },
					  "GOI-BOM-I5472": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 5471.82,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "331"
					},
					"ind": 331,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/sszfa8migBJLzGcMImpgPorGjs/TVXScxgyKbCNfYgwJ+O6jIDNr2+7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5796": {
						"sr": 1,
						"ft": "EP"
					  },
					  "I5472": {
						"sr": 1,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-RPR-AI477*RPR-VTZ-AI651*VTZ-BOM-AI651": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5230,
					"c": "INR",
					"bf": 4300,
					"t": 930,
					"f": 0,
					"d": 0,
					"tf": 5230,
					"ttf": 5230,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.538,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI477": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI651": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "VTZ-BOM-AI651": {
						"at": "Terminal 2"
					  },
					  "DEL-RPR-AI477": {
						"dt": "Terminal 3"
					  },
					  "RPR-VTZ-AI651": {}
					},
					"ottf": 5142.75,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "313"
					},
					"ind": 313,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYyXNeCmntYZ/BtunEA8w77MTZFoCE0wraJQnJcVBr6/cWYbwQj4nlvRECYyx8oLwy1OIb56AYtHXw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI477": {
						"sr": 2,
						"ft": "S"
					  },
					  "AI651": {
						"sr": 2,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-UK811*BLR-BOM-UK866": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.802,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK866": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK811": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK866": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK811": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u804i1680QA75sTFVyTuq7km5uyOzaujphgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK866": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK811": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-AMD-I5795*AMD-BOM-I5359": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3685,
					"c": "INR",
					"bf": 3094,
					"t": 591,
					"f": 0,
					"d": 0,
					"tf": 3685,
					"ttf": 3685,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.678,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5359": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5795": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-AMD-I5795": {
						"dt": "Terminal 3"
					  },
					  "AMD-BOM-I5359": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 3649.54,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H8KNZdmKRzrE6opjOalXrMxc9UTSQAHFQRgyKbCNfYgwO3emwBNeCoj7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5359": {
						"sr": 57,
						"ft": "EP"
					  },
					  "I5795": {
						"sr": 57,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-UK813*BLR-BOM-UK858": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.664,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK813": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK858": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK858": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK813": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u91aWMPAGE5U5sTFVyTuq7kVx8ZFA3pK3VCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK813": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK858": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-IDR-6E2278*IDR-BOM-6E957": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3335,
					"c": "INR",
					"bf": 2686,
					"t": 649,
					"f": 0,
					"d": 0,
					"tf": 3335,
					"ttf": 3335,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.034,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E957": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E2278": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "IDR-BOM-6E957": {
						"at": "Terminal 1"
					  },
					  "DEL-IDR-6E2278": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 3306.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8k2jUYQh9YLHOoxrMyrRO+F/8JfD2pJ4nzLLuUiv5eMNpcaH5FhDYGT4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E957": {
						"sr": 20,
						"ft": "R"
					  },
					  "6E2278": {
						"sr": 20,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK813*BLR-BOM-UK852": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.069,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK813": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK852": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK852": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK813": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u91aWMPAGE5U5sTFVyTuq7kH5vcgAQjXoVgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK813": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK852": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-MAA-AI540*MAA-BOM-AI569": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4448,
					"c": "INR",
					"bf": 3555,
					"t": 893,
					"f": 0,
					"d": 0,
					"tf": 4448,
					"ttf": 4448,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.735,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI569": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI540": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-MAA-AI540": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  },
					  "MAA-BOM-AI569": {
						"dt": "Terminal 4",
						"at": "Terminal 2"
					  }
					},
					"ottf": 4375.27,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEfCmAjViqaIvViqpJp4mySNg0tM8AbFpU9gyKbCNfYgwKASgztLw8Uj7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI569": {
						"sr": 9,
						"ft": "T"
					  },
					  "AI540": {
						"sr": 9,
						"ft": "T"
					  }
					}
				  }
				],
				"DEL-HYD-AI560*HYD-BOM-AI620": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.249,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI560": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI620": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-AI620": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-AI560": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYwToEmxkKo9/6LuWTGDycseQmyYSyE3wNgyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI560": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI620": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-I52879*BLR-BOM-I5941": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5157,
					"c": "INR",
					"bf": 4496,
					"t": 661,
					"f": 0,
					"d": 0,
					"tf": 5157,
					"ttf": 5157,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.794,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I52879": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5941": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-I52879": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-I5941": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 5107.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "309"
					},
					"ind": 309,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/p6SWHhhomDNDq53+K/2AdE+nkHwk9NKnLLuUiv5eMNpaPr8hJSxX9Y2zK8M5MquM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I52879": {
						"sr": 54,
						"ft": "EP"
					  },
					  "I5941": {
						"sr": 54,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BOM-6E5328": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.684,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5328": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5328": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u80jy3pHEfT4GCntm6QrZA9eo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5328": {
						"sr": 118,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-RAJ-SG3237*RAJ-BOM-SG436": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5848,
					"c": "INR",
					"bf": 4332,
					"t": 1516,
					"f": 0,
					"d": 0,
					"tf": 5848,
					"ttf": 5848,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.156,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG436": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG3237": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "RAJ-BOM-SG436": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-RAJ-SG3237": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "350"
					},
					"ind": 350,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYyKq0HVl85SnEP/6XSAlyB0hRreGU4zjFnLLuUiv5eMNjpnL/ZNHyuifZk/4Aoz45k=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG436": {
						"sr": 30,
						"ft": "RS"
					  },
					  "SG3237": {
						"sr": 30,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-BOM-G8392": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.571,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8392": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8392": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/2GZ27CRhA9clW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8392": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-AMD-6E161*AMD-BOM-6E5303": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.413,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E161": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E5303": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-AMD-6E161": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "AMD-BOM-6E5303": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H+JknDkWZMXlqopjOalXrMx5HW0uC6K9f7LLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E161": {
						"sr": 53,
						"ft": "R"
					  },
					  "6E5303": {
						"sr": 53,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-IXC-UK706*IXC-BOM-UK652": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 8371,
					"c": "INR",
					"bf": 7371,
					"t": 1000,
					"f": 0,
					"d": 0,
					"tf": 8371,
					"ttf": 8371,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.689,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK706": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK652": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 349,
					  "tdf": 349
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-IXC-UK706": {
						"dt": "Terminal 3"
					  },
					  "IXC-BOM-UK652": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 8052.84,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "502"
					},
					"ind": 502,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8m7rKeV5YBIaQcTBXd5yTzuNW2O4duYRPtgyKbCNfYgwPK8QzQZ14fI7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK706": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK652": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BLR-UK813*BLR-BOM-UK866": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 8.131,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK866": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK813": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK813": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK866": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u91aWMPAGE5U5sTFVyTuq7km5uyOzaujphCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK866": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK813": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-AI560*HYD-BOM-AI616": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.249,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI560": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI616": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-AI560": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-AI616": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYwToEmxkKo9/6LuWTGDycsyi5/eohLjc1gyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI560": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI616": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-AI678": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.57,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI678": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-AI678": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u92HgD/+wYHEslW6QQXu2eRNk9795mi/AFAxnl2QHB1lw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI678": {
						"sr": 8,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-6E5031": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.684,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5031": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5031": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u//r/ShaTAr8yA+oE54nzXGeo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5031": {
						"sr": 139,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-CCU-UK707*CCU-BOM-UK776": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 17.763,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK776": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK707": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK776": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK707": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQpgZsX3lUIu5fT4z+KegT3li9E9THdgGzJgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK776": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK707": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-LKO-6E546*LKO-BOM-6E891": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.413,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E546": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E891": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "LKO-BOM-6E891": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  },
					  "DEL-LKO-6E546": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTnekJwsuO77QePf4zDMLBU03ScbLdRrfJhgyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E546": {
						"sr": 55,
						"ft": "R"
					  },
					  "6E891": {
						"sr": 55,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-VNS-UK671*VNS-BOM-UK622": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 8444,
					"c": "INR",
					"bf": 7441,
					"t": 1003,
					"f": 0,
					"d": 0,
					"tf": 8444,
					"ttf": 8444,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 18,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK622": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK671": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 349,
					  "tdf": 349
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-VNS-UK671": {
						"dt": "Terminal 3"
					  },
					  "VNS-BOM-UK622": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 8124.82,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "506"
					},
					"ind": 506,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7IvjHpZPJB57uU3A/MnhvIXmmLMgGMMwyu496i0dqdgBgyKbCNfYgwCJ5WY9hZCnG7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK622": {
						"sr": 7,
						"ft": "V"
					  },
					  "UK671": {
						"sr": 7,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-GOI-AI883*GOI-BOM-AI664": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5881,
					"c": "INR",
					"bf": 4920,
					"t": 961,
					"f": 0,
					"d": 0,
					"tf": 5881,
					"ttf": 5881,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.462,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI664": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI883": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-AI883": {
						"dt": "Terminal 3"
					  },
					  "GOI-BOM-AI664": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5781.67,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/vtm8x+F5l+/7zGcMImpgPoTRmzcBFAOMJgyKbCNfYgwEYxfCjkivit7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI664": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI883": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-HYD-6E816*HYD-BOM-6E903": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.913,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E903": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E816": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-6E816": {
						"dt": "Terminal 1"
					  },
					  "HYD-BOM-6E903": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAbfsTJKBCmJ9f6LuWTGDycsfWWYlEn1r5FgyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E903": {
						"sr": 31,
						"ft": "R"
					  },
					  "6E816": {
						"sr": 31,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-CCU-UK707*CCU-BOM-UK772": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.829,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK707": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK772": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK772": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK707": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQpgZsX3lUIu5fT4z+KegT3lz8sYXzW1wVNgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK707": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK772": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-LKO-AI411*LKO-BOM-AI626": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5199,
					"c": "INR",
					"bf": 4270,
					"t": 929,
					"f": 0,
					"d": 0,
					"tf": 5199,
					"ttf": 5199,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.153,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI626": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI411": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "LKO-BOM-AI626": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  },
					  "DEL-LKO-AI411": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5112.33,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "311"
					},
					"ind": 311,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTm7U7XthJ8goOPf4zDMLBU0B4n4/jZYjBlgyKbCNfYgwKaz0vDNDl4/7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI626": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI411": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-UK953": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3379,
					"c": "INR",
					"bf": 2746,
					"t": 633,
					"f": 0,
					"d": 0,
					"tf": 3379,
					"ttf": 3379,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.893,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK953": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK953": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+h8Vgey1mCfslW6QQXu2eR9P0lti1RNteDbBOimm1r/w==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK953": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-UK955": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5871,
					"c": "INR",
					"bf": 5120,
					"t": 751,
					"f": 0,
					"d": 0,
					"tf": 5871,
					"ttf": 5871,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 4.456,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK955": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK955": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5691.2,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/ckVZhQtiexclW6QQXu2eR1SzKbjJBnhLctlopdfUqWw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK955": {
						"sr": 9,
						"ft": "Q"
					  }
					}
				  }
				],
				"DEL-BLR-AI504*BLR-BOM-AI640": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5241,
					"c": "INR",
					"bf": 4310,
					"t": 931,
					"f": 0,
					"d": 0,
					"tf": 5241,
					"ttf": 5241,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.178,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI504": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI640": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-AI504": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-AI640": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5153.56,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "314"
					},
					"ind": 314,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9ebgF+4X23g5sTFVyTuq7kzEVSengXhiNgyKbCNfYgwP8PjH7aMk617Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI504": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI640": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-I5482": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2232,
					"t": 549,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.535,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5482": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-I5482": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2754.23,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9pYtUy8SnQhMlW6QQXu2eRFpoT8PVfX8ls/MBZ2wJYrw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5482": {
						"sr": 106,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-NAG-AI641*NAG-BOM-AI630": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4301,
					"c": "INR",
					"bf": 3415,
					"t": 886,
					"f": 0,
					"d": 0,
					"tf": 4301,
					"ttf": 4301,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.514,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI630": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI641": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "NAG-BOM-AI630": {
						"at": "Terminal 2"
					  },
					  "DEL-NAG-AI641": {}
					},
					"ottf": 4231.01,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7dJ1pG5cHoUA+M6/mVpxTQi8Zb7HUh2aehyketvUi0rpgyKbCNfYgwKnvBXMfui447Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI630": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI641": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-AI805": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.537,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI805": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-AI805": {}
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8VgBCmBvAa48lW6QQXu2eRNk9795mi/AFAxnl2QHB1lw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI805": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-LKO-G82509*LKO-BOM-G8307": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5016,
					"c": "INR",
					"bf": 3225,
					"t": 1791,
					"f": 0,
					"d": 0,
					"tf": 5016,
					"ttf": 5016,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.143,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G82509": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8307": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-LKO-G82509": {
						"dt": "Terminal 2"
					  },
					  "LKO-BOM-G8307": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 4912.54,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTlsVE0gkSQxCc4Gv6nKxN1T+VeKHC4DZkvLLuUiv5eMNmx83b+Yj5jF4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G82509": {
						"sr": 3,
						"ft": "GS"
					  },
					  "G8307": {
						"sr": 3,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BLR-UK817*BLR-BOM-UK866": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 8.731,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK866": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK817": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK817": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK866": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8ht0WsSpfGG5sTFVyTuq7km5uyOzaujphCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK866": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK817": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-G8346": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.571,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8346": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8346": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9CHYJEbhCGsslW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8346": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-JLR-SG3029*JLR-BOM-SG2871": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7001,
					"c": "INR",
					"bf": 6202,
					"t": 799,
					"f": 0,
					"d": 0,
					"tf": 7001,
					"ttf": 7001,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.734,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG2871": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG3029": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "JLR-BOM-SG2871": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-JLR-SG3029": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  }
					},
					"ottf": 6899.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "420"
					},
					"ind": 420,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz71gQeS9pQ7+i/dC87EVtiXhD69bHf5s/RwetS/U1bBy/JVukEF7tnkTjklBNSHd5S2sHLbn72JDI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG2871": {
						"sr": 1,
						"ft": "RS"
					  },
					  "SG3029": {
						"sr": 1,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-AMD-6E5072*AMD-BOM-6E6811": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.913,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6811": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E5072": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-AMD-6E5072": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  },
					  "AMD-BOM-6E6811": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H/iMiB1p2VcQwJW46tf7sneaOzOgsQTUtTJVukEF7tnkR6KAc/aQ7UGd7zwRW4jB9M=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6811": {
						"sr": 55,
						"ft": "R"
					  },
					  "6E5072": {
						"sr": 55,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-IDR-AI636*IDR-BOM-AI636": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 5.437,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI636": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "IDR-BOM-AI636": {
						"at": "Terminal 2"
					  },
					  "DEL-IDR-AI636": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8nYwz+0SzAmy5dUzNpmP/xW9ckmjYYWjGVgyKbCNfYgwCL6kX4ySgbu7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI636": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-UK941": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2413,
					"t": 616,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.683,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK941": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK941": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+77BHr69z1N2/EjpPW4k3YZooiaCy8or6SDks8/OROmA==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK941": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-6E5025": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2209,
					"t": 572,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.502,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5025": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5025": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2756.91,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8pMmQU9fbwuGl8UNbXcFHEHT+7NAjbv3J3B/cWMT8z7Q==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5025": {
						"sr": 61,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK813*BLR-BOM-UK846": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.602,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK846": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK813": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK846": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK813": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u91aWMPAGE5U5sTFVyTuq7kv2ybWuQ1xtBgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK846": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK813": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-UK943": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5871,
					"c": "INR",
					"bf": 5120,
					"t": 751,
					"f": 0,
					"d": 0,
					"tf": 5871,
					"ttf": 5871,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 4.423,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK943": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK943": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5691.2,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9HY2KHJCLsy8lW6QQXu2eR1SzKbjJBnhLctlopdfUqWw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK943": {
						"sr": 9,
						"ft": "Q"
					  }
					}
				  }
				],
				"DEL-BOM-UK945": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2413,
					"t": 616,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.75,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK945": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK945": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8m2Jd0WI1X6m/EjpPW4k3YZooiaCy8or6SDks8/OROmA==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK945": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-G8338": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.571,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8338": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8338": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9r+HAWa1l6LslW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8338": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BLR-UK817*BLR-BOM-UK858": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.402,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK858": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK817": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK817": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK858": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8ht0WsSpfGG5sTFVyTuq7kVx8ZFA3pK3VgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK858": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK817": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-HYD-UK899*HYD-BOM-UK876": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5548,
					"c": "INR",
					"bf": 4682,
					"t": 866,
					"f": 0,
					"d": 0,
					"tf": 5548,
					"ttf": 5548,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 18.728,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK876": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK899": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-UK899": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-UK876": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAbx1R4UESTye/6LuWTGDycszz5kLbi0kqdgyKbCNfYgwMHZIKrQKuHy",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK876": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK899": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-HYD-UK899*HYD-BOM-UK874": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5547,
					"c": "INR",
					"bf": 4682,
					"t": 865,
					"f": 0,
					"d": 0,
					"tf": 5547,
					"ttf": 5547,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.495,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK899": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK874": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-UK899": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-UK874": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5267.66,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAbx1R4UESTye/6LuWTGDycsKdyrZm9jJP5gyKbCNfYgwC9I5411vIcM7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK899": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK874": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-G8330": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.671,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8330": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8330": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9hLDS3bBezT8lW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8330": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BOM-G8336": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.538,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8336": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8336": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+KvSlkJ9JMjMlW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8336": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BLR-UK817*BLR-BOM-UK850": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 8.098,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK817": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK850": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK850": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK817": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8ht0WsSpfGG5sTFVyTuq7kwuDPPA3SCzRCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK817": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK850": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-JDH-AI475*JDH-BOM-AI646": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4718,
					"c": "INR",
					"bf": 3700,
					"t": 1018,
					"f": 0,
					"d": 0,
					"tf": 4718,
					"ttf": 4718,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.197,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI646": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI475": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "JDH-BOM-AI646": {
						"at": "Terminal 2"
					  },
					  "DEL-JDH-AI475": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4641.88,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz71gQeS9pQ7+hE3OSqp3UbVUhZ1jVYfsJ7OoYdp2ailnFgyKbCNfYgwDEIdjct/CEG7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI646": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI475": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-G8334": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.538,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8334": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8334": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9m6yZrNvykxMlW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8334": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-MAA-UK835*MAA-BOM-UK828": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 10.523,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK835": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK828": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK828": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK835": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEfO3gaE5xH4g1iqpJp4mySNa9Js+usI3rRCvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK835": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK828": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-AMD-6E6373*AMD-BOM-6E167": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.946,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6373": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E167": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-6E167": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "DEL-AMD-6E6373": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H/ouGTmoRtRvbJLfxynqjQWQ26aqc5peJbLLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6373": {
						"sr": 41,
						"ft": "R"
					  },
					  "6E167": {
						"sr": 41,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-GOI-UK847*GOI-BOM-UK844": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 8627,
					"c": "INR",
					"bf": 7615,
					"t": 1012,
					"f": 0,
					"d": 0,
					"tf": 8627,
					"ttf": 8627,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.61,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK844": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK847": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 349,
					  "tdf": 349
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-UK847": {
						"dt": "Terminal 3"
					  },
					  "GOI-BOM-UK844": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 8305.32,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "517"
					},
					"ind": 517,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/sQTTuJGT4JH7zGcMImpgPoYK3DzyvC7RxgyKbCNfYgwBk268+wICAp7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK844": {
						"sr": 9,
						"ft": "Q"
					  },
					  "UK847": {
						"sr": 9,
						"ft": "Q"
					  }
					}
				  }
				],
				"DEL-MAA-UK835*MAA-BOM-UK826": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 12.723,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK835": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK826": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK826": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK835": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEfO3gaE5xH4g1iqpJp4mySNFvQ+t4TcKT1CvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK835": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK826": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-AI542*HYD-BOM-AI698": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 20.215,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI698": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI542": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-AI698": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-AI542": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAY2oYSRm6wczf6LuWTGDycsE84/rZEc2glgyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI698": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI542": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-RAJ-SG2643*RAJ-BOM-SG224": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5848,
					"c": "INR",
					"bf": 4332,
					"t": 1516,
					"f": 0,
					"d": 0,
					"tf": 5848,
					"ttf": 5848,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.522,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG224": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG2643": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "RAJ-BOM-SG224": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-RAJ-SG2643": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "350"
					},
					"ind": 350,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7FZHq91/QjYzexQUHcYWyt4t3ZmTzHUHTzvClE3FGdSTLLuUiv5eMNjpnL/ZNHyuifZk/4Aoz45k=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG224": {
						"sr": 20,
						"ft": "RS"
					  },
					  "SG2643": {
						"sr": 20,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-HYD-AI839*HYD-BOM-AI698": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.649,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI839": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI698": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-AI698": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-AI839": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAb3yL5b0vnoqv6LuWTGDycsE84/rZEc2glgyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI839": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI698": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-GOI-AI883*GOI-BOM-AI684": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5881,
					"c": "INR",
					"bf": 4920,
					"t": 961,
					"f": 0,
					"d": 0,
					"tf": 5881,
					"ttf": 5881,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.329,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI883": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI684": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-AI883": {
						"dt": "Terminal 3"
					  },
					  "GOI-BOM-AI684": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5781.67,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/vtm8x+F5l+/7zGcMImpgPo3iB9UKV8e1RgyKbCNfYgwEYxfCjkivit7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI883": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI684": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-6E5379": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2209,
					"t": 572,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.535,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5379": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5379": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2756.91,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u86ZjGR1ZRD1mjB9wWmZOmGHT+7NAjbv3J3B/cWMT8z7Q==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5379": {
						"sr": 92,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-6E5012": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2209,
					"t": 572,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.535,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5012": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5012": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2756.91,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/gthIxpVfGNrfl1H2cOBbVHT+7NAjbv3J3B/cWMT8z7Q==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5012": {
						"sr": 127,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-6E5013": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.651,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5013": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5013": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/gthIxpVfGNgmFDeHdIIZ9eo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5013": {
						"sr": 81,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-UK933": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2413,
					"t": 616,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.683,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK933": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK933": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/y/L+FC2ohK2/EjpPW4k3YZooiaCy8or6SDks8/OROmA==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK933": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-LKO-G8268*LKO-BOM-G8396": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5016,
					"c": "INR",
					"bf": 3225,
					"t": 1791,
					"f": 0,
					"d": 0,
					"tf": 5016,
					"ttf": 5016,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.243,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8268": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8396": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-LKO-G8268": {
						"dt": "Terminal 2"
					  },
					  "LKO-BOM-G8396": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 4912.54,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTlzNiOPs4/rcuPf4zDMLBU0S/RYawEyDvdgyKbCNfYgwEr0HNzNE3OC7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8268": {
						"sr": 4,
						"ft": "GS"
					  },
					  "G8396": {
						"sr": 4,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-IXU-AI441*IXU-BOM-AI441": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2784,
					"c": "INR",
					"bf": 2140,
					"t": 644,
					"f": 0,
					"d": 0,
					"tf": 2784,
					"ttf": 2784,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 5.67,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI441": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "IXU-BOM-AI441": {
						"at": "Terminal 2"
					  },
					  "DEL-IXU-AI441": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 2739.71,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8koLvnJmpi4F9Vpe0vOQkaaDyznlOTmikZgyKbCNfYgwCL6kX4ySgbu7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI441": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-IDR-6E5051*IDR-BOM-6E232": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2980,
					"c": "INR",
					"bf": 2348,
					"t": 632,
					"f": 0,
					"d": 0,
					"tf": 2980,
					"ttf": 2980,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 5.921,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E232": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E5051": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-IDR-6E5051": {
						"dt": "Terminal 3"
					  },
					  "IDR-BOM-6E232": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 2954.18,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7Q3KaPknoY8mBgAXDt4khIhoxEOMuIe0QHH83KTL0KJvLLuUiv5eMNrMu6lOiEq/skymwHBX0SUo=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E232": {
						"sr": 2,
						"ft": "R"
					  },
					  "6E5051": {
						"sr": 2,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-JAI-6E736*JAI-BOM-6E6031": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.146,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E736": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E6031": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-JAI-6E736": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "JAI-BOM-6E6031": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz71gQeS9pQ7+hkwiCes/pTGnA18+8NqGTd9UIIkYOflk/LLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E736": {
						"sr": 12,
						"ft": "R"
					  },
					  "6E6031": {
						"sr": 12,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-MAA-UK835*MAA-BOM-UK824": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 15.957,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK824": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK835": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-MAA-UK835": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  },
					  "MAA-BOM-UK824": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEfO3gaE5xH4g1iqpJp4mySNN7JdAQRyhmNCvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK824": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK835": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-MAA-UK835*MAA-BOM-UK822": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.59,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK822": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK835": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-MAA-UK835": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  },
					  "MAA-BOM-UK822": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEfO3gaE5xH4g1iqpJp4mySNz4KJ7+tYKJtCvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK822": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK835": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-GOI-G8286*GOI-BOM-G83006": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 6722,
					"c": "INR",
					"bf": 4849,
					"t": 1873,
					"f": 0,
					"d": 0,
					"tf": 6722,
					"ttf": 6722,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.333,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83006": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8286": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 329,
					  "tdf": 329
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-G8286": {
						"dt": "Terminal 2"
					  },
					  "GOI-BOM-G83006": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 6578.28,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "403"
					},
					"ind": 403,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/sz/E2cLC88DrzGcMImpgPo3iTbD/cw6FnLLuUiv5eMNmT+MZFrzYcZSMMLUxrC9LM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83006": {
						"sr": 3,
						"ft": "GS"
					  },
					  "G8286": {
						"sr": 3,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BOM-UK927": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2413,
					"t": 616,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.683,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK927": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK927": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8afGfcPWSonm/EjpPW4k3YZooiaCy8or6SDks8/OROmA==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK927": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-UK871*HYD-BOM-UK874": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5548,
					"c": "INR",
					"bf": 4682,
					"t": 866,
					"f": 0,
					"d": 0,
					"tf": 5548,
					"ttf": 5548,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 11.162,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK871": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK874": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-UK871": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-UK874": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYv9gngZUaTNf6LuWTGDycsKdyrZm9jJP5gyKbCNfYgwMHZIKrQKuHy",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK871": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK874": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-AMD-SG8913*AMD-BOM-SG636": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3777,
					"c": "INR",
					"bf": 3132,
					"t": 645,
					"f": 0,
					"d": 0,
					"tf": 3777,
					"ttf": 3777,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.7,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG8913": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG636": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-SG636": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-AMD-SG8913": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3722.47,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H9NP7kcREzxCLJLfxynqjQWiXg7xJx3J3XLLuUiv5eMNsVBZRcd7tV+UKgAPjC/U6M=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG8913": {
						"sr": 5,
						"ft": "RS"
					  },
					  "SG636": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-AMD-6E153*AMD-BOM-6E6813": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.846,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6813": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E153": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-AMD-6E153": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "AMD-BOM-6E6813": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H+loufziIQodKopjOalXrMxZ96obKcuShDLLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6813": {
						"sr": 34,
						"ft": "R"
					  },
					  "6E153": {
						"sr": 34,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-JAI-AI491*JAI-BOM-AI612": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3760,
					"c": "INR",
					"bf": 2900,
					"t": 860,
					"f": 0,
					"d": 0,
					"tf": 3760,
					"ttf": 3760,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.289,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI491": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI612": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "JAI-BOM-AI612": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  },
					  "DEL-JAI-AI491": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3700.04,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz71gQeS9pQ7+hzougTxoDtinA18+8NqGTdaeTTGYLSjTRgyKbCNfYgwDFUPR6zK8WY7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI491": {
						"sr": 8,
						"ft": "S"
					  },
					  "AI612": {
						"sr": 8,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-GOI-G83019*GOI-BOM-G83008": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5599,
					"c": "INR",
					"bf": 3780,
					"t": 1819,
					"f": 0,
					"d": 0,
					"tf": 5599,
					"ttf": 5599,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.126,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83008": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G83019": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "GOI-BOM-G83008": {
						"at": "Terminal 1"
					  },
					  "DEL-GOI-G83019": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 5481.78,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "335"
					},
					"ind": 335,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/vzplVzoZHRFhzO8TfE8yKma4mD++0EhLzJVukEF7tnkcpx5eitp0nyuU/oGjz57/U=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83008": {
						"sr": 6,
						"ft": "GS"
					  },
					  "G83019": {
						"sr": 6,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-GOI-G8286*GOI-BOM-G82606": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 6722,
					"c": "INR",
					"bf": 4849,
					"t": 1873,
					"f": 0,
					"d": 0,
					"tf": 6722,
					"ttf": 6722,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.8,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8286": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G82606": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 329,
					  "tdf": 329
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "GOI-BOM-G82606": {
						"at": "Terminal 1"
					  },
					  "DEL-GOI-G8286": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 6578.28,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "403"
					},
					"ind": 403,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/sz/E2cLC88DrzGcMImpgPotxj5xFuAP5vLLuUiv5eMNmT+MZFrzYcZSMMLUxrC9LM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8286": {
						"sr": 3,
						"ft": "GS"
					  },
					  "G82606": {
						"sr": 3,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BHO-AI481*BHO-BOM-AI632": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4099,
					"c": "INR",
					"bf": 3110,
					"t": 989,
					"f": 0,
					"d": 0,
					"tf": 4099,
					"ttf": 4099,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.859,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI481": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI632": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BHO-BOM-AI632": {
						"at": "Terminal 2"
					  },
					  "DEL-BHO-AI481": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4034.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/gYhDFONC/nRPtf0yVbKE2gmpFuODWwzBgyKbCNfYgwCUA1RkKDhyH7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI481": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI632": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-JAI-AI9643*JAI-BOM-AI612": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3718,
					"c": "INR",
					"bf": 2780,
					"t": 938,
					"f": 0,
					"d": 0,
					"tf": 3718,
					"ttf": 3718,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.864,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI612": {
						"hb": "5 Kg",
						"cb": "15 Kilograms"
					  },
					  "AI9643": {
						"hb": "5 Kg",
						"cb": "15 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "JAI-BOM-AI612": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  },
					  "DEL-JAI-AI9643": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3659.97,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz71gQeS9pQ7+jJFn8emNzTxZNfvMiApCzuXnDGeaIDiU7LLuUiv5eMNnQtP/2EdxaG4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI612": {
						"sr": 5,
						"ft": "S"
					  },
					  "AI9643": {
						"sr": 5,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BDQ-6E6245*BDQ-BOM-6E628": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.313,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E628": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E6245": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BDQ-BOM-6E628": {
						"at": "Terminal 1"
					  },
					  "DEL-BDQ-6E6245": {
						"dt": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8LIAMO+yJUJZxIbOlTKuYstdLwuMgx803LLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E628": {
						"sr": 28,
						"ft": "R"
					  },
					  "6E6245": {
						"sr": 28,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BHO-AI481*BHO-BOM-AI634": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4099,
					"c": "INR",
					"bf": 3110,
					"t": 989,
					"f": 0,
					"d": 0,
					"tf": 4099,
					"ttf": 4099,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 14.959,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI481": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI634": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BHO-BOM-AI634": {
						"at": "Terminal 2"
					  },
					  "DEL-BHO-AI481": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4034.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/gYhDFONC/nRPtf0yVbKE2gL8IlmZln31gyKbCNfYgwCUA1RkKDhyH7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI481": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI634": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-G82501": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.571,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G82501": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G82501": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/9OW2jWCGO+SA+oE54nzXGN2XhgojdrJBteuucViHh6w==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G82501": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-HYD-UK871*HYD-BOM-UK876": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5547,
					"c": "INR",
					"bf": 4682,
					"t": 865,
					"f": 0,
					"d": 0,
					"tf": 5547,
					"ttf": 5547,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.395,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK876": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK871": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-UK871": {
						"dt": "Terminal 3"
					  },
					  "HYD-BOM-UK876": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5267.66,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYv9gngZUaTNf6LuWTGDycszz5kLbi0kqdgyKbCNfYgwC9I5411vIcM7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK876": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK871": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BLR-UK807*BLR-BOM-UK858": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 10.431,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK858": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK807": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK807": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK858": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8Guhn6fVS1UJsTFVyTuq7kVx8ZFA3pK3VCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK858": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK807": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-VNS-G8404*VNS-BOM-G8350": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4973,
					"c": "INR",
					"bf": 3184,
					"t": 1789,
					"f": 0,
					"d": 0,
					"tf": 4973,
					"ttf": 4973,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.55,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8404": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8350": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-VNS-G8404": {
						"dt": "Terminal 2"
					  },
					  "VNS-BOM-G8350": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 4870.56,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7IvjHpZPJB56fB+z2Ol8ioHmmLMgGMMwyMzHpKMrIyOtgyKbCNfYgwMI8NSZA6+5p7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8404": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G8350": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-VNS-AI406*VNS-BOM-AI696": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 6144,
					"c": "INR",
					"bf": 5170,
					"t": 974,
					"f": 0,
					"d": 0,
					"tf": 6144,
					"ttf": 6144,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.353,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI406": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI696": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 329,
					  "tdf": 329
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-VNS-AI406": {
						"dt": "Terminal 3"
					  },
					  "VNS-BOM-AI696": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 6039.79,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "368"
					},
					"ind": 368,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7IvjHpZPJB57kUuoMCqi95XmmLMgGMMwyj7QcMaJuRYBgyKbCNfYgwMRZC2sTzz4H7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI406": {
						"sr": 8,
						"ft": "T"
					  },
					  "AI696": {
						"sr": 8,
						"ft": "T"
					  }
					}
				  }
				],
				"DEL-HYD-6E967*HYD-BOM-6E884": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.213,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E967": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E884": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-6E884": {
						"at": "Terminal 1"
					  },
					  "DEL-HYD-6E967": {
						"dt": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAaVmKX/7SLmhP6LuWTGDycsCXNZgX9c8DtgyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E967": {
						"sr": 61,
						"ft": "R"
					  },
					  "6E884": {
						"sr": 61,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-6E5071": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.651,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5071": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5071": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8g0WFl+xNagiA+oE54nzXGeo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5071": {
						"sr": 125,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BBI-I5814*BBI-BOM-I5630": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5593,
					"c": "INR",
					"bf": 4910,
					"t": 683,
					"f": 0,
					"d": 0,
					"tf": 5593,
					"ttf": 5593,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.822,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5814": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5630": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BBI-I5814": {
						"dt": "Terminal 3"
					  },
					  "BBI-BOM-I5630": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 5539.17,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "335"
					},
					"ind": 335,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8RuQgc8dgyQfHoC2YW/qXiNY61Pk2vctBgyKbCNfYgwJQz7DQCguWe7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5814": {
						"sr": 50,
						"ft": "EP"
					  },
					  "I5630": {
						"sr": 50,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BOM-UK995": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2413,
					"t": 616,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.717,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK995": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK995": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+BQ7c5z6LU42/EjpPW4k3YZooiaCy8or6SDks8/OROmA==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK995": {
						"sr": 1,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-NAG-6E774*NAG-BOM-6E297": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.08,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E297": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E774": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "NAG-BOM-6E297": {
						"at": "Terminal 1"
					  },
					  "DEL-NAG-6E774": {
						"dt": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7dJ1pG5cHoUA4UMPbcEEe9S8Zb7HUh2aepvX0Ox8ft2pgyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E297": {
						"sr": 35,
						"ft": "R"
					  },
					  "6E774": {
						"sr": 35,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-MAA-UK833*MAA-BOM-UK828": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4650,
					"c": "INR",
					"bf": 3828,
					"t": 822,
					"f": 0,
					"d": 0,
					"tf": 4650,
					"ttf": 4650,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 15.557,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK833": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK828": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK828": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK833": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 4382.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEcKOFjDfREWiFiqpJp4mySNa9Js+usI3rRCvZmJPFlT1MvWl/VXGHvv4Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK833": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK828": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BLR-UK807*BLR-BOM-UK854": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.369,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK854": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK807": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK854": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK807": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8Guhn6fVS1UJsTFVyTuq7khnsBGurXwalgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK854": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK807": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-ATQ-6E322*ATQ-BOM-6E448": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.88,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E448": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E322": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "ATQ-BOM-6E448": {
						"at": "Terminal 1"
					  },
					  "DEL-ATQ-6E322": {
						"dt": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H876Ld3aZpuamrEfpCptsgJSOHG+YPyYo5gyKbCNfYgwKGwr1ZPqeUb7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E448": {
						"sr": 6,
						"ft": "R"
					  },
					  "6E322": {
						"sr": 6,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK807*BLR-BOM-UK852": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.698,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK807": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK852": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK852": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK807": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8Guhn6fVS1UJsTFVyTuq7kH5vcgAQjXoVCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK807": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK852": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-MAA-UK833*MAA-BOM-UK824": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4651,
					"c": "INR",
					"bf": 3828,
					"t": 823,
					"f": 0,
					"d": 0,
					"tf": 4651,
					"ttf": 4651,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 20.99,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK833": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK824": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK824": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK833": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEcKOFjDfREWiFiqpJp4mySNN7JdAQRyhmNCvZmJPFlT1G+zSvoRZT/wY2zK8M5MquM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK833": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK824": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-SG8701": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2783,
					"c": "INR",
					"bf": 2184,
					"t": 599,
					"f": 0,
					"d": 0,
					"tf": 2783,
					"ttf": 2783,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.57,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG8701": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-SG8701": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2742.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8GnOC/hQhS4iA+oE54nzXGfQ9YZpteyi4DIMrzWa/w+g==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG8701": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-MAA-UK833*MAA-BOM-UK826": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4651,
					"c": "INR",
					"bf": 3828,
					"t": 823,
					"f": 0,
					"d": 0,
					"tf": 4651,
					"ttf": 4651,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 17.757,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK833": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK826": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "MAA-BOM-UK826": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-MAA-UK833": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEcKOFjDfREWiFiqpJp4mySNFvQ+t4TcKT1CvZmJPFlT1G+zSvoRZT/wY2zK8M5MquM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK833": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK826": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-LKO-6E6612*LKO-BOM-6E362": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.58,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E362": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E6612": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "LKO-BOM-6E362": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  },
					  "DEL-LKO-6E6612": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTlzgaxTzqdPC0hPVvgkSwjnuYxn0XnHiz3LLuUiv5eMNmUsXvl4guuVu+J81r6y5e8=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E362": {
						"sr": 104,
						"ft": "R"
					  },
					  "6E6612": {
						"sr": 104,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-PAT-G82511*PAT-BOM-G8352": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5674,
					"c": "INR",
					"bf": 3851,
					"t": 1823,
					"f": 0,
					"d": 0,
					"tf": 5674,
					"ttf": 5674,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.271,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G82511": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8352": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "PAT-BOM-G8352": {
						"at": "Terminal 1"
					  },
					  "DEL-PAT-G82511": {
						"dt": "Terminal 2"
					  }
					},
					"ottf": 5555.02,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "340"
					},
					"ind": 340,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7N/1NW0HAOR7+rcaiwGXUnzvmPu/WHK9O1ngGYiQl99nLLuUiv5eMNl7UmCN3EyKh7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G82511": {
						"sr": 7,
						"ft": "GS"
					  },
					  "G8352": {
						"sr": 7,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BOM-UK993": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3229,
					"c": "INR",
					"bf": 2413,
					"t": 816,
					"f": 0,
					"d": 0,
					"tf": 3229,
					"ttf": 3229,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.683,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK993": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK993": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+Bq+E0/C1ALW/EjpPW4k3Y7s1ESNqmb62SDks8/OROmA==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK993": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-SG8709": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2783,
					"c": "INR",
					"bf": 2184,
					"t": 599,
					"f": 0,
					"d": 0,
					"tf": 2783,
					"ttf": 2783,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.603,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG8709": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-SG8709": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2742.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8GnOC/hQhS4mjB9wWmZOmGfQ9YZpteyi4DIMrzWa/w+g==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG8709": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-UDR-AI471*UDR-BOM-AI644": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4075,
					"c": "INR",
					"bf": 3200,
					"t": 875,
					"f": 0,
					"d": 0,
					"tf": 4075,
					"ttf": 4075,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.745,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI471": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI644": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "UDR-BOM-AI644": {
						"at": "Terminal 2"
					  },
					  "DEL-UDR-AI471": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4009.18,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7bC7vxRRnG9dB5b0E2/IzO7xQJMDNGQItLIVPIqt/AClgyKbCNfYgwADUBgoaUglW7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI471": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI644": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-I5740*BLR-BOM-I5306": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5157,
					"c": "INR",
					"bf": 4496,
					"t": 661,
					"f": 0,
					"d": 0,
					"tf": 5157,
					"ttf": 5157,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.228,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5306": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5740": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-I5740": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-I5306": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 5107.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "309"
					},
					"ind": 309,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+Bk+BXv5M525sTFVyTuq7kOynDTtjbNf1gyKbCNfYgwG8n7NK6oXLi7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5306": {
						"sr": 79,
						"ft": "EP"
					  },
					  "I5740": {
						"sr": 79,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-UK807*BLR-BOM-UK846": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.231,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK846": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK807": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK807": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK846": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8Guhn6fVS1UJsTFVyTuq7kv2ybWuQ1xtBCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK846": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK807": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-AI839*HYD-BOM-AI620": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.115,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI839": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI620": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-AI620": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-AI839": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAb3yL5b0vnoqv6LuWTGDycseQmyYSyE3wNgyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI839": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI620": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-CCU-UK727*CCU-BOM-UK776": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.563,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK776": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK727": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK776": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK727": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQrqnAmczQeYNvT4z+KegT3li9E9THdgGzJgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK776": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK727": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-CCU-UK727*CCU-BOM-UK772": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7549,
					"c": "INR",
					"bf": 6589,
					"t": 960,
					"f": 0,
					"d": 0,
					"tf": 7549,
					"ttf": 7549,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.629,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK727": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK772": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "CCU-BOM-UK772": {
						"at": "Terminal 2"
					  },
					  "DEL-CCU-UK727": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7242.12,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "452"
					},
					"ind": 452,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7+80J8uqBKQrqnAmczQeYNvT4z+KegT3lz8sYXzW1wVNgyKbCNfYgwOCnE9SbpvNQ7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK727": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK772": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-LKO-AI811*LKO-BOM-AI626": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5199,
					"c": "INR",
					"bf": 4270,
					"t": 929,
					"f": 0,
					"d": 0,
					"tf": 5199,
					"ttf": 5199,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.186,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI811": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI626": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-LKO-AI811": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  },
					  "LKO-BOM-AI626": {
						"dt": "Terminal 2",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5112.33,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "311"
					},
					"ind": 311,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7gEVNeXCMgTniuRnVRYtrDePf4zDMLBU0B4n4/jZYjBlgyKbCNfYgwKaz0vDNDl4/7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI811": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI626": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-UK985": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5871,
					"c": "INR",
					"bf": 5120,
					"t": 751,
					"f": 0,
					"d": 0,
					"tf": 5871,
					"ttf": 5871,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 4.389,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK985": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK985": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5691.2,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/mYKW2ZBcynslW6QQXu2eR1SzKbjJBnhLctlopdfUqWw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK985": {
						"sr": 9,
						"ft": "Q"
					  }
					}
				  }
				],
				"DEL-BLR-UK809*BLR-BOM-UK864": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.902,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK864": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK809": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK809": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK864": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9EGN9c6Ahtu5sTFVyTuq7kSqy2XPVrgwNgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK864": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK809": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-GOI-I5775*GOI-BOM-I5472": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 6068,
					"c": "INR",
					"bf": 5363,
					"t": 705,
					"f": 0,
					"d": 0,
					"tf": 6068,
					"ttf": 6068,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 11.807,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5775": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5472": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-I5775": {
						"dt": "Terminal 3"
					  },
					  "GOI-BOM-I5472": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 6009.6,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "364"
					},
					"ind": 364,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/urZ6VM8xX0lLzGcMImpgPorGjs/TVXScxgyKbCNfYgwIAySJgUmXAl7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5775": {
						"sr": 4,
						"ft": "EP"
					  },
					  "I5472": {
						"sr": 4,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-AI504*BLR-BOM-AI604": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5241,
					"c": "INR",
					"bf": 4310,
					"t": 931,
					"f": 0,
					"d": 0,
					"tf": 5241,
					"ttf": 5241,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.978,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI504": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI604": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-AI604": {},
					  "DEL-BLR-AI504": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5153.56,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "314"
					},
					"ind": 314,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9ebgF+4X23g5sTFVyTuq7kgML3cW0GJd9gyKbCNfYgwP8PjH7aMk617Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI504": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI604": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-SXR-G83009*SXR-BOM-G83002": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 9561,
					"c": "INR",
					"bf": 7554,
					"t": 2007,
					"f": 0,
					"d": 0,
					"tf": 9561,
					"ttf": 9561,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.037,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83009": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G83002": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 359,
					  "tdf": 359
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-SXR-G83009": {
						"dt": "Terminal 2"
					  },
					  "SXR-BOM-G83002": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 9350.25,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "573"
					},
					"ind": 573,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy5bCiT7L1Bhw2FcgEEVWGMwlG2Ch1+BERbJVukEF7tnkZ72t3bKBDNgsqWo29Ar9sY=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83009": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G83002": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-BOM-SG8169": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2783,
					"c": "INR",
					"bf": 2184,
					"t": 599,
					"f": 0,
					"d": 0,
					"tf": 2783,
					"ttf": 2783,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.603,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG8169": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-SG8169": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2742.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8mM6s6WG5V72jB9wWmZOmGfQ9YZpteyi4DIMrzWa/w+g==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG8169": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-BOM-I5332": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2232,
					"t": 549,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.535,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5332": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-I5332": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2754.23,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8pBu5uxMxcyslW6QQXu2eRFpoT8PVfX8ls/MBZ2wJYrw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5332": {
						"sr": 119,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BOM-UK981": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2929,
					"c": "INR",
					"bf": 2318,
					"t": 611,
					"f": 0,
					"d": 0,
					"tf": 2929,
					"ttf": 2929,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.623,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK981": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK981": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+RXx4OYnMAWm/EjpPW4k3YK41fjj6brgiDcOmfzf0rDg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK981": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-6E665": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.684,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E665": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E665": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/NPUkyBsoJTslW6QQXu2eRdRi7zbWSMU9th+ZEoUooTw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E665": {
						"sr": 111,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-G8530": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.504,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8530": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8530": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+2X1CGLHwk6MlW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8530": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-SXR-G8191*SXR-BOM-G8287": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 11630,
					"c": "INR",
					"bf": 9525,
					"t": 2105,
					"f": 0,
					"d": 0,
					"tf": 11630,
					"ttf": 11630,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.078,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8287": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8191": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 369,
					  "tdf": 369
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-SXR-G8191": {
						"dt": "Terminal 2"
					  },
					  "SXR-BOM-G8287": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 11370.42,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "697"
					},
					"ind": 697,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy653KNeo6uAKYy5NkBwJfYSTwkpGOns5L9gyKbCNfYgwNo0RAzduSwM",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8287": {
						"sr": 2,
						"ft": "GS"
					  },
					  "G8191": {
						"sr": 2,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-VNS-UK812*VNS-BOM-UK622": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 8444,
					"c": "INR",
					"bf": 7441,
					"t": 1003,
					"f": 0,
					"d": 0,
					"tf": 8444,
					"ttf": 8444,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 19.166,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK622": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK812": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 349,
					  "tdf": 349
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-VNS-UK812": {
						"dt": "Terminal 3"
					  },
					  "VNS-BOM-UK622": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 8124.82,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "506"
					},
					"ind": 506,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7IvjHpZPJB570B+qKzKSCh3mmLMgGMMwyu496i0dqdgBgyKbCNfYgwCJ5WY9hZCnG7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK622": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK812": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-AMD-6E5006*AMD-BOM-6E6813": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.146,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E6813": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E5006": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-6E6813": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "DEL-AMD-6E5006": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H/2RRzAFzhodX91GN6oMdZ2UhnFFrwBePLJVukEF7tnkR6KAc/aQ7UGd7zwRW4jB9M=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E6813": {
						"sr": 47,
						"ft": "R"
					  },
					  "6E5006": {
						"sr": 47,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK809*BLR-BOM-UK852": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.169,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK809": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK852": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK809": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK852": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9EGN9c6Ahtu5sTFVyTuq7kH5vcgAQjXoVgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK809": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK852": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BLR-UK809*BLR-BOM-UK854": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.702,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK854": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK809": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK809": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK854": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9EGN9c6Ahtu5sTFVyTuq7khnsBGurXwalgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK854": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK809": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BHO-AI437*BHO-BOM-AI634": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4099,
					"c": "INR",
					"bf": 3110,
					"t": 989,
					"f": 0,
					"d": 0,
					"tf": 4099,
					"ttf": 4099,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.426,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI437": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI634": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BHO-BOM-AI634": {
						"at": "Terminal 2"
					  },
					  "DEL-BHO-AI437": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4034.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/KS2YcLHgXbRPtf0yVbKE2gL8IlmZln31gyKbCNfYgwCUA1RkKDhyH7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI437": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI634": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-HYD-UK829*HYD-BOM-UK874": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5547,
					"c": "INR",
					"bf": 4682,
					"t": 865,
					"f": 0,
					"d": 0,
					"tf": 5547,
					"ttf": 5547,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 16.562,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK829": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK874": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-UK874": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-UK829": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5267.66,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAZHxilb7bW4Pf6LuWTGDycsKdyrZm9jJP5gyKbCNfYgwC9I5411vIcM7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK829": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK874": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-UK975": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2929,
					"c": "INR",
					"bf": 2318,
					"t": 611,
					"f": 0,
					"d": 0,
					"tf": 2929,
					"ttf": 2929,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 2.557,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK975": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK975": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+wjChWvpwuwm/EjpPW4k3YK41fjj6brgiDcOmfzf0rDg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK975": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-6E2187*HYD-BOM-6E5325": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.98,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E2187": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E5325": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-HYD-6E2187": {
						"dt": "Terminal 2"
					  },
					  "HYD-BOM-6E5325": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYMmwH9E+2z4yQxYSVFDqNDljOS8IztIiXJVukEF7tnkR6KAc/aQ7UGd7zwRW4jB9M=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E2187": {
						"sr": 72,
						"ft": "R"
					  },
					  "6E5325": {
						"sr": 72,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-UK809*BLR-BOM-UK858": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 10.764,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK858": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK809": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK809": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK858": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9EGN9c6Ahtu5sTFVyTuq7kVx8ZFA3pK3VCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK858": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK809": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-HYD-UK829*HYD-BOM-UK876": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5548,
					"c": "INR",
					"bf": 4682,
					"t": 866,
					"f": 0,
					"d": 0,
					"tf": 5548,
					"ttf": 5548,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 21.795,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK876": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK829": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-UK876": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-UK829": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 0,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "332"
					},
					"ind": 332,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAZHxilb7bW4Pf6LuWTGDycszz5kLbi0kqdgyKbCNfYgwMHZIKrQKuHy",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK876": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK829": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BOM-UK977": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5871,
					"c": "INR",
					"bf": 5120,
					"t": 751,
					"f": 0,
					"d": 0,
					"tf": 5871,
					"ttf": 5871,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 4.456,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK977": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK977": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5691.2,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8FWW/LZE7YYMlW6QQXu2eR1SzKbjJBnhLctlopdfUqWw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK977": {
						"sr": 9,
						"ft": "Q"
					  }
					}
				  }
				],
				"DEL-HYD-6E6205*HYD-BOM-6E5312": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3466,
					"c": "INR",
					"bf": 2811,
					"t": 655,
					"f": 0,
					"d": 0,
					"tf": 3466,
					"ttf": 3466,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 6.88,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5312": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "6E6205": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-6E5312": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-6E6205": {
						"dt": "Terminal 1"
					  }
					},
					"ottf": 3435.98,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYbropOAsu+PngkrJOm2h4jQA9Ad62qIgTJVukEF7tnkR6KAc/aQ7UGd7zwRW4jB9M=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5312": {
						"sr": 61,
						"ft": "R"
					  },
					  "6E6205": {
						"sr": 61,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BOM-SG8723": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2783,
					"c": "INR",
					"bf": 2184,
					"t": 599,
					"f": 0,
					"d": 0,
					"tf": 2783,
					"ttf": 2783,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.536,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG8723": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-SG8723": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 2742.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+Ze0wusKd6xQmFDeHdIIZ9fQ9YZpteyi4DIMrzWa/w+g==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG8723": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-BOM-6E993": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2209,
					"t": 572,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.535,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E993": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E993": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2756.91,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+/T1JrMBJASslW6QQXu2eRcVdwf3liqxyCA11aWgRf+A==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E993": {
						"sr": 85,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-BLR-I5721*BLR-BOM-I5941": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5157,
					"c": "INR",
					"bf": 4496,
					"t": 661,
					"f": 0,
					"d": 0,
					"tf": 5157,
					"ttf": 5157,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.561,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5941": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5721": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-I5721": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-I5941": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 5107.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "309"
					},
					"ind": 309,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/a78XZytfTR5sTFVyTuq7kKhSgY5swJYtgyKbCNfYgwG8n7NK6oXLi7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5941": {
						"sr": 94,
						"ft": "EP"
					  },
					  "I5721": {
						"sr": 94,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-I52879*BLR-BOM-I5306": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5157,
					"c": "INR",
					"bf": 4496,
					"t": 661,
					"f": 0,
					"d": 0,
					"tf": 5157,
					"ttf": 5157,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.328,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5306": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I52879": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-I5306": {
						"at": "Terminal 1"
					  },
					  "DEL-BLR-I52879": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 5107.37,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "309"
					},
					"ind": 309,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/p6SWHhhomDNDq53+K/2Ad1cNoQ/HOKszLLuUiv5eMNpaPr8hJSxX9Y2zK8M5MquM=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5306": {
						"sr": 54,
						"ft": "EP"
					  },
					  "I52879": {
						"sr": 54,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BOM-G8323": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2785,
					"c": "INR",
					"bf": 1450,
					"t": 1335,
					"f": 0,
					"d": 0,
					"tf": 2785,
					"ttf": 2785,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.571,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8323": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-G8323": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2730.83,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u97emYPuh9rI8lW6QQXu2eR3TdCjTtWKUl67C4ISsJfgw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8323": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-GOI-I5773*GOI-BOM-I5472": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5525,
					"c": "INR",
					"bf": 4846,
					"t": 679,
					"f": 0,
					"d": 0,
					"tf": 5525,
					"ttf": 5525,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 8.348,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5773": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  },
					  "I5472": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-GOI-I5773": {
						"dt": "Terminal 3"
					  },
					  "GOI-BOM-I5472": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 5471.82,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "331"
					},
					"ind": 331,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7grq7QexN+/vRGFdfKFVdVLzGcMImpgPorGjs/TVXScxgyKbCNfYgwJ+O6jIDNr2+7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5773": {
						"sr": 34,
						"ft": "EP"
					  },
					  "I5472": {
						"sr": 34,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BOM-6E5162": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3029,
					"c": "INR",
					"bf": 2445,
					"t": 584,
					"f": 0,
					"d": 0,
					"tf": 3029,
					"ttf": 3029,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.651,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "6E5162": {
						"hb": "01 Small Handbag under the seat",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-6E5162": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3002.76,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+Ze4jNOFaMGrfl1H2cOBbVeo07cdsZosqSpwdLVV3lJg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "6E5162": {
						"sr": 118,
						"ft": "R"
					  }
					}
				  }
				],
				"DEL-SXR-G8357*SXR-BOM-G83002": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 12918,
					"c": "INR",
					"bf": 10751,
					"t": 2167,
					"f": 0,
					"d": 0,
					"tf": 12918,
					"ttf": 12918,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.851,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G83002": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8357": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 369,
					  "tdf": 369
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-SXR-G8357": {
						"dt": "Terminal 2"
					  },
					  "SXR-BOM-G83002": {
						"at": "Terminal 1"
					  }
					},
					"ottf": 12628.02,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "775"
					},
					"ind": 775,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy5ix9fD+EvDCYy5NkBwJfYS0uWq3bpB7vvLLuUiv5eMNh/X5hWw4Tx84Rw15UG4XzI=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G83002": {
						"sr": 4,
						"ft": "GS"
					  },
					  "G8357": {
						"sr": 4,
						"ft": "GS"
					  }
					}
				  }
				],
				"DEL-MAA-AI439*MAA-BOM-AI569": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4028,
					"c": "INR",
					"bf": 3155,
					"t": 873,
					"f": 0,
					"d": 0,
					"tf": 4028,
					"ttf": 4028,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 15.35,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI569": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI439": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-MAA-AI439": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  },
					  "MAA-BOM-AI569": {
						"dt": "Terminal 4",
						"at": "Terminal 2"
					  }
					},
					"ottf": 3963.07,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7L44e4zIhmEcdA+v9Imoo11iqpJp4mySNg0tM8AbFpU9gyKbCNfYgwHP+b1KOy8lu7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI569": {
						"sr": 6,
						"ft": "S"
					  },
					  "AI439": {
						"sr": 6,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-UK963": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5871,
					"c": "INR",
					"bf": 5120,
					"t": 751,
					"f": 0,
					"d": 0,
					"tf": 5871,
					"ttf": 5871,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 4.456,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK963": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-UK963": {
						"dt": "Terminal 3",
						"at": "Terminal 2"
					  }
					},
					"ottf": 5691.2,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "352"
					},
					"ind": 352,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9qvCNuK9srtclW6QQXu2eR1SzKbjJBnhLctlopdfUqWw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK963": {
						"sr": 9,
						"ft": "Q"
					  }
					}
				  }
				],
				"DEL-NAG-AI641*NAG-BOM-AI628": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4301,
					"c": "INR",
					"bf": 3415,
					"t": 886,
					"f": 0,
					"d": 0,
					"tf": 4301,
					"ttf": 4301,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 9.014,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI628": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI641": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "NAG-BOM-AI628": {
						"at": "Terminal 2"
					  },
					  "DEL-NAG-AI641": {}
					},
					"ottf": 4231.01,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7dJ1pG5cHoUA+M6/mVpxTQi8Zb7HUh2ae7HpsK7SbEv1gyKbCNfYgwKnvBXMfui447Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI628": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI641": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BLR-UK807*BLR-BOM-UK864": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 7392,
					"c": "INR",
					"bf": 6439,
					"t": 953,
					"f": 0,
					"d": 0,
					"tf": 7392,
					"ttf": 7392,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 13.569,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK864": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "UK807": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 339,
					  "tdf": 339
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "BLR-BOM-UK864": {
						"at": "Terminal 2"
					  },
					  "DEL-BLR-UK807": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 7087.3,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "443"
					},
					"ind": 443,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u8Guhn6fVS1UJsTFVyTuq7kSqy2XPVrgwNgyKbCNfYgwE0/tyeuPU/q7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK864": {
						"sr": 9,
						"ft": "V"
					  },
					  "UK807": {
						"sr": 9,
						"ft": "V"
					  }
					}
				  }
				],
				"DEL-BLR-UK809*BLR-BOM-UK846": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5496,
					"c": "INR",
					"bf": 4633,
					"t": 863,
					"f": 0,
					"d": 0,
					"tf": 5496,
					"ttf": 5496,
					"ibf": 0,
					"obf": 0,
					"rt": "NONREFUNDABLE",
					"r": 11.564,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "UK846": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  },
					  "UK809": {
						"hb": "7 Kg",
						"cb": "01 Bag of 15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-UK809": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-UK846": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5217.36,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "329"
					},
					"ind": 329,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9EGN9c6Ahtu5sTFVyTuq7kv2ybWuQ1xtBCvZmJPFlT1FaGvrU+sYN27Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "UK846": {
						"sr": 9,
						"ft": "O"
					  },
					  "UK809": {
						"sr": 9,
						"ft": "O"
					  }
					}
				  }
				],
				"DEL-BOM-SG2871": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3028,
					"c": "INR",
					"bf": 2419,
					"t": 609,
					"f": 0,
					"d": 0,
					"tf": 3028,
					"ttf": 3028,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 3.55,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG2871": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-SG2871": {
						"dt": "Terminal 1D",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2765.49,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u+6zDlg0V1LFiA+oE54nzXGFSMAiBaMgSyNSllgXRj5rg==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG2871": {
						"sr": 3,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-STV-SG8483*STV-BOM-SG2773": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5887,
					"c": "INR",
					"bf": 5140,
					"t": 747,
					"f": 0,
					"d": 0,
					"tf": 5887,
					"ttf": 5887,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 12.766,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "SG2773": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "SG8483": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "STV-BOM-SG2773": {
						"dt": "Terminal 1",
						"at": "Terminal 2"
					  },
					  "DEL-STV-SG8483": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 5774.41,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "353"
					},
					"ind": 353,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7xuAcIi48Uy61glCltJhoihG/xTdFRCIwHq60WJyxi47JVukEF7tnkarHua5pQMMYN4bk1SCBReU=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "SG2773": {
						"sr": 5,
						"ft": "RS"
					  },
					  "SG8483": {
						"sr": 5,
						"ft": "RS"
					  }
					}
				  }
				],
				"DEL-HYD-AI560*HYD-BOM-AI698": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 4359,
					"c": "INR",
					"bf": 3470,
					"t": 889,
					"f": 0,
					"d": 0,
					"tf": 4359,
					"ttf": 4359,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 20.782,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI560": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI698": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "HYD-BOM-AI698": {
						"at": "Terminal 2"
					  },
					  "DEL-HYD-AI560": {
						"dt": "Terminal 3"
					  }
					},
					"ottf": 4287.92,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7MQu9dHBrfAYwToEmxkKo9/6LuWTGDycsE84/rZEc2glgyKbCNfYgwJHswxlMGL4V7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI560": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI698": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-BOM-I5314": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 2781,
					"c": "INR",
					"bf": 2232,
					"t": 549,
					"f": 0,
					"d": 0,
					"tf": 2781,
					"ttf": 2781,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 2.535,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "I5314": {
						"hb": "7 Kg",
						"cb": "15 Kg (01 Piece only)"
					  }
					},
					"cvf": {
					  "df": 350,
					  "tdf": 350
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BOM-I5314": {
						"dt": "Terminal 3",
						"at": "Terminal 1"
					  }
					},
					"ottf": 2754.23,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u9hmkkqtSGUDclW6QQXu2eRFpoT8PVfX8ls/MBZ2wJYrw==",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "I5314": {
						"sr": 136,
						"ft": "EP"
					  }
					}
				  }
				],
				"DEL-BLR-AI506*BLR-BOM-AI610": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 5241,
					"c": "INR",
					"bf": 4310,
					"t": 931,
					"f": 0,
					"d": 0,
					"tf": 5241,
					"ttf": 5241,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 10.178,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "AI506": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  },
					  "AI610": {
						"hb": "7 Kg",
						"cb": "25 Kilograms"
					  }
					},
					"cvf": {
					  "df": 319,
					  "tdf": 319
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "DEL-BLR-AI506": {
						"dt": "Terminal 3"
					  },
					  "BLR-BOM-AI610": {
						"at": "Terminal 2"
					  }
					},
					"ottf": 5153.56,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "314"
					},
					"ind": 314,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7map9ab1i3u/YEg6tctVmk5sTFVyTuq7kicuW7f84gbdgyKbCNfYgwP8PjH7aMk617Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "AI506": {
						"sr": 9,
						"ft": "S"
					  },
					  "AI610": {
						"sr": 9,
						"ft": "S"
					  }
					}
				  }
				],
				"DEL-AMD-G8713*AMD-BOM-G8413": [
				  {
					"providerId": 1044,
					"markup": 0,
					"valid": false,
					"total": 3622,
					"c": "INR",
					"bf": 2297,
					"t": 1325,
					"f": 0,
					"d": 0,
					"tf": 3622,
					"ttf": 3622,
					"ibf": 0,
					"obf": 0,
					"rt": "REFUNDABLE",
					"r": 7.04,
					"pk": "08032021193846992$1044",
					"hbo": false,
					"b": 0,
					"tb": 0,
					"cb": 0,
					"hbf": [],
					"e": 0,
					"bag": {
					  "G8413": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  },
					  "G8713": {
						"hb": "7 Kg",
						"cb": "15 Kg"
					  }
					},
					"cvf": {
					  "df": 309,
					  "tdf": 309
					},
					"icf": {
					  "df": 0
					},
					"tpd": 0,
					"pd": 0,
					"term": {
					  "AMD-BOM-G8413": {
						"dt": "Terminal 1",
						"at": "Terminal 1"
					  },
					  "DEL-AMD-G8713": {
						"dt": "Terminal 2",
						"at": "Terminal 1"
					  }
					},
					"ottf": 3547.6,
					"o": 0,
					"m": 0,
					"op": {
					  "$oa": "300"
					},
					"ind": 300,
					"earnUpto": false,
					"mFareType": "BASIC_FARE",
					"ftk": "R7qCh3m47HV2FhuK4uZfuj6WExu/u/r8OPmbB7AEhYL7YuUJtQB8NQXEjKLhGFz7BMCKbVCX9H8IeszGvfE0M6opjOalXrMx4DV5tA4UIadgyKbCNfYgwKpgjBeJ3PhG7Bf7JSWnO48=",
					"smfs": true,
					"sf": false,
					"fmd": {
					  "G8413": {
						"sr": 9,
						"ft": "GS"
					  },
					  "G8713": {
						"sr": 9,
						"ft": "GS"
					  }
					}
				  }
				]
			  },
			  "groupedFares": {},
			  "offerText": "<span class=\"dynot\" style=\"font:11px/1 system-ui,-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Oxygen,Ubuntu,Cantarell,Open Sans,Helvetica Neue; color: #559B09;\">₹$oa Instant Off</span>",
			  "multiFareTypeEnabled": true,
			  "bannerUrl": "https://images.ixigo.com/image/upload/a9350f56051b7a53ac019060f1ebdaa8-bzqrg.png",
			  "searchId": "08032021193846992",
			  "airlines": [
				"G8",
				"SG",
				"UK",
				"AI",
				"I5",
				"6E"
			  ]
			}
		  }
		}
	  }`

	fmt.Println("Go Resolver Running")

	type BagFaresInternal struct {
		Hb string `json:"hb,omitempty"`
		Cb string `json:"cb,omitempty"`
	}

	type TermFaresInternal struct {
		At string `json:"At,omitempty"`
		Dt string `json:"Dt,omitempty"`
	}

	type FmdFaresInternal struct {
		Sr int    `json:"sr,omitempty"`
		Ft string `json:"ft,omitempty"`
	}

	type FaresInternal struct {
		ProviderID int                         `json:"providerId,omitempty"`
		Markup     int                         `json:"markup,omitempty"`
		Valid      bool                        `json:"valid,omitempty"`
		Total      int                         `json:"total,omitempty"`
		C          string                      `json:"c,omitempty"`
		Bf         int                         `json:"bf,omitempty"`
		T          int                         `json:"t,omitempty"`
		F          int                         `json:"f,omitempty"`
		D          int                         `json:"d,omitempty"`
		Tf         int                         `json:"tf,omitempty"`
		Ttf        int                         `json:"ttf,omitempty"`
		Ibf        int                         `json:"ibf,omitempty"`
		Obf        int                         `json:"obf,omitempty"`
		Rt         string                      `json:"rt,omitempty"`
		R          float64                     `json:"r,omitempty"`
		Pk         string                      `json:"pk,omitempty"`
		Hbo        bool                        `json:"hbo,omitempty"`
		B          int                         `json:"b,omitempty"`
		Tb         int                         `json:"tb,omitempty"`
		Cb         int                         `json:"cb,omitempty"`
		Hbf        []interface{}               `json:"hbf,omitempty"`
		E          int                         `json:"e,omitempty"`
		Bag        map[string]BagFaresInternal `json:"bag,omitempty"`
		Cvf        struct {
			Df  int `json:"df,omitempty"`
			Tdf int `json:"tdf,omitempty"`
		} `json:"cvf,omitempty"`
		Icf struct {
			Df int `json:"df,omitempty"`
		} `json:"icf,omitempty"`
		Tpd  int                          `json:"tpd,omitempty"`
		Pd   int                          `json:"pd,omitempty"`
		Term map[string]TermFaresInternal `json:"term,omitempty"`
		Ottf float64                      `json:"ottf,omitempty"`
		O    int                          `json:"o,omitempty"`
		M    int                          `json:"m,omitempty"`
		Op   struct {
			Oa string `json:"$oa,omitempty"`
		} `json:"op,omitempty"`
		Ind       int                         `json:"ind,omitempty"`
		EarnUpto  bool                        `json:"earnUpto,omitempty"`
		MFareType string                      `json:"mFareType,omitempty"`
		Ftk       string                      `json:"ftk,omitempty"`
		Smfs      bool                        `json:"smfs,omitempty"`
		Sf        bool                        `json:"sf,omitempty"`
		Fmd       map[string]FmdFaresInternal `json:"fmd,omitempty"`
	}

	type OutRComR struct {
		Flights              map[string]string          `json:"flights,omitempty"`
		Fares                map[string][]FaresInternal `json:"fares,omitempty"`
		GroupedFares         struct{}                   `json:"groupedFares,omitempty"`
		OfferText            string                     `json:"offerText,omitempty"`
		MultiFareTypeEnabled bool                       `json:"multiFareTypeEnabled,omitempty"`
		BannerURL            string                     `json:"bannerUrl,omitempty"`
		SearchID             string                     `json:"searchId,omitempty"`
		Airlines             []string                   `json:"airlines,omitempty"`
	}

	type Result struct {
		Providers []string               `json:"providers,omitempty"`
		Results   map[string]interface{} `json:"results,omitempty"`
	}

	resultObj := Result{}

	err := json.Unmarshal([]byte(json2), &resultObj)

	if err != nil {
		return nil, fmt.Errorf("Error Occured")
	}

	providerNumber := resultObj.Providers[0]

	if resultObj.Results[providerNumber].(map[string]interface{})["outR"] != nil {

		outRMap := resultObj.Results[providerNumber].(map[string]interface{})["outR"]

		providerResultMarshal, marshalErr := json.Marshal(outRMap)

		if marshalErr != nil {
			return nil, fmt.Errorf("Error Occured")
		}

		outRJson := new(OutRComR)

		outRErr := json.Unmarshal(providerResultMarshal, outRJson)

		if outRErr != nil {
			return nil, fmt.Errorf("Error Occured")
		}

		toReturn := "Done"

		return &toReturn, nil

	} else if resultObj.Results[providerNumber].(map[string]interface{})["comR"] != nil {

		comRMap := resultObj.Results[providerNumber].(map[string]interface{})["comR"]

		providerResultMarshal, marshalErr := json.Marshal(comRMap)

		if marshalErr != nil {
			return nil, fmt.Errorf("Error Occured")
		}

		comRJson := new(OutRComR)

		outRErr := json.Unmarshal(providerResultMarshal, comRJson)

		if outRErr != nil {
			return nil, fmt.Errorf("Error Occured")
		}

		toReturn := "Done"

		return &toReturn, nil

	} else {

		toReturn := "Done"

		return &toReturn, nil

	}
}
