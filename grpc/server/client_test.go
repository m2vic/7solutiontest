package main

import (
	"context"
	"log"
	"testing"

	pb "gofortest/grpc/proto"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterMeatServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetMeatAmounts(t *testing.T) {
	ctx := context.Background()
	//resolver.SetDefaultScheme("passthrough")

	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufffnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMeatServiceClient(conn)

	testCase := `Dolore ea aliqua tenderloin.  Irure biltong sint, ex proident drumstick andouille.  Cupidatat andouille spare ribs officia.  Mollit dolor flank rump sint pork.  Pastrami in consequat spare ribs meatloaf rump dolore adipisicing enim mollit pork belly eu eiusmod quis.

	Dolore aliqua pork exercitation sint.  Chicken capicola andouille beef ribs drumstick esse pig shank swine reprehenderit.  Irure capicola culpa short ribs labore.  Anim sausage ullamco shoulder, dolor ut ad pariatur short ribs eu ham hock in.  Cupim tempor tri-tip ut, pork loin ut kevin pig consectetur ea ad filet mignon.  Laboris deserunt shank, ut in fugiat prosciutto ham hock kielbasa chislic ham.
	
	Laborum meatball velit porchetta, pig sunt buffalo spare ribs aliqua strip steak beef ribs jerky.  Picanha ground round laborum frankfurter, ut culpa shankle pork pig jowl ribeye porchetta.  T-bone corned beef buffalo ex, meatloaf exercitation pastrami minim biltong do hamburger boudin ribeye occaecat.  Pork belly ad beef bacon.
	
	Ball tip porchetta landjaeger sed non.  Ground round boudin meatloaf ut, turkey bresaola culpa occaecat picanha.  Sint pariatur dolor qui, meatloaf hamburger velit.  Occaecat pariatur sunt, sint doner shankle et velit flank duis.  Irure aliquip swine, pork chop filet mignon capicola strip steak adipisicing qui labore salami.
	
	Sausage shankle flank, tenderloin eiusmod aliqua pork loin reprehenderit ham hock andouille strip steak tail.  Non ut spare ribs filet mignon elit, enim irure proident lorem.  Sed dolore cow officia, tempor exercitation kevin doner filet mignon.  Eiusmod sunt meatball ad landjaeger irure ham.  Andouille cow picanha enim dolore.
	
	Veniam shoulder sunt ut sirloin labore do spare ribs swine pastrami picanha landjaeger officia elit cupidatat.  Boudin minim leberkas, tempor kielbasa ut pig incididunt mollit ribeye sint porchetta kevin shank flank.  Ex lorem culpa exercitation in.  Pastrami ullamco est, fatback ex laboris kielbasa short ribs ut irure non tempor biltong shoulder in.  Eu sed t-bone salami, strip steak turducken eiusmod.  Kielbasa doner boudin salami ut nostrud buffalo irure.
	
	Meatball biltong nisi ball tip magna in incididunt commodo anim.  Pork belly leberkas pork filet mignon labore, spare ribs qui beef.  Jowl porchetta chuck venison corned beef tempor alcatra consectetur picanha buffalo andouille et short loin.  Mollit ham hock short ribs tri-tip shank reprehenderit id labore shoulder.
	
	In eiusmod ex, pig minim kevin do burgdoggen nisi reprehenderit swine short ribs beef nulla cupim.  Ut pastrami laborum, bacon turkey brisket dolor commodo.  Frankfurter magna enim sunt ground round.  Jerky brisket buffalo sausage fugiat, anim frankfurter pork chop mollit.  Kielbasa ham est pariatur, deserunt alcatra excepteur turducken ullamco lorem.  Meatball capicola consectetur laboris est labore deserunt sunt sausage.  Ham hock lorem commodo, shank est non eu pork sunt salami bacon eiusmod corned beef.
	
	T-bone aliqua cupim irure chuck deserunt.  Ground round salami laboris capicola pancetta chicken corned beef kielbasa id eiusmod leberkas.  Short ribs jowl ullamco in id.  Velit capicola spare ribs exercitation fatback doner.  Sed beef ribs eu tongue ut.  Kevin aute nostrud aliquip ut labore, strip steak filet mignon quis.
	
	Drumstick proident laboris dolore t-bone, salami beef flank.  Landjaeger ham duis rump short ribs.  Jerky ut magna do.  Aliqua do excepteur prosciutto.  Spare ribs consequat mollit chuck turkey occaecat beef rump do jowl corned beef capicola laboris velit in.  Jowl chislic duis sunt, chicken strip steak aliqua pork belly non laborum leberkas pork chop.
	
	Nulla ut ipsum ad duis brisket meatloaf beef ribs magna cillum shankle bresaola consequat.  Elit nulla biltong ut, ad occaecat qui in voluptate consequat porchetta doner cupim.  Ham hock ullamco dolore, brisket officia et beef ribs leberkas turkey.  Sed short loin dolore swine, consectetur brisket doner in.  Consequat jerky est dolor brisket prosciutto tenderloin t-bone.  Corned beef ullamco kevin, buffalo swine hamburger incididunt tempor anim alcatra ipsum.
	
	Eiusmod fugiat corned beef doner.  Nostrud pork loin aliquip hamburger dolor.  Et shoulder andouille, pork belly aute leberkas dolore bresaola shankle sausage incididunt ground round esse.  Officia shankle pariatur, meatball quis picanha voluptate ad kevin chicken turducken.
	
	Tenderloin aliquip officia consectetur adipisicing tempor est corned beef buffalo dolore short ribs cupim picanha t-bone ribeye.  Sausage beef tenderloin, culpa non tempor spare ribs short ribs incididunt burgdoggen pork loin cillum in occaecat.  Pancetta drumstick tail salami non fugiat.  Turducken ut ad meatball, boudin flank sunt sint est non corned beef prosciutto in mollit labore.
	
	Laborum pancetta pariatur qui consequat prosciutto ribeye pork belly leberkas laboris irure jerky sint bacon.  Chuck cupidatat velit proident esse brisket irure kielbasa prosciutto do.  Rump occaecat beef ribs picanha consectetur est bacon cupim.  Deserunt venison ea commodo sausage.
	
	Velit doner shoulder, capicola shankle elit sed ut filet mignon venison.  Commodo alcatra turkey drumstick, beef ribs ea lorem biltong ad shoulder strip steak.  Excepteur salami minim ribeye corned beef in.  Bresaola beef ribs short loin meatball ut dolore, ball tip exercitation adipisicing boudin veniam in fugiat.  Et chuck shank jowl rump proident fatback cupidatat kielbasa consectetur.  Bacon beef sirloin ut bresaola consequat strip steak leberkas minim.
	
	Drumstick cupidatat doner corned beef, esse consectetur tongue jowl.  Enim incididunt deserunt, pariatur officia shankle tempor magna irure rump voluptate turducken et.  Eiusmod beef ribs chislic, voluptate aliqua pig laborum.  Consequat mollit kevin exercitation, deserunt beef jerky laboris capicola t-bone biltong turducken ullamco strip steak landjaeger.  Corned beef prosciutto esse jowl.  Tongue aliqua bacon cillum boudin eiusmod tri-tip do dolor.
	
	Sint pastrami beef pariatur.  Anim meatball et landjaeger.  Beef ribs pastrami fugiat ipsum laboris.  Cupim dolore commodo voluptate capicola ullamco ham hock doner.
	
	Aliquip pork belly frankfurter salami.  Drumstick filet mignon aliqua chicken.  Capicola excepteur leberkas in, reprehenderit in kielbasa veniam brisket.  Cupidatat turkey lorem, quis spare ribs venison ullamco alcatra pariatur cillum esse ex chicken.  Meatloaf aliquip elit id fatback ad consectetur.  Id meatloaf enim cow, consectetur ribeye prosciutto filet mignon spare ribs bresaola shoulder culpa buffalo ball tip.
	
	Do prosciutto in, bacon aliquip lorem jerky.  Labore duis proident non laboris t-bone in aliquip tongue bresaola hamburger meatloaf venison officia.  Capicola biltong voluptate jowl, hamburger et leberkas.  Drumstick adipisicing hamburger fatback boudin, ribeye strip steak proident cupidatat nisi chicken.  Andouille beef sausage duis, in in kielbasa aliquip sed dolore officia t-bone.  Sed sausage cillum salami buffalo elit landjaeger pig in shank ad jowl strip steak id.
	
	Ipsum cillum ut, shank nisi reprehenderit pariatur pork picanha tenderloin minim ut venison jerky eiusmod.  Ut pig pancetta, alcatra voluptate prosciutto occaecat reprehenderit andouille laboris magna capicola boudin.  Bacon kevin meatball, irure picanha capicola fugiat shankle brisket.  Velit voluptate beef, ullamco irure cupidatat picanha buffalo ribeye qui landjaeger biltong aliquip pig kevin.
	
	Jowl ut eiusmod, ea kevin short ribs adipisicing.  Ham hock est boudin nostrud cupidatat.  Ut strip steak culpa deserunt chislic in, shank ribeye cillum esse ham cupidatat eiusmod aliquip.  Dolor chicken in sirloin.  Sint frankfurter pork loin commodo ad bacon.  Irure pig biltong hamburger id.  Magna adipisicing salami ball tip in, shoulder fatback andouille ad porchetta.
	
	Rump pastrami dolor ex incididunt.  Ut rump labore, porchetta ea shankle incididunt drumstick tongue chuck tail est.  Bresaola quis boudin bacon id strip steak eiusmod short ribs.  Do pig turducken quis boudin, magna biltong bacon laboris spare ribs sausage ullamco.  Tri-tip turkey ullamco, reprehenderit mollit ad anim.
	
	Sunt nostrud incididunt aliqua pariatur.  Do kevin nostrud, meatloaf rump t-bone frankfurter nisi reprehenderit pig jowl aliquip.  Bresaola tempor pariatur in.  Shank id swine kevin consectetur qui t-bone enim sed.
	
	Meatball excepteur proident qui aliquip porchetta.  Enim aliquip est shankle sed sint nostrud laborum ut ipsum.  Cupidatat tempor nulla voluptate pork bresaola.  Sunt beef ribs bacon, bresaola t-bone laborum pariatur swine reprehenderit tail.
	
	Burgdoggen consequat in ham hock consectetur labore.  Ipsum kielbasa chicken excepteur, boudin burgdoggen ball tip cupim ribeye id pig.  Ball tip excepteur quis anim, pork magna id sed sint laboris sirloin.  Beef ribs ground round in jowl.  Cillum tongue strip steak proident anim fugiat shank labore.  Sed swine prosciutto ribeye consectetur veniam.  Labore ut nulla boudin, culpa occaecat dolore prosciutto.
	
	Strip steak turducken eiusmod doner picanha.  Landjaeger dolore short loin quis, eu adipisicing picanha reprehenderit mollit.  Ipsum bacon irure pork non aliquip sed picanha prosciutto tenderloin id in qui ham sint.  Voluptate aliqua boudin non chislic et.  Chuck commodo tempor quis anim.  Biltong andouille prosciutto pork cow chicken.
	
	Dolor duis culpa spare ribs lorem.  Alcatra ut nisi, tongue bacon pork drumstick.  Turducken ham chicken ullamco quis mollit, enim kevin veniam chislic frankfurter biltong labore.  Fugiat ribeye pork loin, dolor tenderloin hamburger commodo.  Jerky non capicola cillum elit, pork loin landjaeger lorem pork belly eiusmod doner duis short ribs shankle spare ribs.  Andouille consequat eu ut beef ribs shankle, capicola non pig rump quis sunt leberkas.  Quis et laboris kevin enim anim nostrud exercitation shankle ham hock.
	
	Frankfurter lorem cillum salami mollit ham.  Cillum et ad doner mollit, commodo labore.  Adipisicing duis bresaola est, biltong tempor laborum laboris in short ribs tongue consequat.  Flank labore dolor meatball short loin minim, aliqua tail porchetta pork exercitation chicken.  Dolore shankle frankfurter, prosciutto non culpa officia.  Do short ribs aliquip ham hock officia capicola pork belly pariatur pork.
	
	Picanha burgdoggen frankfurter brisket.  Porchetta consequat capicola andouille aliqua.  Labore aliqua incididunt chislic ball tip irure short loin pariatur.  Proident ham ground round aliquip et ut pig pork chop.  Cillum reprehenderit exercitation occaecat alcatra shankle, tri-tip tempor jowl lorem.  Shankle biltong buffalo, cow shank drumstick ham consectetur tri-tip strip steak turducken.  Pig andouille jerky cow eu.
	
	Ut sirloin filet mignon veniam, ea pork chicken picanha turducken t-bone duis.  Hamburger boudin flank, duis beef bresaola cupim minim.  Shankle alcatra t-bone, esse sint duis brisket non cow aliqua ipsum.  Nulla bacon magna brisket strip steak, ut occaecat.  Capicola flank sunt, rump shank laborum aute pig voluptate.  Rump jowl in, short loin sint nisi spare ribs beef ribs labore salami ut burgdoggen incididunt.  Pariatur minim sunt meatball pancetta landjaeger aute capicola incididunt corned beef.
	
	Shoulder ad dolore consequat mollit id, aliquip picanha porchetta rump pariatur frankfurter ground round consectetur ut.  Enim dolore meatloaf drumstick sed.  Nostrud tri-tip ex beef ribs cillum picanha, frankfurter drumstick qui kevin dolore hamburger.  Et lorem tri-tip short loin burgdoggen dolor.
	
	Pork esse tri-tip, sint quis ground round doner aute nulla eiusmod boudin frankfurter.  Eu capicola chuck leberkas.  Reprehenderit jerky ex turkey culpa strip steak brisket voluptate sed velit shank spare ribs proident aliquip eiusmod.  Cillum andouille consectetur in tongue nisi deserunt ball tip pork chop shank bresaola.  Magna ham proident hamburger.
	
	T-bone pork turkey fugiat, eu hamburger ball tip.  Tri-tip ea qui in, reprehenderit chicken ut ribeye ham hock kielbasa pig aute.  Leberkas occaecat jerky sunt landjaeger turducken quis ball tip.  Filet mignon picanha jowl brisket ribeye short loin leberkas in id mollit chuck in ullamco.  Est aliquip in veniam id pig.
	
	Ball tip magna fugiat adipisicing.  Sunt short ribs cupidatat pork belly, mollit culpa sirloin anim ea.  Qui officia short loin quis leberkas strip steak irure.  In aliquip nostrud doner sirloin ullamco salami tongue aute jerky dolore id rump.
	
	Corned beef chuck dolore est short loin bresaola.  Biltong ut kielbasa, ham laborum aliqua capicola fatback sint.  Flank id tri-tip ground round short ribs short loin pork belly leberkas pancetta occaecat tempor ham.  Kevin pariatur sausage frankfurter.
	
	Voluptate ex cupidatat landjaeger tenderloin jowl elit occaecat duis et burgdoggen sirloin shoulder.  Boudin tongue eu, enim esse incididunt ut shoulder irure turducken nisi.  Esse eu ad, turducken landjaeger ex adipisicing occaecat dolore.  Chuck cupim proident jerky enim occaecat.  Eiusmod chicken spare ribs turducken ea dolore short loin fugiat non beef ribs.  Burgdoggen pork frankfurter, filet mignon sed drumstick dolore.  Turkey ullamco ea hamburger pancetta, jowl pastrami.
	
	Lorem pork capicola, shank cow reprehenderit exercitation chislic.  Non meatball jerky et.  Pig ut tail beef.  Pork loin consequat landjaeger id elit venison, bresaola ea pariatur cow ut.  Nostrud leberkas capicola pariatur shoulder reprehenderit pork chuck.  Fugiat shank tempor rump.  Shankle bresaola bacon ad chicken cillum velit burgdoggen deserunt est doner sint boudin do duis.
	
	Chislic dolore aliquip, jowl sint voluptate consectetur brisket duis eu cillum strip steak.  Tri-tip capicola ut salami, aliquip do in.  Ham alcatra picanha doner ut adipisicing hamburger kielbasa meatloaf anim deserunt excepteur.  Salami elit chuck, ribeye shankle duis doner ground round id officia jowl.  Ea laboris swine fugiat tongue turducken deserunt commodo strip steak jowl.
	
	Id tail mollit hamburger est.  Enim in bacon boudin, leberkas pork chop pork labore pork loin.  Filet mignon bresaola brisket, hamburger corned beef aliquip turkey lorem in shoulder.  Do eu porchetta, strip steak ut ullamco capicola.  Fugiat tenderloin sausage burgdoggen alcatra non jowl.  Quis lorem in, magna dolor chicken jowl ad shank t-bone sunt pig pork loin chuck.
	
	Voluptate pork belly doner in cupim.  Commodo cow elit, exercitation pork chop aliqua salami in.  Ut do excepteur pork belly turkey jowl occaecat ball tip shoulder laborum andouille pig pariatur.  Fatback rump pork chop, pastrami adipisicing est venison.  Bresaola ground round dolore, tail mollit in spare ribs bacon in boudin fugiat.  Exercitation prosciutto enim, capicola ad buffalo bresaola pariatur excepteur pig.  Short loin ad strip steak, pork belly cillum culpa ground round deserunt alcatra jerky duis.
	
	Biltong voluptate non landjaeger shankle dolore ham hock pastrami.  Tri-tip boudin et, corned beef shankle chislic fugiat veniam aliquip.  Frankfurter exercitation beef ribs consequat hamburger culpa, cupim turkey capicola shoulder veniam.  Consequat brisket pork chop tri-tip cupidatat short ribs boudin t-bone meatloaf cow chislic leberkas.  Tenderloin tongue ullamco capicola burgdoggen hamburger turkey.
	
	Short ribs commodo bresaola ex culpa meatloaf pariatur fugiat capicola turducken veniam landjaeger pork belly.  Leberkas duis strip steak, tenderloin pastrami kielbasa tri-tip.  Sirloin minim aliquip alcatra.  Aute leberkas pork occaecat chislic rump ham hock spare ribs swine.  Enim hamburger laboris, shankle fugiat eu doner tri-tip veniam lorem pig.  Bresaola tempor picanha, eu brisket pastrami consectetur nulla enim aliqua id chicken in aliquip tongue.  Nulla meatball nisi velit.
	
	Jowl eu magna tri-tip landjaeger hamburger commodo kevin deserunt.  Cupidatat est bacon nisi, anim doner tempor kielbasa turducken chuck landjaeger mollit.  Sausage reprehenderit sint, beef ribs capicola aliquip ipsum cupim.  Short loin picanha in, flank pork belly sausage proident aliqua.  Sint est capicola pancetta velit filet mignon non enim et ad swine meatball fatback exercitation.  Shank sausage sirloin esse in aliqua enim excepteur nostrud labore prosciutto ipsum.  In flank salami elit.
	
	Pancetta rump consequat lorem velit brisket tenderloin meatloaf kielbasa chuck tempor corned beef pariatur et.  Elit strip steak incididunt, tail bacon swine voluptate capicola dolore meatloaf adipisicing exercitation cillum sint.  Dolor et dolore, pork chop exercitation cillum strip steak doner in do sausage ipsum reprehenderit ribeye tail.  Venison picanha consequat flank, pariatur laboris reprehenderit leberkas bresaola strip steak officia bacon.
	
	Alcatra ipsum nostrud velit ad doner.  T-bone est cillum nulla non.  Meatball pork chop nulla, corned beef cillum in elit ball tip tongue aliqua pastrami anim brisket deserunt.  Ut ham pork belly, adipisicing ullamco magna tempor irure fatback mollit jowl cillum dolore chicken cupim.  Ribeye buffalo cupidatat shank elit anim.  Short ribs pork chop pork loin tri-tip burgdoggen.
	
	Chislic pork tongue salami magna commodo pancetta turkey nulla.  Short loin exercitation quis laborum eu, hamburger pig incididunt cillum aliquip chuck duis.  Jerky chislic excepteur tongue pork buffalo aliquip turkey.  Capicola ground round voluptate meatloaf consectetur adipisicing biltong corned beef.
	
	Nulla picanha et meatloaf shank.  Labore sausage frankfurter commodo in.  Nostrud veniam ham, ut esse consequat buffalo short loin occaecat chicken andouille shank qui alcatra laborum.  Adipisicing ex beef, landjaeger biltong shoulder minim burgdoggen.  Tongue tenderloin deserunt beef ribs elit chislic hamburger in pig brisket burgdoggen ut irure.  Eiusmod sirloin pancetta doner, duis dolore nulla.  Brisket meatloaf dolore minim.
	
	T-bone velit tail voluptate, andouille lorem ut sirloin ham.  Pastrami velit shank, dolor ham cupim ullamco fugiat voluptate prosciutto.  Cupidatat tail short ribs, ground round leberkas salami porchetta elit laboris.  Magna occaecat voluptate brisket officia.  Ham tri-tip et pig, bacon cillum jowl.
	
	Biltong dolor turkey, ex filet mignon reprehenderit cupim dolore sed minim eu kevin.  Tempor sausage pork belly, nulla meatball laboris id flank aute.  Ham hock excepteur sausage sirloin andouille ut chicken burgdoggen elit irure biltong.  In prosciutto officia sint, qui pastrami do laboris sunt spare ribs chislic.  Alcatra ad tongue, andouille beef reprehenderit velit capicola quis shoulder beef ribs sint consectetur tempor.  Doner sunt voluptate exercitation kevin deserunt ea ut consectetur quis.
	
	Chicken dolore hamburger fugiat, sausage ipsum proident.  Sint qui ground round sausage, spare ribs incididunt velit elit aute consectetur pork loin.  Shankle pork chop turducken, quis boudin ham reprehenderit short loin pancetta bacon rump ham hock ground round.  Ad et chislic proident ullamco boudin corned beef sirloin laborum cillum prosciutto excepteur sunt fatback.  Shankle cillum chislic ea brisket officia tongue commodo dolor ribeye id.  Drumstick nulla veniam nisi occaecat enim.
	
	Filet mignon ribeye ex corned beef enim.  Nulla culpa veniam, flank jerky pariatur esse.  Jowl picanha minim venison consectetur spare ribs flank cillum elit commodo aliquip doner kevin.  Bacon anim brisket in aliquip shankle.  Aute non andouille landjaeger meatloaf dolor.  Chuck sausage ut pig irure buffalo, bresaola shoulder dolore quis short ribs.  Ground round pig cupim reprehenderit leberkas chicken.
	
	Spare ribs swine meatloaf, kevin drumstick eiusmod pastrami ex mollit tongue enim.  Ex jowl consectetur nulla, beef tenderloin eu.  Filet mignon meatloaf frankfurter, veniam fatback ham hock shank strip steak porchetta picanha pariatur quis.  Sint meatball turkey, velit fugiat culpa shank biltong brisket ut chuck porchetta chicken rump.  Turducken veniam incididunt proident ex flank.
	
	Ut ball tip burgdoggen dolor.  Beef ribs consectetur meatloaf id.  Venison brisket shoulder shank strip steak kielbasa.  Ipsum mollit ut, cow doner jowl non reprehenderit brisket tenderloin pariatur ground round.
	
	Shankle in shoulder tri-tip hamburger meatloaf nostrud, labore ut pastrami culpa dolore tongue incididunt.  Consectetur aliqua velit, sed esse proident pork belly shankle ut frankfurter in tongue ut leberkas.  Cupim jowl turkey chislic.  Exercitation quis id meatball aliqua pork loin.
	
	Ut ut duis, ex sed biltong pariatur pork chop lorem ham hock spare ribs dolor capicola.  Commodo ut biltong do qui landjaeger.  Tempor fatback fugiat duis venison tongue sausage.  Anim deserunt chicken nisi.
	
	Velit tail fugiat sausage ut ground round magna et eiusmod ullamco.  Non bacon cow, drumstick ham hock shank turducken burgdoggen.  Tenderloin nisi rump consequat nostrud pastrami aliqua magna dolore veniam incididunt velit doner hamburger.  Pig reprehenderit strip steak shoulder pancetta sunt jerky drumstick commodo swine aute.  Commodo ut tenderloin cillum, meatball ea hamburger tail sausage beef ribs duis exercitation pig.  Chislic pancetta chicken pastrami ipsum est andouille ham, hamburger culpa.
	
	Do burgdoggen esse sunt chicken excepteur exercitation anim pork cupim.  Adipisicing picanha laboris boudin consequat rump ribeye ullamco pariatur pork belly in jowl in spare ribs.  Commodo bresaola ea magna enim ut, shank pork loin chicken doner sirloin lorem.  Andouille ex meatloaf dolor turducken mollit cow tenderloin jerky shank kielbasa beef.  Meatball anim ipsum ea dolor biltong.
	
	Ipsum proident shank do commodo tenderloin.  Mollit esse frankfurter pig est, pork loin leberkas quis shankle chicken bresaola.  Consequat kevin shank, tempor consectetur venison bresaola.  Frankfurter meatloaf salami, pork chop lorem pastrami aliqua qui brisket ribeye cupim commodo.  Buffalo quis cupidatat excepteur velit, turkey culpa leberkas et sausage.  Magna quis filet mignon et chuck ham hock commodo enim rump buffalo hamburger drumstick labore.  Porchetta shank proident ex tempor eu minim t-bone irure lorem pig sint veniam.
	
	Andouille leberkas filet mignon ut t-bone ullamco frankfurter.  Aliquip mollit meatball enim, exercitation aliqua chislic tri-tip.  Incididunt id tenderloin shoulder.  Prosciutto chuck laboris ut sint officia bresaola aliquip chicken.  Ribeye pancetta ut, chislic ea id venison pig capicola incididunt dolore.  Venison shankle mollit, swine sirloin kielbasa alcatra incididunt enim salami t-bone pork chop.  Commodo eu bresaola strip steak, meatball porchetta reprehenderit kevin.
	
	Shank sausage meatloaf id reprehenderit eiusmod boudin.  Consequat velit buffalo ut, exercitation ut cillum leberkas minim tongue dolore.  Eu short ribs doner, capicola deserunt nulla irure esse pancetta.  Corned beef short ribs laborum reprehenderit shank filet mignon.  Hamburger eu magna, duis pork chop cillum proident sirloin ut.  Bresaola consequat nisi ipsum shoulder spare ribs commodo consectetur pork loin prosciutto.
	
	Leberkas aliquip turducken laborum, ribeye picanha adipisicing ad short loin prosciutto ullamco laboris.  Fatback drumstick laborum id beef, mollit frankfurter proident irure eu biltong flank.  Turducken nulla officia, hamburger leberkas pork loin kevin cupim ut ground round tri-tip laboris elit salami.  Cow cupim porchetta sirloin eiusmod reprehenderit bresaola, aliquip hamburger commodo andouille meatloaf filet mignon.
	
	Dolore incididunt ribeye pastrami jowl ball tip frankfurter quis pariatur shoulder dolor occaecat bacon.  T-bone ea chicken beef ribs swine.  Rump ipsum ex, excepteur aliqua corned beef elit frankfurter eu quis tail prosciutto consequat.  Pork enim exercitation ad, brisket aliqua ground round tri-tip irure rump.  Aute pork chop kielbasa laborum magna irure in filet mignon shoulder consectetur eu flank.  Ball tip voluptate ham hock pariatur irure pork chop landjaeger consequat magna.
	
	Prosciutto labore minim cillum culpa dolore.  Hamburger commodo chicken pancetta.  Bacon picanha meatloaf cupidatat.  Dolore meatloaf tail laborum ball tip boudin tongue tri-tip ribeye biltong ut aliquip consequat.  Est chuck nulla velit cupim.
	
	In ut ground round labore, tenderloin fatback excepteur adipisicing do.  Eu tenderloin ribeye sunt veniam short loin meatloaf.  Velit culpa porchetta minim, nisi qui in et fugiat swine beef ribs cupim id venison pastrami.  Drumstick ullamco commodo flank ground round.  Bresaola turducken ham fatback hamburger swine leberkas ullamco nisi culpa drumstick chuck andouille.  Tempor ball tip jowl, magna leberkas pork belly labore nisi.  Porchetta aliquip filet mignon turducken rump.
	
	Porchetta quis boudin ut sint magna dolore ullamco laborum flank kielbasa hamburger.  Irure excepteur hamburger, tail sirloin aliqua brisket ham esse laboris velit.  Velit beef ribs ipsum rump cow burgdoggen commodo, pastrami esse labore leberkas pork chop kevin.  Velit ut ut, sunt consectetur salami eu cillum consequat chislic sausage nisi capicola sint corned beef.  Dolor andouille pork in ground round venison minim ea chuck aute labore sint culpa.
	
	Jerky cupim pork corned beef.  Dolor beef ribs consectetur cow, ad hamburger fatback chislic short loin ipsum landjaeger pork loin porchetta pig.  Est jowl esse veniam, flank boudin andouille sirloin cow.  Burgdoggen pariatur occaecat buffalo bacon rump cillum tempor nulla.
	
	Shankle tail ham kielbasa drumstick.  Pastrami adipisicing bacon t-bone, dolor ham nisi.  Reprehenderit pork loin lorem, qui kevin velit beef ribs est minim adipisicing burgdoggen frankfurter.  Spare ribs officia picanha aute cillum tenderloin id.  Qui eu rump, commodo nostrud shoulder laborum excepteur laboris.
	
	Cupim nisi laboris magna landjaeger bresaola andouille.  Meatloaf leberkas duis lorem qui reprehenderit landjaeger ullamco boudin.  Brisket mollit prosciutto, adipisicing ribeye esse leberkas.  Nisi esse fugiat laboris, pork leberkas ham hock mollit dolor bresaola sunt beef ribs.  Shank labore doner boudin sed.
	
	Id jerky consequat, filet mignon shankle quis andouille in fatback et venison pork.  T-bone exercitation aliqua et enim, est flank tri-tip swine drumstick biltong brisket esse incididunt nostrud.  Veniam aliquip swine cow dolor.  Landjaeger aute occaecat, sausage excepteur alcatra in enim capicola leberkas.  Eiusmod beef ribs meatloaf, aliquip ham hock adipisicing in.  Laboris fugiat do qui excepteur t-bone doner turducken laborum pariatur pastrami venison.
	
	Reprehenderit tempor leberkas sint buffalo lorem burgdoggen velit id ut frankfurter pork loin salami ipsum strip steak.  Beef ribs pork short ribs qui, ham hock esse ullamco bacon labore buffalo.  Landjaeger fatback qui tongue cow.  Chicken buffalo venison excepteur, in qui salami shankle alcatra ribeye frankfurter.  Dolore sunt brisket, jerky sirloin buffalo bresaola aliqua corned beef chuck doner id ground round shankle ribeye.
	
	Minim chicken boudin biltong venison pork magna nisi shoulder enim in.  Capicola laboris proident in minim andouille.  Laboris meatball elit, pork belly chislic lorem meatloaf bresaola t-bone ipsum short loin aute ground round ex tempor.  Qui quis flank, consectetur buffalo pancetta cupidatat ad strip steak cow tail in non turkey consequat.  Ball tip labore chicken shankle duis.  Kevin quis veniam drumstick ham filet mignon.  Nulla filet mignon buffalo strip steak magna drumstick pork anim kielbasa ham.
	
	Aute aliqua deserunt est anim aliquip burgdoggen ea turducken, ribeye ad.  Boudin ball tip et strip steak cupim tongue dolore nisi.  Laboris dolor beef sint t-bone excepteur dolore nisi biltong est occaecat sunt ut fugiat.  Nostrud dolore pig drumstick salami pork loin qui shank.  Boudin ut buffalo, fugiat in aliquip hamburger.  Cow pancetta biltong strip steak buffalo id hamburger magna esse et nisi spare ribs non excepteur.
	
	Hamburger beef ribs esse, ad nostrud filet mignon biltong brisket fatback reprehenderit turducken.  Nulla jowl sausage ut.  Cupim occaecat beef ribs deserunt consequat mollit ut ribeye ham ham hock elit tenderloin.  Qui tempor nulla aliqua, tri-tip anim in nostrud tenderloin dolor sed ham hock pork belly lorem jerky.
	
	Pancetta tri-tip sint elit boudin fugiat sirloin tongue.  Tenderloin filet mignon tongue doner prosciutto ut nisi laboris kevin landjaeger et pig pariatur aute.  Corned beef chicken tongue meatball doner spare ribs alcatra occaecat id sunt, meatloaf tenderloin laboris sirloin beef ribs.  Cillum tri-tip aliqua dolore, pork belly adipisicing kielbasa chislic ad dolor t-bone doner pork chop.  Ullamco rump ipsum cupidatat proident dolor capicola anim corned beef dolore exercitation occaecat.  Sausage frankfurter chicken fugiat jowl non aliqua tail id eiusmod bresaola reprehenderit pariatur.  Sirloin tri-tip enim andouille.
	
	Jerky anim fatback, doner alcatra tri-tip tempor tenderloin meatball short ribs.  Aute labore pork loin eiusmod pork belly consequat nostrud occaecat.  Filet mignon pastrami ham hock jerky voluptate tri-tip sint, adipisicing id quis hamburger dolore.  Nulla brisket meatloaf non.
	
	Tempor enim buffalo venison.  Sirloin esse reprehenderit, in sausage lorem ea.  Tail shank aliqua occaecat boudin commodo deserunt ham pancetta est pork loin.  Kevin tempor strip steak kielbasa velit.
	
	Deserunt ham cillum ball tip nulla do turkey ad jerky.  Officia nulla duis turducken venison ex cillum ground round frankfurter.  Andouille turducken sausage picanha.  Prosciutto consectetur salami, dolore jerky pariatur exercitation ground round meatloaf burgdoggen.  Tenderloin est in tongue duis filet mignon id aliqua labore ut cow prosciutto laboris.  Eu elit ham short ribs, doner ham hock laboris kielbasa lorem turkey in sed.  Swine jowl eiusmod tongue pork loin.
	
	Kielbasa pork loin proident pancetta velit ut quis pastrami veniam flank frankfurter bacon.  Eu anim magna, turducken burgdoggen porchetta cupim ball tip.  Velit proident venison burgdoggen ipsum boudin turducken filet mignon.  Labore shankle fatback ribeye do.  Alcatra salami reprehenderit nisi fugiat.
	
	Landjaeger ipsum beef ribs tempor alcatra ut drumstick porchetta.  Pork belly sunt sausage, veniam do meatball aliquip commodo eiusmod est venison in landjaeger alcatra.  Ad prosciutto tempor, andouille nulla meatloaf voluptate jerky filet mignon labore frankfurter buffalo irure.  Cillum leberkas shank buffalo in meatloaf non.  Irure ut tenderloin rump capicola.  Duis id eiusmod enim pork belly ribeye turkey capicola aliqua consectetur frankfurter picanha.
	
	Proident swine beef cupim.  Hamburger lorem sed meatloaf ball tip pastrami reprehenderit jerky aliqua venison voluptate non sirloin dolor deserunt.  Corned beef brisket tail sirloin turkey shankle fugiat.  Landjaeger leberkas fugiat esse magna quis chicken ea ut in kevin venison.
	
	Turkey meatloaf biltong sausage ground round tri-tip lorem shoulder excepteur, filet mignon ham.  Lorem nulla corned beef pariatur pork loin ball tip sirloin.  Ad et jowl, tongue culpa burgdoggen strip steak ham aute eu flank shoulder.  Dolore ribeye prosciutto dolore ut jowl pancetta cupim ut tail.
	
	Velit id cupidatat, culpa reprehenderit turducken ham hock in.  Turducken flank mollit commodo dolore proident.  Ex officia in incididunt, dolore non ham do flank consectetur irure nulla chislic meatball.  Occaecat cupim pastrami ut.
	
	Pancetta sed labore, irure consectetur pork belly beef ribs voluptate turkey sint tempor meatloaf adipisicing incididunt dolore.  In exercitation kevin qui in.  Proident ham hock short loin ham enim ribeye minim sirloin cow commodo.  Qui ball tip boudin porchetta, nostrud kevin sausage ham.  Tempor boudin minim, ad pastrami flank pork chop incididunt quis id sunt capicola chislic lorem.  Chicken id ex sint ut tail cupim venison shoulder dolor dolore jerky pig.  Dolor dolore pork, tempor biltong boudin ut elit voluptate ad pork loin alcatra kielbasa prosciutto lorem.
	
	Tempor biltong ad drumstick mollit reprehenderit.  Kevin meatloaf ut sed in turkey.  Turducken sunt qui ea corned beef exercitation in tri-tip kevin ut.  Non fatback nulla buffalo proident ham.  Landjaeger frankfurter rump hamburger ground round occaecat short ribs pastrami bresaola laboris.  Laboris ea pork aliquip cillum qui, tail ex spare ribs in ut mollit beef ribs nulla.
	
	Drumstick prosciutto nisi frankfurter aliqua swine.  Pig ex ham, burgdoggen buffalo chuck filet mignon tri-tip cupim boudin strip steak swine pork doner chislic.  Cillum burgdoggen incididunt, ground round pig cupidatat doner spare ribs est pork chop quis ham tri-tip.  Jerky ham hock leberkas ut, beef ribs eu venison t-bone.  Voluptate corned beef filet mignon, sunt shank ea porchetta dolore ut.  Pork bacon kevin qui, pork belly reprehenderit dolor beef.
	
	Minim biltong spare ribs cupidatat swine in culpa in short ribs meatloaf fugiat kevin salami.  Ham aliquip lorem, ad drumstick jerky consectetur buffalo pig t-bone short loin rump.  Ham hock minim eiusmod sausage chuck in, beef ribs turducken.  Beef est ut, eiusmod burgdoggen leberkas aliqua chislic.  Salami labore commodo esse cow.  Ut non elit ham frankfurter leberkas spare ribs minim eiusmod nostrud ut cow mollit kevin.  Chislic filet mignon swine cupidatat, velit ribeye beef.
	
	Rump pork chop short loin kevin, et officia duis in veniam ut occaecat pork belly.  Chuck meatloaf swine, jerky id ball tip rump irure consectetur ut pork belly meatball.  Tri-tip venison ham hock, mollit kevin bresaola pastrami short loin bacon anim consectetur short ribs shoulder jowl pork chop.  Jowl biltong aliquip shoulder irure proident kielbasa frankfurter pariatur capicola laborum bacon ullamco.  Aliqua lorem beef ribs, sirloin turducken spare ribs dolore ball tip doner porchetta pig biltong elit.
	
	Turducken sed quis nostrud ad.  Proident boudin do sed nisi venison aliqua dolore shank kielbasa ball tip beef ribs strip steak.  Frankfurter velit ipsum, quis consequat eiusmod sirloin.  Enim duis mollit turkey, est pork belly chuck meatloaf ground round irure deserunt jerky.  Ad strip steak nulla pork chop cupidatat laborum.
	
	Ea turducken ad, hamburger laboris tail ut cillum fatback.  Pork salami shoulder exercitation, reprehenderit id tail.  Ham enim mollit filet mignon officia sirloin shankle non.  Anim andouille pork chop t-bone, in bresaola doner ribeye landjaeger commodo ball tip.  In boudin meatloaf, dolore aute irure hamburger shoulder.  Quis officia chuck excepteur dolore meatball.
	
	Nisi ex ground round tempor.  Dolor turducken buffalo ball tip eiusmod sint magna veniam sunt swine chuck in.  Salami pastrami ex spare ribs, occaecat proident exercitation velit aliqua meatloaf andouille.  Dolore cillum shoulder mollit.  Do biltong ribeye, ad in tempor porchetta.
	
	Dolore incididunt bacon ut pork loin exercitation chicken pastrami meatloaf.  Reprehenderit mollit doner, cow ex buffalo ut dolore.  Dolor reprehenderit prosciutto occaecat cupim chuck, pancetta quis.  Rump tail duis chicken.
	
	Culpa mollit ham hock leberkas pig esse short loin dolore fugiat capicola landjaeger pariatur.  Flank meatloaf leberkas nisi, cillum in tail cupidatat ut.  Elit ground round salami pancetta drumstick sausage burgdoggen reprehenderit.  Lorem voluptate spare ribs do.  Boudin ham hock adipisicing dolore venison excepteur id fugiat enim ut.  Filet mignon sunt do, ea kevin bacon flank prosciutto id ut venison meatloaf culpa tongue laboris.
	
	Esse prosciutto qui pork, in sausage pancetta.  Meatloaf beef esse, eu drumstick cupidatat aliquip sint pastrami ground round minim eiusmod et andouille culpa.  Fatback leberkas ham ribeye filet mignon pastrami et drumstick doner dolore.  Eu pastrami tongue ball tip, bresaola venison sirloin buffalo cupim.  Pancetta sint corned beef incididunt duis consequat andouille tempor esse anim shank voluptate id.  Turducken enim ribeye velit elit reprehenderit short ribs ham hock prosciutto corned beef.  Flank buffalo aliquip do meatball sint commodo deserunt.
	
	Pancetta velit sausage pig.  Tempor commodo minim salami cow officia, ham hock jowl short loin sed.  Landjaeger aute et fugiat.  Sausage swine capicola, pastrami doner minim lorem ut cupim tempor est do quis veniam buffalo.  Minim ham excepteur turducken alcatra duis porchetta commodo id landjaeger boudin pork irure.
	
	Corned beef duis pariatur fugiat pork belly buffalo biltong.  Reprehenderit ullamco andouille meatball velit exercitation pork jerky turkey officia.  Elit andouille quis voluptate pork belly irure pancetta meatloaf pastrami hamburger, jerky aliqua.  Picanha occaecat cupim drumstick ribeye consequat frankfurter veniam porchetta brisket chislic ipsum pork belly.  Bresaola spare ribs nulla anim beef ribs fugiat picanha shankle tongue.
	
	Beef buffalo magna, turducken eu leberkas filet mignon veniam hamburger elit pork chop.  Est pig spare ribs elit proident, occaecat hamburger ham irure.  Doner pork alcatra ut irure swine incididunt velit veniam do.  Pork chop in cow chuck fugiat pig laboris sed consequat.
	
	Alcatra tail anim turkey jowl tempor in non mollit sausage duis swine burgdoggen.  Porchetta in sed chuck, aliqua ut velit prosciutto.  Ut chislic tri-tip ipsum.  Adipisicing corned beef filet mignon, in culpa pastrami do cupidatat in tail rump non.
	
	Ex bacon excepteur enim beef turducken tenderloin sirloin in dolore reprehenderit commodo labore pork belly.  Venison beef officia ground round chuck.  Laboris culpa shankle turducken proident, salami pastrami ex dolore exercitation aliquip in quis.  Adipisicing laborum lorem velit magna eu tempor ham hock.  Sausage pork belly drumstick pastrami ball tip, jerky tri-tip chislic est laborum cupidatat occaecat.  Fugiat consequat lorem, turkey voluptate hamburger porchetta labore.  Reprehenderit id cupidatat, venison laboris dolor filet mignon occaecat andouille.
	
	Bacon bresaola leberkas corned beef exercitation veniam beef eiusmod cupim.  Pig ground round cillum occaecat, meatloaf jowl ut cow excepteur incididunt beef shankle eu reprehenderit.  Turkey ham hock rump leberkas.  Spare ribs sed reprehenderit, laboris labore salami ut short ribs tri-tip excepteur shankle leberkas.  Short loin pig flank ham hock enim buffalo.  Ham hock sirloin shank esse.  Burgdoggen sausage ut non, pastrami rump cillum tail ground round in in aliqua.

	`

	req := &pb.MeatRequest{MeatList: testCase}
	resp, err := client.GetMeatAmounts(ctx, req)
	if err != nil {
		t.Fatalf("GetMeatAmounts failed: %v", err)
	}

	expected := map[string]int32{"ad": 43, "adipisicing": 30, "alcatra": 28, "aliqua": 48, "aliquip": 49, "andouille": 41, "anim": 30, "aute": 23, "bacon": 39, "ball": 33, "beef": 107, "belly": 37, "biltong": 40, "boudin": 46, "bresaola": 43, "brisket": 38, "buffalo": 42, "burgdoggen": 34, "capicola": 48, "chicken": 41, "chislic": 34, "chop": 29, "chuck": 35, "cillum": 45, "commodo": 42, "consectetur": 39, "consequat": 37, "corned": 37, "cow": 30, "culpa": 32, "cupidatat": 31, "cupim": 39, "deserunt": 26, "do": 35, "dolor": 40, "dolore": 71, "doner": 46, "drumstick": 36, "duis": 36, "ea": 29, "eiusmod": 37, "elit": 34, "enim": 40, "esse": 35, "est": 37, "et": 36, "eu": 39, "ex": 33, "excepteur": 32, "exercitation": 36, "fatback": 25, "filet": 42, "flank": 33, "frankfurter": 39, "fugiat": 42, "ground": 39, "ham": 86, "hamburger": 49, "hock": 41, "id": 50, "in": 104, "incididunt": 32, "ipsum": 29, "irure": 41, "jerky": 38, "jowl": 45, "kevin": 41, "kielbasa": 31, "labore": 40, "laboris": 45, "laborum": 27, "landjaeger": 37, "leberkas": 51, "loin": 57, "lorem": 40, "magna": 35, "meatball": 32, "meatloaf": 53, "mignon": 42, "minim": 31, "mollit": 41, "nisi": 32, "non": 37, "nostrud": 25, "nulla": 35, "occaecat": 40, "officia": 28, "pancetta": 30, "pariatur": 41, "pastrami": 44, "picanha": 34, "pig": 49, "porchetta": 33, "pork": 135, "proident": 35, "prosciutto": 38, "qui": 34, "quis": 44, "reprehenderit": 50, "ribeye": 42, "ribs": 111, "round": 39, "rump": 40, "salami": 37, "sausage": 44, "sed": 31, "shank": 42, "shankle": 42, "short": 60, "shoulder": 36, "sint": 41, "sirloin": 35, "spare": 39, "steak": 41, "strip": 41, "sunt": 34, "swine": 32, "t-bone": 32, "tail": 32, "tempor": 46, "tenderloin": 32, "tip": 33, "tongue": 37, "tri-tip": 39, "turducken": 46, "turkey": 33, "ullamco": 33, "ut": 115, "velit": 44, "veniam": 31, "venison": 36, "voluptate": 33}

	for name, amount := range expected {
		if resp.MeatItems[name] != amount {
			t.Errorf("expected %v: %v, got %v", name, amount, resp.MeatItems[name])
		}
	}
}
