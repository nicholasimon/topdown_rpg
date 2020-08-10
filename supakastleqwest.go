package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ( // MARK: var

	// fade variations
	fade1   = float32(0.3)
	fade1on bool
	// info stats tabs
	statson, statshover, infohover, statsflash, infoflash bool
	infoon                                                = true
	armorcount, otherdamagecount                          int
	playerlevel                                           = 1
	attackcount                                           = 1
	hpcount                                               = 10
	hungercount                                           = 10

	// loot
	coins                              int
	chestblock, chestitemnumber        int
	chestitem1, chestitem2, chestitem3 items
	// weather day night
	nightshadow                = rl.NewRectangle(674, 162, 350, 350)
	nighton                    bool
	weathercount, weathertimer int
	weathercurrent             string
	days                       int
	// inventory
	selectedslotnumber                                                                                                                                                                                   int
	placeweapon1, placeweapon2, placehelmet, placenecklace, placearmor, placering1, placering2, placegloves, placebelt, placeboots                                                                       bool
	weaponslot1on, weaponslot2on, helmetsloton, necklacesloton, armorsloton, beltsloton, ring1sloton, ring2sloton, glovessloton, bootssloton                                                             bool
	weaponslot1equipped, weaponslot2equipped, helmetslotequipped, necklaceslotequipped, armorslotequipped, beltslotequipped, ring1slotequipped, ring2slotequipped, glovesslotequipped, bootsslotequipped bool
	invitemstructactive                                                                                                                                                                                  items
	itemactiveblock                                                                                                                                                                                      int
	invitemactive                                                                                                                                                                                        string
	invitemactiveon                                                                                                                                                                                      bool
	// level
	levelnum = 1
	// maps
	questmap               = make([]string, levela)
	equippeditemsstructmap = make([]items, 10)
	charactersmap          = make([]string, levela)
	effectsmap             = make([]string, levela)
	enemiesmap             = make([]string, levela)
	enemiesstructmap       = make([]items, levela)
	itemsstructmap         = make([]items, levela)
	weathermap             = make([]string, gridw2*(gridh2*2))
	invxmap                = make([]int32, 10)
	inventoryfullmap       = make([]bool, 20)
	inventorymap           = make([]string, 20)
	inventorystructsmap    = make([]items, 20)
	itemactionsmap         = make([]string, levela)
	items3map              = make([]string, levela)
	items2map              = make([]string, levela)
	itemsmap               = make([]string, levela)
	extrasexteriormap      = make([]string, levela)
	extrasinteriormap      = make([]string, levela)
	leveltilesmap          = make([]string, levela)
	levelmap               = make([]string, levela)
	words1                 = []string{"rusty ", "old ", "worn ", "chipped ", "dirty ", "cracked ", "rusted ", "damaged ", "crumbling ", "primitive ", "archaic ", "antiquated ", "flawed ", "bent ", "defective ", "imperfect ", "faulty "}
	words2                 = []string{"exquisite ", "finest ", "fine ", "first-rate ", "skillful ", "sterling ", "magnificent ", "priceless ", "supreme ", "tiptop ", "champion ", "first-class ", "elite ", "splendid ", "refined ", "polished ", "superb "}
	words3                 = []string{"sparkling ", "gleaming ", "glowing ", "brilliant ", "shimmering ", "shining ", "radiant ", "glittering ", "dazzling ", "luminous ", "shiny ", "silvery ", "flashing ", "twinkling ", "phosphorescent ", "incandescent "}
	wordscolors            = []string{"red ", "blue ", "green ", "silver ", "gold ", "purple ", "pink ", "magenta ", "blue ", "turquoise ", "black ", "obsidian ", "emerald ", "ruby ", "white ", "pearl ", "yellow ", "mustard ", "brown ", "lilac ", "beige ", "grey "}
	wordsmaterial          = []string{"copper ", "iron ", "steel ", "bronze ", "diamond ", "emerald ", "ruby ", "silver ", "nickel ", "gold ", "titanium ", "platinum ", "tungsten ", "osmium ", "palladium ", "rhodium "}
	wordsjewels            = []string{"diamond", "ruby", "emerald", "sapphire", "alexandrite", "topaz", "jade", "amber", "garnet", "pearl", "opal", "zircon", "tourmaline", "hematite", "onyx", "quartz", "moonstone", "malachite"}
	wordspotions           = []string{"potion", "tonic", "elixir", "medicine", "brew", "cordial", "dose", "libation", "remedy", "filter", "tonic", "mixture"}
	wordsswords            = []string{"blade", "sword", "scimitar", "rapier", "longsword", "broadsword", "sabre", "gladius", "falchion", "claymore", "cutlass", "greatsword", "xiphos", "kopis", "dao"}
	wordsaxes              = []string{"axe", "hatchet", "maul", "tomahawk", "double bit axe", "viking axe", "tactical axe", "battle axe", "grub axe", "broad axe", "hunters axe", "carpenters axe", "hewing axe", "firemans axe"}
	wordsmaces             = []string{"mace", "bludgeon", "war hammer", "flanged mace", "medieval mace", "pegged mace", "spiked mace", "battle mace", "morning star"}
	wordscrossbows         = []string{"crossbow", "recurve crossbow", "repeating crossbow", "compound crossbow", "rifle crossbow", "hunting crossbow", "heavy crossbow", "bullet crossbow"}
	wordsstaff             = []string{"staff", "stave", "battle staff", "quarterstaff", "pole", "pikestaff"}
	wordsclub              = []string{"club", "cudgel", "war club ", "pickaxe handle", "knobkerrie", "kubotan", "baton", "truncheon"}
	wordssickle            = []string{"sickle", "scythe", "war scythe", "war sickle", "bagging hook", "billhook", "harpe", "kama", "khopesh"}
	wordsbow               = []string{"bow", "hunters bow", "recurve bow", "longbow", "takedown bow", "compound bow", "reflex bow", "flatbow", "viking bow"}
	// imgs
	// player outline
	playeroutline = rl.NewRectangle(558, 278, 101, 230)
	// characters
	playerrunR  = rl.NewRectangle(0, 306, 16, 16)
	playerrunL  = rl.NewRectangle(96, 306, 16, 16)
	playeridleR = rl.NewRectangle(0, 325, 16, 16)
	playeridleL = rl.NewRectangle(96, 325, 16, 16)
	char1       = rl.NewRectangle(223, 164, 16, 15)
	char2       = rl.NewRectangle(247, 163, 16, 16)
	char3       = rl.NewRectangle(271, 164, 16, 15)
	char4       = rl.NewRectangle(295, 163, 16, 16)
	char5       = rl.NewRectangle(321, 165, 17, 16)
	char6       = rl.NewRectangle(345, 165, 16, 15)
	// effects
	smoke1 = rl.NewRectangle(21, 147, 19, 18)
	smoke2 = rl.NewRectangle(42, 148, 17, 16)
	smoke3 = rl.NewRectangle(61, 150, 14, 13)
	// extras
	xcross      = rl.NewRectangle(5, 251, 18, 17)
	xskullbones = rl.NewRectangle(28, 248, 19, 23)
	xskulls     = rl.NewRectangle(55, 252, 12, 16)
	xvase       = rl.NewRectangle(76, 253, 15, 16)
	xfishbones  = rl.NewRectangle(97, 254, 18, 13)
	xbonecross  = rl.NewRectangle(121, 251, 18, 17)
	xknight     = rl.NewRectangle(146, 254, 12, 16)
	xcastle     = rl.NewRectangle(162, 255, 10, 14)
	xbone1      = rl.NewRectangle(176, 255, 12, 12)
	xskull1     = rl.NewRectangle(191, 256, 13, 10)
	xskull2     = rl.NewRectangle(208, 255, 11, 12)
	xcandle     = rl.NewRectangle(223, 256, 8, 14)
	xskull3     = rl.NewRectangle(237, 258, 13, 12)
	xskull4     = rl.NewRectangle(250, 256, 16, 12)
	xbone2      = rl.NewRectangle(271, 252, 18, 18)
	xbone3      = rl.NewRectangle(296, 253, 16, 16)
	xbone4      = rl.NewRectangle(318, 256, 17, 11)
	xpotplant   = rl.NewRectangle(341, 251, 11, 17)
	xsign1      = rl.NewRectangle(357, 253, 16, 16)
	xsign2      = rl.NewRectangle(376, 253, 16, 16)
	xlamp       = rl.NewRectangle(398, 255, 18, 13)
	xjug        = rl.NewRectangle(422, 252, 18, 16)
	xvase2      = rl.NewRectangle(446, 252, 18, 15)
	xhanger     = rl.NewRectangle(469, 253, 17, 14)
	xanchor     = rl.NewRectangle(491, 251, 17, 14)
	// chest monster
	chestmonster1 = rl.NewRectangle(1, 224, 16, 16)
	chestmonster2 = rl.NewRectangle(17, 224, 16, 16)
	chestmonster3 = rl.NewRectangle(33, 224, 16, 16)
	chestmonster4 = rl.NewRectangle(49, 224, 16, 16)
	// keys
	key1 = rl.NewRectangle(685, 37, 14, 17)
	key2 = rl.NewRectangle(708, 37, 14, 18)
	key3 = rl.NewRectangle(733, 37, 14, 18)
	key4 = rl.NewRectangle(757, 38, 15, 16)
	key5 = rl.NewRectangle(685, 61, 15, 18)
	key6 = rl.NewRectangle(710, 62, 10, 17)
	key7 = rl.NewRectangle(735, 62, 10, 17)
	key8 = rl.NewRectangle(756, 63, 16, 14)
	// armor
	boots1  = rl.NewRectangle(614, 33, 16, 16)
	boots2  = rl.NewRectangle(630, 33, 16, 16)
	boots3  = rl.NewRectangle(614, 49, 16, 16)
	boots4  = rl.NewRectangle(630, 49, 16, 16)
	gloves1 = rl.NewRectangle(646, 33, 16, 16)
	gloves2 = rl.NewRectangle(662, 33, 16, 16)
	gloves3 = rl.NewRectangle(646, 49, 16, 16)
	gloves4 = rl.NewRectangle(662, 49, 16, 16)
	armor1  = rl.NewRectangle(613, 19, 15, 14)
	armor2  = rl.NewRectangle(628, 19, 16, 14)
	armor3  = rl.NewRectangle(644, 19, 16, 14)
	armor4  = rl.NewRectangle(660, 19, 16, 14)
	armor5  = rl.NewRectangle(676, 19, 16, 14)
	helmet1 = rl.NewRectangle(614, 0, 14, 16)
	helmet2 = rl.NewRectangle(630, 0, 14, 16)
	helmet3 = rl.NewRectangle(646, 0, 14, 16)
	helmet4 = rl.NewRectangle(662, 0, 14, 16)
	helmet5 = rl.NewRectangle(678, 0, 14, 16)
	helmet6 = rl.NewRectangle(694, 0, 14, 16)
	helmet7 = rl.NewRectangle(710, 0, 14, 16)
	shield1 = rl.NewRectangle(725, 0, 14, 16)
	shield2 = rl.NewRectangle(741, 0, 14, 16)
	shield3 = rl.NewRectangle(757, 0, 14, 16)
	shield4 = rl.NewRectangle(773, 0, 14, 16)
	shield5 = rl.NewRectangle(725, 16, 14, 16)
	shield6 = rl.NewRectangle(741, 16, 14, 16)
	shield7 = rl.NewRectangle(757, 16, 14, 16)
	shield8 = rl.NewRectangle(773, 16, 14, 16)
	belt    = rl.NewRectangle(694, 17, 16, 16)
	// food
	apple       = rl.NewRectangle(272, 206, 16, 16)
	pear        = rl.NewRectangle(292, 204, 15, 18)
	pumpkin     = rl.NewRectangle(313, 206, 18, 17)
	icecream    = rl.NewRectangle(336, 207, 16, 16)
	gingerbread = rl.NewRectangle(358, 205, 18, 18)
	pizza       = rl.NewRectangle(381, 207, 15, 16)
	chocolate   = rl.NewRectangle(402, 207, 15, 16)
	// weapons
	axe1      = rl.NewRectangle(468, 32, 16, 16)
	axe2      = rl.NewRectangle(484, 32, 16, 16)
	axe3      = rl.NewRectangle(500, 32, 16, 16)
	axe4      = rl.NewRectangle(516, 32, 16, 16)
	axe5      = rl.NewRectangle(452, 48, 16, 16)
	axe6      = rl.NewRectangle(468, 48, 16, 16)
	axe7      = rl.NewRectangle(484, 48, 16, 16)
	axe8      = rl.NewRectangle(502, 50, 14, 12)
	axe9      = rl.NewRectangle(518, 51, 12, 11)
	bomb      = rl.NewRectangle(518, 74, 14, 13)
	crossbow1 = rl.NewRectangle(452, 0, 16, 16)
	crossbow2 = rl.NewRectangle(468, 0, 16, 16)
	crossbow3 = rl.NewRectangle(484, 0, 16, 16)
	bow1      = rl.NewRectangle(452, 16, 16, 16)
	bow2      = rl.NewRectangle(468, 16, 16, 16)
	bow3      = rl.NewRectangle(484, 16, 16, 16)
	bow4      = rl.NewRectangle(500, 16, 16, 16)
	bow5      = rl.NewRectangle(516, 16, 16, 16)
	club1     = rl.NewRectangle(533, 0, 16, 16)
	club2     = rl.NewRectangle(549, 0, 16, 16)
	club3     = rl.NewRectangle(565, 0, 16, 16)
	club4     = rl.NewRectangle(597, 32, 16, 16)
	club5     = rl.NewRectangle(597, 48, 16, 16)
	club6     = rl.NewRectangle(581, 48, 16, 16)
	sickle1   = rl.NewRectangle(581, 0, 16, 16)
	sickle2   = rl.NewRectangle(597, 0, 16, 16)
	mace1     = rl.NewRectangle(452, 33, 16, 16)
	mace2     = rl.NewRectangle(581, 48, 16, 16)
	mace3     = rl.NewRectangle(597, 48, 16, 16)
	mace4     = rl.NewRectangle(565, 32, 16, 16)
	staff1    = rl.NewRectangle(533, 16, 16, 16)
	staff2    = rl.NewRectangle(549, 31, 16, 16)
	staff3    = rl.NewRectangle(565, 16, 16, 16)
	staff4    = rl.NewRectangle(581, 16, 16, 16)
	staff5    = rl.NewRectangle(597, 16, 16, 16)
	staff6    = rl.NewRectangle(533, 32, 16, 16)
	staff7    = rl.NewRectangle(549, 32, 16, 16)
	staff8    = rl.NewRectangle(565, 32, 16, 16)
	staff9    = rl.NewRectangle(581, 32, 16, 16)
	staff10   = rl.NewRectangle(533, 48, 16, 16)
	staff11   = rl.NewRectangle(549, 48, 16, 16)
	staff12   = rl.NewRectangle(565, 48, 16, 16)
	uzzi      = rl.NewRectangle(503, 66, 12, 10)
	shotgun   = rl.NewRectangle(456, 64, 16, 15)
	sword1    = rl.NewRectangle(535, 66, 12, 12)
	sword2    = rl.NewRectangle(551, 66, 12, 12)
	sword3    = rl.NewRectangle(567, 66, 12, 12)
	sword4    = rl.NewRectangle(583, 66, 12, 12)
	sword5    = rl.NewRectangle(599, 66, 12, 12)
	sword6    = rl.NewRectangle(533, 80, 16, 16)
	sword7    = rl.NewRectangle(549, 80, 16, 16)
	sword8    = rl.NewRectangle(565, 80, 16, 16)
	sword9    = rl.NewRectangle(581, 80, 16, 16)
	sword10   = rl.NewRectangle(597, 80, 16, 16)
	// potions
	potion1 = rl.NewRectangle(3, 206, 16, 16)
	potion2 = rl.NewRectangle(28, 206, 16, 16)
	potion3 = rl.NewRectangle(52, 205, 16, 16)
	potion4 = rl.NewRectangle(75, 206, 16, 16)
	potion5 = rl.NewRectangle(99, 204, 16, 17)
	potion6 = rl.NewRectangle(123, 204, 16, 17)
	potion7 = rl.NewRectangle(146, 205, 16, 16)
	potion8 = rl.NewRectangle(170, 206, 16, 16)
	potion9 = rl.NewRectangle(194, 204, 17, 17)
	// scrolls
	scroll1 = rl.NewRectangle(217, 207, 14, 14)
	scroll2 = rl.NewRectangle(232, 206, 16, 16)
	scroll3 = rl.NewRectangle(249, 206, 16, 16)
	// loot
	coin        = rl.NewRectangle(19, 131, 14, 14)
	ring1       = rl.NewRectangle(469, 89, 14, 14)
	ring2       = rl.NewRectangle(485, 89, 14, 14)
	ring3       = rl.NewRectangle(501, 89, 14, 14)
	ring4       = rl.NewRectangle(517, 89, 14, 14)
	necklace1   = rl.NewRectangle(469, 106, 14, 13)
	necklace2   = rl.NewRectangle(485, 106, 14, 13)
	necklace3   = rl.NewRectangle(501, 106, 14, 13)
	necklace4   = rl.NewRectangle(517, 106, 14, 13)
	necklace5   = rl.NewRectangle(469, 121, 14, 14)
	necklace6   = rl.NewRectangle(485, 121, 14, 14)
	necklace7   = rl.NewRectangle(501, 121, 14, 14)
	necklace8   = rl.NewRectangle(517, 121, 14, 14)
	jewel1      = rl.NewRectangle(539, 100, 16, 14)
	jewel2      = rl.NewRectangle(538, 123, 17, 14)
	jewel3      = rl.NewRectangle(564, 103, 14, 11)
	jewel4      = rl.NewRectangle(563, 122, 15, 16)
	chestclosed = rl.NewRectangle(35, 188, 14, 13)
	chestopen   = rl.NewRectangle(51, 187, 14, 14)
	// weather
	weatherimg rl.Rectangle
	moon       = rl.NewRectangle(66, 167, 15, 18)
	sun        = rl.NewRectangle(88, 167, 18, 18)
	rain       = rl.NewRectangle(111, 167, 18, 18)
	wind       = rl.NewRectangle(134, 167, 18, 18)
	storm      = rl.NewRectangle(159, 167, 18, 18)
	snow       = rl.NewRectangle(46, 170, 16, 16)
	snow2      = rl.NewRectangle(0, 169, 15, 15)
	snow3      = rl.NewRectangle(15, 170, 15, 15)
	snow4      = rl.NewRectangle(30, 170, 15, 15)
	// menu
	foodmenu         = rl.NewRectangle(102, 76, 17, 14)
	heart            = rl.NewRectangle(87, 79, 12, 10)
	inmenu1, inmenu2 bool
	mousel           = rl.NewRectangle(46, 77, 10, 14)
	mouser           = rl.NewRectangle(62, 77, 10, 14)
	mousem           = rl.NewRectangle(75, 77, 10, 14)
	mouse            = rl.NewRectangle(30, 77, 10, 14)
	// actions
	dug = rl.NewRectangle(1, 148, 18, 17)
	// tools
	torch   = rl.NewRectangle(474, 65, 14, 16)
	pickaxe = rl.NewRectangle(488, 65, 14, 14)
	spade   = rl.NewRectangle(0, 131, 14, 15)
	// doors
	door1 = rl.NewRectangle(0, 92, 16, 16)
	// exterior terrain & animals
	flower1 = rl.NewRectangle(1, 109, 16, 16)
	flower2 = rl.NewRectangle(25, 109, 16, 16)
	flower3 = rl.NewRectangle(48, 109, 16, 16)
	tree1   = rl.NewRectangle(72, 109, 16, 16)
	tree2   = rl.NewRectangle(96, 109, 16, 16)
	tree3   = rl.NewRectangle(121, 109, 16, 16)
	tree4   = rl.NewRectangle(138, 109, 16, 16)
	tree5   = rl.NewRectangle(162, 108, 16, 17)
	tree6   = rl.NewRectangle(187, 109, 16, 16)
	tree7   = rl.NewRectangle(275, 111, 16, 16)
	tree8   = rl.NewRectangle(291, 111, 16, 16)
	tree9   = rl.NewRectangle(307, 111, 16, 16)
	tree10  = rl.NewRectangle(323, 111, 16, 16)
	tree11  = rl.NewRectangle(339, 111, 16, 16)
	tree12  = rl.NewRectangle(355, 111, 16, 16)
	tree13  = rl.NewRectangle(372, 111, 16, 16)
	tree14  = rl.NewRectangle(387, 111, 16, 16)
	tree15  = rl.NewRectangle(405, 111, 16, 16)
	tree16  = rl.NewRectangle(421, 111, 16, 16)
	grass1  = rl.NewRectangle(0, 45, 16, 16)
	grass2  = rl.NewRectangle(212, 108, 16, 16)
	grass3  = rl.NewRectangle(235, 108, 16, 16)
	grass4  = rl.NewRectangle(259, 108, 16, 16)

	duckon, unicornon, caton, butterflyon, snailon, manon, snakeon, pigeonon, turtleon, rabbiton, pigon bool
	duck                                                                                                = rl.NewRectangle(19, 59, 16, 18)
	unicorn                                                                                             = rl.NewRectangle(35, 59, 20, 18)
	cat                                                                                                 = rl.NewRectangle(55, 59, 18, 18)
	butterfly                                                                                           = rl.NewRectangle(73, 59, 19, 17)
	snail                                                                                               = rl.NewRectangle(92, 60, 17, 15)
	man                                                                                                 = rl.NewRectangle(109, 58, 19, 18)
	snake                                                                                               = rl.NewRectangle(128, 58, 18, 18)
	pigeon                                                                                              = rl.NewRectangle(146, 59, 19, 18)
	turtle                                                                                              = rl.NewRectangle(166, 60, 16, 16)
	rabbit                                                                                              = rl.NewRectangle(184, 58, 19, 20)
	pig                                                                                                 = rl.NewRectangle(0, 61, 18, 15)
	// wall floor images
	floor1  = rl.NewRectangle(0, 29, 16, 16)
	floor2  = rl.NewRectangle(17, 28, 16, 16)
	floor3  = rl.NewRectangle(33, 28, 16, 16)
	floor4  = rl.NewRectangle(49, 28, 16, 16)
	floor5  = rl.NewRectangle(65, 28, 16, 16)
	floor6  = rl.NewRectangle(81, 28, 16, 16)
	floor7  = rl.NewRectangle(97, 28, 16, 16)
	floor8  = rl.NewRectangle(113, 28, 16, 16)
	floor9  = rl.NewRectangle(129, 28, 16, 16)
	floor10 = rl.NewRectangle(145, 28, 16, 16)
	floor11 = rl.NewRectangle(164, 28, 16, 16)
	floor12 = rl.NewRectangle(188, 28, 16, 16)
	floor13 = rl.NewRectangle(209, 28, 16, 16)
	floor14 = rl.NewRectangle(229, 28, 16, 16)
	floor15 = rl.NewRectangle(250, 28, 16, 16)
	wall1   = rl.NewRectangle(0, 13, 16, 16)
	wall2   = rl.NewRectangle(18, 12, 16, 16)
	wall3   = rl.NewRectangle(37, 12, 16, 16)
	wall4   = rl.NewRectangle(61, 12, 16, 16)
	wall5   = rl.NewRectangle(82, 12, 16, 16)
	wall6   = rl.NewRectangle(101, 12, 16, 16)
	wall7   = rl.NewRectangle(121, 12, 16, 16)
	wall8   = rl.NewRectangle(141, 12, 16, 16)
	// menu options
	colorson bool
	// menu
	selectinv1on, selectinv2on, selectinv3on, selectinv4on, selectinv5on, selectinv6on, selectinv7on, selectinv8on, selectinv9on, selectinv10on, selectinv11on, selectinv12on, selectinv13on, selectinv14on, selectinv15on, selectinv16on, selectinv17on, selectinv18on, selectinv19on, selectinv20on bool
	inv1on, inv2on, inv3on, inv4on, inv5on, inv6on, inv7on, inv8on, inv9on, inv10on, inv11on, inv12on, inv13on, inv14on, inv15on, inv16on, inv17on, inv18on, inv19on, inv20on                                                                                                                         bool
	optionsmenuclosehighlighton                                                                                                                                                                                                                                                                       bool
	optionsmenutabx, optionsmenutabx2, optionsmenutaby, optionsmenutaby2                                                                                                                                                                                                                              int32
	optionsmenuh, optionsmenuw, optionsmenuhspace, optionsmenuwspace                                                                                                                                                                                                                                  int32
	optionsmenuon, optionshighlighton, savehighlighton, movemenuhighlighton, movemenu2highlighton                                                                                                                                                                                                     bool
	movemenu2on, movemenuon                                                                                                                                                                                                                                                                           bool
	menutabx, menutaby, menutabx2, menutaby2, menu2tabx, menu2taby, menu2tabx2, menu2taby2                                                                                                                                                                                                            float32
	menux, menuy, menuw, menuh, menu2x, menu2y, menu2w, menu2h                                                                                                                                                                                                                                        int32
	mousetileinfo                                                                                                                                                                                                                                                                                     string
	movemenuimg                                                                                                                                                                                                                                                                                       = rl.NewRectangle(0, 76, 14, 14)
	closemenuimg                                                                                                                                                                                                                                                                                      = rl.NewRectangle(16, 77, 13, 15)
	// camera
	zoomlevel int
	// player
	player, playerh, playerv      int
	playerx, playery              int32
	playermoving, playerdirection bool
	// level
	tiletype                                                                                                                                                                                                                               string
	passagemin                                                                                                                                                                                                                             = 12
	passagemax                                                                                                                                                                                                                             = 25
	block, blockh, blockv, door, blockholder, roomnum, roomblockholder, roomlholder, roomwholder, roomaholder, rooml, roomh, rooma, passagel1, passagel2, passagea1, passagea2, passage1block, passage2block, passage3block, passage4block int
	path1, path2, path3, path4                                                                                                                                                                                                             bool
	drawgrida                                                                                                                                                                                                                              int
	drawblock, drawblocknext, drawblocknexth, drawblocknextv                                                                                                                                                                               int
	levelw                                                                                                                                                                                                                                 = 1280
	levelh                                                                                                                                                                                                                                 = 720
	levela                                                                                                                                                                                                                                 = levelw * levelh // set level size = 1280 * 720 (16:9) = 921 600
	minlevelh                                                                                                                                                                                                                              = levelh / 10
	minlevelw                                                                                                                                                                                                                              = levelw / 10
	// mouse
	cursoritem                                    rl.Rectangle
	cursorimg                                     = rl.NewRectangle(0, 0, 11, 11)
	selectxint                                    = int32(0)
	selectyint                                    = int32(0)
	cursorxint                                    = int32(0)
	cursoryint                                    = int32(0)
	selectedblocklines                            = int32(16)
	selectedblock, selectedblockh, selectedblockv int
	mouseblocklines                               = int32(16)
	mouseblock                                    int
	mousepos                                      rl.Vector2
	// core
	gridw, gridh, grida, gridw2, gridh2, grida2, gridw3, gridh3, grida3, gridw4, gridh4, grida4 int
	monh32, monw32                                                                              int32
	monitorh, monitorw                                                                          int
	debugon                                                                                     bool
	framecount                                                                                  int
	imgs                                                                                        rl.Texture2D
	camera                                                                                      rl.Camera2D
	cameracursor                                                                                rl.Camera2D
	camerainventory                                                                             rl.Camera2D
	cameraweather                                                                               rl.Camera2D
)

type items struct {
	itemtype, itemtype2, itemname                                    string
	coins, hp, damage, magic, durability, duration, random1, random2 int
	itemimg                                                          rl.Rectangle
}

/* notes

set level size = 1280 * 720 (16:9) = 921 600

*/
func raylib() { // MARK: raylib
	rl.InitWindow(monw32, monh32, "supa kastle qwest")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(30)
	rl.HideCursor()
	//	rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {
		mousepos = rl.GetMousePosition()
		framecount++
		if framecount%9000 == 0 {
			days++
		}
		if framecount%4500 == 0 {
			if nighton {
				if weathercurrent == "night" {
					weathercurrent = "sunny"
					weatherimg = sun
				}
				nighton = false
			} else {
				nighton = true
				if weathercurrent == "sunny" {
					weathercurrent = "night"
					weatherimg = moon
				}
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		drawx := int32(0)
		drawy := int32(0)
		count := 0
		drawblock = drawblocknext
		switch zoomlevel {
		case 1:
			drawgrida = grida
		case 2:
			drawgrida = grida2
		case 3:
			drawgrida = grida3
		case 4:
			drawgrida = grida4
		}

		for a := 0; a < drawgrida; a++ {
			checkitems := itemsmap[drawblock]
			checkitems2 := items2map[drawblock]
			checkitems3 := items3map[drawblock]
			checklevel := levelmap[drawblock]
			checkleveltiles := leveltilesmap[drawblock]
			checkextrasexterior := extrasexteriormap[drawblock]
			checkextrasinterior := extrasinteriormap[drawblock]
			checkitemactions := itemactionsmap[drawblock]
			checkenemies := enemiesmap[drawblock]
			checkeffects := effectsmap[drawblock]
			checkcharacters := charactersmap[drawblock]
			checkquest := questmap[drawblock]
			switch checkextrasexterior {
			case "grass":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, grass1, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, grass1, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree1, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree1, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree2, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree2, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree3, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree3, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree4, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree4, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree5":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree5, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree5, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree6":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree6, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree6, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree7":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree7, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree7, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree8":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree8, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree8, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree9":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree9, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree9, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree10":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree10, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree10, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree11":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree11, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree11, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree12":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree12, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree12, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree13":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree13, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree13, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree14":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree14, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree14, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree15":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree15, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree15, v2, rl.Fade(rl.White, 0.7))
				}
			case "tree16":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, tree16, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, tree16, v2, rl.Fade(rl.White, 0.7))
				}
			case "flower1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, flower1, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, flower1, v2, rl.Fade(rl.White, 0.7))
				}
			case "flower2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, flower2, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, flower2, v2, rl.Fade(rl.White, 0.7))
				}
			case "flower3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, flower3, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, flower3, v2, rl.Fade(rl.White, 0.7))
				}
			case "grass2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, grass2, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, grass2, v2, rl.Fade(rl.White, 0.7))
				}
			case "grass3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, grass3, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, grass3, v2, rl.Fade(rl.White, 0.7))
				}
			case "grass4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, grass4, v2, rl.Fade(rl.Green, 0.7))
				} else {
					rl.DrawTextureRec(imgs, grass4, v2, rl.Fade(rl.White, 0.7))
				}
			case "pig":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if pigon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, pig, v2, rl.Fade(rl.Pink, 0.7))
				} else {
					rl.DrawTextureRec(imgs, pig, v2, rl.Fade(rl.White, 0.7))
				}
			case "duck":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if duckon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, duck, v2, rl.Fade(rl.Yellow, 0.7))
				} else {
					rl.DrawTextureRec(imgs, duck, v2, rl.Fade(rl.White, 0.7))
				}
			case "unicorn":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if unicornon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, unicorn, v2, rl.Fade(rl.SkyBlue, 0.7))
				} else {
					rl.DrawTextureRec(imgs, unicorn, v2, rl.Fade(rl.White, 0.7))
				}
			case "cat":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if caton {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, cat, v2, rl.Fade(rl.Orange, 0.7))
				} else {
					rl.DrawTextureRec(imgs, cat, v2, rl.Fade(rl.White, 0.7))
				}
			case "butterfly":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if butterflyon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, butterfly, v2, rl.Fade(rl.Magenta, 0.7))
				} else {
					rl.DrawTextureRec(imgs, butterfly, v2, rl.Fade(rl.White, 0.7))
				}
			case "snail":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if snailon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, snail, v2, rl.Fade(rl.Brown, 0.7))
				} else {
					rl.DrawTextureRec(imgs, snail, v2, rl.Fade(rl.White, 0.7))
				}
			case "man":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if manon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, man, v2, rl.Fade(rl.Blue, 0.7))
				} else {
					rl.DrawTextureRec(imgs, man, v2, rl.Fade(rl.White, 0.7))
				}
			case "snake":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if snakeon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, snake, v2, rl.Fade(rl.Lime, 0.7))
				} else {
					rl.DrawTextureRec(imgs, snake, v2, rl.Fade(rl.White, 0.7))
				}
			case "pigeon":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if pigeonon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, pigeon, v2, rl.Fade(rl.Gray, 0.7))
				} else {
					rl.DrawTextureRec(imgs, pigeon, v2, rl.Fade(rl.White, 0.7))
				}
			case "turtle":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if turtleon {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, turtle, v2, rl.Fade(rl.DarkGreen, 0.7))
				} else {
					rl.DrawTextureRec(imgs, turtle, v2, rl.Fade(rl.White, 0.7))
				}
			case "rabbit":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if rabbiton {
					v2 = rl.NewVector2(float32(drawx), float32(drawy))
				} else {
					v2 = rl.NewVector2(float32(drawx), float32(drawy-2))
				}
				if colorson {
					rl.DrawTextureRec(imgs, rabbit, v2, rl.Fade(rl.Purple, 0.7))
				} else {
					rl.DrawTextureRec(imgs, rabbit, v2, rl.Fade(rl.White, 0.7))
				}
			}
			switch checklevel {
			case ".":
				//	rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.White, 0.1))

			case "test":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Red, 0.8))
			}
			switch checkleveltiles {
			case "floor", "center":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor1, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor1, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor2, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor2, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor3, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor3, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor4, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor4, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor5":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor5, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor5, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor6":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor6, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor6, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor7":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor7, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor7, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor8":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor8, v2, rl.Fade(rl.Maroon, 0.2))
				} else {
					rl.DrawTextureRec(imgs, floor8, v2, rl.Fade(rl.White, 0.1))
				}
			case "floor9":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor9, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor9, v2, rl.Fade(rl.White, fade1))
				}
			case "floor10":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor10, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor10, v2, rl.Fade(rl.White, fade1))
				}
			case "floor11":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor11, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor11, v2, rl.Fade(rl.White, fade1))
				}
			case "floor12":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor12, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor12, v2, rl.Fade(rl.White, fade1))
				}
			case "floor13":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor13, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor13, v2, rl.Fade(rl.White, fade1))
				}
			case "floor14":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor14, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor14, v2, rl.Fade(rl.White, fade1))
				}
			case "floor15":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if colorson {
					rl.DrawTextureRec(imgs, floor15, v2, rl.Fade(rl.Maroon, fade1))
				} else {
					rl.DrawTextureRec(imgs, floor15, v2, rl.Fade(rl.White, fade1))
				}
			case "wall":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall1, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall1, v2, rl.Orange)
				} else {
					rl.DrawTextureRec(imgs, wall1, v2, rl.White)
				}
			case "wall2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall2, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall2, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall2, v2, rl.White)
				}
			case "wall3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall3, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall3, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall3, v2, rl.White)
				}
			case "wall4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall4, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall4, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall4, v2, rl.White)
				}
			case "wall5":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall5, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall5, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall5, v2, rl.White)
				}
			case "wall6":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall6, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall6, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall6, v2, rl.White)
				}
			case "wall7":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall7, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall7, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall7, v2, rl.White)
				}
			case "wall8":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				v2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, wall8, v2shadow, rl.DarkGray)
				if colorson {
					rl.DrawTextureRec(imgs, wall8, v2, rl.Beige)
				} else {
					rl.DrawTextureRec(imgs, wall8, v2, rl.White)
				}
			case "door":
				//	rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Orange, 0.4))
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, door1, v2, rl.White)

			}
			switch checkextrasinterior {
			case "xcross":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xcross, v2, rl.Fade(rl.White, 0.4))
			case "xskullbones":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xskullbones, v2, rl.Fade(rl.White, 0.4))
			case "xskulls":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xskulls, v2, rl.Fade(rl.White, 0.4))
			case "xvase":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xvase, v2, rl.Fade(rl.White, 0.4))
			case "xfishbones":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xfishbones, v2, rl.Fade(rl.White, 0.4))
			case "xbonecross":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xbonecross, v2, rl.Fade(rl.White, 0.4))
			case "xknight":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xknight, v2, rl.Fade(rl.White, 0.4))
			case "xcastle":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xcastle, v2, rl.Fade(rl.White, 0.4))
			case "xbone1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xbone1, v2, rl.Fade(rl.White, 0.4))
			case "xskull1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xskull1, v2, rl.Fade(rl.White, 0.4))
			case "xskull2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xskull2, v2, rl.Fade(rl.White, 0.4))
			case "xcandle":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xcandle, v2, rl.Fade(rl.White, 0.4))
			case "xskull3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xskull3, v2, rl.Fade(rl.White, 0.4))
			case "xskull4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xskull4, v2, rl.Fade(rl.White, 0.4))
			case "xbone2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xbone2, v2, rl.Fade(rl.White, 0.4))
			case "xbone3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xbone3, v2, rl.Fade(rl.White, 0.4))
			case "xbone4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xbone4, v2, rl.Fade(rl.White, 0.4))
			case "xpotplant":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xpotplant, v2, rl.Fade(rl.White, 0.4))
			case "xsign1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xsign1, v2, rl.Fade(rl.White, 0.4))
			case "xsign2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xsign2, v2, rl.Fade(rl.White, 0.4))
			case "xlamp":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xlamp, v2, rl.Fade(rl.White, 0.4))
			case "xjug":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xjug, v2, rl.Fade(rl.White, 0.4))
			case "xvase2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xvase2, v2, rl.Fade(rl.White, 0.4))
			case "xhanger":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xhanger, v2, rl.Fade(rl.White, 0.4))
			case "xanchor":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, xanchor, v2, rl.Fade(rl.White, 0.4))

			}
			switch checkitemactions {
			case "dig":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, dug, v2, rl.White)
			}
			if checkitems != "" {
				checkitem := itemsstructmap[drawblock]
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, checkitem.itemimg, v2, rl.White)
			}
			if checkitems2 != "" {
				checkitem := itemsstructmap[drawblock]
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, checkitem.itemimg, v2, rl.White)
			}
			switch checkeffects {
			case "smoke1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, smoke1, v2, rl.White)
			case "smoke2":
				v2 := rl.NewVector2(float32(drawx+2), float32(drawy+2))
				rl.DrawTextureRec(imgs, smoke2, v2, rl.White)
			case "smoke3":
				v2 := rl.NewVector2(float32(drawx+4), float32(drawy+4))
				rl.DrawTextureRec(imgs, smoke3, v2, rl.White)
			}
			switch checkitems3 {
			case "chest":
				//	rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Red, 0.8))
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, chestclosed, v2, rl.White)
			case "chestopen":
				//	rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Red, 0.8))
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, chestopen, v2, rl.White)
			}

			// draw quest
			switch checkquest {
			case "quest":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Red, 0.8))
			case "tree":
				rl.DrawText("a tree", drawx+16, drawy, 10, rl.White)
			case "grass":
				rl.DrawText("some grass", drawx+16, drawy, 10, rl.White)
			case "flower":
				rl.DrawText("a flower", drawx+16, drawy, 10, rl.White)
			}
			// draw characters`
			switch checkcharacters {
			case "char1":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, char1, v2, rl.White)
			case "char2":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, char2, v2, rl.White)
			case "char3":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, char3, v2, rl.White)
			case "char4":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, char4, v2, rl.White)
			case "char5":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, char5, v2, rl.White)
			case "char6":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, char6, v2, rl.White)
			}

			// draw enemies
			if checkenemies != "" {
				checkenemy := enemiesstructmap[drawblock]
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, checkenemy.itemimg, v2, rl.White)
			}

			// draw player
			if player == drawblock {
				//	rl.DrawRectangle(drawx, drawy, 16, 16, rl.Fade(rl.Blue, 0.2))
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				if playermoving {
					if selectedblockv < playerv {
						rl.DrawTextureRec(imgs, playerrunL, v2, rl.White)
					} else {
						rl.DrawTextureRec(imgs, playerrunR, v2, rl.White)
					}
				} else {
					rl.DrawTextureRec(imgs, playeridleR, v2, rl.White)
				}
				playerx = drawx
				playery = drawy
			}
			// draw mouseblock
			if mouseblock == drawblock {
				if !optionsmenuon && !movemenu2on && !movemenuon {
					rl.DrawRectangle(drawx+cursorxint, drawy+cursoryint, mouseblocklines, mouseblocklines, rl.Fade(rl.White, 0.2))
					rl.DrawRectangleLines(drawx+cursorxint, drawy+cursoryint, mouseblocklines, mouseblocklines, rl.Fade(rl.White, 0.8))
					if framecount%4 == 0 {
						mouseblocklines -= 2
						cursoryint++
						cursorxint++
						if mouseblocklines == 2 {
							mouseblocklines = 16
							cursoryint = 0
							cursorxint = 0
						}
					}
				}
			}
			if selectedblock == drawblock && selectedblock != player {
				rl.DrawRectangleLines(drawx+selectxint, drawy+selectyint, selectedblocklines, selectedblocklines, rl.Fade(rl.White, 0.8))
				if framecount%4 == 0 {
					selectedblocklines -= 2
					selectxint++
					selectyint++
					if selectedblocklines == 2 {
						selectedblocklines = 16
						selectxint = 0
						selectyint = 0
					}
				}
			}
			// update animation blocks
			if framecount%2 == 0 {
				if checkeffects == "smoke1" {
					effectsmap[drawblock] = "smoke2"
				} else if checkeffects == "smoke2" {
					effectsmap[drawblock] = "smoke3"
				} else if checkeffects == "smoke3" {
					effectsmap[drawblock] = ""
				}
			}

			if framecount%10 == 0 {
				if checkenemies != "" {
					choose := rInt(1, 9)
					switch choose {
					case 1:
						if levelmap[drawblock-levelw-1] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock-levelw-1] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock-levelw-1] = "monster"
						}
					case 2:
						if levelmap[drawblock-levelw] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock-levelw] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock-levelw] = "monster"
						}
					case 3:
						if levelmap[drawblock-levelw+1] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock-levelw+1] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock-levelw+1] = "monster"
						}
					case 4:
						if levelmap[drawblock+1] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock+1] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock+1] = "monster"
						}
					case 5:
						if levelmap[drawblock+levelw+1] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock+levelw+1] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock+levelw+1] = "monster"
						}
					case 6:
						if levelmap[drawblock+levelw] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock+levelw] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock+levelw] = "monster"
						}
					case 7:
						if levelmap[drawblock+levelw-1] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock+levelw-1] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock+levelw-1] = "monster"
						}
					case 8:
						if levelmap[drawblock-1] == "floor" {
							enemiesmap[drawblock] = ""
							enemiesstructmap[drawblock-1] = enemiesstructmap[drawblock]
							enemiesstructmap[drawblock] = items{}
							enemiesmap[drawblock-1] = "monster"
						}
					}
				}
				if checkcharacters != "" {
					choose := rInt(1, 9)
					characterholder := charactersmap[drawblock]
					switch choose {
					case 1:
						if levelmap[drawblock-levelw-1] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock-levelw-1] = characterholder
						}
					case 2:
						if levelmap[drawblock-levelw] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock-levelw] = characterholder
						}
					case 3:
						if levelmap[drawblock-levelw+1] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock-levelw+1] = characterholder
						}
					case 4:
						if levelmap[drawblock+1] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock+1] = characterholder
						}
					case 5:
						if levelmap[drawblock+levelw+1] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock+levelw+1] = characterholder
						}
					case 6:
						if levelmap[drawblock+levelw] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock+levelw] = characterholder
						}
					case 7:
						if levelmap[drawblock+levelw-1] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock+levelw-1] = characterholder
						}
					case 8:
						if levelmap[drawblock-1] == "floor" {
							charactersmap[drawblock] = ""
							charactersmap[drawblock-1] = characterholder
						}
					}

				}

			}

			count++
			drawx += 16
			drawblock++
			switch zoomlevel {
			case 1:
				if count == gridw {
					count = 0
					drawx = 0
					drawy += 16
					drawblock += levelw - gridw
				}
			case 2:
				if count == gridw2 {
					count = 0
					drawx = 0
					drawy += 16
					drawblock += levelw - gridw2
				}
			case 3:
				if count == gridw3 {
					count = 0
					drawx = 0
					drawy += 16
					drawblock += levelw - gridw3
				}
			case 4:
				if count == gridw4 {
					count = 0
					drawx = 0
					drawy += 16
					drawblock += levelw - gridw4
				}
			}

		}
		rl.EndMode2D() // MARK: draw no camera

		// day night
		if nighton {
			if zoomlevel == 1 {
				v2 := rl.NewVector2(float32(playerx-175), float32(playery-175))
				rl.DrawRectangle(0, 0, monw32, playery-175, rl.Fade(rl.Black, 0.7))                          // top rectangle
				rl.DrawRectangle(0, (playery)+175, monw32, monh32/3+5, rl.Fade(rl.Black, 0.7))               // bottom rectangle
				rl.DrawRectangle(0, (playery)-175, (monw32/2)-175, 350, rl.Fade(rl.Black, 0.7))              // left rectangle
				rl.DrawRectangle((monw32/2)+175, (playery)-175, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7)) // right rectangle
				rl.DrawTextureRec(imgs, nightshadow, v2, rl.White)
			} else if zoomlevel == 2 {
				v2 := rl.NewVector2(float32((playerx*2)-160), float32((playery*2)-165))
				rl.DrawRectangle(0, 0, monw32, (playery*2)-165, rl.Fade(rl.Black, 0.7))                        // top rectangle
				rl.DrawRectangle(0, (playery*2)+185, monw32, monh32/3, rl.Fade(rl.Black, 0.7))                 // bottom rectangle
				rl.DrawRectangle(0, (playery*2)-165, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7))              // left rectangle
				rl.DrawRectangle((monw32/2)+190, (playery*2)-165, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7)) // right rectangle
				rl.DrawTextureRec(imgs, nightshadow, v2, rl.White)
			} else if zoomlevel == 3 {
				v2 := rl.NewVector2(float32((playerx*3)-160), float32((playery*3)-165))
				rl.DrawRectangle(0, 0, monw32, (playery*3)-165, rl.Fade(rl.Black, 0.7))                        // top rectangle
				rl.DrawRectangle(0, (playery*3)+185, monw32, monh32/3+10, rl.Fade(rl.Black, 0.7))              // bottom rectangle
				rl.DrawRectangle(0, (playery*3)-165, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7))              // left rectangle
				rl.DrawRectangle((monw32/2)+190, (playery*3)-165, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7)) // right rectangle
				rl.DrawTextureRec(imgs, nightshadow, v2, rl.White)
			} else if zoomlevel == 4 {
				v2 := rl.NewVector2(float32((playerx*4)-160), float32((playery*4)-165))
				rl.DrawRectangle(0, 0, monw32, (playery*4)-165, rl.Fade(rl.Black, 0.7))                        // top rectangle
				rl.DrawRectangle(0, (playery*4)+185, monw32, monh32/3+25, rl.Fade(rl.Black, 0.7))              // bottom rectangle
				rl.DrawRectangle(0, (playery*4)-165, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7))              // left rectangle
				rl.DrawRectangle((monw32/2)+190, (playery*4)-165, (monw32/2)-160, 350, rl.Fade(rl.Black, 0.7)) // right rectangle
				rl.DrawTextureRec(imgs, nightshadow, v2, rl.White)
			}

		}

		// draw weather
		rl.BeginMode2D(cameraweather)
		count2 := 0
		drawx2 := float32(0)
		drawy2 := float32(0)
		block := gridw2 * (gridh2 / 2)
		for a := 0; a < grida2; a++ {
			checkweather := weathermap[block]
			switch checkweather {
			case "snow":
				v2 := rl.NewVector2(drawx2, drawy2)
				rl.DrawTextureRec(imgs, snow, v2, rl.Fade(rl.White, 0.7))
			case "snow2":
				v2 := rl.NewVector2(drawx2, drawy2)
				rl.DrawTextureRec(imgs, snow2, v2, rl.Fade(rl.White, 0.7))
			case "snow3":
				v2 := rl.NewVector2(drawx2, drawy2)
				rl.DrawTextureRec(imgs, snow3, v2, rl.Fade(rl.White, 0.7))
			case "snow4":
				v2 := rl.NewVector2(drawx2, drawy2)
				rl.DrawTextureRec(imgs, snow4, v2, rl.Fade(rl.White, 0.7))
			}
			count2++
			block++
			drawx2 += 32
			if count2 == gridw2 {
				drawx2 = 0
				drawy2 += 32
				count2 = 0
			}
		}
		rl.EndMode2D()

		input()
		updateall()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
func createlevel() { // MARK: createlevel
	for a := 0; a < levela; a++ {
		levelmap[a] = "."
		if rolldice()+rolldice()+rolldice() >= 17 {
			if flipcoin() {
				extrasexteriormap[a] = "grass"
			} else {
				choose := rInt(1, 23)
				switch choose {
				case 1:
					extrasexteriormap[a] = "tree1"
				case 2:
					extrasexteriormap[a] = "tree2"
				case 3:
					extrasexteriormap[a] = "tree3"
				case 4:
					extrasexteriormap[a] = "tree4"
				case 5:
					extrasexteriormap[a] = "tree5"
				case 6:
					extrasexteriormap[a] = "tree6"
				case 7:
					extrasexteriormap[a] = "tree7"
				case 8:
					extrasexteriormap[a] = "tree8"
				case 9:
					extrasexteriormap[a] = "tree9"
				case 10:
					extrasexteriormap[a] = "tree10"
				case 11:
					extrasexteriormap[a] = "tree11"
				case 12:
					extrasexteriormap[a] = "tree12"
				case 13:
					extrasexteriormap[a] = "tree13"
				case 14:
					extrasexteriormap[a] = "tree14"
				case 15:
					extrasexteriormap[a] = "tree15"
				case 16:
					extrasexteriormap[a] = "tree16"
				case 17:
					extrasexteriormap[a] = "flower1"
				case 18:
					extrasexteriormap[a] = "flower2"
				case 19:
					extrasexteriormap[a] = "flower3"
				case 20:
					extrasexteriormap[a] = "grass2"
				case 21:
					extrasexteriormap[a] = "grass3"
				case 22:
					extrasexteriormap[a] = "grass4"
				}

			}
		}
		if rolldice()+rolldice()+rolldice()+rolldice() == 24 {
			choose := rInt(1, 12)
			switch choose {
			case 1:
				extrasexteriormap[a] = "pig"
			case 2:
				extrasexteriormap[a] = "duck"
			case 3:
				extrasexteriormap[a] = "unicorn"
			case 4:
				extrasexteriormap[a] = "cat"
			case 5:
				extrasexteriormap[a] = "butterfly"
			case 6:
				extrasexteriormap[a] = "snail"
			case 7:
				extrasexteriormap[a] = "man"
			case 8:
				extrasexteriormap[a] = "snake"
			case 9:
				extrasexteriormap[a] = "pigeon"
			case 10:
				extrasexteriormap[a] = "turtle"
			case 11:
				extrasexteriormap[a] = "rabbit"
			}
		}
	}
	// player position relative to starting block
	block = levelw * (levelh / 2)
	block += levelw / 2
	player = block
	player += 4
	player += levelw * 4

	var newitem items
	newitem.itemtype2 = "pickaxe"
	newitem.itemtype = "pickaxe"
	newitem.itemimg = pickaxe
	newitem.coins = rolldice()
	newitem.hp = rolldice()
	newitem.damage = rolldice()
	newitem.itemname = createname("pickaxe", newitem.damage)
	newitem.magic = rolldice()
	newitem.durability = rInt(50, 101)
	newitem.duration = rInt(30, 121)
	newitem.random1 = rolldice() + rolldice()
	newitem.random2 = rolldice() + rolldice()

	items2map[player+2] = "pickaxe"
	itemsstructmap[player+2] = newitem

	var newitem2 items
	itemimg := chooseimg("monster")
	newitem2.itemtype = "monster"
	newitem2.itemtype2 = "monster"
	newitem2.itemname = "monster name here"
	newitem2.itemimg = itemimg
	newitem2.coins = rolldice()
	newitem2.hp = rolldice()
	newitem2.damage = rolldice()
	newitem2.magic = rolldice()
	newitem2.durability = rInt(50, 101)
	newitem2.duration = rInt(30, 121)
	newitem2.random1 = rolldice() + rolldice()
	newitem2.random2 = rolldice() + rolldice()

	enemiesmap[player+4] = "monster"
	enemiesstructmap[player+4] = newitem2

	drawblocknext = player
	drawblocknext -= (gridh2 / 2) * levelw
	drawblocknext -= gridw2 / 2

	// create room layout
	roomnum = 1000
	for {
		createroom()
		if roomnum <= 0 {
			break
		}
	}
	createwalls()
	createtiles()
	createcharacters()
}
func createcharacters() { // MARK: createcharacters

	charactercount := 100

	for {

		charblock := rInt(levelw*4, levela-(levelw*4))

		if levelmap[charblock] == "floor" {
			switch rolldice() {
			case 1:
				charactersmap[charblock] = "char1"
			case 2:
				charactersmap[charblock] = "char2"
			case 3:
				charactersmap[charblock] = "char3"
			case 4:
				charactersmap[charblock] = "char4"
			case 5:
				charactersmap[charblock] = "char5"
			case 6:
				charactersmap[charblock] = "char6"
			}

			charactercount--

		}

		if charactercount == 0 {
			break
		}
	}

}
func createwalls() { // MARK: createwalls
	for a := 0; a < levela; a++ {
		if levelmap[a] == "floor" && levelmap[a-levelw] == "." {
			levelmap[a-levelw] = "wall"
			leveltilesmap[a-levelw] = "wall"
		}
		if levelmap[a] == "floor" && levelmap[a+1] == "." {
			levelmap[a+1] = "wall"
			leveltilesmap[a+1] = "wall"
		}
		if levelmap[a] == "floor" && levelmap[a-1] == "." {
			levelmap[a-1] = "wall"
			leveltilesmap[a-1] = "wall"
		}
		if levelmap[a] == "floor" && levelmap[a+levelw] == "." {
			levelmap[a+levelw] = "wall"
			leveltilesmap[a+levelw] = "wall"
		}
	}
	for a := 0; a < levela; a++ {
		if levelmap[a] == "wall" && levelmap[a+1] == "." && levelmap[a+1+levelw] == "wall" {
			levelmap[a+1] = "wall"
			leveltilesmap[a+1] = "wall"
		}
		if levelmap[a] == "wall" && levelmap[a-levelw] == "." && levelmap[a+1-levelw] == "wall" {
			levelmap[a-levelw] = "wall"
			leveltilesmap[a-levelw] = "wall"
		}
		if levelmap[a] == "wall" && levelmap[a+1] == "." && levelmap[a-levelw] == "floor" && levelmap[a+1-levelw] == "wall" {
			levelmap[a+1] = "wall"
			leveltilesmap[a+1] = "wall"
		}
		if levelmap[a] == "wall" && levelmap[a-1] == "." && levelmap[a+1] == "floor" && levelmap[a-1-levelw] == "wall" {
			levelmap[a-1] = "wall"
			leveltilesmap[a-1] = "wall"
		}
	}
	for a := 0; a < levela; a++ {
		if levelmap[a] == "wall" && levelmap[a-levelw] == "floor" && levelmap[a-1] == "." && levelmap[a+1] == "wall" && levelmap[a+levelw] == "." {
			levelmap[a-1] = "wall"
			leveltilesmap[a-1] = "wall"
		}
	}
	for a := 0; a < levela; a++ {
		if leveltilesmap[a] == "center" {
			choose := rInt(1, 5)
			checkblock := a
			switch choose {
			case 1:
				for b := 0; b < minlevelh/2; b++ {
					if levelmap[checkblock] == "wall" {
						levelmap[checkblock] = "door"
						leveltilesmap[checkblock] = "door"
						break
					}
					checkblock -= levelw
					checkblockh := checkblock / levelw
					if checkblockh < 5 {
						break
					}
				}
			case 2:
				for b := 0; b < minlevelw/2; b++ {
					if levelmap[checkblock] == "wall" {
						levelmap[checkblock] = "door"
						leveltilesmap[checkblock] = "door"
						break
					}
					checkblock++
					checkblockv := checkblock % levelw
					if checkblockv > levelw-5 {
						break
					}
				}
			case 3:
				for b := 0; b < minlevelh/2; b++ {
					if levelmap[checkblock] == "wall" {
						levelmap[checkblock] = "door"
						leveltilesmap[checkblock] = "door"
						break
					}
					checkblock += levelw
					checkblockh := checkblock / levelw
					if checkblockh > levelh-5 {
						break
					}
				}
			case 4:
				for b := 0; b < minlevelw/2; b++ {
					if levelmap[checkblock] == "wall" {
						levelmap[checkblock] = "door"
						leveltilesmap[checkblock] = "door"
						break
					}
					checkblock--
					checkblockv := checkblock % levelw
					if checkblockv < 5 {
						break
					}
				}
			}
		}
	}

	// interior walls
	count := 200
	for {
		wallblock := rInt(0, levela)
		if levelmap[wallblock] == "floor" {
			count2 := 10
			wallblockholder := wallblock
			obstruction := false
			for a := 0; a < 100; a++ {
				if levelmap[wallblock] != "floor" {
					obstruction = true
				}
				wallblock++
				count2--
				if count2 == 0 {
					count2 = 10
					wallblock += levelw - 10
				}
			}
			wallblock = wallblockholder
			if !obstruction {

				choose := rolldice()

				wall1l := rInt(5, 11)
				wall2l := rInt(5, 11)

				switch choose {
				case 1: // cross walls
					choosetile("wall")
					wallblock += levelw * (wall1l / 2)
					for a := 0; a < wall1l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					wallblock = wallblockholder
					wallblock += wall1l / 2
					for a := 0; a < wall1l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
				case 2: // room door up
					// draw different floor tiles
					choosetile("roomfloor")
					floora := wall1l * wall2l
					count := 0
					for a := 0; a < floora; a++ {
						leveltilesmap[wallblock] = tiletype
						wallblock++
						count++
						if count == wall2l {
							count = 0
							wallblock += levelw - wall2l
						}
					}
					wallblock = wallblockholder
					choosetile("wall")
					for a := 0; a < wall1l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
					for a := 0; a < wall2l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					for a := 0; a < wall1l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock -= levelw
					}
					for a := 0; a < wall2l-2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock--
					}
					wallblock += rInt(2, wall1l-2) * levelw
					wallblock += rInt(0, wall2l-2)
					levelmap[wallblock] = "chest"
					items3map[wallblock] = "chest" // loot chest
				case 3: // room door right
					// draw different floor tiles
					choosetile("roomfloor")
					floora := wall1l * wall2l
					count := 0
					wallblock++
					for a := 0; a < floora; a++ {
						leveltilesmap[wallblock] = tiletype
						wallblock++
						count++
						if count == wall1l {
							count = 0
							wallblock += levelw - wall1l
						}
					}
					wallblock = wallblockholder
					choosetile("wall")
					for a := 0; a < wall1l+1; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					wallblock = wallblockholder
					for a := 0; a < wall2l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
					for a := 0; a < wall1l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					for a := 0; a < wall2l-2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock -= levelw
					}
					wallblock -= rInt(2, wall1l-2)
					wallblock += rInt(0, wall2l-2) * levelw
					levelmap[wallblock] = "chest"
					items3map[wallblock] = "chest" // loot chest
				case 4: // room door down
					// draw different floor tiles
					choosetile("roomfloor")
					floora := wall1l * wall2l
					count := 0
					wallblock += levelw
					for a := 0; a < floora; a++ {
						leveltilesmap[wallblock] = tiletype
						wallblock++
						count++
						if count == wall1l {
							count = 0
							wallblock += levelw - wall1l
						}
					}
					wallblock = wallblockholder
					choosetile("wall")
					for a := 0; a < wall1l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					for a := 0; a < wall2l+1; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
					wallblock = wallblockholder
					for a := 0; a < wall2l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
					for a := 0; a < wall1l-2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					wallblock -= rInt(2, wall2l-2) * levelw
					wallblock -= rInt(0, wall2l-2)
					levelmap[wallblock] = "chest"
					items3map[wallblock] = "chest" // loot chest
				case 5: // room door left
					// draw different floor tiles
					choosetile("roomfloor")
					floora := (wall1l + 1) * wall2l
					count := 0
					for a := 0; a < floora; a++ {
						leveltilesmap[wallblock] = tiletype
						wallblock++
						count++
						if count == wall1l+1 {
							count = 0
							wallblock += levelw - (wall1l + 1)
						}
					}
					wallblock = wallblockholder
					choosetile("wall")
					for a := 0; a < wall1l+1; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					for a := 0; a < wall2l; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
					for a := 0; a < wall1l+1; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock--
					}
					for a := 0; a < wall2l-2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock -= levelw
					}
					wallblock += rInt(2, wall1l-2)
					wallblock += rInt(0, wall2l-2) * levelw
					levelmap[wallblock] = "chest"
					items3map[wallblock] = "chest" // loot chest
				case 6: // z wall
					choosetile("wall")
					for a := 0; a < wall1l/2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					for a := 0; a < wall2l/2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock += levelw
					}
					for a := 0; a < wall1l/2; a++ {
						levelmap[wallblock] = "wall"
						leveltilesmap[wallblock] = tiletype
						wallblock++
					}
					count--
				}
			}
		}
		if count <= 0 {
			break
		}
	}

}
func createroom() { // MARK: createroom
	roomtype := rInt(1, 6)
	maxl := 45
	minl := 10
	rooml = rInt(minl, maxl)
	roomh = rInt(minl, maxl)
	rooma = roomh * rooml
	count := 0

	switch roomtype {
	case 1: // square
		rooma = rooml * rooml
		centerblock := block + (rooml / 2)
		centerblock += (rooml / 2) * levelw
		for a := 0; a < rooma; a++ {
			levelmap[block] = "floor"
			leveltilesmap[block] = "floor"
			extrasexteriormap[block] = ""
			block++
			count++
			if count == rooml {
				count = 0
				block += levelw - rooml
			}
		}
		leveltilesmap[centerblock] = "center"
	case 2: // rectangle
		centerblock := block + (rooml / 2)
		centerblock += (roomh / 2) * levelw
		for a := 0; a < rooma; a++ {
			levelmap[block] = "floor"
			leveltilesmap[block] = "floor"
			extrasexteriormap[block] = ""
			block++
			count++
			if count == rooml {
				count = 0
				block += levelw - rooml
			}
		}
		leveltilesmap[centerblock] = "center"

	case 3: // diamond
		centerblock := block + ((rooml / 2) * levelw)
		rooml2 := 1
		for {
			for a := 0; a < rooml2; a++ {
				levelmap[block] = "floor"
				leveltilesmap[block] = "floor"
				extrasexteriormap[block] = ""
				block++
			}
			rooml2 += 2
			block += levelw
			block -= rooml2 - 1
			if rooml2 >= rooml {
				break
			}
		}
		rooml2 = rooml
		for {
			for a := 0; a < rooml2; a++ {
				levelmap[block] = "floor"
				leveltilesmap[block] = "floor"
				extrasexteriormap[block] = ""
				block++
			}
			rooml2 -= 2
			block += levelw
			block -= rooml2 + 1
			if rooml2 <= 1 {
				break
			}
		}
		leveltilesmap[centerblock] = "center"
	case 4: // triangle
		centerblock := block + ((rooml / 2) * levelw)
		rooml2 := 1
		for {
			for a := 0; a < rooml2; a++ {
				levelmap[block] = "floor"
				leveltilesmap[block] = "floor"
				extrasexteriormap[block] = ""
				block++
			}
			rooml2 += 2
			block += levelw
			block -= rooml2 - 1
			if rooml2 >= rooml {
				break
			}
		}
		rooml2 = rooml
		for {
			for a := 0; a < rooml2; a++ {
				levelmap[block] = "floor"
				leveltilesmap[block] = "floor"
				extrasexteriormap[block] = ""
				block++
			}
			rooml2 -= 2
			block += levelw
			block -= rooml2 - 1
			if rooml2 <= 1 {
				break
			}
		}
		leveltilesmap[centerblock] = "center"
	case 5: // hollow square
		rooma = rooml * rooml
		centerblock := block + (rooml / 2)
		centerblock += +2 * levelw
		blockholder = block
		for a := 0; a < rooma; a++ {
			levelmap[block] = "floor"
			leveltilesmap[block] = "floor"
			extrasexteriormap[block] = ""
			block++
			count++
			if count == rooml {
				count = 0
				block += levelw - rooml
			}
		}
		rooml = rooml / 2
		rooma = rooml * rooml
		block = blockholder
		block += rooml / 2
		block += (rooml / 2) * levelw
		for a := 0; a < rooma; a++ {
			levelmap[block] = "."
			block++
			count++
			if count == rooml {
				count = 0
				block += levelw - rooml
			}
		}
		leveltilesmap[centerblock] = "center"

	} // end switch roomtype

	roomnum--

	// next room starting block position
	nextblock := rInt(0, levela)
	block = nextblock
	horizvert()
	if blockh <= 5 {
		block += (rInt(maxl, maxl*2)) * levelw
		horizvert()
	}
	if blockh > levelh-minlevelh {
		block -= (rInt(maxl, maxl*2)) * levelw
		horizvert()
	}
	if blockv <= 5 {
		block += rInt(maxl, maxl*2)
		horizvert()
	}
	if blockv > levelw-minlevelw {
		block -= rInt(maxl, maxl*2)
		horizvert()
	}

}
func createtiles() { // MARK: createtiles

	for a := 0; a < levela; a++ {
		if leveltilesmap[a] == "floor" {
			if rolldice()+rolldice()+rolldice() == 18 {
				choosetile("floor")
				leveltilesmap[a] = tiletype
			}
			if rolldice()+rolldice()+rolldice() == 17 {
				choose := rInt(1, 26)

				switch choose {
				case 1:
					extrasinteriormap[a] = "xcross"
				case 2:
					extrasinteriormap[a] = "xskullbones"
				case 3:
					extrasinteriormap[a] = "xskulls"
				case 4:
					extrasinteriormap[a] = "xvase"
				case 5:
					extrasinteriormap[a] = "xfishbones"
				case 6:
					extrasinteriormap[a] = "xbonecross"
				case 7:
					extrasinteriormap[a] = "xknight"
				case 8:
					extrasinteriormap[a] = "xcastle"
				case 9:
					extrasinteriormap[a] = "xbone1"
				case 10:
					extrasinteriormap[a] = "xskull1"
				case 11:
					extrasinteriormap[a] = "xskull2"
				case 12:
					extrasinteriormap[a] = "xcandle"
				case 13:
					extrasinteriormap[a] = "xskull3"
				case 14:
					extrasinteriormap[a] = "xskull4"
				case 15:
					extrasinteriormap[a] = "xbone2"
				case 16:
					extrasinteriormap[a] = "xbone3"
				case 17:
					extrasinteriormap[a] = "xbone4"
				case 18:
					extrasinteriormap[a] = "xpotplant"
				case 19:
					extrasinteriormap[a] = "xsign1"
				case 20:
					extrasinteriormap[a] = "xsign2"
				case 21:
					extrasinteriormap[a] = "xlamp"
				case 22:
					extrasinteriormap[a] = "xjug"
				case 23:
					extrasinteriormap[a] = "xvase2"
				case 24:
					extrasinteriormap[a] = "xhanger"
				case 25:
					extrasinteriormap[a] = "xanchor"
				}
			}
		}

	}

}
func choosetile(blocktype string) { // MARK: choosetile
	switch blocktype {
	case "floor":
		choose := rInt(1, 8)
		switch choose {
		case 1:
			tiletype = "floor2"
		case 2:
			tiletype = "floor3"
		case 3:
			tiletype = "floor4"
		case 4:
			tiletype = "floor5"
		case 5:
			tiletype = "floor6"
		case 6:
			tiletype = "floor7"
		case 7:
			tiletype = "floor8"
		}
	case "roomfloor":
		choose := rInt(1, 8)
		switch choose {
		case 1:
			tiletype = "floor9"
		case 2:
			tiletype = "floor10"
		case 3:
			tiletype = "floor11"
		case 4:
			tiletype = "floor12"
		case 5:
			tiletype = "floor13"
		case 6:
			tiletype = "floor14"
		case 7:
			tiletype = "floor15"
		}
	case "wall":
		choose := rInt(1, 8)
		switch choose {
		case 1:
			tiletype = "wall2"
		case 2:
			tiletype = "wall3"
		case 3:
			tiletype = "wall4"
		case 4:
			tiletype = "wall5"
		case 5:
			tiletype = "wall6"
		case 6:
			tiletype = "wall7"
		case 7:
			tiletype = "wall8"
		}
	}

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides info window
	initialize()
	raylib()
}
func weather() { // MARK: weather

	if framecount%30 == 0 {
		weathercount++
	}

	if weathercount == weathertimer {
		for a := 0; a < len(weathermap); a++ {
			weathermap[a] = ""
		}
		if rolldice() < 5 {
			if nighton {
				weathercurrent = "night"
				weatherimg = moon
			} else {
				weathercurrent = "sunny"
				weatherimg = sun
			}
		} else {
			choose := rInt(1, 5)
			switch choose {
			case 1:
				weathercurrent = "snowing"
				weatherimg = snow
			case 2:
				weathercurrent = "raining"
				weatherimg = rain
			case 3:
				weathercurrent = "windy"
				weatherimg = wind
			case 4:
				weathercurrent = "stormy"
				weatherimg = storm
			}
		}

		weathertimer = rInt(10, 20)
		weathercount = 0
		switch weathercurrent {
		case "snowing":

			for a := 0; a < len(weathermap); a++ {
				weathermap[a] = ""
			}
			for a := 0; a < len(weathermap); a++ {
				if rolldice() == 6 {
					choose := rInt(1, 5)
					switch choose {
					case 1:
						weathermap[a] = "snow"
					case 2:
						weathermap[a] = "snow2"
					case 3:
						weathermap[a] = "snow3"
					case 4:
						weathermap[a] = "snow4"

					}

				}
			}

		}
	}

	if framecount%6 == 0 {
		switch weathercurrent {
		case "snowing":
			for a := len(weathermap) - 1; a >= 0; a-- {
				if weathermap[a] != "" {
					if a < (len(weathermap) - gridw2*2) {
						holder := weathermap[a]
						weathermap[a] = ""
						weathermap[a+gridw2] = holder
					} else {
						weathermap[a] = ""
					}
				}
			}
		}
		for a := 0; a < gridw2; a++ {
			if rolldice() == 6 {
				choose := rInt(1, 5)
				switch choose {
				case 1:
					weathermap[a] = "snow"
				case 2:
					weathermap[a] = "snow2"
				case 3:
					weathermap[a] = "snow3"
				case 4:
					weathermap[a] = "snow4"

				}

			}
		}
	}

}
func updateall() { // MARK: updateall
	horizvert()
	getmouseblock()
	getmousepos()
	moveplayer()
	pickup()
	weather()
	movecamera()
	fx()
	menus()
	inventory()
	if debugon {
		debug()
	}
	if optionsmenuon {
		menuoptions()
	}
	animations()
	cursor()
}
func useitem() { // MARK: useitem

	switch invitemactive {
	case "spade":
		if levelmap[itemactiveblock] == "floor" {
			itemactionsmap[itemactiveblock] = "dig"
			itemaction("dig")
		}
	case "pickaxe":
		if levelmap[itemactiveblock] == "wall" {
			itemactionsmap[itemactiveblock] = "breakblock"
			itemaction("breakblock")
		}
	}

}
func itemaction(action string) { // MARK: itemaction

	switch action {
	case "breakblock":
		if levelmap[itemactiveblock] == "wall" {
			levelmap[itemactiveblock] = "floor"
			leveltilesmap[itemactiveblock] = "floor"
			effectsmap[itemactiveblock] = "smoke1"
		}
	case "dig":
		if rolldice()+rolldice()+rolldice()+rolldice()+rolldice()+rolldice() == 36 {
			items2map[itemactiveblock+1] = "test"
		} else {
			if rolldice()+rolldice()+rolldice()+rolldice()+rolldice() == 30 {
				items2map[itemactiveblock+1] = "test"
			} else if rolldice()+rolldice()+rolldice()+rolldice() == 24 {
				items2map[itemactiveblock+1] = "test"
			} else if rolldice()+rolldice()+rolldice() == 18 {
				items2map[itemactiveblock+1] = "test"
			} else if rolldice()+rolldice() == 12 {
				items2map[itemactiveblock+1] = "test"
			}
		}

	}

}
func inventory() { // MARK: inventory
	if invitemactiveon && invitemstructactive.itemtype2 == "boots" && bootssloton {
		placeboots = true
	} else {
		placeboots = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "gloves" && glovessloton {
		placegloves = true
	} else {
		placegloves = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "belt" && beltsloton {
		placebelt = true
	} else {
		placebelt = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "armor" && armorsloton {
		placearmor = true
	} else {
		placearmor = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "ring" && ring2sloton {
		placering2 = true
	} else {
		placering2 = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "ring" && ring1sloton {
		placering1 = true
	} else {
		placering1 = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "armor" && armorsloton {
		placearmor = true
	} else {
		placearmor = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "necklace" && necklacesloton {
		placenecklace = true
	} else {
		placenecklace = false
	}
	if invitemactiveon && invitemstructactive.itemtype2 == "helmet" && helmetsloton {
		placehelmet = true
	} else {
		placehelmet = false
	}
	if invitemactiveon && invitemstructactive.itemtype == "weapon" && weaponslot1on || invitemactiveon && invitemstructactive.itemtype2 == "shield" && weaponslot1on {
		placeweapon1 = true
	} else {
		placeweapon1 = false
	}
	if invitemactiveon && invitemstructactive.itemtype == "weapon" && weaponslot2on || invitemactiveon && invitemstructactive.itemtype2 == "shield" && weaponslot2on {
		placeweapon2 = true
	} else {
		placeweapon2 = false
	}

	for a := 0; a < len(inventorymap); a++ {
		if inventorymap[a] != "" {
			drawitem := rl.NewRectangle(0, 0, 0, 0)
			checkitem := inventorystructsmap[a]
			drawitem = checkitem.itemimg

			rl.BeginMode2D(camerainventory)
			if a < 10 {
				xpos := float32(invxmap[a])
				xpos = xpos / 2
				xpos++
				ypos := float32((menu2y + 4) / 2)
				ypos++
				v2 := rl.NewVector2(xpos, ypos)
				rl.DrawTextureRec(imgs, drawitem, v2, rl.White)
			} else {
				xpos := float32(invxmap[a-10])
				xpos = xpos / 2
				xpos++
				ypos := float32((menu2y + 40) / 2)
				ypos++
				v2 := rl.NewVector2(xpos, ypos)
				rl.DrawTextureRec(imgs, drawitem, v2, rl.White)
			}
			rl.EndMode2D()
		}
	}

	if selectinv1on {
		if inventoryfullmap[0] {
			invitemactive = inventorymap[0]
			invitemstructactive = inventorystructsmap[0]
			selectedslotnumber = 0
			invitemactiveon = true
		}
	} else if selectinv2on {
		if inventoryfullmap[1] {
			invitemactive = inventorymap[1]
			invitemstructactive = inventorystructsmap[1]
			selectedslotnumber = 1
			invitemactiveon = true
		}
	} else if selectinv3on {
		if inventoryfullmap[2] {
			invitemactive = inventorymap[2]
			invitemstructactive = inventorystructsmap[2]
			selectedslotnumber = 2
			invitemactiveon = true
		}
	} else if selectinv4on {
		if inventoryfullmap[3] {
			invitemactive = inventorymap[3]
			invitemstructactive = inventorystructsmap[3]
			invitemactiveon = true
		}
	} else if selectinv5on {
		if inventoryfullmap[4] {
			invitemactive = inventorymap[4]
			invitemstructactive = inventorystructsmap[4]
			selectedslotnumber = 4
			invitemactiveon = true
		}
	} else if selectinv6on {
		if inventoryfullmap[5] {
			invitemactive = inventorymap[5]
			invitemstructactive = inventorystructsmap[5]
			selectedslotnumber = 5
			invitemactiveon = true
		}
	} else if selectinv7on {
		if inventoryfullmap[6] {
			invitemactive = inventorymap[6]
			invitemstructactive = inventorystructsmap[6]
			selectedslotnumber = 6
			invitemactiveon = true
		}
	} else if selectinv8on {
		if inventoryfullmap[7] {
			invitemactive = inventorymap[7]
			invitemstructactive = inventorystructsmap[7]
			selectedslotnumber = 7
			invitemactiveon = true
		}
	} else if selectinv9on {
		if inventoryfullmap[8] {
			invitemactive = inventorymap[8]
			invitemstructactive = inventorystructsmap[8]
			selectedslotnumber = 8
			invitemactiveon = true
		}
	} else if selectinv10on {
		if inventoryfullmap[9] {
			invitemactive = inventorymap[9]
			invitemstructactive = inventorystructsmap[9]
			selectedslotnumber = 9
			invitemactiveon = true
		}
	} else if selectinv11on {
		if inventoryfullmap[10] {
			invitemactive = inventorymap[10]
			invitemstructactive = inventorystructsmap[10]
			selectedslotnumber = 10
			invitemactiveon = true
		}
	} else if selectinv12on {
		if inventoryfullmap[11] {
			invitemactive = inventorymap[11]
			invitemstructactive = inventorystructsmap[11]
			selectedslotnumber = 11
			invitemactiveon = true
		}
	} else if selectinv13on {
		if inventoryfullmap[12] {
			invitemactive = inventorymap[12]
			invitemstructactive = inventorystructsmap[12]
			selectedslotnumber = 12
			invitemactiveon = true
		}
	} else if selectinv14on {
		if inventoryfullmap[13] {
			invitemactive = inventorymap[13]
			invitemstructactive = inventorystructsmap[13]
			selectedslotnumber = 13
			invitemactiveon = true
		}
	} else if selectinv15on {
		if inventoryfullmap[14] {
			invitemactive = inventorymap[14]
			invitemstructactive = inventorystructsmap[14]
			selectedslotnumber = 14
			invitemactiveon = true
		}
	} else if selectinv16on {
		if inventoryfullmap[15] {
			invitemactive = inventorymap[15]
			invitemstructactive = inventorystructsmap[15]
			selectedslotnumber = 15
			invitemactiveon = true
		}
	} else if selectinv17on {
		if inventoryfullmap[16] {
			invitemactive = inventorymap[16]
			invitemstructactive = inventorystructsmap[16]
			selectedslotnumber = 16
			invitemactiveon = true
		}
	} else if selectinv18on {
		if inventoryfullmap[17] {
			invitemactive = inventorymap[17]
			invitemstructactive = inventorystructsmap[17]
			selectedslotnumber = 17
			invitemactiveon = true
		}
	} else if selectinv19on {
		if inventoryfullmap[18] {
			invitemactive = inventorymap[18]
			invitemstructactive = inventorystructsmap[18]
			selectedslotnumber = 18
			invitemactiveon = true
		}
	} else if selectinv20on {
		if inventoryfullmap[19] {
			invitemactive = inventorymap[19]
			invitemstructactive = inventorystructsmap[19]
			selectedslotnumber = 19
			invitemactiveon = true
		}
	}

	if invitemactiveon {
		cursoritem = invitemstructactive.itemimg
	}

	if inv1on {
		if inventoryfullmap[0] {
			checkitem := inventorystructsmap[0]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[0]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[0]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[0]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[0]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv2on {
		if inventoryfullmap[1] {
			checkitem := inventorystructsmap[1]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[1]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[1]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[1]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[1]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv3on {
		if inventoryfullmap[2] {
			checkitem := inventorystructsmap[2]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[2]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[2]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[2]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[2]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv4on {
		if inventoryfullmap[3] {
			checkitem := inventorystructsmap[3]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[3]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[3]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[3]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[3]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv5on {
		if inventoryfullmap[4] {
			checkitem := inventorystructsmap[4]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[4]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[4]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[4]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[4]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv6on {
		if inventoryfullmap[5] {
			checkitem := inventorystructsmap[5]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[5]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[5]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[5]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[5]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv7on {
		if inventoryfullmap[6] {
			checkitem := inventorystructsmap[6]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[6]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[6]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[6]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[6]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv8on {
		if inventoryfullmap[7] {
			checkitem := inventorystructsmap[7]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[7]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[7]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[7]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[7]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv9on {
		if inventoryfullmap[8] {
			checkitem := inventorystructsmap[8]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[8]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[8]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[8]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[8]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv10on {
		if inventoryfullmap[9] {
			checkitem := inventorystructsmap[9]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[9]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[9]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[9]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[9]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv11on {
		if inventoryfullmap[10] {
			checkitem := inventorystructsmap[10]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[0]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[0]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[0]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[0]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv12on {
		if inventoryfullmap[11] {
			checkitem := inventorystructsmap[11]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[1]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[1]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[1]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[1]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv13on {
		if inventoryfullmap[12] {
			checkitem := inventorystructsmap[12]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[2]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[2]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[2]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[2]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv14on {
		if inventoryfullmap[13] {
			checkitem := inventorystructsmap[13]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[3]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[3]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[3]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[3]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv15on {
		if inventoryfullmap[14] {
			checkitem := inventorystructsmap[14]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[4]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[4]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[4]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[4]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv16on {
		if inventoryfullmap[15] {
			checkitem := inventorystructsmap[15]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[5]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[5]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[5]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[5]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv17on {
		if inventoryfullmap[16] {
			checkitem := inventorystructsmap[16]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[6]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[6]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[6]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[6]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv18on {
		if inventoryfullmap[17] {
			checkitem := inventorystructsmap[17]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[7]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[7]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[7]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[7]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv19on {
		if inventoryfullmap[18] {
			checkitem := inventorystructsmap[18]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[8]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[8]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[8]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[8]-10, menu2y+80, 20, rl.White)
			}
		}
	} else if inv20on {
		if inventoryfullmap[19] {
			checkitem := inventorystructsmap[19]
			if menu2y > monh32/2 {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[9]-14, menu2y-31, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[9]-10, menu2y-30, 20, rl.White)
			} else {
				length := int32(len(checkitem.itemtype2))
				length = length * 12
				rl.DrawRectangle(invxmap[9]-14, menu2y+79, length, 22, rl.Black)
				rl.DrawText(checkitem.itemtype2, invxmap[9]-10, menu2y+80, 20, rl.White)
			}
		}
	}

}
func getmousepos() { // MARK: getmousepos

	// in menu positions
	if mousepos.X > float32(menux) && mousepos.X < float32(menux+menuw) && mousepos.Y > float32(menuy) && mousepos.Y < float32(menuy+menuh) {
		inmenu1 = true
	} else {
		inmenu1 = false
	}
	if mousepos.X > float32(menu2x) && mousepos.X < float32(menu2x+menu2w) && mousepos.Y > float32(menu2y) && mousepos.Y < float32(menu2y+menu2h) {
		inmenu2 = true
	} else {
		inmenu2 = false
	}

	// stats info tabs positions
	if mousepos.X > float32(menux) && mousepos.X < float32(menux+menuw/2) && mousepos.Y > float32(menuy+(menuh/2)-2) && mousepos.Y < float32(menuy+(menuh/2)+22) {
		infohover = true
	} else {
		infohover = false
	}
	if mousepos.X > float32(menux+menuw/2) && mousepos.X < float32(menux+menuw) && mousepos.Y > float32(menuy+(menuh/2)-2) && mousepos.Y < float32(menuy+(menuh/2)+22) {
		statshover = true
	} else {
		statshover = false
	}
	// equipment slot positions
	if mousepos.X > float32(menux+39) && mousepos.X < float32(menux+71) && mousepos.Y > float32(menuy+12) && mousepos.Y < float32(menuy+44) {
		helmetsloton = true
	} else {
		helmetsloton = false
	}
	if mousepos.X > float32(menux+39) && mousepos.X < float32(menux+71) && mousepos.Y > float32(menuy+48) && mousepos.Y < float32(menuy+80) {
		necklacesloton = true
	} else {
		necklacesloton = false
	}
	if mousepos.X > float32(menux+39) && mousepos.X < float32(menux+71) && mousepos.Y > float32(menuy+84) && mousepos.Y < float32(menuy+116) {
		armorsloton = true
	} else {
		armorsloton = false
	}
	if mousepos.X > float32(menux+39) && mousepos.X < float32(menux+71) && mousepos.Y > float32(menuy+120) && mousepos.Y < float32(menuy+152) {
		beltsloton = true
	} else {
		beltsloton = false
	}
	if mousepos.X > float32(menux+5) && mousepos.X < float32(menux+37) && mousepos.Y > float32(menuy+84) && mousepos.Y < float32(menuy+116) {
		ring1sloton = true
	} else {
		ring1sloton = false
	}
	if mousepos.X > float32(menux+77) && mousepos.X < float32(menux+109) && mousepos.Y > float32(menuy+84) && mousepos.Y < float32(menuy+116) {
		ring2sloton = true
	} else {
		ring2sloton = false
	}
	if mousepos.X > float32(menux+5) && mousepos.X < float32(menux+37) && mousepos.Y > float32(menuy+120) && mousepos.Y < float32(menuy+152) || mousepos.X > float32(menux+77) && mousepos.X < float32(menux+109) && mousepos.Y > float32(menuy+120) && mousepos.Y < float32(menuy+152) {
		glovessloton = true
	} else {
		glovessloton = false
	}
	if mousepos.X > float32(menux+5) && mousepos.X < float32(menux+37) && mousepos.Y > float32(menuy+156) && mousepos.Y < float32(menuy+188) {
		weaponslot1on = true
	} else {
		weaponslot1on = false
	}
	if mousepos.X > float32(menux+77) && mousepos.X < float32(menux+109) && mousepos.Y > float32(menuy+156) && mousepos.Y < float32(menuy+188) {
		weaponslot2on = true
	} else {
		weaponslot2on = false
	}
	if mousepos.X > float32(menux+16) && mousepos.X < float32(menux+48) && mousepos.Y > float32(menuy+192) && mousepos.Y < float32(menuy+224) || mousepos.X > float32(menux+68) && mousepos.X < float32(menux+100) && mousepos.Y > float32(menuy+192) && mousepos.Y < float32(menuy+224) {
		bootssloton = true
	} else {
		bootssloton = false
	}

	// inventory slot positions
	if mousepos.X > float32(invxmap[0]) && mousepos.X < float32(invxmap[9]+32) && mousepos.Y > float32(menu2y+4) && mousepos.Y < float32(menu2y+40) {

		if mousepos.X > float32(invxmap[0]) && mousepos.X < float32(invxmap[1]-2) {
			inv1on = true
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[1]) && mousepos.X < float32(invxmap[2]-2) {
			inv1on = false
			inv2on = true
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[2]) && mousepos.X < float32(invxmap[3]-2) {
			inv1on = false
			inv2on = false
			inv3on = true
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[3]) && mousepos.X < float32(invxmap[4]-2) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = true
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[4]) && mousepos.X < float32(invxmap[5]-2) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = true
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[5]) && mousepos.X < float32(invxmap[6]-2) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = true
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[6]) && mousepos.X < float32(invxmap[7]-2) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = true
			inv8on = false
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[7]) && mousepos.X < float32(invxmap[8]-2) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = true
			inv9on = false
			inv10on = false
		} else if mousepos.X > float32(invxmap[8]) && mousepos.X < float32(invxmap[9]-2) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = true
			inv10on = false
		} else if mousepos.X > float32(invxmap[9]) && mousepos.X < float32(invxmap[9]+32) {
			inv1on = false
			inv2on = false
			inv3on = false
			inv4on = false
			inv5on = false
			inv6on = false
			inv7on = false
			inv8on = false
			inv9on = false
			inv10on = true
		}
	} else {
		inv1on = false
		inv2on = false
		inv3on = false
		inv4on = false
		inv5on = false
		inv6on = false
		inv7on = false
		inv8on = false
		inv9on = false
		inv10on = false
	}

	if mousepos.X > float32(invxmap[0]) && mousepos.X < float32(invxmap[9]+32) && mousepos.Y > float32(menu2y+40) && mousepos.Y < float32(menu2y+76) {

		if mousepos.X > float32(invxmap[0]) && mousepos.X < float32(invxmap[1]-2) {
			inv11on = true
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[1]) && mousepos.X < float32(invxmap[2]-2) {
			inv11on = false
			inv12on = true
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[2]) && mousepos.X < float32(invxmap[3]-2) {
			inv11on = false
			inv12on = false
			inv13on = true
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[3]) && mousepos.X < float32(invxmap[4]-2) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = true
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[4]) && mousepos.X < float32(invxmap[5]-2) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = true
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[5]) && mousepos.X < float32(invxmap[6]-2) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = true
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[6]) && mousepos.X < float32(invxmap[7]-2) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = true
			inv18on = false
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[7]) && mousepos.X < float32(invxmap[8]-2) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = true
			inv19on = false
			inv20on = false
		} else if mousepos.X > float32(invxmap[8]) && mousepos.X < float32(invxmap[9]-2) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = true
			inv20on = false
		} else if mousepos.X > float32(invxmap[9]) && mousepos.X < float32(invxmap[9]+32) {
			inv11on = false
			inv12on = false
			inv13on = false
			inv14on = false
			inv15on = false
			inv16on = false
			inv17on = false
			inv18on = false
			inv19on = false
			inv20on = true
		}
	} else {
		inv11on = false
		inv12on = false
		inv13on = false
		inv14on = false
		inv15on = false
		inv16on = false
		inv17on = false
		inv18on = false
		inv19on = false
		inv20on = false
	}

	if int32(mousepos.X) > optionsmenutabx && int32(mousepos.X) < optionsmenutabx2 && int32(mousepos.Y) > optionsmenutaby && int32(mousepos.Y) < optionsmenutaby2 {
		optionsmenuclosehighlighton = true
	} else {
		optionsmenuclosehighlighton = false
	}

	if int32(mousepos.X) > menux && int32(mousepos.X) < menux+(menuw/2+16) && int32(mousepos.Y) > menuy+(menuh-38) && int32(mousepos.Y) < menuy+menuh {
		optionshighlighton = true
	} else {
		optionshighlighton = false
	}
	if int32(mousepos.X) > menux+(menuw/2+16) && int32(mousepos.X) < menux+menuw && int32(mousepos.Y) > menuy+(menuh-32) && int32(mousepos.Y) < menuy+menuh {
		savehighlighton = true
	} else {
		savehighlighton = false
	}
	if mousepos.X > menu2tabx && mousepos.X < menu2tabx2 && mousepos.Y > menu2taby && mousepos.Y < menu2taby2 {
		movemenu2highlighton = true
	} else {
		movemenu2highlighton = false
	}
	if mousepos.X > menutabx && mousepos.X < menutabx2 && mousepos.Y > menutaby && mousepos.Y < menutaby2 {
		movemenuhighlighton = true
	} else {
		movemenuhighlighton = false
	}

}
func getmouseblock() { // MARK: getmouseblock

	block := drawblocknext
	xchange := float32(0)
	ychange := float32(0)

	switch zoomlevel {
	case 1:
		for a := 0; a < gridh; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 16+ychange {
				for b := 0; b < gridw; b++ {
					if mousepos.X > 0+xchange && mousepos.X < 16+xchange {
						mouseblock = block
					}
					block++
					xchange += 16
				}
			}
			block += levelw
			xchange = 0
			ychange += 16
		}
	case 2:
		for a := 0; a < gridh2; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 32+ychange {
				for b := 0; b < gridw2; b++ {
					if mousepos.X > 0+xchange && mousepos.X < 32+xchange {
						mouseblock = block
					}
					block++
					xchange += 32
				}
			}
			block += levelw
			xchange = 0
			ychange += 32
		}
	case 3:
		for a := 0; a < gridh2; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 48+ychange {
				for b := 0; b < gridw2; b++ {
					if mousepos.X > 0+xchange && mousepos.X < 48+xchange {
						mouseblock = block
					}
					block++
					xchange += 48
				}
			}
			block += levelw
			xchange = 0
			ychange += 48
		}
	case 4:
		for a := 0; a < gridh4; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 64+ychange {
				for b := 0; b < gridw4; b++ {
					if mousepos.X > 0+xchange && mousepos.X < 64+xchange {
						mouseblock = block
					}
					block++
					xchange += 64
				}
			}
			block += levelw
			xchange = 0
			ychange += 64
		}
	}

	checkmouseitems1 := itemsmap[mouseblock]
	checkmouseitems2 := items2map[mouseblock]
	checkmouseitems3 := items3map[mouseblock]
	checkmouseenemies := enemiesmap[mouseblock]

	if checkmouseenemies != "" || checkmouseitems1 != "" || checkmouseitems2 != "" || checkmouseitems3 != "" {

		if checkmouseenemies != "" {
			checkblock := enemiesstructmap[mouseblock]
			length := int32(len(checkblock.itemtype))
			length = length * 11
			rl.DrawRectangle(int32(mousepos.X-51), int32(mousepos.Y-49), length, 22, rl.Black)
			rl.DrawText(checkblock.itemtype, int32(mousepos.X-50), int32(mousepos.Y-48), 20, rl.White)
		} else if checkmouseitems1 != "" {
			checkblock := itemsstructmap[mouseblock]
			length := int32(len(checkblock.itemtype))
			length = length * 11
			rl.DrawRectangle(int32(mousepos.X-51), int32(mousepos.Y-49), length, 22, rl.Black)
			rl.DrawText(checkblock.itemtype, int32(mousepos.X-50), int32(mousepos.Y-48), 20, rl.White)
		} else if checkmouseitems2 != "" {
			checkblock := itemsstructmap[mouseblock]
			length := int32(len(checkblock.itemname))
			length = length * 11
			rl.DrawRectangle(int32(mousepos.X-51), int32(mousepos.Y-49), length, 22, rl.Black)
			rl.DrawText(checkblock.itemname, int32(mousepos.X-50), int32(mousepos.Y-48), 20, rl.White)
		} else if checkmouseitems3 != "" {
			checkblock := itemsstructmap[mouseblock]
			length := int32(len(checkblock.itemtype))
			length = length * 11
			rl.DrawRectangle(int32(mousepos.X-51), int32(mousepos.Y-49), length, 22, rl.Black)
			rl.DrawText(checkblock.itemtype, int32(mousepos.X-50), int32(mousepos.Y-48), 20, rl.White)
		}

	}

}
func cursor() { // MARK: cursor
	mousev2 := rl.NewVector2(mousepos.X/2, mousepos.Y/2)
	mousev2shadow := rl.NewVector2(mousepos.X/2-2, mousepos.Y/2+2)
	rl.BeginMode2D(cameracursor)
	if invitemactiveon {
		rl.DrawTextureRec(imgs, cursoritem, mousev2, rl.White)
	} else {
		rl.DrawTextureRec(imgs, cursorimg, mousev2shadow, rl.Black)
		rl.DrawTextureRec(imgs, cursorimg, mousev2, rl.White)
	}
	rl.EndMode2D()
}
func movecamera() { // MARK: movecamera
	horizvert()
	switch zoomlevel {
	case 1:
		if playerh > drawblocknexth+gridh/2 {
			playermoving = true
			if drawblocknexth < levelh-gridh {
				drawblocknext += levelw
			}
		}
		if playerh < drawblocknexth+gridh/2 {
			playermoving = true
			if drawblocknexth > 0 {
				drawblocknext -= levelw
			}
		}
		if playerv > drawblocknextv+gridw/2 {
			playermoving = true
			if drawblocknextv < levelw-gridw {
				drawblocknext++
			}
		}
		if playerv < drawblocknextv+gridw/2 {
			playermoving = true
			if drawblocknextv > 0 {
				drawblocknext--
			}
		}
	case 2:
		if playerh > drawblocknexth+gridh2/2 {
			playermoving = true
			if drawblocknexth < levelh-gridh2 {
				drawblocknext += levelw
			}
		}
		if playerh < drawblocknexth+gridh2/2 {
			playermoving = true
			if drawblocknexth > 0 {
				drawblocknext -= levelw
			}
		}
		if playerv > drawblocknextv+gridw2/2 {
			playermoving = true
			if drawblocknextv < levelw-gridw2 {
				drawblocknext++
			}
		}
		if playerv < drawblocknextv+gridw2/2 {
			playermoving = true
			if drawblocknextv > 0 {
				drawblocknext--
			}
		}
	case 3:
		if playerh > drawblocknexth+gridh3/2 {
			playermoving = true
			if drawblocknexth < levelh-gridh3 {
				drawblocknext += levelw
			}
		}
		if playerh < drawblocknexth+gridh3/2 {
			playermoving = true
			if drawblocknexth > 0 {
				drawblocknext -= levelw
			}
		}
		if playerv > drawblocknextv+gridw3/2 {
			playermoving = true
			if drawblocknextv < levelw-gridw3 {
				drawblocknext++
			}
		}
		if playerv < drawblocknextv+gridw3/2 {
			playermoving = true
			if drawblocknextv > 0 {
				drawblocknext--
			}
		}
	case 4:
		if playerh > drawblocknexth+gridh4/2 {
			playermoving = true
			if drawblocknexth < levelh-gridh4 {
				drawblocknext += levelw
			}
		}
		if playerh < drawblocknexth+gridh4/2 {
			playermoving = true
			if drawblocknexth > 0 {
				drawblocknext -= levelw
			}
		}
		if playerv > drawblocknextv+gridw4/2 {
			playermoving = true
			if drawblocknextv < levelw-gridw4 {
				drawblocknext++
			}
		}
		if playerv < drawblocknextv+gridw4/2 {
			playermoving = true
			if drawblocknextv > 0 {
				drawblocknext--
			}
		}
	}
	if selectedblock == player {
		playermoving = false
	}
}
func pickup() { // MARK: pickup

	if itemsmap[player] != "" {
		for a := 0; a < len(inventorymap); a++ {
			if inventoryfullmap[a] == false {
				pickedup := itemsmap[player]
				itemsmap[player] = ""
				if pickedup == "coins" {
					checkitem := itemsstructmap[player]
					coins += checkitem.coins
					itemsstructmap[player] = items{}
				} else {
					inventorymap[a] = pickedup
					inventorystructsmap[a] = itemsstructmap[player]
					itemsstructmap[player] = items{}
					inventoryfullmap[a] = true
					break
				}
			}
		}
	}
	if items2map[player] != "" {
		for a := 0; a < len(inventorymap); a++ {
			if inventoryfullmap[a] == false {
				pickedup := items2map[player]
				items2map[player] = ""
				if pickedup == "coins" {
					checkitem := itemsstructmap[player]
					coins += checkitem.coins
					itemsstructmap[player] = items{}
				} else {

					inventorymap[a] = pickedup
					inventorystructsmap[a] = itemsstructmap[player]
					itemsstructmap[player] = items{}
					inventoryfullmap[a] = true
					break
				}
			}
		}
	}

}
func chooseimg(itemtype string) rl.Rectangle { // MARK: chooseimg

	itemimg := rl.NewRectangle(0, 0, 0, 0)

	switch itemtype {
	case "key":
		choose := rInt(1, 9)
		switch choose {
		case 1:
			itemimg = key1
		case 2:
			itemimg = key2
		case 3:
			itemimg = key3
		case 4:
			itemimg = key4
		case 5:
			itemimg = key5
		case 6:
			itemimg = key6
		case 7:
			itemimg = key7
		case 8:
			itemimg = key8
		}
	case "bow":
		choose := rInt(1, 6)
		switch choose {
		case 1:
			itemimg = bow1
		case 2:
			itemimg = bow2
		case 3:
			itemimg = bow3
		case 4:
			itemimg = bow4
		case 5:
			itemimg = bow5
		}
	case "sickle":
		choose := rInt(1, 3)
		switch choose {
		case 1:
			itemimg = sickle1
		case 2:
			itemimg = sickle2
		}
	case "club":
		choose := rInt(1, 7)
		switch choose {
		case 1:
			itemimg = club1
		case 2:
			itemimg = club2
		case 3:
			itemimg = club3
		case 4:
			itemimg = club4
		case 5:
			itemimg = club5
		case 6:
			itemimg = club6
		}
	case "staff":
		choose := rInt(1, 13)
		switch choose {
		case 1:
			itemimg = staff1
		case 2:
			itemimg = staff2
		case 3:
			itemimg = staff3
		case 4:
			itemimg = staff4
		case 5:
			itemimg = staff5
		case 6:
			itemimg = staff6
		case 7:
			itemimg = staff7
		case 8:
			itemimg = staff8
		case 9:
			itemimg = staff9
		case 10:
			itemimg = staff10
		case 11:
			itemimg = staff11
		case 12:
			itemimg = staff12
		}
	case "axe":
		choose := rInt(1, 10)
		switch choose {
		case 1:
			itemimg = axe1
		case 2:
			itemimg = axe2
		case 3:
			itemimg = axe3
		case 4:
			itemimg = axe4
		case 5:
			itemimg = axe5
		case 6:
			itemimg = axe6
		case 7:
			itemimg = axe7
		case 8:
			itemimg = axe8
		case 9:
			itemimg = axe9
		}
	case "crossbow":
		choose := rInt(1, 4)
		switch choose {
		case 1:
			itemimg = crossbow1
		case 2:
			itemimg = crossbow2
		case 3:
			itemimg = crossbow3
		}
	case "mace":
		choose := rInt(1, 5)
		switch choose {
		case 1:
			itemimg = mace1
		case 2:
			itemimg = mace2
		case 3:
			itemimg = mace3
		case 4:
			itemimg = mace4
		}
	case "potion":
		choose := rInt(1, 10)
		switch choose {
		case 1:
			itemimg = potion1
		case 2:
			itemimg = potion2
		case 3:
			itemimg = potion3
		case 4:
			itemimg = potion4
		case 5:
			itemimg = potion5
		case 6:
			itemimg = potion6
		case 7:
			itemimg = potion7
		case 8:
			itemimg = potion8
		case 9:
			itemimg = potion9
		}
	case "scroll":
		choose := rInt(1, 4)
		switch choose {
		case 1:
			itemimg = scroll1
		case 2:
			itemimg = scroll2
		case 3:
			itemimg = scroll3
		}
	case "sword":
		choose := rInt(1, 11)
		switch choose {
		case 1:
			itemimg = sword1
		case 2:
			itemimg = sword2
		case 3:
			itemimg = sword3
		case 4:
			itemimg = sword4
		case 5:
			itemimg = sword5
		case 6:
			itemimg = sword6
		case 7:
			itemimg = sword7
		case 8:
			itemimg = sword8
		case 9:
			itemimg = sword9
		case 10:
			itemimg = sword10
		}
	case "jewel":
		choose := rInt(1, 5)
		switch choose {
		case 1:
			itemimg = jewel1
		case 2:
			itemimg = jewel2
		case 3:
			itemimg = jewel3
		case 4:
			itemimg = jewel4
		}
	case "ring":
		choose := rInt(1, 5)
		switch choose {
		case 1:
			itemimg = ring1
		case 2:
			itemimg = ring2
		case 3:
			itemimg = ring3
		case 4:
			itemimg = ring4
		}
	case "necklace":
		choose := rInt(1, 9)
		switch choose {
		case 1:
			itemimg = necklace1
		case 2:
			itemimg = necklace2
		case 3:
			itemimg = necklace3
		case 4:
			itemimg = necklace4
		case 5:
			itemimg = necklace5
		case 6:
			itemimg = necklace6
		case 7:
			itemimg = necklace7
		case 8:
			itemimg = necklace8
		}
	case "helmet":
		choose := rInt(1, 8)
		switch choose {
		case 1:
			itemimg = helmet1
		case 2:
			itemimg = helmet2
		case 3:
			itemimg = helmet3
		case 4:
			itemimg = helmet4
		case 5:
			itemimg = helmet5
		case 6:
			itemimg = helmet6
		case 7:
			itemimg = helmet7
		}
	case "shield":
		choose := rInt(1, 9)
		switch choose {
		case 1:
			itemimg = shield1
		case 2:
			itemimg = shield2
		case 3:
			itemimg = shield3
		case 4:
			itemimg = shield4
		case 5:
			itemimg = shield5
		case 6:
			itemimg = shield6
		case 7:
			itemimg = shield7
		case 8:
			itemimg = shield8
		}
	case "armor":
		choose := rInt(1, 6)
		switch choose {
		case 1:
			itemimg = armor1
		case 2:
			itemimg = armor2
		case 3:
			itemimg = armor3
		case 4:
			itemimg = armor4
		case 5:
			itemimg = armor5
		}
	case "boots":
		choose := rInt(1, 5)
		switch choose {
		case 1:
			itemimg = boots1
		case 2:
			itemimg = boots2
		case 3:
			itemimg = boots3
		case 4:
			itemimg = boots4
		}
	case "gloves":
		choose := rInt(1, 5)
		switch choose {
		case 1:
			itemimg = gloves1
		case 2:
			itemimg = gloves2
		case 3:
			itemimg = gloves3
		case 4:
			itemimg = gloves4
		}
	case "monster":
		choose := rInt(1, 5)
		switch choose {
		case 1:
			itemimg = chestmonster1
		case 2:
			itemimg = chestmonster2
		case 3:
			itemimg = chestmonster3
		case 4:
			itemimg = chestmonster4
		}
	}

	return itemimg

}
func createchestitems() { // MARK: createchestitems

	count := 1

	for {
		var newitem items
		choose := rInt(1, 12)
		switch choose {
		case 1:
			newitem.itemtype = "potion"
		case 2:
			newitem.itemtype = "weapon"
		case 3:
			newitem.itemtype = "tool"
		case 4:
			newitem.itemtype = "loot"
		case 5:
			newitem.itemtype = "scroll"
		case 6:
			newitem.itemtype = "food"
		case 7:
			newitem.itemtype = "armor"
		case 8:
			newitem.itemtype = "trap"
		case 9:
			newitem.itemtype = "monster"
		case 10:
			newitem.itemtype = "key"
		}
		if newitem.itemtype == "weapon" {
			choose := rInt(1, 12)
			switch choose {
			case 1:
				newitem.itemtype2 = "mace"
				newitem.itemimg = chooseimg("mace")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("mace", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "crossbow"
				newitem.itemimg = chooseimg("crossbow")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("crossbow", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "axe"
				newitem.itemimg = chooseimg("axe")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("axe", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				newitem.itemtype2 = "staff"
				newitem.itemimg = chooseimg("staff")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("staff", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 5:
				newitem.itemtype2 = "club"
				newitem.itemimg = chooseimg("club")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("club", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 6:
				newitem.itemtype2 = "sickle"
				newitem.itemimg = chooseimg("sickle")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("sickle", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 7:
				newitem.itemtype2 = "bow"
				newitem.itemimg = chooseimg("bow")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("bow", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 8:
				newitem.itemtype2 = "shotgun"
				newitem.itemimg = shotgun
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 9:
				newitem.itemtype2 = "uzzi"
				newitem.itemimg = uzzi
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 10:
				newitem.itemtype2 = "bomb"
				newitem.itemimg = bomb
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 11:
				newitem.itemtype2 = "sword"
				newitem.itemimg = chooseimg("sword")
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("sword", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "potion" {
			choose := rInt(1, 9)
			itemimg := chooseimg("potion")
			itemname := createname("potion", 0)
			switch choose {
			case 1:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 5:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 6:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 7:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 8:
				newitem.itemtype2 = "potion"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = itemname
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "tool" {
			choose := rInt(1, 4)
			switch choose {
			case 1:
				newitem.itemtype2 = "spade"
				newitem.itemimg = spade
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("spade", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "pickaxe"
				newitem.itemimg = pickaxe
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("pickaxe", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "torch"
				newitem.itemimg = torch
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("torch", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "loot" {
			choose := rInt(1, 5)
			switch choose {
			case 1:
				newitem.itemtype2 = "coins"
				newitem.itemimg = coin
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("coins", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				itemimg := chooseimg("jewel")
				newitem.itemtype2 = "jewel"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.itemname = createname("jewel", newitem.damage)
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				itemimg := chooseimg("necklace")
				newitem.itemtype2 = "necklace"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				itemimg := chooseimg("ring")
				newitem.itemtype2 = "ring"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "scroll" {
			choose := rInt(1, 6)
			itemimg := chooseimg("scroll")
			switch choose {
			case 1:
				newitem.itemtype2 = "fireball"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "iceball"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "poison gas"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				newitem.itemtype2 = "teleport"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 5:
				newitem.itemtype2 = "find secrets"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 6:
				newitem.itemtype2 = "identify"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "food" {
			choose := rInt(1, 8)
			switch choose {
			case 1:
				newitem.itemtype2 = "apple"
				newitem.itemimg = apple
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "pear"
				newitem.itemimg = pear
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "pumpkin"
				newitem.itemimg = pumpkin
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				newitem.itemtype2 = "ice cream"
				newitem.itemimg = icecream
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 5:
				newitem.itemtype2 = "gingerbread"
				newitem.itemimg = gingerbread
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 6:
				newitem.itemtype2 = "chocolate"
				newitem.itemimg = chocolate
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 7:
				newitem.itemtype2 = "pizza"
				newitem.itemimg = pizza
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "armor" {
			choose := rInt(1, 7)
			switch choose {
			case 1:
				itemimg := chooseimg("helmet")
				newitem.itemtype2 = "helmet"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				itemimg := chooseimg("shield")
				newitem.itemtype2 = "shield"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				itemimg := chooseimg("armor")
				newitem.itemtype2 = "armor"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				itemimg := chooseimg("boots")
				newitem.itemtype2 = "boots"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 5:
				itemimg := chooseimg("gloves")
				newitem.itemtype2 = "gloves"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 6:
				newitem.itemtype2 = "belt"
				newitem.itemimg = belt
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "trap" {
			choose := rInt(1, 5)
			switch choose {
			case 1:
				newitem.itemtype2 = "fireball"
				newitem.itemimg = potion1
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "poison gas"
				newitem.itemimg = potion1
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "freeze"
				newitem.itemimg = potion1
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				newitem.itemtype2 = "darkness"
				newitem.itemimg = potion1
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "key" {
			choose := rInt(1, 8)
			itemimg := chooseimg("key")
			switch choose {
			case 1:
				newitem.itemtype2 = "bronze key"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 2:
				newitem.itemtype2 = "copper key"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 3:
				newitem.itemtype2 = "gold key"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 4:
				newitem.itemtype2 = "silver key"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 5:
				newitem.itemtype2 = "emerald key"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			case 6:
				newitem.itemtype2 = "stone key"
				newitem.itemimg = itemimg
				newitem.coins = rolldice()
				newitem.hp = rolldice()
				newitem.damage = rolldice()
				newitem.magic = rolldice()
				newitem.durability = rInt(50, 101)
				newitem.duration = rInt(30, 121)
				newitem.random1 = rolldice() + rolldice()
				newitem.random2 = rolldice() + rolldice()
			}
		} else if newitem.itemtype == "monster" {
			itemimg := chooseimg("monster")
			newitem.itemtype2 = "monster"
			newitem.itemname = "monster name here"
			newitem.itemimg = itemimg
			newitem.coins = rolldice()
			newitem.hp = rolldice()
			newitem.damage = rolldice()
			newitem.magic = rolldice()
			newitem.durability = rInt(50, 101)
			newitem.duration = rInt(30, 121)
			newitem.random1 = rolldice() + rolldice()
			newitem.random2 = rolldice() + rolldice()

		}

		if count == 1 {
			chestitem1 = newitem
		} else if count == 2 {
			chestitem2 = newitem
		} else if count == 3 {
			chestitem3 = newitem
		}

		if count == chestitemnumber {
			break
		}

		count++

	}

}
func createname(itemtype string, damage int) string { // MARK: createname

	itemname := "item name here"

	switch itemtype {
	case "jewel":
		choose := rInt(0, len(words3))
		word1 := words3[choose]
		choose = rInt(0, len(wordsjewels))
		word2 := wordsjewels[choose]
		itemname = word1 + word2
	case "coins":
		word1 := strconv.Itoa(damage)
		if damage == 1 {
			itemname = word1 + " coin"
		} else {
			itemname = word1 + " coins"
		}
	case "spade":
		choose := rInt(0, len(words1))
		word1 := words1[choose]
		itemname = word1 + "spade"
	case "pickaxe":
		choose := rInt(0, len(words1))
		word1 := words1[choose]
		itemname = word1 + "pickaxe"
	case "torch":
		choose := rInt(0, len(words3))
		word1 := words3[choose]
		itemname = word1 + "torch"
	case "potion":
		choose := rInt(0, len(words3))
		word1 := words3[choose]
		choose = rInt(0, len(wordscolors))
		word2 := wordscolors[choose]
		choose = rInt(0, len(wordspotions))
		word3 := wordspotions[choose]
		itemname = word1 + word2 + word3
	case "sword":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsswords))
			word3 := wordsswords[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsswords))
			word3 := wordsswords[choose]
			itemname = word1 + word2 + word3
		}
	case "axe":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsaxes))
			word3 := wordsaxes[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsaxes))
			word3 := wordsaxes[choose]
			itemname = word1 + word2 + word3
		}
	case "mace":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsmaces))
			word3 := wordsmaces[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsmaces))
			word3 := wordsmaces[choose]
			itemname = word1 + word2 + word3
		}
	case "crossbow":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordscrossbows))
			word3 := wordscrossbows[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordscrossbows))
			word3 := wordscrossbows[choose]
			itemname = word1 + word2 + word3
		}
	case "staff":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsstaff))
			word3 := wordsstaff[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsstaff))
			word3 := wordsstaff[choose]
			itemname = word1 + word2 + word3
		}
	case "club":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsclub))
			word3 := wordsclub[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsclub))
			word3 := wordsclub[choose]
			itemname = word1 + word2 + word3
		}
	case "sickle":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordssickle))
			word3 := wordssickle[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordssickle))
			word3 := wordssickle[choose]
			itemname = word1 + word2 + word3
		}
	case "bow":
		if damage > 3 {
			choose := rInt(0, len(words2))
			word1 := words2[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsbow))
			word3 := wordsbow[choose]
			itemname = word1 + word2 + word3
		} else {
			choose := rInt(0, len(words1))
			word1 := words1[choose]
			choose = rInt(0, len(wordsmaterial))
			word2 := wordsmaterial[choose]
			choose = rInt(0, len(wordsbow))
			word3 := wordsbow[choose]
			itemname = word1 + word2 + word3
		}
	}

	return itemname
}
func openchest() { // MARK: openchest

	chestitemnumber = rInt(1, 4)
	createchestitems()

	for {
		choose := rInt(1, 9)
		switch choose {
		case 1:
			if levelmap[chestblock-1-levelw] == "floor" && items2map[chestblock-1-levelw] == "" && chestblock-1-levelw != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock-1-levelw] = chestitem1.itemname
							enemiesstructmap[chestblock-1-levelw] = chestitem1
						}

					} else {
						items2map[chestblock-1-levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock-1-levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock-1-levelw] = chestitem2.itemname
							enemiesstructmap[chestblock-1-levelw] = chestitem2
						}

					} else {
						items2map[chestblock-1-levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock-1-levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock-1-levelw] = chestitem3.itemname
							enemiesstructmap[chestblock-1-levelw] = chestitem3
						}
					} else {
						items2map[chestblock-1-levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock-1-levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 2:
			if levelmap[chestblock-1] == "floor" && items2map[chestblock-1] == "" && chestblock-1 != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock-1] = chestitem1.itemname
							enemiesstructmap[chestblock-1] = chestitem1
						}
					} else {
						items2map[chestblock-1] = chestitem1.itemtype2
						itemsstructmap[chestblock-1] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock-1] = chestitem2.itemname
							enemiesstructmap[chestblock-1] = chestitem2
						}
					} else {
						items2map[chestblock-1] = chestitem2.itemtype2
						itemsstructmap[chestblock-1] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock-1] = chestitem3.itemname
							enemiesstructmap[chestblock-1] = chestitem3
						}
					} else {
						items2map[chestblock-1] = chestitem3.itemtype2
						itemsstructmap[chestblock-1] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 3:
			if levelmap[chestblock-levelw] == "floor" && items2map[chestblock-levelw] == "" && chestblock-levelw != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock-levelw] = chestitem1.itemname
							enemiesstructmap[chestblock-levelw] = chestitem1
						}
					} else {
						items2map[chestblock-levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock-levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock-levelw] = chestitem2.itemname
							enemiesstructmap[chestblock-levelw] = chestitem2
						}
					} else {
						items2map[chestblock-levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock-levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock-levelw] = chestitem3.itemname
							enemiesstructmap[chestblock-levelw] = chestitem3
						}
					} else {
						items2map[chestblock-levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock-levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 4:
			if levelmap[chestblock+1-levelw] == "floor" && items2map[chestblock+1-levelw] == "" && chestblock+1-levelw != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock+1-levelw] = chestitem1.itemname
							enemiesstructmap[chestblock+1-levelw] = chestitem1
						}
					} else {
						items2map[chestblock+1-levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock+1-levelw] = chestitem2.itemname
							enemiesstructmap[chestblock+1-levelw] = chestitem2
						}
					} else {
						items2map[chestblock+1-levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock+1-levelw] = chestitem3.itemname
							enemiesstructmap[chestblock+1-levelw] = chestitem3
						}
					} else {
						items2map[chestblock+1-levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 5:
			if levelmap[chestblock+1] == "floor" && items2map[chestblock+1] == "" && chestblock+1 != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock+1] = chestitem1.itemname
							enemiesstructmap[chestblock+1] = chestitem1
						}
					} else {
						items2map[chestblock+1-levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock+1] = chestitem2.itemname
							enemiesstructmap[chestblock+1] = chestitem2
						}
					} else {
						items2map[chestblock+1-levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock+1] = chestitem3.itemname
							enemiesstructmap[chestblock+1] = chestitem3
						}
					} else {
						items2map[chestblock+1-levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 6:
			if levelmap[chestblock+1+levelw] == "floor" && items2map[chestblock+1+levelw] == "" && chestblock+1+levelw != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock+1+levelw] = chestitem1.itemname
							enemiesstructmap[chestblock+1+levelw] = chestitem1
						}
					} else {
						items2map[chestblock+1+levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock+1+levelw] = chestitem2.itemname
							enemiesstructmap[chestblock+1+levelw] = chestitem2
						}
					} else {
						items2map[chestblock+1+levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock+1+levelw] = chestitem3.itemname
							enemiesstructmap[chestblock+1+levelw] = chestitem3
						}
					} else {
						items2map[chestblock+1+levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock+1-levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 7:
			if levelmap[chestblock+levelw] == "floor" && items2map[chestblock+levelw] == "" && chestblock+levelw != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock+levelw] = chestitem1.itemname
							enemiesstructmap[chestblock+levelw] = chestitem1
						}
					} else {
						items2map[chestblock+levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock+levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock+levelw] = chestitem2.itemname
							enemiesstructmap[chestblock+levelw] = chestitem2
						}
					} else {
						items2map[chestblock+levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock+levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock+levelw] = chestitem3.itemname
							enemiesstructmap[chestblock+levelw] = chestitem3
						}
					} else {
						items2map[chestblock+levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock+levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		case 8:
			if levelmap[chestblock-1+levelw] == "floor" && items2map[chestblock-1+levelw] == "" && chestblock-1+levelw != player {
				if chestitemnumber == 1 {
					if chestitem1.itemtype == "monster" || chestitem1.itemtype == "trap" {
						if chestitem1.itemtype == "monster" {
							enemiesmap[chestblock-1+levelw] = chestitem1.itemname
							enemiesstructmap[chestblock-1+levelw] = chestitem1
						}
					} else {
						items2map[chestblock-1+levelw] = chestitem1.itemtype2
						itemsstructmap[chestblock-1+levelw] = chestitem1
					}
				} else if chestitemnumber == 2 {
					if chestitem2.itemtype == "monster" || chestitem2.itemtype == "trap" {
						if chestitem2.itemtype == "monster" {
							enemiesmap[chestblock-1+levelw] = chestitem2.itemname
							enemiesstructmap[chestblock-1+levelw] = chestitem2
						}
					} else {
						items2map[chestblock-1+levelw] = chestitem2.itemtype2
						itemsstructmap[chestblock-1+levelw] = chestitem2
					}
				} else if chestitemnumber == 3 {
					if chestitem3.itemtype == "monster" || chestitem3.itemtype == "trap" {
						if chestitem3.itemtype == "monster" {
							enemiesmap[chestblock-1+levelw] = chestitem3.itemname
							enemiesstructmap[chestblock-1+levelw] = chestitem3
						}
					} else {
						items2map[chestblock-1+levelw] = chestitem3.itemtype2
						itemsstructmap[chestblock-1+levelw] = chestitem3
					}
				}
				chestitemnumber--
			}
		}

		if chestitemnumber <= 0 {
			break
		}
	}

}
func moveplayer() { // MARK: moveplayer

	if selectedblock > 0 {

		if selectedblock != player {
			if playerh < selectedblockh {
				if levelmap[player+levelw] == "floor" || levelmap[player+levelw] == "." || levelmap[player+levelw] == "door" {
					player += levelw
				} else if levelmap[player-levelw] == "floor" || levelmap[player-levelw] == "." || levelmap[player-levelw] == "door" {
					player -= levelw
				}
			} else if playerh > selectedblockh {
				if levelmap[player-levelw] == "floor" || levelmap[player-levelw] == "." || levelmap[player-levelw] == "door" {
					player -= levelw
				} else if levelmap[player+levelw] == "floor" || levelmap[player+levelw] == "." || levelmap[player+levelw] == "door" {
					player += levelw
				}
			}

			if playerv < selectedblockv {
				if levelmap[player+1] == "floor" || levelmap[player+1] == "." || levelmap[player+1] == "door" {
					player++
				}
			} else if playerv > selectedblockv {
				if levelmap[player-1] == "floor" || levelmap[player-1] == "." || levelmap[player-1] == "door" {
					player--
				}
			}
		}

	}

}
func menuoptions() { // MARK: optionsmenu

	optionsmenuwspace = (monw32 - optionsmenuw) / 2
	optionsmenuhspace = (monh32 - optionsmenuh) / 2
	optionsmenutabx = optionsmenuwspace + (optionsmenuw - 20)
	optionsmenutabx2 = optionsmenuwspace + optionsmenuw
	optionsmenutaby = optionsmenuhspace
	optionsmenutaby2 = optionsmenuhspace + 20

	rl.DrawRectangle(optionsmenuwspace, optionsmenuhspace, optionsmenuw, optionsmenuh, rl.Fade(rl.Black, 0.7)) // border lines
	rl.DrawRectangleLines(optionsmenuwspace, optionsmenuhspace, optionsmenuw, optionsmenuh, rl.White)          // border lines
	rl.DrawRectangleLines(optionsmenuwspace-1, optionsmenuhspace-1, optionsmenuw+2, optionsmenuh+2, rl.White)  // border lines
	if optionsmenuclosehighlighton {
		rl.DrawRectangle(optionsmenuwspace+(optionsmenuw-20), optionsmenuhspace, 20, 20, rl.White)                   // close window rec highlight
		rl.DrawRectangle(optionsmenuwspace+(optionsmenuw-20), optionsmenuhspace-1, 21, 21, rl.Fade(rl.SkyBlue, 0.6)) // close window rec highlight
		rl.DrawRectangleLines(optionsmenuwspace+(optionsmenuw-20), optionsmenuhspace-1, 21, 21, rl.SkyBlue)          // close window rec highlight
	} else {
		rl.DrawRectangle(optionsmenuwspace+(optionsmenuw-20), optionsmenuhspace, 20, 20, rl.White) // close window rec
	}
	// close window icon
	v2 := rl.NewVector2(float32(optionsmenuwspace+(optionsmenuw-16)), float32(optionsmenuhspace+3))
	rl.DrawTextureRec(imgs, closemenuimg, v2, rl.White)

}
func menus() { // MARK: menu

	if movemenuon {
		menux = int32(mousepos.X) - (menuw + 10)
		menuy = int32(mousepos.Y) - 10
	}
	if movemenu2on {
		menu2x = int32(mousepos.X) - (menu2w + 10)
		menu2y = int32(mousepos.Y) - 10
	}
	menutabx = float32(menux + menuw)
	menutaby = float32(menuy - 1)
	menutabx2 = float32(menutabx + 20)
	menutaby2 = float32(menutaby + 20)

	menu2tabx = float32(menu2x + menu2w)
	menu2taby = float32(menu2y - 1)
	menu2tabx2 = float32(menu2tabx + 20)
	menu2taby2 = float32(menu2taby + 20)

	check := levelmap[mouseblock]
	switch check {
	case "floor":
		mousetileinfo = "slightly damp floor"
	case "wall":
		mousetileinfo = "cold stone wall"
	}

	// menu 1
	rl.DrawRectangle(menux, menuy, menuw, menuh, rl.Fade(rl.Black, 0.7)) // background
	rl.DrawRectangleLines(menux, menuy, menuw, menuh, rl.White)          // border lines
	rl.DrawRectangleLines(menux-1, menuy-1, menuw+2, menuh+2, rl.White)  // border lines

	// hp hunger bars
	hpy := menuy + 210
	for a := 0; a < 10; a++ {
		rl.DrawRectangleLines(menux+130, hpy, 18, 18, rl.White)
		hpy -= 19
	}
	hpy = menuy + 210
	for a := 0; a < hpcount; a++ {
		rl.DrawRectangle(menux+130, hpy, 18, 18, rl.White)
		hpy -= 19
	}
	foody := menuy + 210
	for a := 0; a < 10; a++ {
		rl.DrawRectangleLines(menux+166, foody, 18, 18, rl.White)
		foody -= 19
	}
	foody = menuy + 210
	for a := 0; a < hungercount; a++ {
		rl.DrawRectangle(menux+166, foody, 18, 18, rl.White)
		foody -= 19
	}

	// player outline image
	v2 := rl.NewVector2(float32(menux+5), float32(menuy+5))
	rl.DrawTextureRec(imgs, playeroutline, v2, rl.White)   // player outline
	rl.DrawRectangle(menux+39, menuy+12, 32, 32, rl.Black) // helmet slot
	if helmetsloton {
		rl.DrawRectangle(menux+39, menuy+12, 32, 32, rl.Fade(rl.White, 0.2)) // helmet hover
	}
	rl.DrawRectangleLines(menux+39, menuy+12, 32, 32, rl.White) // helmet slot
	rl.DrawRectangle(menux+39, menuy+48, 32, 32, rl.Black)      // necklace slot
	if necklacesloton {
		rl.DrawRectangle(menux+39, menuy+48, 32, 32, rl.Fade(rl.White, 0.2)) // necklace hover
	}
	rl.DrawRectangleLines(menux+39, menuy+48, 32, 32, rl.White) // necklace slot
	rl.DrawRectangle(menux+40, menuy+84, 32, 32, rl.Black)      // armor slot
	if armorsloton {
		rl.DrawRectangle(menux+40, menuy+84, 32, 32, rl.Fade(rl.White, 0.2)) // armor hover
	}
	rl.DrawRectangleLines(menux+40, menuy+84, 32, 32, rl.White) // armor slot
	rl.DrawRectangle(menux+41, menuy+120, 32, 32, rl.Black)     // belt slot
	if beltsloton {
		rl.DrawRectangle(menux+41, menuy+120, 32, 32, rl.Fade(rl.White, 0.2)) // belt hover
	}
	rl.DrawRectangleLines(menux+41, menuy+120, 32, 32, rl.White) // belt slot
	rl.DrawRectangle(menux+5, menuy+84, 32, 32, rl.Black)        // left ring slot
	if ring1sloton {
		rl.DrawRectangle(menux+5, menuy+84, 32, 32, rl.Fade(rl.White, 0.2)) // left ring hover
	}
	rl.DrawRectangleLines(menux+5, menuy+84, 32, 32, rl.White)  // left ring slot
	rl.DrawRectangle(menux+5, menuy+120, 32, 32, rl.Black)      // left gloves slot
	rl.DrawRectangleLines(menux+5, menuy+120, 32, 32, rl.White) // left gloves slot
	rl.DrawRectangle(menux+5, menuy+156, 32, 32, rl.Black)      // left weapon slot
	if weaponslot1on {
		rl.DrawRectangle(menux+5, menuy+156, 32, 32, rl.Fade(rl.White, 0.2)) // left weapon hover
	}
	rl.DrawRectangleLines(menux+5, menuy+156, 32, 32, rl.White) // left weapon slot
	rl.DrawRectangle(menux+77, menuy+84, 32, 32, rl.Black)      // right ring slot
	if ring2sloton {
		rl.DrawRectangle(menux+77, menuy+84, 32, 32, rl.Fade(rl.White, 0.2)) // right ring hover
	}
	rl.DrawRectangleLines(menux+77, menuy+84, 32, 32, rl.White) // right ring slot
	rl.DrawRectangle(menux+77, menuy+120, 32, 32, rl.Black)     // right gloves slot
	if glovessloton {
		rl.DrawRectangle(menux+77, menuy+120, 32, 32, rl.Fade(rl.White, 0.2)) // gloves hover
		rl.DrawRectangle(menux+5, menuy+120, 32, 32, rl.Fade(rl.White, 0.2))  // gloves hover
	}
	rl.DrawRectangleLines(menux+77, menuy+120, 32, 32, rl.White) // right gloves slot
	rl.DrawRectangle(menux+77, menuy+156, 32, 32, rl.Black)      // right weapon slot
	if weaponslot2on {
		rl.DrawRectangle(menux+77, menuy+156, 32, 32, rl.Fade(rl.White, 0.2)) // right weapon hover
	}
	rl.DrawRectangleLines(menux+77, menuy+156, 32, 32, rl.White) // right weapon slot
	rl.DrawRectangle(menux+16, menuy+192, 32, 32, rl.Black)      // left boots slot
	rl.DrawRectangleLines(menux+16, menuy+192, 32, 32, rl.White) // left boots slot
	rl.DrawRectangle(menux+68, menuy+192, 32, 32, rl.Black)      // right boots slot
	rl.DrawRectangleLines(menux+68, menuy+192, 32, 32, rl.White) // right boots slot
	if bootssloton {
		rl.DrawRectangle(menux+16, menuy+192, 32, 32, rl.Fade(rl.White, 0.2)) // left boots hover
		rl.DrawRectangle(menux+68, menuy+192, 32, 32, rl.Fade(rl.White, 0.2)) // right boots hover
	}

	// draw slot icons equipped items icons
	rl.BeginMode2D(camerainventory)

	// hp food icons
	hpv2 := rl.NewVector2(float32((menux+128)/2), float32((menuy+10)/2))
	rl.DrawTextureRec(imgs, heart, hpv2, rl.White)
	foodv2 := rl.NewVector2(float32((menux+160)/2), float32((menuy+6)/2))
	rl.DrawTextureRec(imgs, foodmenu, foodv2, rl.White)

	helmetv2 := rl.NewVector2(float32(((menux+39)/2)+1), float32((menuy+12)/2))
	if helmetslotequipped {
		checkitem := equippeditemsstructmap[0]
		rl.DrawTextureRec(imgs, checkitem.itemimg, helmetv2, rl.White) // helmet image
	} else {
		rl.DrawTextureRec(imgs, helmet2, helmetv2, rl.Fade(rl.White, 0.4)) // helmet background image
	}
	necklacev2 := rl.NewVector2(float32(((menux+39)/2)+1), float32((menuy+52)/2))
	if necklaceslotequipped {
		checkitem := equippeditemsstructmap[1]
		rl.DrawTextureRec(imgs, checkitem.itemimg, necklacev2, rl.White) // necklace image
	} else {
		rl.DrawTextureRec(imgs, necklace3, necklacev2, rl.Fade(rl.White, 0.4)) // necklace background image
	}
	ring1v2 := rl.NewVector2(float32(((menux+5)/2)+1), float32((menuy+84)/2)+1)
	if ring1slotequipped {
		checkitem := equippeditemsstructmap[2]
		rl.DrawTextureRec(imgs, checkitem.itemimg, ring1v2, rl.White) // left ring image
	} else {
		rl.DrawTextureRec(imgs, ring2, ring1v2, rl.Fade(rl.White, 0.4)) // left ring background image
	}
	armorv2 := rl.NewVector2(float32(((menux+39)/2)+2), float32((menuy+84)/2)+1)
	if armorslotequipped {
		checkitem := equippeditemsstructmap[3]
		rl.DrawTextureRec(imgs, checkitem.itemimg, armorv2, rl.White) // armor image
	} else {
		rl.DrawTextureRec(imgs, armor1, armorv2, rl.Fade(rl.White, 0.4)) // armor background image
	}
	ring2v2 := rl.NewVector2(float32(((menux+78)/2)+1), float32((menuy+84)/2)+1)
	if ring2slotequipped {
		checkitem := equippeditemsstructmap[4]
		rl.DrawTextureRec(imgs, checkitem.itemimg, ring2v2, rl.White) // right ring image
	} else {
		rl.DrawTextureRec(imgs, ring2, ring2v2, rl.Fade(rl.White, 0.4)) // right ring background image
	}
	gloves1v2 := rl.NewVector2(float32(((menux + 5) / 2)), float32((menuy+120)/2))
	gloves2v2 := rl.NewVector2(float32(((menux + 78) / 2)), float32((menuy+120)/2))
	if glovesslotequipped {
		checkitem := equippeditemsstructmap[5]
		rl.DrawTextureRec(imgs, checkitem.itemimg, gloves1v2, rl.White) // gloves image
		rl.DrawTextureRec(imgs, checkitem.itemimg, gloves2v2, rl.White) // gloves image
	} else {
		rl.DrawTextureRec(imgs, gloves1, gloves1v2, rl.Fade(rl.White, 0.4)) // left gloves background image
		rl.DrawTextureRec(imgs, gloves1, gloves2v2, rl.Fade(rl.White, 0.4)) // right gloves background image
	}
	beltv2 := rl.NewVector2(float32(((menux + 42) / 2)), float32((menuy+120)/2))
	if beltslotequipped {
		checkitem := equippeditemsstructmap[6]
		rl.DrawTextureRec(imgs, checkitem.itemimg, beltv2, rl.White) // belt image
	} else {
		rl.DrawTextureRec(imgs, belt, beltv2, rl.Fade(rl.White, 0.4)) // belt background image
	}
	weapon1v2 := rl.NewVector2(float32(((menux + 5) / 2)), float32((menuy+156)/2))
	if weaponslot1equipped {
		checkitem := equippeditemsstructmap[7]
		rl.DrawTextureRec(imgs, checkitem.itemimg, weapon1v2, rl.White) // left weapon image
	} else {
		rl.DrawTextureRec(imgs, axe1, weapon1v2, rl.Fade(rl.White, 0.4)) // left weapon background image
	}
	weapon2v2 := rl.NewVector2(float32(((menux + 78) / 2)), float32((menuy+156)/2))
	if weaponslot2equipped {
		checkitem := equippeditemsstructmap[8]
		rl.DrawTextureRec(imgs, checkitem.itemimg, weapon2v2, rl.White) // right weapon background image
	} else {
		rl.DrawTextureRec(imgs, shield1, weapon2v2, rl.Fade(rl.White, 0.4)) // right weapon background image
	}
	boots1v2 := rl.NewVector2(float32(((menux + 16) / 2)), float32((menuy+192)/2))
	boots2v2 := rl.NewVector2(float32(((menux + 68) / 2)), float32((menuy+192)/2))
	if bootsslotequipped {
		checkitem := equippeditemsstructmap[9]
		rl.DrawTextureRec(imgs, checkitem.itemimg, boots1v2, rl.White) // left boots image
		rl.DrawTextureRec(imgs, checkitem.itemimg, boots2v2, rl.White) // right boots image
	} else {
		rl.DrawTextureRec(imgs, boots4, boots1v2, rl.Fade(rl.White, 0.4)) // left boots background image
		rl.DrawTextureRec(imgs, boots4, boots2v2, rl.Fade(rl.White, 0.4)) // right boots background image
	}
	rl.EndMode2D()

	// draw slot text
	if helmetsloton && helmetslotequipped == false {
		length := int32(6 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("helmet", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if necklacesloton && necklaceslotequipped == false {
		length := int32(8 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("necklace", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if armorsloton && armorslotequipped == false {
		length := int32(5 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("armor", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if ring1sloton && ring1slotequipped == false {
		length := int32(4 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("ring", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if ring2sloton && ring2slotequipped == false {
		length := int32(4 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("ring", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if beltsloton && beltslotequipped == false {
		length := int32(4 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("belt", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if glovessloton && glovesslotequipped == false {
		length := int32(6 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("gloves", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if weaponslot1on && weaponslot1equipped == false {
		length := int32(13 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("weapon/shield", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if weaponslot2on && weaponslot2equipped == false {
		length := int32(13 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("weapon/shield", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}
	if bootssloton && bootsslotequipped == false {
		length := int32(5 * 12)
		rl.DrawRectangle(int32(mousepos.X+32), int32(mousepos.Y+20), length, 22, rl.Black)
		rl.DrawText("boots", int32(mousepos.X+32), int32(mousepos.Y+20), 20, rl.White)
	}

	if optionshighlighton {
		rl.DrawRectangle(menux+3, menuy+(menuh-38), menuw/2+16, 35, rl.White) // menu highlight background
		rl.DrawText("OPTIONS", menux+15, menuy+(menuh-28), 20, rl.Black)      // menu text
	} else {
		rl.DrawText("OPTIONS", menux+15, menuy+(menuh-28), 20, rl.White) // menu text
	}
	if savehighlighton {
		rl.DrawRectangle((menux+3)+(menuw/2+16), menuy+(menuh-38), menuw/2-22, 35, rl.White) // menu highlight background
		rl.DrawText("SAVE", menux+menuw-72, menuy+(menuh-28), 20, rl.Black)                  // menu text
	} else {
		rl.DrawText("SAVE", menux+menuw-72, menuy+(menuh-28), 20, rl.White) // menu text
	}
	if movemenuhighlighton {
		rl.DrawRectangle(menux+menuw, menuy-1, 20, 20, rl.White)                 // move tab highlight
		rl.DrawRectangle(menux+menuw, menuy-1, 21, 21, rl.Fade(rl.SkyBlue, 0.6)) // move tab highlight
		rl.DrawRectangle(menux+menuw, menuy-1, 21, 21, rl.SkyBlue)               // move tab highlight highlight
	} else {
		rl.DrawRectangle(menux+menuw, menuy-1, 20, 20, rl.White) // move tab
	}
	v2 = rl.NewVector2(float32(menux+menuw+3), float32(menuy+2)) // move menu img
	rl.DrawTextureRec(imgs, movemenuimg, v2, rl.White)           // move menu img

	// tip window

	rl.DrawRectangle(menux, menuy+((menuh/2)+21), menuw, 2, rl.White)

	if statson == false {
		rl.DrawRectangle(menux, menuy+((menuh/2)-2), menuw/2, 24, rl.White)
		rl.DrawText("info", menux+20, menuy+(menuh/2), 20, rl.Black)
		rl.DrawRectangleLines(menux+menuw/2, menuy+((menuh/2)-2), menuw/2, 24, rl.White)
		rl.DrawText("stats", menux+((menuw/2)+20), menuy+(menuh/2), 20, rl.Fade(rl.White, 0.6))
	} else {
		rl.DrawRectangleLines(menux, menuy+((menuh/2)-2), menuw/2, 24, rl.White)
		rl.DrawText("info", menux+20, menuy+(menuh/2), 20, rl.Fade(rl.White, 0.6))
		rl.DrawRectangle(menux+menuw/2, menuy+((menuh/2)-2), menuw/2, 24, rl.White)
		rl.DrawText("stats", menux+((menuw/2)+20), menuy+(menuh/2), 20, rl.Black)
	}
	if statshover && statson == false {
		if statsflash {
			rl.DrawRectangle(menux+menuw/2+3, menuy+((menuh/2)-2), menuw/2-3, 24, rl.White)
			rl.DrawText("stats", menux+((menuw/2)+20), menuy+(menuh/2), 20, rl.Black)
		}
		if framecount%9 == 0 {
			if statsflash {
				statsflash = false
			} else {
				statsflash = true
			}
		}
	} else if infohover && infoon == false {
		if infoflash {
			rl.DrawRectangle(menux, menuy+((menuh/2)-2), menuw/2-3, 24, rl.White)
			rl.DrawText("info", menux+20, menuy+(menuh/2), 20, rl.Black)
		}
		if framecount%9 == 0 {
			if infoflash {
				infoflash = false
			} else {
				infoflash = true
			}
		}
	}

	if infoon {
		if invitemactiveon { // active item information
			switch invitemactive {
			case "spade":
				rl.DrawText("spade", menux+10, menuy+((menuh/2)+46), 20, rl.White)
				rl.DrawText("dig for treasure", menux+10, menuy+((menuh/2)+70), 20, rl.White)
				rl.BeginMode2D(camerainventory)
				xpos := float32((menux + 28) / 2)
				ypos := float32((menuy + ((menuh / 2) + 100)) / 2)
				v2 := rl.NewVector2(xpos, ypos)
				v3 := rl.NewVector2(xpos+30, ypos)
				v4 := rl.NewVector2(xpos+60, ypos)
				rl.DrawTextureRec(imgs, mousel, v2, rl.White)
				rl.DrawTextureRec(imgs, mouser, v3, rl.White)
				rl.DrawTextureRec(imgs, mousem, v4, rl.White)
				rl.EndMode2D()
				rl.DrawText("use", menux+22, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("exit", menux+80, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("drop", menux+135, menuy+((menuh/2)+135), 20, rl.White)
			case "pickaxe":
				rl.DrawText("pickaxe", menux+10, menuy+((menuh/2)+46), 20, rl.White)
				rl.DrawText("demolish walls", menux+10, menuy+((menuh/2)+70), 20, rl.White)
				rl.BeginMode2D(camerainventory)
				xpos := float32((menux + 28) / 2)
				ypos := float32((menuy + ((menuh / 2) + 100)) / 2)
				v2 := rl.NewVector2(xpos, ypos)
				v3 := rl.NewVector2(xpos+30, ypos)
				v4 := rl.NewVector2(xpos+60, ypos)
				rl.DrawTextureRec(imgs, mousel, v2, rl.White)
				rl.DrawTextureRec(imgs, mouser, v3, rl.White)
				rl.DrawTextureRec(imgs, mousem, v4, rl.White)
				rl.EndMode2D()
				rl.DrawText("use", menux+22, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("exit", menux+80, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("drop", menux+135, menuy+((menuh/2)+135), 20, rl.White)
			case "shotgun":
				rl.DrawText("shotgun", menux+10, menuy+((menuh/2)+46), 20, rl.White)
				rl.DrawText("pump action", menux+10, menuy+((menuh/2)+70), 20, rl.White)
				rl.BeginMode2D(camerainventory)
				xpos := float32((menux + 28) / 2)
				ypos := float32((menuy + ((menuh / 2) + 100)) / 2)
				v2 := rl.NewVector2(xpos, ypos)
				v3 := rl.NewVector2(xpos+30, ypos)
				v4 := rl.NewVector2(xpos+60, ypos)
				rl.DrawTextureRec(imgs, mousel, v2, rl.White)
				rl.DrawTextureRec(imgs, mouser, v3, rl.White)
				rl.DrawTextureRec(imgs, mousem, v4, rl.White)
				rl.EndMode2D()
				rl.DrawText("shoot", menux+22, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("exit", menux+80, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("drop", menux+135, menuy+((menuh/2)+135), 20, rl.White)
			case "food":
				rl.DrawText("food", menux+10, menuy+((menuh/2)+46), 20, rl.White)
				rl.DrawText("eat it", menux+10, menuy+((menuh/2)+70), 20, rl.White)
				rl.BeginMode2D(camerainventory)
				xpos := float32((menux + 28) / 2)
				ypos := float32((menuy + ((menuh / 2) + 100)) / 2)
				v2 := rl.NewVector2(xpos, ypos)
				v3 := rl.NewVector2(xpos+30, ypos)
				v4 := rl.NewVector2(xpos+60, ypos)
				rl.DrawTextureRec(imgs, mousel, v2, rl.White)
				rl.DrawTextureRec(imgs, mouser, v3, rl.White)
				rl.DrawTextureRec(imgs, mousem, v4, rl.White)
				rl.EndMode2D()
				rl.DrawText("eat", menux+22, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("exit", menux+80, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("drop", menux+135, menuy+((menuh/2)+135), 20, rl.White)
			case "potion":
				rl.DrawText("potion", menux+10, menuy+((menuh/2)+46), 20, rl.White)
				rl.DrawText("drink it", menux+10, menuy+((menuh/2)+70), 20, rl.White)
				rl.BeginMode2D(camerainventory)
				xpos := float32((menux + 28) / 2)
				ypos := float32((menuy + ((menuh / 2) + 100)) / 2)
				v2 := rl.NewVector2(xpos, ypos)
				v3 := rl.NewVector2(xpos+30, ypos)
				v4 := rl.NewVector2(xpos+60, ypos)
				rl.DrawTextureRec(imgs, mousel, v2, rl.White)
				rl.DrawTextureRec(imgs, mouser, v3, rl.White)
				rl.DrawTextureRec(imgs, mousem, v4, rl.White)
				rl.EndMode2D()
				rl.DrawText("drink", menux+22, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("exit", menux+80, menuy+((menuh/2)+135), 20, rl.White)
				rl.DrawText("drop", menux+135, menuy+((menuh/2)+135), 20, rl.White)
			}
		} else { // current weather
			rl.DrawText(weathercurrent, menux+((menuw/4)+5), menuy+((menuh/2)+46), 20, rl.White)
			rl.DrawText("day", (menux + 12), menuy+((menuh/2)+86), 20, rl.White)
			daynum := strconv.Itoa(days)
			coinnum := strconv.Itoa(coins)
			rl.DrawText(daynum, (menux + 56), menuy+((menuh/2)+86), 20, rl.White)                  // days
			rl.DrawText(coinnum, (menux + ((menuw / 2) + 40)), menuy+((menuh/2)+86), 20, rl.White) // coins
			v2 := rl.NewVector2(float32((menux+10)/2), float32((menuy+(menuh/2)+40)/2))            // weather img v2
			v3 := rl.NewVector2(float32((menux+(menuw/2)+6)/2), float32((menuy+(menuh/2)+80)/2))   // coin img v2
			rl.BeginMode2D(camerainventory)
			rl.DrawTextureRec(imgs, weatherimg, v2, rl.White)
			rl.DrawTextureRec(imgs, coin, v3, rl.White)
			rl.EndMode2D()
			//	rl.DrawText("help & information", menux+10, menuy+((menuh/2)+46), 20, rl.White)
			//	rl.DrawText("will display here", menux+10, menuy+((menuh/2)+70), 20, rl.White)

			rl.DrawText("no active quest", (menux + 14), menuy+((menuh/2)+124), 20, rl.White)
		}
	} else { // draw stats

		armorcountTEXT := strconv.Itoa(armorcount)
		attackcountTEXT := strconv.Itoa(attackcount)
		playerlevelTEXT := strconv.Itoa(playerlevel)
		rl.DrawText("armor", menux+12, menuy+((menuh/2)+46), 20, rl.White)
		rl.DrawText(armorcountTEXT, menux+menuw/2, menuy+((menuh/2)+46), 20, rl.White)
		rl.DrawText("attack", menux+12, menuy+((menuh/2)+70), 20, rl.White)
		rl.DrawText(attackcountTEXT, menux+menuw/2, menuy+((menuh/2)+70), 20, rl.White)
		rl.DrawText("level", menux+12, menuy+((menuh/2)+94), 20, rl.White)
		rl.DrawText(playerlevelTEXT, menux+menuw/2, menuy+((menuh/2)+94), 20, rl.White)
	}

	// menu 2
	rl.DrawRectangle(menu2x, menu2y, menu2w, menu2h, rl.Fade(rl.Black, 0.7)) // background
	//	rl.DrawRectangle(menux, menuy, menuw, 20, rl.White)
	rl.DrawRectangleLines(menu2x, menu2y, menu2w, menu2h, rl.White)                    // border lines
	rl.DrawRectangleLines(menu2x-1, menu2y-1, menu2w+2, menu2h+2, rl.White)            // border lines
	rl.DrawRectangle(menu2x+2, menu2y+menu2h+2, menu2w, 2, rl.DarkGray)                // bottom shadow
	rl.DrawRectangle(menu2x+2, menu2y+menu2h+2, menu2w, 2, rl.Fade(rl.Black, 0.5))     // bottom shadow
	rl.DrawRectangle(menu2x+menu2w+2, menu2y+18, 2, menu2h-14, rl.DarkGray)            // right shadow
	rl.DrawRectangle(menu2x+menu2w+2, menu2y+18, 2, menu2h-14, rl.Fade(rl.Black, 0.5)) // right shadow
	if movemenu2highlighton {
		rl.DrawRectangle(menu2x+menu2w, menu2y-1, 20, 20, rl.White)                 // move tab highlight
		rl.DrawRectangle(menu2x+menu2w, menu2y-1, 21, 21, rl.Fade(rl.SkyBlue, 0.6)) // move tab highlight
		rl.DrawRectangle(menu2x+menu2w, menu2y-1, 21, 21, rl.SkyBlue)               // move tab highlight
	} else {
		rl.DrawRectangle(menu2x+menu2w, menu2y-1, 20, 20, rl.White) // move tab
	}
	v22 := rl.NewVector2(float32(menu2x+menu2w+3), float32(menu2y+2)) // move menu2 img
	rl.DrawTextureRec(imgs, movemenuimg, v22, rl.White)               // move menu2 img
	// tile info text
	rl.DrawText(mousetileinfo, menu2x+20, menu2y+20, 20, rl.White)
	rl.DrawText("secret chance 90%", menu2x+20, menu2y+40, 20, rl.White)
	// icon row 1
	changex := menu2x + menu2w - 35
	for a := 0; a < 10; a++ {
		rl.DrawRectangleLines(changex-int32(a*36), menu2y+4, 32, 32, rl.White)
		invxmap[9-a] = changex - int32(a*36)
	}
	// icon row 2
	changex = menu2x + menu2w - 35
	for a := 0; a < 10; a++ {
		rl.DrawRectangleLines(changex-int32(a*36), menu2y+40, 32, 32, rl.White)
	}

	if inv1on { // background select rectangles
		rl.DrawRectangle(invxmap[0], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv2on {
		rl.DrawRectangle(invxmap[1], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv3on {
		rl.DrawRectangle(invxmap[2], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv4on {
		rl.DrawRectangle(invxmap[3], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv5on {
		rl.DrawRectangle(invxmap[4], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv6on {
		rl.DrawRectangle(invxmap[5], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv7on {
		rl.DrawRectangle(invxmap[6], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv8on {
		rl.DrawRectangle(invxmap[7], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv9on {
		rl.DrawRectangle(invxmap[8], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv10on {
		rl.DrawRectangle(invxmap[9], menu2y+4, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv11on {
		rl.DrawRectangle(invxmap[0], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv12on {
		rl.DrawRectangle(invxmap[1], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv13on {
		rl.DrawRectangle(invxmap[2], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv14on {
		rl.DrawRectangle(invxmap[3], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv15on {
		rl.DrawRectangle(invxmap[4], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv16on {
		rl.DrawRectangle(invxmap[5], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv17on {
		rl.DrawRectangle(invxmap[6], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv18on {
		rl.DrawRectangle(invxmap[7], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv19on {
		rl.DrawRectangle(invxmap[8], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	} else if inv20on {
		rl.DrawRectangle(invxmap[9], menu2y+38, 32, 32, rl.Fade(rl.White, 0.3))
	}

}
func animations() { // MARK: animations

	if framecount%6 == 0 {
		if fade1on {
			fade1 += 0.02
			if fade1 >= 0.4 {
				fade1on = false
			}
		} else {
			fade1 -= 0.02
			if fade1 <= 0.1 {
				fade1on = true
			}
		}
	}

	if framecount%3 == 0 {

		playerrunR.X += 16
		if playerrunR.X > 90 {
			playerrunR.X = 0
		}
		playerrunL.X += 16
		if playerrunL.X > 180 {
			playerrunL.X = 0
		}
		playeridleR.X += 16
		if playeridleR.X > 90 {
			playeridleR.X = 0
		}
		playeridleL.X += 16
		if playeridleL.X > 180 {
			playeridleL.X = 96
		}
	}
	if framecount%6 == 0 {

		if duckon {
			duckon = false
		} else {
			duckon = true
		}
		if unicornon {
			unicornon = false
		} else {
			unicornon = true
		}
		if caton {
			caton = false
		} else {
			caton = true
		}
		if butterflyon {
			butterflyon = false
		} else {
			butterflyon = true
		}
		if snailon {
			snailon = false
		} else {
			snailon = true
		}
		if manon {
			manon = false
		} else {
			manon = true
		}
		if snakeon {
			snakeon = false
		} else {
			snakeon = true
		}
		if pigeonon {
			pigeonon = false
		} else {
			pigeonon = true
		}
		if turtleon {
			turtleon = false
		} else {
			turtleon = true
		}
		if rabbiton {
			rabbiton = false
		} else {
			rabbiton = true
		}
		if pigon {
			pigon = false
		} else {
			pigon = true
		}
	}
}
func fx() { // MARK: fx
	scany := int32(0)
	for a := 0; a < monitorh; a++ {
		rl.DrawLine(0, scany, monw32, scany, rl.Fade(rl.Black, 0.4))
		scany += 2
		a++
	}
}
func horizvert() { // MARK: horzivert
	drawblocknexth, drawblocknextv = drawblocknext/levelw, drawblocknext%levelw
	playerh, playerv = player/levelw, player%levelw
	selectedblockh, selectedblockv = selectedblock/levelw, selectedblock%levelw
	blockh, blockv = block/levelw, block%levelw
}
func initialize() {
	rl.InitWindow(monw32, monh32, "supakastleqwest")
	setscreen()
	rl.CloseWindow()
	createmaps()
	createlevel()
	startsettings()
}
func startsettings() { // MARK:startsettings

	//	player = 15742
	// camera
	camera.Zoom = 2.0
	cameracursor.Zoom = 2.0
	cameraweather.Zoom = 2.0
	zoomlevel = 2
	// menus
	menux = 10
	menuy = 10
	menuw = 200
	menuh = 500
	menu2x = 250
	menu2y = 10
	menu2w = 800
	menu2h = 76
	optionsmenuh = 600
	optionsmenuw = 1200

	weathertimer = 10 // testing
	weathercurrent = "sunny"
	weatherimg = sun
	days = 1
}
func createmaps() { // MARK:createmaps
	questmap = make([]string, levela)
	equippeditemsstructmap = make([]items, 10)
	charactersmap = make([]string, levela)
	enemiesstructmap = make([]items, levela)
	enemiesmap = make([]string, levela)
	levelmap = make([]string, levela)
	leveltilesmap = make([]string, levela)
	extrasinteriormap = make([]string, levela)
	extrasexteriormap = make([]string, levela)
	itemsmap = make([]string, levela)
	items2map = make([]string, levela)
	items3map = make([]string, levela)
	itemactionsmap = make([]string, levela)
	weathermap = make([]string, gridw2*(gridh2*2))
	itemsstructmap = make([]items, levela)
}
func setscreen() { // MARK: setscreen
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)

	camera.Zoom = 1.0
	camera.Target.X = 0
	camera.Target.Y = 0

	camerainventory.Zoom = 2.0

	gridw = monitorw/16 + 1
	gridh = monitorh/16 + 1
	grida = gridw * gridh

	gridw2 = gridw / 2
	gridh2 = gridh / 2
	grida2 = gridw2 * gridh2

	gridw3 = gridw / 3
	gridh3 = gridh / 3
	grida3 = gridw2 * gridh2

	gridw4 = gridw / 4
	gridh4 = gridh / 4
	grida4 = gridw4 * gridh4

}
func clearactiveitem() { // MARK: clearactiveitem
	invitemactive = ""
	invitemactiveon = false
	selectinv1on, selectinv2on, selectinv3on, selectinv4on, selectinv5on, selectinv6on, selectinv7on, selectinv8on, selectinv9on, selectinv10on, selectinv11on, selectinv12on, selectinv13on, selectinv14on, selectinv15on, selectinv16on, selectinv2on, selectinv2on, selectinv2on, selectinv2on = false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false
	selectedslotnumber = 0
}
func interactexterior(itemblock int) { // MARK: interactexterior

	questmap[itemblock] = "quest"

	switch extrasexteriormap[itemblock] {

	case "tree1", "tree2", "tree3", "tree4", "tree5", "tree6", "tree7", "tree8", "tree9", "tree10", "tree11", "tree12", "tree13", "tree14", "tree15", "tree16":

		questmap[itemblock] = "tree"
	case "flower1", "flower2", "flower3":
		questmap[itemblock] = "flower"
	case "grass2", "grass3", "grass4":
		questmap[itemblock] = "grass"
	}

}
func input() { // MARK: input
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if !movemenu2on && !movemenuon {
			if movemenu2highlighton {
				movemenu2on = true
			} else if movemenuhighlighton {
				movemenuon = true
			} else if optionshighlighton {
				optionsmenuon = true
			} else if optionsmenuclosehighlighton {
				optionsmenuon = false
			} else if extrasexteriormap[mouseblock] != "" && !invitemactiveon {
				interactexterior(mouseblock)
			} else if inv1on && !invitemactiveon {
				selectinv1on = true
			} else if inv2on && !invitemactiveon {
				selectinv2on = true
			} else if inv3on && !invitemactiveon {
				selectinv3on = true
			} else if inv4on && !invitemactiveon {
				selectinv4on = true
			} else if inv5on && !invitemactiveon {
				selectinv5on = true
			} else if inv6on && !invitemactiveon {
				selectinv6on = true
			} else if inv7on && !invitemactiveon {
				selectinv7on = true
			} else if inv8on && !invitemactiveon {
				selectinv8on = true
			} else if inv9on && !invitemactiveon {
				selectinv9on = true
			} else if inv10on && !invitemactiveon {
				selectinv10on = true
			} else if inv11on && !invitemactiveon {
				selectinv11on = true
			} else if inv12on && !invitemactiveon {
				selectinv12on = true
			} else if inv13on && !invitemactiveon {
				selectinv13on = true
			} else if inv14on && !invitemactiveon {
				selectinv14on = true
			} else if inv15on && !invitemactiveon {
				selectinv15on = true
			} else if inv16on && !invitemactiveon {
				selectinv16on = true
			} else if inv17on && !invitemactiveon {
				selectinv17on = true
			} else if inv18on && !invitemactiveon {
				selectinv18on = true
			} else if inv19on && !invitemactiveon {
				selectinv19on = true
			} else if inv20on && !invitemactiveon {
				selectinv20on = true
			} else if invitemactiveon && inmenu1 == false && inmenu2 == false {
				itemactiveblock = mouseblock
				useitem()
			} else if items3map[mouseblock] == "chest" {
				items3map[mouseblock] = "chestopen"
				chestblock = mouseblock
				openchest()
			} else if statshover {
				statson = true
				infoon = false
			} else if infohover {
				infoon = true
				statson = false
			} else if placehelmet {
				equippeditemsstructmap[0] = invitemstructactive
				helmetslotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[0]
				armorcount += checkitem.damage
				clearactiveitem()
			} else if placenecklace {
				equippeditemsstructmap[1] = invitemstructactive
				necklaceslotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				clearactiveitem()
			} else if placering1 {
				equippeditemsstructmap[2] = invitemstructactive
				ring1slotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				clearactiveitem()
			} else if placearmor {
				equippeditemsstructmap[3] = invitemstructactive
				armorslotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[3]
				armorcount += checkitem.damage
				clearactiveitem()
			} else if placering2 {
				equippeditemsstructmap[4] = invitemstructactive
				ring2slotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				clearactiveitem()
			} else if placegloves {
				equippeditemsstructmap[5] = invitemstructactive
				glovesslotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[5]
				armorcount += checkitem.damage
				clearactiveitem()
			} else if placebelt {
				equippeditemsstructmap[6] = invitemstructactive
				beltslotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[6]
				armorcount += checkitem.damage
				clearactiveitem()
			} else if placeweapon1 {
				equippeditemsstructmap[7] = invitemstructactive
				weaponslot1equipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[7]
				if checkitem.itemtype == "armor" {
					armorcount += checkitem.damage
				}
				clearactiveitem()
			} else if placeweapon2 {
				equippeditemsstructmap[8] = invitemstructactive
				weaponslot2equipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[8]
				if checkitem.itemtype == "armor" {
					armorcount += checkitem.damage
				}
				clearactiveitem()
			} else if placeboots {
				equippeditemsstructmap[9] = invitemstructactive
				bootsslotequipped = true
				inventoryfullmap[selectedslotnumber] = false
				inventorystructsmap[selectedslotnumber] = items{}
				inventorymap[selectedslotnumber] = ""
				checkitem := equippeditemsstructmap[9]
				armorcount += checkitem.damage
				clearactiveitem()
			} else if helmetsloton && helmetslotequipped {
				invitemstructactive = equippeditemsstructmap[0]
				equippeditemsstructmap[0] = items{}
				helmetslotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
				armorcount -= invitemstructactive.damage
			} else if necklacesloton && necklaceslotequipped {
				invitemstructactive = equippeditemsstructmap[1]
				equippeditemsstructmap[1] = items{}
				necklaceslotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
			} else if ring1sloton && ring1slotequipped {
				invitemstructactive = equippeditemsstructmap[2]
				equippeditemsstructmap[2] = items{}
				ring1slotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
			} else if armorsloton && armorslotequipped {
				invitemstructactive = equippeditemsstructmap[3]
				equippeditemsstructmap[3] = items{}
				armorslotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
				armorcount -= invitemstructactive.damage
			} else if ring2sloton && ring2slotequipped {
				invitemstructactive = equippeditemsstructmap[4]
				equippeditemsstructmap[4] = items{}
				ring2slotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
			} else if glovessloton && glovesslotequipped {
				invitemstructactive = equippeditemsstructmap[5]
				equippeditemsstructmap[5] = items{}
				glovesslotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
				armorcount -= invitemstructactive.damage
			} else if beltsloton && beltslotequipped {
				invitemstructactive = equippeditemsstructmap[6]
				equippeditemsstructmap[6] = items{}
				beltslotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
				armorcount -= invitemstructactive.damage
			} else if weaponslot1on && weaponslot1equipped {
				invitemstructactive = equippeditemsstructmap[7]
				equippeditemsstructmap[7] = items{}
				weaponslot1equipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
			} else if weaponslot2on && weaponslot2equipped {
				invitemstructactive = equippeditemsstructmap[8]
				equippeditemsstructmap[8] = items{}
				weaponslot2equipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
			} else if bootssloton && bootsslotequipped {
				invitemstructactive = equippeditemsstructmap[9]
				equippeditemsstructmap[9] = items{}
				bootsslotequipped = false
				invitemactive = invitemstructactive.itemtype
				invitemactiveon = true
				armorcount -= invitemstructactive.damage
			} else if invitemactiveon && inv1on && inventoryfullmap[0] == false {
				inventorystructsmap[0] = invitemstructactive
				inventoryfullmap[0] = true
				inventorymap[0] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv2on && inventoryfullmap[1] == false {
				inventorystructsmap[1] = invitemstructactive
				inventoryfullmap[1] = true
				inventorymap[1] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv3on && inventoryfullmap[2] == false {
				inventorystructsmap[2] = invitemstructactive
				inventoryfullmap[2] = true
				inventorymap[2] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv4on && inventoryfullmap[3] == false {
				inventorystructsmap[3] = invitemstructactive
				inventoryfullmap[3] = true
				inventorymap[3] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv5on && inventoryfullmap[4] == false {
				inventorystructsmap[4] = invitemstructactive
				inventoryfullmap[4] = true
				inventorymap[4] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv6on && inventoryfullmap[5] == false {
				inventorystructsmap[5] = invitemstructactive
				inventoryfullmap[5] = true
				inventorymap[5] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv7on && inventoryfullmap[6] == false {
				inventorystructsmap[6] = invitemstructactive
				inventoryfullmap[6] = true
				inventorymap[6] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv8on && inventoryfullmap[7] == false {
				inventorystructsmap[7] = invitemstructactive
				inventoryfullmap[7] = true
				inventorymap[7] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv9on && inventoryfullmap[8] == false {
				inventorystructsmap[8] = invitemstructactive
				inventoryfullmap[8] = true
				inventorymap[8] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv10on && inventoryfullmap[9] == false {
				inventorystructsmap[9] = invitemstructactive
				inventoryfullmap[9] = true
				inventorymap[9] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv11on && inventoryfullmap[10] == false {
				inventorystructsmap[10] = invitemstructactive
				inventoryfullmap[10] = true
				inventorymap[10] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv12on && inventoryfullmap[11] == false {
				inventorystructsmap[11] = invitemstructactive
				inventoryfullmap[11] = true
				inventorymap[11] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv13on && inventoryfullmap[12] == false {
				inventorystructsmap[12] = invitemstructactive
				inventoryfullmap[12] = true
				inventorymap[12] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv14on && inventoryfullmap[13] == false {
				inventorystructsmap[13] = invitemstructactive
				inventoryfullmap[13] = true
				inventorymap[13] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv15on && inventoryfullmap[14] == false {
				inventorystructsmap[14] = invitemstructactive
				inventoryfullmap[14] = true
				inventorymap[14] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv16on && inventoryfullmap[15] == false {
				inventorystructsmap[15] = invitemstructactive
				inventoryfullmap[15] = true
				inventorymap[15] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv17on && inventoryfullmap[16] == false {
				inventorystructsmap[16] = invitemstructactive
				inventoryfullmap[16] = true
				inventorymap[16] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv18on && inventoryfullmap[17] == false {
				inventorystructsmap[17] = invitemstructactive
				inventoryfullmap[17] = true
				inventorymap[17] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv19on && inventoryfullmap[18] == false {
				inventorystructsmap[18] = invitemstructactive
				inventoryfullmap[18] = true
				inventorymap[18] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else if invitemactiveon && inv20on && inventoryfullmap[19] == false {
				inventorystructsmap[19] = invitemstructactive
				inventoryfullmap[19] = true
				inventorymap[19] = invitemstructactive.itemtype
				invitemstructactive = items{}
				invitemactive = ""
				invitemactiveon = false
			} else {
				if !optionsmenuon && !inmenu1 && !inmenu2 {
					selectedblock = mouseblock
				}
			}

		} else if movemenuon {
			menux = int32(mousepos.X) - menuw
			if menux <= 1 {
				menux = 2
			} else if menux > monw32-(menuw+22) {
				menux = monw32 - (menuw + 22)
			}
			menuy = int32(mousepos.Y)
			if menuy > monh32-(menuh+6) {
				menuy = monh32 - (menuh + 6)
			} else if menuy < 1 {
				menuy = 1
			}
			movemenuon = false
		} else if movemenu2on {
			menu2x = int32(mousepos.X) - menu2w
			if menu2x <= 1 {
				menu2x = 2
			} else if menu2x > monw32-(menu2w+22) {
				menu2x = monw32 - (menu2w + 22)
			}
			menu2y = int32(mousepos.Y)
			if menu2y > monh32-(menu2h+6) {
				menu2y = monh32 - (menu2h + 6)
			} else if menu2y < 1 {
				menu2y = 1
			}
			movemenu2on = false
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		if invitemactiveon {
			clearactiveitem()
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseMiddleButton) {
		if invitemactiveon {
			if levelmap[mouseblock] == "floor" || levelmap[mouseblock] == "." {
				if selectinv1on {
					inventorymap[0] = ""
					inventoryfullmap[0] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv2on {
					inventorymap[1] = ""
					inventoryfullmap[1] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv3on {
					inventorymap[2] = ""
					inventoryfullmap[2] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv4on {
					inventorymap[3] = ""
					inventoryfullmap[3] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv5on {
					inventorymap[4] = ""
					inventoryfullmap[4] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv6on {
					inventorymap[5] = ""
					inventoryfullmap[5] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv7on {
					inventorymap[6] = ""
					inventoryfullmap[6] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv8on {
					inventorymap[7] = ""
					inventoryfullmap[7] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv9on {
					inventorymap[8] = ""
					inventoryfullmap[8] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv10on {
					inventorymap[9] = ""
					inventoryfullmap[9] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv11on {
					inventorymap[10] = ""
					inventoryfullmap[10] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv12on {
					inventorymap[11] = ""
					inventoryfullmap[11] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv13on {
					inventorymap[12] = ""
					inventoryfullmap[12] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv14on {
					inventorymap[13] = ""
					inventoryfullmap[13] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv15on {
					inventorymap[14] = ""
					inventoryfullmap[14] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv16on {
					inventorymap[15] = ""
					inventoryfullmap[15] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv17on {
					inventorymap[16] = ""
					inventoryfullmap[16] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv18on {
					inventorymap[17] = ""
					inventoryfullmap[17] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv19on {
					inventorymap[18] = ""
					inventoryfullmap[18] = false
					itemsmap[mouseblock] = invitemactive
				} else if selectinv20on {
					inventorymap[19] = ""
					inventoryfullmap[19] = false
					itemsmap[mouseblock] = invitemactive
				}

				clearactiveitem()
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if colorson {
			colorson = false
		} else {
			colorson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
			zoomlevel = 2
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 3.0
			zoomlevel = 3
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 4.0
			zoomlevel = 4
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 4.0 {
			camera.Zoom = 3.0
			zoomlevel = 3
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 2.0
			zoomlevel = 2
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
			zoomlevel = 1
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Blue, 0.4))
	rl.DrawFPS(monw32-290, monh32-100)

	gridwTEXT := strconv.Itoa(gridw)
	gridhTEXT := strconv.Itoa(gridh)
	gridaTEXT := strconv.Itoa(grida)
	drawblocknextTEXT := strconv.Itoa(drawblocknext)
	playerTEXT := strconv.Itoa(player)
	playervTEXT := strconv.Itoa(playerv)
	playerhTEXT := strconv.Itoa(playerh)
	mousexTEXT := fmt.Sprintf("%g", mousepos.X)
	mouseyTEXT := fmt.Sprintf("%g", mousepos.Y)
	mouseblockTEXT := strconv.Itoa(mouseblock)
	levelhTEXT := strconv.Itoa(levelh)
	levelwTEXT := strconv.Itoa(levelw)
	levelaTEXT := strconv.Itoa(levela)
	zoomlevelTEXT := strconv.Itoa(zoomlevel)
	optionsmenuclosehighlightonTEXT := strconv.FormatBool(optionsmenuclosehighlighton)
	path1TEXT := strconv.FormatBool(path1)
	path2TEXT := strconv.FormatBool(path2)
	path3TEXT := strconv.FormatBool(path3)
	path4TEXT := strconv.FormatBool(path4)
	invxmap9TEXT := fmt.Sprint(invxmap[9])
	invitemactiveonTEXT := strconv.FormatBool(invitemactiveon)
	selectinv1onTEXT := strconv.FormatBool(selectinv1on)
	playerxTEXT := strconv.Itoa(int(playerx))
	playeryTEXT := strconv.Itoa(int(playery))
	helmetslotonTEXT := strconv.FormatBool(helmetsloton)
	necklaceslotonTEXT := strconv.FormatBool(necklacesloton)
	armorslotonTEXT := strconv.FormatBool(armorsloton)
	beltslotonTEXT := strconv.FormatBool(beltsloton)
	ring1slotonTEXT := strconv.FormatBool(ring1sloton)
	ring2slotonTEXT := strconv.FormatBool(ring2sloton)
	weaponslot1onTEXT := strconv.FormatBool(weaponslot1on)
	weaponslot2onTEXT := strconv.FormatBool(weaponslot2on)
	bootsslotonTEXT := strconv.FormatBool(bootssloton)
	glovesslotonTEXT := strconv.FormatBool(glovessloton)
	statshoverTEXT := strconv.FormatBool(statshover)
	infohoverTEXT := strconv.FormatBool(infohover)
	placeweapon1TEXT := strconv.FormatBool(placeweapon1)
	weaponslot1equippedTEXT := strconv.FormatBool(weaponslot1equipped)
	inmenu1TEXT := strconv.FormatBool(inmenu1)
	inmenu2TEXT := strconv.FormatBool(inmenu2)
	checkinvitem := inventorystructsmap[0]
	inv1onTEXT := strconv.FormatBool(inv1on)

	rl.DrawText(gridwTEXT, monw32-290, 10, 10, rl.White)
	rl.DrawText("gridw", monw32-200, 10, 10, rl.White)
	rl.DrawText(gridhTEXT, monw32-290, 20, 10, rl.White)
	rl.DrawText("gridh", monw32-200, 20, 10, rl.White)
	rl.DrawText(gridaTEXT, monw32-290, 30, 10, rl.White)
	rl.DrawText("grida", monw32-200, 30, 10, rl.White)
	rl.DrawText(levelhTEXT, monw32-290, 40, 10, rl.White)
	rl.DrawText("levelh", monw32-200, 40, 10, rl.White)
	rl.DrawText(levelwTEXT, monw32-290, 50, 10, rl.White)
	rl.DrawText("levelw", monw32-200, 50, 10, rl.White)
	rl.DrawText(levelaTEXT, monw32-290, 60, 10, rl.White)
	rl.DrawText("levela", monw32-200, 60, 10, rl.White)
	rl.DrawText(drawblocknextTEXT, monw32-290, 70, 10, rl.White)
	rl.DrawText("drawblocknext", monw32-200, 70, 10, rl.White)
	rl.DrawText(playerTEXT, monw32-290, 80, 10, rl.White)
	rl.DrawText("player", monw32-200, 80, 10, rl.White)
	rl.DrawText(playervTEXT, monw32-290, 90, 10, rl.White)
	rl.DrawText("playerv", monw32-200, 90, 10, rl.White)
	rl.DrawText(playerhTEXT, monw32-290, 100, 10, rl.White)
	rl.DrawText("playerh", monw32-200, 100, 10, rl.White)
	rl.DrawText(mousexTEXT, monw32-290, 110, 10, rl.White)
	rl.DrawText("mousex", monw32-200, 110, 10, rl.White)
	rl.DrawText(mouseyTEXT, monw32-290, 120, 10, rl.White)
	rl.DrawText("mousey", monw32-200, 120, 10, rl.White)
	rl.DrawText(mouseblockTEXT, monw32-290, 130, 10, rl.White)
	rl.DrawText("mouseblock", monw32-200, 130, 10, rl.White)
	rl.DrawText(zoomlevelTEXT, monw32-290, 140, 10, rl.White)
	rl.DrawText("zoomlevel", monw32-200, 140, 10, rl.White)
	rl.DrawText(optionsmenuclosehighlightonTEXT, monw32-290, 150, 10, rl.White)
	rl.DrawText("optionsmenuclosehighlighton", monw32-200, 150, 10, rl.White)
	rl.DrawText(path1TEXT, monw32-290, 160, 10, rl.White)
	rl.DrawText("path1", monw32-200, 160, 10, rl.White)
	rl.DrawText(path2TEXT, monw32-290, 170, 10, rl.White)
	rl.DrawText("path2", monw32-200, 170, 10, rl.White)
	rl.DrawText(path3TEXT, monw32-290, 180, 10, rl.White)
	rl.DrawText("path3", monw32-200, 180, 10, rl.White)
	rl.DrawText(path4TEXT, monw32-290, 190, 10, rl.White)
	rl.DrawText("path4", monw32-200, 190, 10, rl.White)
	rl.DrawText(invxmap9TEXT, monw32-290, 200, 10, rl.White)
	rl.DrawText("invxmap9", monw32-200, 200, 10, rl.White)
	rl.DrawText(invitemactiveonTEXT, monw32-290, 210, 10, rl.White)
	rl.DrawText("invitemactiveon", monw32-200, 210, 10, rl.White)
	rl.DrawText(selectinv1onTEXT, monw32-290, 220, 10, rl.White)
	rl.DrawText("selectinv1on", monw32-200, 220, 10, rl.White)
	rl.DrawText("selectinv1on", monw32-200, 220, 10, rl.White)
	rl.DrawText(invitemactive, monw32-290, 230, 10, rl.White)
	rl.DrawText("invitemactive", monw32-200, 230, 10, rl.White)
	rl.DrawText(playerxTEXT, monw32-290, 240, 10, rl.White)
	rl.DrawText("playerx", monw32-200, 240, 10, rl.White)
	rl.DrawText(playeryTEXT, monw32-290, 250, 10, rl.White)
	rl.DrawText("playery", monw32-200, 250, 10, rl.White)
	rl.DrawText(helmetslotonTEXT, monw32-290, 260, 10, rl.White)
	rl.DrawText("helmetsloton", monw32-200, 260, 10, rl.White)
	rl.DrawText(glovesslotonTEXT, monw32-290, 270, 10, rl.White)
	rl.DrawText("glovessloton", monw32-200, 270, 10, rl.White)
	rl.DrawText(necklaceslotonTEXT, monw32-290, 280, 10, rl.White)
	rl.DrawText("necklacesloton", monw32-200, 280, 10, rl.White)
	rl.DrawText(armorslotonTEXT, monw32-290, 290, 10, rl.White)
	rl.DrawText("armorsloton", monw32-200, 290, 10, rl.White)
	rl.DrawText(beltslotonTEXT, monw32-290, 300, 10, rl.White)
	rl.DrawText("beltsloton", monw32-200, 300, 10, rl.White)
	rl.DrawText(ring1slotonTEXT, monw32-290, 310, 10, rl.White)
	rl.DrawText("ring1sloton", monw32-200, 310, 10, rl.White)
	rl.DrawText(ring2slotonTEXT, monw32-290, 320, 10, rl.White)
	rl.DrawText("ring2sloton", monw32-200, 320, 10, rl.White)
	rl.DrawText(weaponslot1onTEXT, monw32-290, 330, 10, rl.White)
	rl.DrawText("weaponslot1on", monw32-200, 330, 10, rl.White)
	rl.DrawText(weaponslot2onTEXT, monw32-290, 340, 10, rl.White)
	rl.DrawText("weaponslot2on", monw32-200, 340, 10, rl.White)
	rl.DrawText(bootsslotonTEXT, monw32-290, 350, 10, rl.White)
	rl.DrawText("bootssloton", monw32-200, 350, 10, rl.White)
	rl.DrawText(infohoverTEXT, monw32-290, 360, 10, rl.White)
	rl.DrawText("infohover", monw32-200, 360, 10, rl.White)
	rl.DrawText(statshoverTEXT, monw32-290, 370, 10, rl.White)
	rl.DrawText("statshover", monw32-200, 370, 10, rl.White)
	rl.DrawText(placeweapon1TEXT, monw32-290, 380, 10, rl.White)
	rl.DrawText("placeweapon1", monw32-200, 380, 10, rl.White)
	rl.DrawText(weaponslot1equippedTEXT, monw32-290, 390, 10, rl.White)
	rl.DrawText("weaponslot1equipped", monw32-200, 390, 10, rl.White)
	rl.DrawText(inmenu1TEXT, monw32-290, 400, 10, rl.White)
	rl.DrawText("inmenu1", monw32-200, 400, 10, rl.White)
	rl.DrawText(inmenu2TEXT, monw32-290, 410, 10, rl.White)
	rl.DrawText("inmenu2", monw32-200, 410, 10, rl.White)
	rl.DrawText(checkinvitem.itemtype, monw32-290, 420, 10, rl.White)
	rl.DrawText("inven item 0", monw32-200, 420, 10, rl.White)
	rl.DrawText(inv1onTEXT, monw32-290, 430, 10, rl.White)
	rl.DrawText("inv1on", monw32-200, 430, 10, rl.White)

}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
