package eckcdsa

import (
	"testing"

	"github.com/RyuaNerin/go-krypto/internal"
)

func Test_Testcase_P256_SHA256(t *testing.T) {
	testVerify(t, testCases_P256_SHA256)
}

var (
	testCases_P256_SHA256 = []testCase{
		////////////////////////////////////////////////////////////////////////////////////////////////////
		// EC-KCDSA_(P-256)(SHA-256)_SVT.txt
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`9E2E7F3D44E9D01E2C8CEB76F07F7558DEC520EC3FB9A7225DB2FACF531F07EB65250C49013DCB880FE89F84B8F4C6144F61CE59836E986683D1C19233AAF9282C320CD79EB02893C1BF2B39066477C6A20F2156A3F3F7C233D932243EE448000B54C35C02CE83A74A38548155A4466F135FA08AE7A7E7121907D727D1F8BA65`),
			Qx:    internal.HI(`1200CFC0B312B25F6A4D40003625A043B31D533FD1D1A3A90DFDC8D4549350AC`),
			Qy:    internal.HI(`E8EBD1A68CA3F201774500CDA2C0B628FBCA0298793BB13902165CF9F4514AAB`),
			R:     internal.HI(`761B5C36706BD4148AF7DB729872C824A07987B0A4D0AE1FA2489787A93AFC87`),
			S:     internal.HI(`6E35B838499497F970D6B8EBFEF90B5D3F7ADF4D7C0DF19BA6708D8B6B3FA1DA`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`B2FEB3DCCA8114FAA56AB2E531FCFA210E5C036592901A2AC1C2A511B81838CA9A111EA380D112697E4189BF31683897FDED3AF56424879F918E0DE99C373F3E82D848967428C27763392AB47161E48737740E3C74DCEC002E5C3A10C6F128F00AB846BBD0422B4B1E38225B68C69F4BEE98C4A4C1F9C51373D86C01002B895E`),
			Qx:    internal.HI(`34983E6E19B8FD1E69A9C6BED5625767C4C110967DA6BA845032993D89413698`),
			Qy:    internal.HI(`88B51848A8AE9ECEA4423EFDD04BAB8EADAC5F1C539973377964BA4AD77838F1`),
			R:     internal.HI(`DA62575F94E4B87106FDCFC83E546EF42553BA60CA3C38907EF7025E8A2F7049`),
			S:     internal.HI(`FF0F363C40E99506319B992125BD073CCAF9D505CF4971039A2D6C4189A0B26F`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`9D91E269E53D559EAD6818327EF4857E8A0C5FF4C576F165334C02BCCEDF4D3AFB0DA1BADED14EFD8D5711967A551D6803FE695C914CF43127EB2012CF90CCE7CF0456E603A6347B57B53800AF1B4758431884F0B77BC02AF8C32FC529EC2F3E9C2B2ABA3165508011AF46861E8A7B2011423C50971F40CF6FB5CC3FE02815A5`),
			Qx:    internal.HI(`CDB2CA9A859A8D85796EB2B9A41A8371F93E7BA822538A8D2BC2136CA1A2E068`),
			Qy:    internal.HI(`CA566AE0CE903139A9C50A6AEFE7D03ABF7D7434D196CACBCC7420EA2B6CC40A`),
			R:     internal.HI(`7874BEE81F5031C1191176D7CAD5A462FE42575592CFCA63AFEC7367202CE9D1`),
			S:     internal.HI(`C594498417A1771F7EA8C10D0FD59957AE1D3E8B5931E687E2B746B7304308D3`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`9B0AA94D9E319D49DABCE6CFCA7CB9BFCAE2C49B3BA9C3DF085C90FD6164AADED9AECD45C9B08B1733E286CFEE54482468933A34BB75D221700C72969F490503CF93EC0E45EC2A171FC97337F15417A23A2F4CDB756AF29904D146C5CF3B9483CF0270A55B4D82716B4A6225B316971FC864371324933F31466718ADC6F65D8A`),
			Qx:    internal.HI(`DB5CA350FFCD24C21B4EA96B346F718EC5FEFB171DF25D9AA462FC17E4404755`),
			Qy:    internal.HI(`0D87B5480C33405123FA0EAEA4066D332BBA4A04A28F61D4455C31E503B51D7F`),
			R:     internal.HI(`362E3B6EAC12457C6BF51D768386436A3D6A4CF1395DC3A934042F5A14F9E303`),
			S:     internal.HI(`B1072B6A658918D10C7BAEC758397768D528106AABDAE627223E95F9926F2F3E`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`096D0F3B82C5D34C06B2C2007D640E01A2E5646B43FE5A1E53977A44818FCD477E9E5BFD844A726065673C89B140653C26E35FD8179B10DFE21D5FEF2A80896B01689CA193930FAB9930CBE82416A537F26B9EDDF2AAD44028AD975387AF80F7F8D5F10F3548ED9B9A61AD9F91F9D71BD05BC2D800E8B16304EF66E875C30146`),
			Qx:    internal.HI(`7E5C2B5C91EF3F1FAEBB9CA339340B05B07C67732F9DFEC72BE900451EE73761`),
			Qy:    internal.HI(`F10B19CF2C12B2BE2FFB25B00FA37F52C1A42E8DECC489D4D0243AE062351E27`),
			R:     internal.HI(`6366B11769B1404C62AC80D8B39140C789A69EE075F24641CE760422059E12D5`),
			S:     internal.HI(`80452846360A0316D49DB3ECDFF743ABDDF54AE8CAD3D21D0EB59A16B0FDEABE`),
			Fail:  true,
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`14F037531FFD455A1EE09B66974BEAA9AB2220F39D1066E3D0AA3CA0422317390AE12B53EFF0C15489ADD0A90994315B24A0278C2F1A8B18EB4763E81BDA2D3D5593A27E27F33DC814694CDA69EAD28A834D086B7766B9CABB3254B2C710E20B4063C39BB08A67F05493D1F40B933343E70A5EFB743D53182BEEC71B2C51D313`),
			Qx:    internal.HI(`FC62E2DD0F1A5A3F0BB1173D4D9C1C023D42832EB189DD8B8DF7902EB9DAF7D9`),
			Qy:    internal.HI(`F7AC755E1E44AD4A869E9EC6CDF7E3144275A50A5D86A32293843E6493C78B0D`),
			R:     internal.HI(`A5712E589E245307FFC1EE6E0C55F738DE7FCE71189F862D8AF154CE6541A556`),
			S:     internal.HI(`54DD9EEF98E6A54F38B43D8EAB9214FB11AE139AF92CB353F48A6D1CB359056D`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`B69FBB99C552CD87B409376760245818BFADF149721FE6CB06588A59B1018C6DFA2E351D3A7C279259852D22ABDB64845F6348425F50E0D50F8FAC84CDD9738EA44FAA9A07CA8F041BC1A90552F8C944E2D243D5E67F7C2BA2D8780CCA2D910D31CB009AD38A5E11991EC403BF82B3184A37709E71ADF756648B2B775338737F`),
			Qx:    internal.HI(`39F25B7D263CD3EDEBDF6B4CE8F13C5ED2FA934BC7AA3128FA230F91DE549A96`),
			Qy:    internal.HI(`4DFFB6D472475C642D4C8C0EDBC2A2C040A56787387F7E0D6440565C98A52EBC`),
			R:     internal.HI(`DDA4A24EA06D8FAEF7F39F89DEADA4BCAAC3931DE86D389B18A4BE7ED86A13B8`),
			S:     internal.HI(`68D7FB003435AE2112792B1E75143DB9C76F09CEA5C8FA6D94373938FE5AE3DA`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`5268CDCB70B5977AE6E037E1E71466A0B7A162F442558800E54A749C9D711D117D5A0CA6997D1B6D4BA6E33100A73FF65CF30CE710629B280CA5D96368AACE88CA9557A38A66CC2A680A71AF09EF59641BC9B34F50B510990F4FA6EB0B041A27A221D88831E97F703E4376E4AC36150DD0D091B4A3A91016E3A80000622E3793`),
			Qx:    internal.HI(`039836EBE248114D6058D1EBAA1C58152B02FD38F36A13525A875E975D76EA4B`),
			Qy:    internal.HI(`BFA5766627E1D6F2F74AFA39D094E22312D16F6BF10E2ACDA34D4F78B38159FA`),
			R:     internal.HI(`AE1C16CE132CDF77884F43BEEE7A07C93DC05658D93812C300D2CE9EDC297F12`),
			S:     internal.HI(`B0ED66D25CB9EA1E3287333B35E4854BD48F7C194EF4162951262E361FD60C5E`),
			Fail:  true,
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`61EFDEEEB73C2E018B63701882377E4DEE6DA6D446DF1A601F45D8FAC7F88B1039C13508803E1B1A49EDDAC5AA520095EC54178F614DA20B00DBC33EE9F11A387843C0B3127AAE5C524E96FDBE0C8337AE7C2D2488CB6EA3006CB1928284DCCB67DF851B2403E30720596AA549046C4F1E60B57556C07F5CBD5FAF138223B65B`),
			Qx:    internal.HI(`E86AF96A98D0164830CE1BFBA291946F5ADFE3552E631CF2326309290ECAFFE2`),
			Qy:    internal.HI(`4F9945786EFC54A6BC07E853930DA9A1CBCD9DAEC87E36F8678AB23A3EB90BD4`),
			R:     internal.HI(`9C7696726B1EFF5ECB7CBA3AF9137F2CEB83616A993D73CBC09630A45E6AB6D6`),
			S:     internal.HI(`E31F55844AF6BC1739F9B720E52C0FE22CEACFAF459883EA1CA5100F0F24D835`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`2513D73899C4467509A70F9C8122E371A6FECA313FE2561C435243A0EA965DDBCE7FE296A05E2B5A7F26EFD37418DD221AE5D54701AA954CCEB50C599E713458BA69A34F26CBB022074767D173C712F5B0CBAF8A69F309C2DF7EC3F6A8442674F1F6E63DE027B86513F3FC40BEAF3EF5859BD8B97EDE863C7DD822841E1465DE`),
			Qx:    internal.HI(`40F8E7638937707785358F6F784B78CE8A5DD62D8E7C24EAAA0CFC0B7EE21E07`),
			Qy:    internal.HI(`52FE3AE45395790C6A54CC695883005412E43532B995B163E1420DD80E5F6865`),
			R:     internal.HI(`DD41240FFBBB53C2FB4235C4F541C3DA7BC97DE9533005E2FF39FA590F452CBD`),
			S:     internal.HI(`2E98CAFF7A2086F327DA77890F14320D1D8B557E36CFBB42473CD0D7A3609840`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`D5E0002D9C0B073D999E933D5B0578DDC63B59F399CCBFD5784DDBBAEF078F2D3C2E0011595D314932A4067F5BF218264826B52255809B3F9ACE99E1EB0DC70BE9D66D3B944754203ECC9ABD497BB9F3A355E30C4F3164390298269D23600DE7234442008D4B2AD6C53F1BD4C6FB8702E0E8C05AC954F869FD06B0A2BFCE3091`),
			Qx:    internal.HI(`20F411996BFD10954C9C42F4DC3EBB26170B2DB75E9DE601F52F1351E0DD8FE4`),
			Qy:    internal.HI(`DEAD2083C2F0E634DE3D194E334B555E331558E0ACFD33C73364FD13455A98ED`),
			R:     internal.HI(`BA8B2A6193955EC3C5DC65599B6A8D74F80D18F80BB59ADA93E2E64ECA7BB2E2`),
			S:     internal.HI(`1C52ACC5290BBA9F28D61B79D77DFB3B4531BC8E05D0C97287B700E4203DAEAC`),
		},
		////////////////////////////////////////////////////////////////////////////////////////////////////
		// EC-KCDSA_(P-256)(SHA-256)_SGT.txt
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`93B38ECE8AA1C0E5633B04676E0FACED97B17F217521A782190720B537CD90ABE21DE45B6ACDA867A3CB78651BF8721A899821760148AB8D2309CC5CF0766AA163903A8B75435D91C71F4E30815025DE9E5976F247A5D4D76A633F442DD089C46955FB62DEA52FAB344892E886A457BD8C0E6E0E22F0375E387B887DF5FCC7CB`),
			Qx:    internal.HI(`7A010CCE56D2CDF9BA5B8483F4641914650541FEED4B41A056688F199204D4BF`),
			Qy:    internal.HI(`BED08099040972E8372C3A63DF73B3C8F3D94CB2242C258770BA5DB031D0FADF`),
			R:     internal.HI(`A34D0B562DFEFE0BBE71D465AD94A6ABC40003B413EB9BB8BB079B89815DD313`),
			S:     internal.HI(`1FBC4CAC222E8B9A70367CB4E5BC27F89A2378EBCADB458C5E879E32BFB18D05`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`0F5B96CB21C2FB671701569257AB27A3D37BFC2AAF0B74859E5DBF039C4682D43739879B028D22DD65ABB71B6A4438E1B80B1E8D57FDB01B73C003B6BB1FA3678D304D8E8D238F8FDA4FD75B7D75B6344BD2DA69CAF89A5359017A6234249B728326724816C8B466FB9EE20F976D56BE638C3DD4005368CDB82655B83115624E`),
			Qx:    internal.HI(`FE70B944ACC94D5A77082031F9BF9DA5E942735EAFE5C1510307281BBB29D8C5`),
			Qy:    internal.HI(`6674F43B5B3016E0424E068584F2C0B2FDBBCDF46B403E23C063AF528F83C19A`),
			R:     internal.HI(`E0E2321CF8B618AEDAFAEC0E0629741F36E14FF93FF7BDB9A44C608A83E522A9`),
			S:     internal.HI(`5DE51BD7BA26FBCC349E4D1202E8EF47DAFE7B6D1C51E0BF36ED18D71672913E`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`55AC885538EA8BB315F40BDE0383540B9E9C5CD03790D88BC2DCCB5223DF78B804F92729D81C962E3D508D545A35740D640D5544F95732FC6F791BB3C593E6D85BECF221201CE0FA505DADE97729031FE3944C00157842F481FB8C44186B3B3DEF0D917F86CDFC9EF26CB4FB822FD505118B6FBCA3EC5FB17F066E5766CB91DF`),
			Qx:    internal.HI(`5CED7DF123CECCBC5B9FA2CC622C2D67C9AEC0A61BD71EC116CC3F27D5ADD1F9`),
			Qy:    internal.HI(`97C7E7CB7E5E654FFE712B4F41A7235F3B56C3C9F5B549E773773EE7427FBDAB`),
			R:     internal.HI(`7175BBBFCAEED5F905591BFD1D7FEFE067C6F718F303DA896992D8E477809F4C`),
			S:     internal.HI(`548A60B9E97464AB8488E766B9A2249140F3C7DBCAE38B405DC2ECBD36F2D3BA`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`188CB671C43C030D9077F6CFE83C47E9AFF7F41803D468D8BAE91827413E871BFB401A0ADF9F979F5F1FCF945E703C604183179DDC79C3754D98E9D6827446BA82A97E462251E4155DAEA25FE310216219831FBA1E495EFE16B3496E4C4A7EE95FEEAC0A21D79B984C15DA2FBC8CEA554BF158C701E0B050C07EA6E008C26742`),
			Qx:    internal.HI(`28625C04651555038277EAFBB00E414DC5569153F1EAF7E6A8F94735A9AC5801`),
			Qy:    internal.HI(`5CC4B15DFEE6DBCA7BDC8C54CBA3D7DA55D4AE274A3589B721A67B2A5AE93848`),
			R:     internal.HI(`65FD026975082B809FA96A958C0AE4763D7EB728E20CA064BA5E1700E285D59B`),
			S:     internal.HI(`E8E55A59B0DB5C56933D0E9FF6BE2618CBB06637646589D542A3C469DDCB1BC2`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`0EDE7521B8DDF7B9BCEFEBEA787A1302B772170505FCB7AFBAE679076805C2C1D1F3B2400B39BA72007C505EEB98A29E0451BA9CF286F8CA3E8241A56768D7CFB44A450386E62F2435A68C4035CE21C2A182A79CD79084B44D8E84634464773B88AB17EDDC0A27973FFE2A30B929A773C4A04AFC0D51EFECAFF1D2D58C9CF93B`),
			Qx:    internal.HI(`179FC5373742ED240CD7060AFDF38880958AA346BC49D1CE46BCDE9A2050D90B`),
			Qy:    internal.HI(`C5609943B9EB931975D787AD62A3A873ABA8FBA86D71A4CD004BD417CC570B63`),
			R:     internal.HI(`18CFD393A31B69BF09A876951B00E9262F61D7DE2C08FF03C34788D5095FD167`),
			S:     internal.HI(`3CA13902FF26C8F40B61CD7E13C0896F6BE7033B484B1032EC1F4C54D417F124`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`EA87196B09F1FDFCCDBFBFB228E0CE1A6CF01A9D332A5856F638C3740EDB3E6E3AF445D1510F93ED52CAE4367552BB8C605B924730A4663E7999F6A2E612AEDDA6B39B5B4000546B0CA83E11E0061A032F7739A9366F465C5AF112A9755D3AF51E2A262CAA8B31DF008B7881EDA92022307E9C5DBB64B0CA82C5C6BB65005A8E`),
			Qx:    internal.HI(`25A2EE072525F200764B72BEE97F2EF12A1A643838AAAFE5FC03DAE37B5890EA`),
			Qy:    internal.HI(`455192C94D5EE47AB187408A34282804F46C49600F5E9720A61DDC8038D8BEFA`),
			R:     internal.HI(`ABA51AD187D29AA020B31B6BE14C28813FDE3D02359DDC6A2FE1CC796ADEFACE`),
			S:     internal.HI(`C064F44F204FB55705B9EF7837F37279BF8DF6E9F546B5F7647AA9A4F9069EF0`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`5F6AF652AA9AA619F74C44AC6B138AF4815450E27F83E10FA144CAF3A7610EE7EB2A25C0A544B5538B6D5FA070429AED0A85F29F8BF59F173043DD527516DEA60CC9D45245C3E72F16188C54595D1FE8784428E62D0B3A38703FC6C352DADBDCD54F2DCC807D50B4BE1F96A6CCB26B27446D9FEFFE3D872E6B5D5517078E9E00`),
			Qx:    internal.HI(`11251D555E71A6D85F6DE6A4BC25A2CC9CC01643F1F1064A66473B220CD7BDA6`),
			Qy:    internal.HI(`5E27EA67EFBFCAA12FEBAEDC03B8DD436F85B2A535364524994DD35696858D54`),
			R:     internal.HI(`7380271EFE6DF72EAF999D3DE345824E6FCF59D733F9EA4D5BE888856207BE10`),
			S:     internal.HI(`7716781C3752B1C3D725CA46F2E2027235E34935E751D48050A10E80EF1C2B8F`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`226B60DA900089546FFA505BB7B75C55AB830DD8DE2BE51FEF6D6108A53C46EE9576A711FAFCB4E9DECB96224F0C5586B5B22FABF59F399798E2CA3987197BF0996F44EC89527DB3885C4A901477433530CEC957B288F38CC4DB76344F7D6FB561FD7FCF5105155BB11E5924C9E69B45B352AAB5CC00085B9E1E556BE6EDD952`),
			Qx:    internal.HI(`FCAEE0D2FC34E1C8F88730489E520985173E0313838F10CFAB4AA4CA0F02ABF6`),
			Qy:    internal.HI(`86B706D6551D954C96436268BA60BE4DD6A9DFAAC66A51D1E24D00A617BF123A`),
			R:     internal.HI(`546CC00AD4C22EF065DAC6E7C28BC83E573819152650FC7CFD92BC677FE664C9`),
			S:     internal.HI(`B71C233DB04A46AA599EFDC53332AB225094259C1E573D721C3ECE22BD77FC02`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`E86FAB07AE4237F2672CB6437E7058009D61A6844546F7C916175E377E11FA48EFBE1FC8445C25F27F465B3D8854FE1B15C89B6D63C4C703E4DC90DB91BE9A7E03893F2D00D2AA3B94D64D4684F79BAF0AF96F00B70A059E8A2BF482E0EC09427618723B124617170DEE967E59EAC34132110EB418D1C897506A973C77C0204C`),
			Qx:    internal.HI(`A4656B7069810314868C2C67A96B3C9F0D204C61E7544621ED2106A86C882779`),
			Qy:    internal.HI(`3E2B97F002C75C7009DA6510F8A975D7FD3F6463BD6DD13C76C04C8E261DBAA4`),
			R:     internal.HI(`5A514B674CAEC73135D895FAB7C98A0CFE08585C4FB45DB09288764A5686E0BD`),
			S:     internal.HI(`8C03F2020AEE6EDCE75021FD909E5A8DF8554968307B27427D73CADFCBA59B00`),
		},
		{
			curve: p256,
			hash:  hashSHA256,
			M:     internal.HB(`63592BDEF987473615474AE935E191BA0CD26EEAA6F8AD5348A59204A5833FBAACE5E2E978889CB2A24284778DBDAB6FDFA98DEACA8ADD9F499404BC06A84E14FDFC191A9B66010C70EA68FD1E833B1ABAA96FE432B504B0F592142F7ACABE48C8845913B665E82C04F02039F061F8DE738D21F0D6D45924B5A6F20E2CAB86AF`),
			Qx:    internal.HI(`05D5337826EDE7EBD32FE552826770C88ECE4D32E4882BAF1CC8C79C7033E270`),
			Qy:    internal.HI(`A8CEDF333BC68E8F35139D4360062A26AA3AB7497120236988B2A8F6EA9B70EF`),
			R:     internal.HI(`E2CDA28C3A5A4B61DD92683A6E73CC10145481E13C04E82A4B8EBC15BA3A346D`),
			S:     internal.HI(`63E1E6F63D515B83820CAB41FE04A4E197502B191C88740363802026E638A03A`),
		},
	}
)
