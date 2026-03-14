Introduction
In May 2016, I watched something impossible happen. A start up —
with no CEO, no board of directors, no employees — raised $150 million
in thirty days. The DAO, they called it. The Decentralized Autonomous
Organization. And like a fool, I quit my job and never looked back.
Within weeks, a hacker drained $60 million from the supposedly
"unhackable" smart contract. Three respected experts published
analyses of what happened. Surprisingly, they disagreed — not on the
mechanics of the attack, but on basic facts: How much money was
stolen? How quickly did the money disappear? One of them proudly
remarked that it had “only taken twelve hours to extract the last few
days of data.”
Twelve hours to read three days’ data. On a system that was famous for
being perfectly transparent.
"How is that possible?" I asked myself. "The blockchain is supposed to
be 18-decimal-place accurate that anyone can see."
That was the moment TrueBlocks was born.
The seventy-ﬁve essays collected in this volume — along with a handful
of cartoons and technical specs— span nearly a decade of my life, from
May 2016 through March 2026, and they document something I never
intended to document: the birth, adolescence, and troubled
maturation of an idea called "decentralization."
I did not set out to write a book. I set out to build software. The essays
were dispatches from the trenches — frustrated notes about problems

I was encountering, excited sharing of solutions I had found, and
increasingly pointed critiques of an ecosystem that seemed
determined to forget its own founding principles.
Reading them now, arranged chronologically, I see themes I didn't fully
recognize at the time:
The problem was never the blockchain. The problem was access to
the data on the blockchain. The Ethereum node software was designed
to validate transactions. It answers the question "Is this transaction
valid?" brilliantly. It fails catastrophically at the question, "What has
happened to this address over time?" After ten years, it still cannot
answer this question well. The entire "Web3" industry — worth
hundreds of billions of dollars — is built on top of a database that isn’t
a database. Every centralized API provider has jumped in to ﬁll that gap.
If they don’t already, these growing behemoths will eventually track
your wallet addresses. They’ll know your transaction patterns. They
could even censor your access. This is the dirty secret of decentralized
ﬁnance (DeFi): it is not decentralized.
We built TrueBlocks to provide an index. The Unchained Index — a
permissionless, open-source index of address appearances published
to IPFS and discoverable via smart contract — represents ten years of
work to deliver what the node software should have provided from day
one. With it, any user can query their own transaction history locally,
privately, instantly, with 18-decimal-place accuracy, without asking
anyone's permission. All three experts would have arrived at the same
conclusion.
The technical problems were never the hard part. The hard part was
convincing people that running their own infrastructure matters. "It's
too hard," they say. "Why should I run my own node when I can get it for
free?" "Who cares about privacy?" And so, the revolution became an
API. Web3 became Web 2.5, then Web 2.1, then — let's be honest —
just Web2 with extra friction and near-worthless tokens.

I have spent ten years building tools for a future that may never arrive.
These essays are the receipts.
You'll ﬁnd technical deep-dives here — explorations of the diciculty
bomb, gas mechanics, token issuance, and the architecture of indexing
systems. You'll ﬁnd progress reports, sometimes optimistic,
sometimes exhausted. You'll ﬁnd rants about the inadequacy of the
RPC, the tragedy of centralized "dApps," and the endless frustration of
watching an industry abandon its principles for short-term
convenience.
Other of our books contain the formal speciﬁcations for our work — the
engineering blueprints for those who want to understand, or rebuild, or
improve what we've made.
But underneath the technical apparatus, there is a simpler story: a
small group of engineers who believed that "trust but verify" should
mean something. Who believed that users should own their own data.
Who believed that the blockchain's promise of transparency should
extend to ordinary people, not just those with the resources to run
infrastructure or the willingness to surrender their privacy to corporate
APIs.
We built TrueBlocks to solve technical problems. We proved that
solutions to the dicicult problem of decentralization were possible.
Now we wait, a beacon standing on an island alone, pointing to one
possible path to a safer harbor.
Revolutions are not tied up in a bow. They don't end with victory
parades. They end with the slow realization that the world has changed
in ways you didn't expect, and that the thing you were ﬁghting for has
become something else entirely.
But the code remains. The Unchained Index works. The tools exist for
anyone who wants to use them.

And maybe that's enough. Maybe the point was never to win, but to
build something that makes winning possible for the dreamers who
come next — whoever still believes that the original promise was worth
keeping.
This collection of essays is for them.

Thomas Jay Rush
Philadelphia, PA, USA
January 2026


The DAO’s First Big Decision
Friday, May 6, 2016

“The DAO” has been alive for about a week, and besides the obvious
pending decision concerning slock.it’s proposal for the USN1, another
major decision looms over the group. This decision will test the ability
of as many as 10,000 people to come together, anonymously, and make
a long-lasting decision.
Let’s just say I’m skeptical.
The locus of discussion concerning the “The DAO” is http://forum.
daohub.org2, which is the oLicial-but-not-oLicial home of “The DAO”,
an organization-but-not-an-organization.
As you can tell — I’m having diLiculty ﬁguring out exactly what the DAO
is. It’s diLicult because nothing like this has ever existed before. There
are simply no words to describe it. If I slip into imprecise language,
forgive me.
The most active discuss is about what to call
The DAO. There have been many names
proposed for the organization, from The DAO to
motherDAO to DAWN to DAOy McDAOFace,
and the conversation is very lively.

slock.it was the group that wrote The DAO smart contract, and it was expected that
they would make the ﬁrst proposal for the USN: Universal Services Network. It never
came to pass.
2
Now defunct.
1

As soon as someone suggests a name, someone responds. Many
times, in support, but also many times in opposition. The discussion is
cordial but disorganized.
To get a grasp on what’s going on, I created a summary of the discussion
which I’ve posted online at “Name This DAO”3. If nothing else, the
summary proves that in any large group of people, many are talented.
Some of the proposed names are amazing. Some of the supporting
arguments are even better.
My question is this: How will this group
ever arrive at consensus?
Hmmm…If only there was something that
could automatically arrive at consensus
between a disparate group of anonymous
actors in a mathematically certain way…I
think that might be good. I wonder…
There have been suggestions that a smart contract should be written,
and it should be submitted as a proposal to the DAO for the token
holders to vote on a name and subsequently pay to design a logo.
Others have suggested an online poll (to which others have responded
that it will get spammed). Others have hinted that someone should grab
The DAO by the horns, declare the discussion over, pick a name, change
the website, and pay someone to make a logo.
I’m not sure how it will work itself out, but I am very interested in the
spectacle. Stay tuned.

3

Now defunct.

What’s Ahead?
Here’s a list of a few issues that the DAO will be wrestling with over the
next few months. I present them in no particular order:
•

Who’s in charge. There is no-one in charge. Of course, this is by
design, but I’m anxious to see how this aLects the group’s ability to
conclude its deliberations. How will the group help itself to move
forward and resolve conﬂicts?

•

Noise-to-information ratio. The noise-to-information ratio on the
forums is very high. Trying to extract a clear idea from such a large
and diverse group of people, each with their own very wellreasoned arguments, is diLicult. Perhaps one of the early proposals
to the DAO should be the development of a discussion-based
software system4 tailored to DAOs. If the DAO outsources the
creation of this software, perhaps the software could later be sold
or licensed to other DAOs as part of the income stream. The system
should reﬂect each participant’s current token holdings but allow
for both per-token weighted and per-participant weighted voting.

•

Antagonism. In any interaction between people annoying human
emotions arise. While the discussions I’ve seen on the forums have
been congenial and polite, there’s always the danger of
antagonism. There are some posters who are “trolling” others. The
above-mentioned DAO-centric discussion system might beneﬁt
from being able to “vote trolls oL the island.”

•

Is the Crowd Wise? It is not clear to me that crowds make good
decisions. Even if the group does arrive at consensus, who’s to say
that their choice is a good one. Is the wisdom of the crowd —
especially the wisdom of 10,000 people — really the best? There is

4

This idea was put forward on the forums. Sorry for the lack of attribution.

an obvious danger of group think, which seems to frequently lead
to suboptimal decisions.
These are just a few of the future challenges I see. I’m sure there will be
many others. The biggest problem I see, though, is that the inevitable
disappointment people will feel if their position is not accepted by the
group. Maybe this causes people to split from the DAO. Maybe it causes
people to become more open minded and learn how to better defend
their positions. Maybe it leads to civil war.
I’m not sure where I come down on the naming decision. Personally, I
like DAWN — the Decentralized Autonomous World Network, but at this
point given the fact that everyone already calls this organization-thatis-not-an-organization “The DAO,” it may be too late to change.


Smart Contracts are Immutable
That's Amazing and It Sucks
Friday, May 13, 2016

Apparently, there are two types of software engineers in the world. One
writes code, pushes it out into the world to see how
it works, keeps track of the bugs, and then goes
back to drawing board, re-writes, ﬁxes the bugs, and
re-releases. Write-release-ﬁx, write-release-ﬁx. A
never-ending circle.
The other seems to be able to write code once, and because it was
carefully planned and carefully implemented, it runs correctly forever.
Write-release-ﬁnito.
I should probably not admit this, but I am not this second type of
programmer. At least, I didn’t used to be. Now that I’ve started learning
to write Solidity (the smart contract programming language), I may be
forced to become one.

EEAR
Recently, I wrote a smart contract for what I believe to be the world’s
ﬁrst Ethereum Early Adopters Registry (EEAR — pronounced “ear”).
What is EEAR? Have you ever driven down the streets of your town and
seen signs in store windows saying, “Doing business since...”? That’s
the idea behind the EEAR. Twenty years from now — hell — two hundred

years from now — you will be able to prove to your customers that your
company was an early adopter of Ethereum1.
How do I know you will be able to come back two hundred years from
now and the EEAR will still be running? This is one of the most
interesting things about smart contracts. They are unchangeable. They
never die (unless they are murdered). Once a smart contract is
deployed, it will run at the same address for as long as Ethereum exists.
That’s not the only thing. The smart contract code running at that
address cannot change. How is this possible?

It’s All in the Address
In Ethereum there are two types of accounts. If you own any Ether at all,
you’re already familiar with the ﬁrst type of account. It’s where your
Ether is. This type of account is called an Externally Owned Account or
EOA. The other type of Ethereum account is a smart contract. A smart
contract can hold Ether in the same way an EOA does, but it also has
software code associated with it. All accounts (both EOAs and smart
contracts) have their own unique addresses on Ethereum.
When you’ve completed your smart contract code, you
deploy it onto the Ethereum network. Deployment means
the contract is compiled and then “sent to the network,” so
it can run. The smart contract is located at its address, now
and forever.
The address of the newly deployed smart contract is created by hashing
together the address of the deploying account and that account’s
transaction counter called its nonce2. The nonce is incremented each
time an account initiates a transaction.

1
2

EEAR is not defunct because smart contracts on Ethereum live forever.
There are other things involved unimportant to the current discussion.

Because the nonce changes with each transaction (a contract
deployment is a transaction) the resulting Ethereum addresses of each
new deployment are unique. In fact, the way the hash function works,
each new address is wildly diQerent than the one before it. This is why I
say that the smart contract runs forever and cannot be changed. Even
if you wanted to (and I almost always do, given the way I write software),
there is simply no way to re-deploy a contract to the same address.
If the Internet worked the way Ethereum does, and I changed my
website http://trueblocks.io3, the re-deployment would move my
website to a new address. Who knows, maybe http://no-one-will-everﬁnd-me-again.com4.

Why is this Amazing?
What makes smart contracts so amazing, in my opinion, is exactly what
I’ve described above. Once a smart contract is in place it will stay there
forever and the code at that location will also stay there forever. This
allows an end user to trust the software perfectly. No-one will ever
again be tricked into using software that does something untoward
behind the scenes. There are no scenes. The software is what the
software does, Mr. Gump.
For this reason, before you agree to interact with a smart contract, you
should insist on seeing its source code. The counterparty should be
happy to provide it to you. If not, do not do business with them. Given
the source code, you can compile it yourself and then pull the bytecode
from the chain and make sure it’s identical. If the compiled code is not
identical to the bytecode found at that address — by jot or a tiddle —
don’t interact with it. They’re lying to you.

3
4

Now defunct.
Never “funct.”

I’ve visited many of dApps (distributed applications) over the last few
months (see The State of the Dapps5). Some of these sites make their
source code very accessible. And while it is not easy, you can conﬁrm
that the code they say is running is running. These types of sites provide
all the information you need to generate the byte code and verify that
the code behind the address is what it says it is. (This process should
be easier than it is, but that’s a diQerent story.) Other dApps provide no
information at all about their source code. I would not interact with
such a dApp. Who knows what they do behind the scenes?
Assuming you can verify a contract, and you can read the source code,
then you then know exactly what will happen when you sign up. Once
you start interacting with the contract, that contract will always be
there, and it will always behave exactly as it’s programmed.
That’s amazing. That’s something totally new that never existed before.
But it also sucks. Here’s why.

Why Does it Suck?
When I ﬁrst started writing my EEAR code, I made a mistake. My
contract keeps a list of registrants (obviously — it’s a registry), and to be
generous and share some of the income from the contract (of which
there has been about $3.50 US), I wrote a bit of code that looked like
this:
1
2
3
4
5
6
7

myShare = msg.value;
owner.send(myShare);
uint amount = (msg.value-myShare) /
max(1, numRecords);
for (uint k = 0; k < numRecords; k++) {
records[keys[k]].owner.send(amount);
}

This code contains at least two serious errors. Can you spot them?

5

Now defunct.

First, it sends all the Ether to me — I am owner because I deployed the
contract. Second, it spins through each existing registrant and gives
them nothing. One of the original selling points was we would share
revenue with registrants.
The ﬁrst error is apparent at line 1 where myShare = msg.value. This
is why end users need to read the source code and verify that it does
what it proports to do. Read the ﬁne print, Forest!
In fact, this was an honest mistake on my part. I meant to multiple the
value by 50%, but I forgot. I would have ﬁxed it except for the immutable
nature of smart contracts described above. By the time I noticed this, I
already had users.
To make matters worse, I had no way of getting in touch with my users.
I didn’t ask them for an email address when they signed up. Smart
contracts are permissionless. There’s a belief that it’s none of my
business who’s using my smart contract.
Worse yet even, I couldn’t change the code because the modiﬁed
contract would be deployed to a diQerent address and my users would
never ﬁnd me. I was stuck.

Worse Yet Even
But then I saw a worse problem, and I fully realized why
immutable/immortal code — for all its amazing beneﬁts — sucks.
If you look carefully at the above code, you will see that as the list of
registrants grows, the number of times the ‘send’ function in line 6 is
called with a zero value grows in proportion.
My lovely smart contract, of which I was so proud, was a total idiot!
Every time the ‘send’ function is called, it costs money in the form of
gas. Ethereum runs on gas. Gas is why the miners mine. My brilliant,

amazing, never-to-be-changed, never-to-be-removed, smart contract
was doomed to cost my users an increasingly larger amount of money
to run6.
So now what was I supposed to do? My only recourse was to list each
of the users’ addresses, the time and date they joined the registry, the
amount they paid to join, and the comment they made to posterity. I
then had to go back to the drawing board and modify my code to ﬁx the
problems. I then re-deployed my contract to a new Ethereum address.
Of course, being a brand-new contract, the new registry was empty. I
had to hand enter the ten records, and luckily, I built in functions to
transfer the ownership of the records to the original registrants’
account.
This still didn’t solve my problem entirely because the old contract was
still running at its old address. Luckily, I had built in a ‘self-destruct’
function in the original code. So, after I had everything transferred, I
called the kill function to remove the old contract.

Lessons Learned
I learned a couple of useful lessons from this misadventure:
•

Make sure to get your Solidity code right the ﬁrst time.

•

Think carefully if you need a ‘self-destruct’ function in your code. If
you don’t have one and the contract is what I would call an “eater of
houses and homes,” it will fail long term.

Thinking about this ten years later, this may have been a feature, not a bug. If the
reward sharing had worked, early adopters would have been rewarded with higher and
higher rewards as the contract become more popular.

6

•

Consider capturing email or other contact information from your
clients in case you must bring the contract down. In this way, you
can at least let them know of the new re-deployment address.

•

Build in an export / import capability so that any possible future
transition to a new contract is easier7.

I hope this short tale of my misadventures into the world of Solidity was
entertaining. Visit my new and improved EEAR 2.0 when it becomes
available. I ﬁxed the revenue sharing. Do you believe me?

This is bad advice since it suBers from the same problem as the loop above. Instead,
today, ten years later, I would say expect to have an oB-chain monitoring system such
as TrueBlocks to watch your contract and constantly keep an oB chain copy of the state.
7


What the F’ is a Finney?
Monday, May 23, 2016

When I was nine, my dad took me to a Phillies baseball game for my
birthday. I don’t remember who won, but I do remember that he yelled
at me for not being careful with my money.
We were in the hot-dog line, and he asked me if I still had the money he
gave me earlier, so I pulled four crumpled one dollar bills out of my
pocket. Two of the bills fell to the ground. Someone walking past
accidentally kicked them, and I had to scramble to get them back.
“You don’t know how to take care of money,” he yelled. It was an
overreaction. It’s the only thing I remember from that day.
To make a blatant understatement — money is important.
As early as 15,000 years ago, people
used Anatolian obsidian as money.
Since then, people have used copper,
silver, gold, grains of wheat, and even
sheep and cattle as money. Early laws
such as The Code of Hammurabi laid
down rules for money such as the
amount of interest one might charge, ﬁnes for wrongdoing, and even
levels of compensation for various jobs. One culture even tried to use
cow manure as money.
“Hey brother, can you spare some money?”

“Sure. Here you go. Have a turd.”
Certain materials obviously operate better
as currency than others.
One important attribute of ancient
currencies was scarcity. Anatolian obsidian
is harder to get your hands on than grains of wheat. Yes, you can pick
them both up oP the ground (Dad!), but to get Anatolian obsidian, you
had to go to Anatolia.
A Finney is a turd to Ether’s pile of shit.

So, What the ‘F’ is a Finney?
There are four formal names for the sub-divisions of Ether: the Wei, the
!

Szabo, the Finney, and, of course, the Ether. A Finney is !""" of an Ether.
!

A Wei is the most basic unit in the Ethereum world. It is worth !"!" of an
Ether. That’s about one God-gillionth of an Ether.
I know this is a little strange, but I looked this up — if a Wei was a single
grain of wheat, and a million grains of wheat make up a bushel, there
would be one billion bushels of Wei in a Finney.

A Szabo is a trillion times larger than a Wei, or about a million bushels
of wheat. According to the National Association of Wheat Growers
(which has a surprisingly informative website), “[a] bushel of wheat
yields 42…loaves of white bread or about 90…loaves of whole wheat.”
(Who knew?)
Averaging these numbers (half your sandwiches are wheat, and the
other half are white), a Szabo yields about 61,000,000 loaves of bread.
At twenty-two slices per loaf and two slices of bread per sandwich, it
would cost only ten szabos to make a peanut butter and jelly sandwich
for every person living on the earth today (about ﬁve billion).

A Finney is a thousand times larger than a Szabo and a thousand times
smaller than an Ether. Or about one hundred and thirty-four peanut
butter and jelly sandwiches for every person on earth.
If Wei was a grain of wheat, then an Ether would be 134,200 peanut
butter and jelly sandwiches for every person on earth.

Or, to put it one ﬁnal way, if Wei was a grain of wheat, an Ether would be
a single peanut butter and jelly sandwich the size of an average kitchen
for every person on earth.

But Where Did the Name Finney Come From?
SuPice to say that you can ﬁnd Mr. Finney’s academic papers from the
1990s and early 2000s on the Internet, Mr. Szabo’s many academic
papers are available on his website, and Mr. Wei’s academic output is
available as well. The names are in homage to a few of the original
thinkers in the crypto-currency space. I would discuss in-depth the
mathematical material in those papers, if I wanted to, but I would much
rather focus on silly things like giant peanut butter and jelly
sandwiches.

Conclusion
Just to bring this post to its absurdly logical conclusion. Think about
what happens to bread after it’s been eaten? (See the discussion above
about early forms of money.) I hope that’s not what’s will happen to my
Finneys. I have some concern that it might.
See Dad. I do care about money!


Downloading the DAO
Friday, June 10, 2016

I am obsessed with the DAO. I spend a ridiculously large amount of time
thinking about it, reading about it, and trying to ﬁgure it out. The most
interesting thing to me is the data. The fact that the smallest
transactions—both into and out of the DAO — are recorded forever is
very interesting to me.
Ever since I got into Ethereum, I’ve been trying to ﬁgure out how to get
at all that data. I’ve been hacking for weeks. I’ve diven (doven? dove?)
into the depths of the core code. I’ve studied geth, eth, Mist, and
web3.js. I’ve learned how to use RESTful APIs and JSON. I had no idea
what I was doing, but I knew what I wanted — I wanted that data.
I ﬁnally have it. In the end, I cobbled together ethslurp (a block chain
scraper) which is now available and open sourced on GitHub. (Watch
the video explaining how it works: http://ethslurp.com/videos.html1).
In ethslurp’s most basic use, one simply tells it to “slurp” an
Ethereum address. The program returns JSON data:
> ethslurp
0xbb9bc244d798123fde783fcc1c72d3bb8c189413

When I run this command from a Mac command line, ethslurp
downloads, imports, and then exports every transaction to or from that
address. (You may notice this address as “The DAO.” There are more

1

Now defunct.

than 100,000 transactions to the DAO, so it takes a long time to
download — If you try it, don’t say I didn’t warn you.)
If, instead, I run this command:
> ethslurp -l --fmt:csv --output transactions.csv

and then open the ﬁle transactions.csv, MS Excel pops up, and I
can see all 100,000 transactions in a spreadsheet. (-l means slurp last
address).
It took me weeks, but I ﬁnally had the data. I’m so pleased, that I’ve
decided to share it with the world: http://daodeepdive.com/cgibin/site.cgi?cmd=Data2.

Now What?
Given this data, the ﬁrst thing I did was to separate out the ‘Token
Creation’ transactions from the transactions that occurred since the
creation period.
The Mist browser tells me that the token creation period ended at
timestamp 146442600.
Every transaction on the Ethereum blockchain has a timestamp. So, I
simply went into the spreadsheet and found the last transaction that
happened before “closing time” and deleted the rest.
I’ve since added an option to EthSlurp to specify the earliest and/or
latest timestamp, blockNumber, or date. I can now run this
command:
./ethslurp --fmt:csv --timeStamp:1464426000 \
-l –output creation.csv

2

Now defunct.

which says ‘using the last (-l) address slurped and storing the data in
comma separated values format, read from the earliest timestamp to
latest inclusive and store the results in creation.csv. In other
words — slurp The DAO’s creation period.
Next, I wanted to verify that the total Ether that was being reported by
ethslurp was accurate. I used three di[erent sources to check my
data: (1) the daohub.org website3, (2) etherscan.io (which, if you look at
the internals of ethslurp is what the program is slurping), and (3) the
Mist browser4, which I consider the authoritative source of information
on the DAO.
What I did was this…
I opened the ﬁle creation.csv in Excel and summed the column
labeled value. My expectation was that this total would agree with the
Mist browser, http://etherscan.io, and the http://daohub.org website.
Before I did that, because the value column is in Wei, I converted it to
Ether. I did this by adding a new column and dividing value by 10¹⁸ (see
“What the F is a Finney?”). I then summed the Ether column.
According to my downloaded list of transactions the total Ether taken
in during the creation was 11,980,287.5826, which was higher than the
number being reported by the Mist browser by 2.17%. Interestingly,
though, this value was closer than the value reported by the
DAOHub.org website. (I leave ﬁguring that out as an exercise . The
Etherscan.io website agrees with the Mist browser.

3
4

Now defunct.
Now defunct.

Where to Next?
I’m currently interested in the data ﬁeld called input which contains
what I believe to be the actual function calls and parameters for each
interaction with the DAO.
I am also interested two ﬁelds called gasUsed and gas, which, in some
cases, are equal. My ﬁrst thought about these ﬁelds is that being equal
may indicate an “out of gas” condition, but it turns out this is not quite
true.
I’ve hand-checked a number of these transactions, and I’m ﬁnding
three di[erent types of gas=gasUsed transactions:
•

Transactions that do not have enough gas to even get started and,
therefore, fail on their ﬁrst opcode,

•

Transactions that have exactly the right amount of gas, and ﬁnish
their execution with zero gas remaining, and

•

A very odd type of transaction that fails when it encounters an
opcode that charges about ten God-gillion gas. 76 digits worth of
gas.

That is way more atoms than there are in the Ether, let alone Ether in
anyone’s wallet. I’m not sure what’s going on there5.
5 I’ve since learned that this large number is just a negative amount of gas remaining.
Thanks Yoichi Hirai.

5

Who Cares?
If you’re like me you care because it’s cool as hell, and all information
is good information. If you’re not like me, you got bored a long time ago,
and you are no longer reading, so you can go pound sand.


A Clue About the DAO Attacker's Location
Tuesday, June 21, 2016

I used ethslurp to scrape every transaction from The DAO since its
inception. This ﬁrst chart shows those transactions distributed by hour
of the day. There are more than 121,000 transactions:

As you can see, the transactions are relatively evenly distributed
throughout the day. This is to be expected because the DAO token
holders are presumably distributed evenly across the globe. There is a
slight dip between midnight and 4:00 AM UTC — when Western Europe
is sleeping — but otherwise, the distribution is ﬂat. Keep that dip in
mind.

So What?
Recently, a user called kumavis on The DAO slack published a list of
accounts that he thought may have been related to the attack on the
DAO. Here is that list along with names kumavis provided for each
account:

Let’s call this Suspect Set1. These ﬁve accounts combined have made
795 transactions on the Ethereum blockchain.
In an unrelated but similar post, Johannes PfeXer provided a larger list
of possibly related accounts in the very informative blog post “The
Attack Story.” PfeXer’s larger list includes each of the accounts from
Suspect Set1 plus six other accounts. PfeXer used his own names, but
I marked the accounts on both lists in red. (The community never did
standardize names for these accounts, which remains part of the
problem.)

Combined, there were 1,084 transactions on these eleven accounts. I
charted the same hourly distribution on both Suspect Set1 and
Suspect Set2, and while I cannot say for certain, I think I see a pattern.

The red and green bars represent the two suspect sets. Nearly all
activity falls between 2:00 AM and 8:00 AM UTC, with a sharp peak
around 4:00 AM. By afternoon UTC, the accounts go almost completely

silent. Whatever else this means, it's not random — someone was
keeping a schedule.
One possibility is that whoever these accounts belong to gets up each
morning, works on the blockchain all day, and then goes out to party at
night. The activity window — roughly midnight to 10:00 AM UTC —
suggests a time zone oXset of UTC+5 to UTC+7, which would put them
somewhere in the Indian subcontinent or Central Asia.

However, there’s another explanation — as many readers have pointed
out. Perhaps this person gets oX work at 6:00 PM, sits down to work on
Ethereum around 7:00 PM local time, and then goes to bed around 1:00
or 2:00 AM. This would put them somewhere near the west coast of the
United States or Canada.
Of course, as with all things crypto, to know for sure is impossible. But
I'll note one thing: remember that dip in the ﬁrst chart? The suspect
accounts are busiest during exactly those hours — when the rest of the
DAO community is quietest. Make of that what you will.


Knowing the Future and Proving You Know It
How the Blockchain Helped Me Predict the Future
Wednesday, July 13, 2016

I’m building a piece of software called ethslurp. It’s a blockchain data
scraper. It runs from the command line and takes various parameters
such as a smart contract address, an output format, a date and/or
block range.
The most salient feature of ethslurp is its ability to convert a
transaction’s input data (which is usually seen as unreadable
hexadecimal gobbledygook) back into human readable text.
Every Ethereum transaction carries with it an arbitrary amount of data
in a ﬁeld called input. The input ﬁeld is one of the nuggets of true
genius behind what Vitalik Buterin et. al. did — Bitcoin also carries input
data in each transaction, but it’s only 80 bytes long. With Ethereum, the
input data can be as long as needed. Ethereum sends function calls
to smart contracts in the input ﬁeld. In fact, the byte code of the smart
contract itself is sent in the input ﬁeld of the transaction that deploys
the contract.

Converting Input Hex Back into English
When a smart contract is compiled into byte code, an ABI (Application
Binary Interface) ﬁle is created. Using that ABI, which contains a list of
all the function calls with each of their parameters, ethslurp can
convert the input back into a human readable form.

For example, this input from a typical transaction on The DAO smart
contract:
0xc9d27afe
0000000000000000000000000000000000000000000000000000000000000044
0000000000000000000000000000000000000000000000000000000000000001

represents this function call:
vote(68,“Yes”)

While this input:
0x23b872dd
000000000000000000000000bf6cc4a882ace8e3a19bb8f5372280b33b2f6c9e
000000000000000000000000fbb1b73c4f0bda4f67dca266ce6ef42f520fbb98
0000000000000000000000000000000000000000000000056bc75e2d63100000

represents the function call:
transferFrom(
“0xbf6cc4a882ace8e3a19bb8f5372280b33b2f6c9e”,
“0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98”,
“-1”)

That is, send all DOA tokens from the ﬁrst account to the second (-1
means inﬁnite).
I’ve been using ethslurp to scrape The DAO every day since it ﬁrst
went live back in late April. I’ve stored that data, representing more than
135,000 transactions, here: http://daodeepdive.com/data.html1. I’ve
separated the transactions in various ways including day-by-day and
function-by-function.

So How Did I Predict the Future?
In addition to The DAO smart contract, I also scrape other smart
contracts associated with The DAO, among them the DarkDAO and The
DOA Curator’s MultiSig contract.

1

Now defunct.

Yesterday, around noon EST, I noticed a bunch of unusual transactions.
For the ﬁrst time since soon after the hack, there were transactions on
the DarkDAO, and we can see the function calls, I noticed that the new
transaction represented a changeAllowedRecipients call. So, I
scraped the transaction2 and got the following function call:
changeAllowedRecipients(
“da4a4626d3e16e094de3225a751aab7128e96526”,
1)

I then scraped that address and recognized it as The DAO Curator’s
MultiSig contract. So now I knew that somehow the DAO Curator’s
MultiSig was put on the Dark Dao’s whitelist.

Oh My God! I Know a Secret!
When I ﬁrst saw this, I was pretty sure I was one of the few people other
than the White Hat group (the assumed initiators of the transaction)
and perhaps the Dark DAO Dick™ (as I call herm) who knew about it.
I decided to document this in true “we-now-have-an-immutable-neverchanging-database” fashion by sending a transaction to myself3.
So, what’s the big deal? Why does this prove anything? Open that
transaction on Etherscan. You will see the following message which I
forever recorded on the blockchain:
on July 11, 2016, this contract (0xb656b2a9c3b2416437a811e07466c
a712f5a5b5a) which is the DAO creator contract sent a transaction to
the DarkDOA (0x304a554a310c7e546dfe434669c62820b7d83490)
that added the DAO’s curator to the DarkDAO’s whitelist. Here’s the
transaction ID: 0x9257c2e0a11de7b7427d4607f5908d6448278070b
b73500139387930826fedc0. I’m writing this to the blockchain before
https://etherscan.io/tx/0x9257c2e0a11de7b7427d4607f5908d6448278070bb73500
139387930826fedc0
3
https://etherscan.io/tx/0xa152016a73c308e2ecdf9A31d6a36f717604472a589f21e9
98e96c3e9500344
2

anything happens to demonstrate the power of http://ethslurp.com
software by Great Hill.
That was last night, July 11, 2016, at 8:58:05 PM UTC. This morning at
8:25 AM UTC this tweet appeared from one of the DAO’s curators,
Stephen Taul: https://twitter.com/stephantual/status/752886658211
676160. I knew this, and proved so4, 12 hours before it was announced.

Help Me Help You
I’ve seen a lot of commentary over the last few weeks about the many
ways that the Ethereum community can make smart contracts smarter
and more immune to hacks. Most of that commentary is focused
(properly) on the front half of the equation — that is, the source code.
There are many, many things we can do to make the source code better
and protect us from bugs.
But I think another interesting area that can be pursued is to learn how
to better watch what happens on a contract. I’m trying to convince
myself that somewhere in these ~135,000 transactions on The DAO
there are hints that we might have seen before the theft happened5. I
haven’t found anything yet. I’m not sure I’d tell anyone if I did, but I might
establish proof that I knew it.

Who needs copyright protection? It’s built in!
I’m proposing here the idea of active monitoring. So, not only do we need to be able to
do deep forensic (after the fact) analysis of what happens on smart contracts, but we
(the community) should be doing constant active monitoring from the moment a smart
contract is deployed.

4
5


A Too Often Neglected Aspect of Smart
Contract Security Auditability
Thursday, August 25, 2016

In the wake of The Great DAO
Debacle of 2016™, there have been
many articles and blog posts
concerning the need for the
community to write more secure Solidity smart contracts.
One such article is “Writing More Robust Smart Contracts” by Elena
Dimitrova of Colony.io1. Dimitrova discusses the use of function
modiﬁers to verify — prior to a function’s execution — both the initiator
of the transaction and the input to the transaction. The gist of the
article, if I may summarize, is that by using readable and easily
understandable modiﬁers, one may increase readability and help to
ensure that the function will execute only under certain conditions. This
is an excellent idea, obviously, and one that all of us should follow.
Another truly excellent article is “Onward with Ethereum Smart
Contract Security” by Manuel Aráoz an editor at BitCorps. The author
discusses numerous well thought out ideas on how to better secure
one’s contracts: fail in a function as early as possible, favor pull over
push payments, limit the amount of funds in a contract, and many other
useful suggestions. You should read that article. It’s accessible and
quite well done.

1

https://blog.colony.io/writing-more-robust-smart-contracts-99ad0a11e948/

At one point, Aráoz’s says,
Security comes from a match between our intention and what our
code [is] actually [allowed] to do. This is very hard to verify, especially
if the code is huge and messy. That’s why it’s important to write simple
and modular code.
I would make an addition to this comment thus:
…it’s important to write simple and modular code…and make it easy
for someone to later audit what happened…
In this article, I will discuss a too-often neglected aspect of smart
contract security: auditability.

The Great DAO Debacle of 2016™
It’s not necessary for me to spell out in detail what happened to The
DAO. Let’s just say that the un-hackable, impenetrable, immutable
DAO was both hacked and penetrated and then stolen from (by a bad
guy), stolen from again (by a group of good guys), refunded (through the
withdraw contract and a hard fork) and now lies nearly dormant.
Unattributed pseudo-quote from Ethereum’s founding fathers:
The Ethereum blockchain is immutable (not in the sense that once
a transaction is processed it will never change, but more in the
sense that this is how we would like you to think of it…)
All of this is water under the bridge, but ever since the hack, I’ve been
trying to re-create what happened with the DAO. I’ve been trying to
verify that the number of tokens purchased was correctly recorded. I’ve
been monitoring the DAO Withdraw contract to see how much of the

community’s Ether has been successfully returned2. I’ve also been
monitoring The DAO’s curators.
My two sources of information are the transactions initiated against the
DAO, and the events generated from those transactions. I’m ﬁnding
recreating what happened to be very di^icult because there appears to
be a signiﬁcant amount of missing or di^icult-to-come-by information.
These smart contracts are unauditable.

The Need for Auditability
Every user-initiated interaction with a smart contract creates a
transaction on the Ethereum blockchain. These transactions are easily
tracked on popular websites such as http://etherscan.io. A transaction
is created even if a function’s execution fails (for example, because the
transaction ran out of gas).
During the execution of many of the transactions, events are also
generated; however, as we shall see, this is not always the case.
Using a purposefully silly example, say you had a function that paid out
Ether to anyone who asked3. A naive implementation might look
something like this:
// Do not use!
function payout(address to, uint amount) {
to.send(amount)
logPaidOut(to, amount)
}

In this case, a record of the transaction is laid on the blockchain when
the transaction is sent as well as an event with identical information.
Notice, however, that there is a mistake here. Besides the obvious fact
that anyone can simply take all the Ether, the function writes a log entry

2
3

Not as much as you might think.
Obviously, a ridiculous idea, intended only to entertain.

to the blockchain even if the send call fails. The contract’s author has
neglected to check the return value of send. Before you sco^ at this
example, look at this line4 from the White Hat ETC Refund contract.
Before it was ﬁxed, it did exactly this.
This is one type of auditing error that we’ve seen as we’ve tried to
understand exactly what’s been going on with the DAO. I’ll present a few
more in the remainder of this article.

Too Much Information
As I mentioned above, for each function invocation, a transaction is
written to the blockchain even when there’s an error. An example of
providing too much information in a function would be this:
// Do not use!
function payout(address to, uint amount) {
if (balance[to] < amount) {
logFailedPayout(to, amount);
throw;
}
if (!to.send(amount)) {
logFailedPayout(to, amount);
throw;
}
balance[to] -= amount;
logPaidOut(to, amount);
return;
}

In this case, the contract’s author is over-reporting. No new information
is given by the two logFailedPayout logs. The transaction already
has the to and the amount values, and the absence of the
logPaidOut and the isError ﬂag indicates a failure5. There’s a
balance to be won. Try not to pollute your code by providing redundant
(or non-existent) information. You’ll just have to sift through it later.

4
5

We lost the location of this code. Sorry.
Not to mention that logs are not written to the chain in the case of an errored return.

Furthermore, be careful with logging in modiﬁers as suggested by
Dimitrova. If your modiﬁers contain too much logging, they might make
the audit trail harder to understand. Finally, a contract’s state is not
changed when an exception is thrown. Perhaps it makes sense to
primarily generate log entries when the state is changed and refrain
otherwise.

Not Enough Information
Another way you might mishandle the auditing aspect of your smart
contract is to not provide enough information. For example, if
logPaidOut only recorded the to account and neglected to record
amount, that would be an example of providing too little information.
While it would be possible for an auditor to ﬁgure out the amount from
the transaction, it would be much easier to simply include it in the event
to begin with.
Not providing enough information is not as bad as not providing any
information at all. In the above example, if the logPaidOut entry was
removed, there would be no way for the auditor to ﬁgure out if the
transaction succeeded or failed in sending the Ether except by querying
balances, which is expensive.

Incorrect Information
The worst type of auditing error is providing incorrect information. One
possible example appears in a familiar contract that many of you may
have already interacted with: The DAO post-fork withdraw contract.
Here’s code from that contract:

function withdraw() {
uint balance = dai.balanceOf(msg.sender)
if (!dao.transferFrom(msg.sender, this, balance) ||
!msg.sender.send(balance)) {
throw;

}
}

When I ﬁrst looked at this, I was convinced this was an example of
reporting incorrect log information, but I have since been corrected by
Mr. Nick Johnson from ENS fame. I initially thought that if the
transferFrom call succeeded on the DAO but the
msg.sender.send failed an event would have been written during the
transferFrom call to the DAO’s log. Obviously, this would have been
bad since the Ether was not actually sent.
Mr. Johnson informs me that in this case, because of the throw, the
entire set of state changes (including parts of the transaction initiated
in another contract — even if that contract call succeeded) are wiped
out. This is a good thing.
So, what would be an example of an incorrect information?
There’s the obvious bug scenario where, for example, you have two
events with similar names, and you incorrectly use one over the other.
I also consider missing or misleading log entries as errors.
The above examples illustrate where incorrect or incoherent logs can
be laid down on the blockchain making it nearly impossible for an
auditor to recover exactly what happened.

Conclusion
I’ve heard it said that any contract is worthless when everything is
working as it’s supposed to. How many times, for example, have you
read your pre-nuptial agreement? I don’t read mine — not to mention
the consequences if my wife found me reading it.
A contract matters when things fall apart. Being able to audit something
takes on more urgency when one it trying to ﬁgure out exactly what fell
apart.

When you’re designing your smart contracts, in addition to adhering to
the recommendations from the many security-related articles, you
should take great care in designing your logging system. Think about
how you will audit your contract’s behavior after the fact if something
goes wrong. Don’t be redundant; don’t be overly terse; but most of all
make sure to lay down correct information.
Good luck. Stay safe.


A Eulogy for the DAO
Or How I Learned to Love the Input Data Field
Friday, September 9, 2016

Introduction
In this two-part series, we analyze the nearly 170,000 transactions
made against the ill-fated smart contract The DOA since its inception
nearly ﬁve months ago. The DAO is dead. This series may serve as part
of its eulogy.
The DAO lived through four distinct time periods, each characterized by
diDerent user behavior, being (1) the Creation Period, (2) the
Operational Period, (3) the Post-Hack Period, and (4) the Recovery
Period. This ﬁrst part of our series gives a broad overview of these
periods and is followed in the coming weeks with further installments
presenting further details.
It is important to note, before we begin, that our analysis is a “work-inprogress.” There are three prominent areas of missing information: the
moment-by-moment Ether balance of the contract, the inclusion of
internal transactions, and an analysis of in-error transactions. This
information is available; but, to the best of our knowledge, it’s not
available quickly or automatically. The primary focus of our project is
the delivery of smart contract analysis and monitoring data — for any
Ethereum address — in as quick and as automatic a way as possible.
Delivering fast, interactive data for any smart contract is fundamental

to what we are working on. Once this missing data is available quickly,
we’ll bring it to you.
In the meantime, on with the analysis…

Four Periods of “The DAO”
The start of each of the four periods is signaled by an uptick in activity
as can be seen by the green- and orange-shaded areas above. This
analysis documents these spikes and proposes explanations for each.
Following is a summary of each of the four periods we’ve identiﬁed.
The ﬁrst period of activity (the pink section above) we’ve called the
Creation Period. Noteworthy here is the spike in token creation on May
14th which was initially reported to be the last day of the 1:100 token
price. A slightly smaller spike appears the next day, May 15th. The
existence of this second ‘fantom’ spike is, we believe, the result of a
mistake on the DAOHub website. Understandably, due to the
excitement generated by the huge amount of Ether attracted, a further
spike in activity is seen at the end of the period.
A huge spike in activity, immediately after the end of the creation
period, marks the beginning of the Operational Period. Many of the
transactions in this initial spike come from three or four accounts
(assumed to be exchanges). Notice the very small, blue area near the
bottom of the chart during this period. This is newProposal and voting

activity (for both split and non-split proposals). Proposal related activity
and voting distinguish this period from the rest and are further
discussed on the website.
While signiﬁcantly smaller, the next spike of activity is a direct response
to the hack on June 17th. This marks the beginning of the Post Hack
period. After the initial ﬂurry of panic selling, the period is characterized
by relative quiescence and declining activity due, we believe, to
uncertainty and an early direct instruction from the White Hat Group to
refrain from interacting with The DAO. Many of the non-transfer
transactions during this period were initiated by the White Hat Group.
The Post Hack period is also discussed further on the website.
The ﬁnal period we’ve identiﬁed, the Recovery Period, is marked by
another spike in activity as well as the advent of the DAO Withdraw
Contract. The start of this period corresponds with the hard fork. The
interactivity pattern results from users reclaiming their ether through a
two-step process of approval and then withdraw. After an initial spike,
this withdraw activity quickly dies down as is discussed elsewhere.
A possible ﬁfth activity spike seems to have occurred on or around
August 17th; however, we do not discuss this ﬁfth period in this
analysis.

Coming Attractions
In the coming weeks we will be expanding on and completing this series
of articles delivering our eulogy for The DAO.
To be honest we pushed this ﬁrst article out a bit early because we
wanted to ﬁnish it before DevCon2 in Shanghai. (We’re going!) If you
wish to talk with us about our project, please contact us. In the
meantime, here are two screen shots of coming attractions. We are
busy analyzing and writing about each individual time period. We’ll
have it out to you as soon as we can.

This is a hint at the ﬁrst, Creation Period charts.

1

And, on the next page, we present a preliminary version of the
Operational Period.

The e%ect of the hack on the price of Ether is obvious in this ﬁrst chart, as is — a few
days after the hard fork — the e%ect of Ether Classic not dying a peaceful death.
Additionally, these charts do not include “in-error” transactions or “internal”
transactions—see version 2.0 J.

1

Conclusion
One may say one thing with certainty about The DAO: It was an amazing
ride. The initial excitement of a purely un-managed piece of software
raising $150,000,000 US dollars in 30 days — despite there being noone in charge — was exhilarating. Then, the disappointment of such
meager voting. And then, on the morning of the hack, waking up to ﬁnd
that the price of Ether had plummeted (along with my naive belief that
it was possible to write an immutable, un-hackable code).
Later, the way the community came together (sort of) and carried out
the hard fork was inspiring, that inspiration only to be dashed by the
zombie-like resurrection of the lesser chain we all incorrectly assumed
would die a quiet death. It’s been an amazing ride.

On a personal note, I’ve enjoyed tracking the data generated by The
DAO but I’m glad it’s over. My desire to “get at this data” gave me the
impetus to develop the tools I’ve needed to feed my obsession (i.e. the
QuickBlocks website and the ethslurp software). I hope you’ve
enjoyed reading. Stay tuned.


A Eulogy for The DAO Part II
Friday, September 9, 2016

Introduction
In this second installment of this two-part series analyzing the nearly
170,000 interactions with “The DAO” smart contract, we discuss the
contract’s “Creation Period.” Previously, we presented an overview of
The DAO’s lifespan. In that overview, we identiﬁed four distinct
timeframes: the Creation Period, the Operational Period, the Post-Hack
Period, and the Recovery Period.
In this installment, we focus our attention on the Creation Period.

Creation Period
The DAO’s creation period lasted for 28 days from April 30, 2016,
01:42:58 AM (block 1428757) when it was deployed to May 28, 2016,
09:00 (block 1599205). According to the Etherscan website, more than
one billion DAO tokens were created.
Because much of the smart contract’s functionality was disabled
during creation, three functions dominate the logs during the period:
function, createTokenProxy, and approve.

We ﬁnd that nearly two-thirds (66.35%) of the 62,873 interactions
during the Creation period were calls to the default function. A call to
the default function, function(), occurs when a contract is sent
Ether directly. The code of the default function is quite simple: either it
creates tokens by calling createTokenProxy(), or, if called after
closingDate but before the end of the grace period, it sends the Ether
back to the caller, or, if called after the end of the grace period, it keeps
the Ether.
Together, the two functions createTokenProxy() and function()
account for 91.86% of all interactions during creation. 8.12% of the
remaining interactions were approve() calls, which are discussed
below. 0.02% of interactions were with other, infrequently used
functions.

Token Creation
In the bar chart below, we present the number of times each of the three
functions was called each day. The initial ﬂurry of activity on the ﬁrst
day, followed by a lessening, followed by a spike near the midpoint of
the period, then another lessening, leading, ﬁnally, to a period of

growing interest, is easily explained, as is the change in the usage of the
createTokenProxy() call relative to direct sends of Ether.

Many initial purchasers of the DAO tokens had clearly been anticipating
its arrival and were ready to act on the ﬁrst day. These users, one
presumes, were more technically adept, which is hinted at by the
signiﬁcantly higher frequency of direct sends (orange bars) compared
to createTokenProxy() (black bars) during the ﬁrst few days of
operation.
One may notice that as more purchasers appeared, the di^erence
between
the
number
function()
calls
relative
to
createTokenProxy() calls lessens. We believe this is due to the use
of tutorials explaining how to purchase tokens. Many tutorials
instructed users to use the Mist browser (as opposed to say the geth
command line interface) and from there to call into the
createTokenProxy() interface. Of course, this is only conjecture,
but it helps explain the growing relative use of createTokenProxy().

A Simple Mistake Leads to User Confusion
At the start of the creation period a single Ether purchased 100 DAO
tokens. Midway through the period, that ratio began a steady ten-day

decrease until ﬁnally reaching 100 DAO tokens per 1.5 ether, where it
remained for the last four days of the period. This behavior is explained
well elsewhere and is further illustrated in the table below:

This table was generated by analyzing the following code. In the code,
createTokenProxy() calls into a function called divisor(), the
result of which, when multiplied by 20, gives the token/wei ratio. This
factor is then multiplied by the amount of Ether sent to the function
(msg.value) to arrive at the number of tokens to award to the caller.

function createTokenProxy(address _tokenHolder) {
...
uint token = (msg.value * 20) / divisor();
extraBalance.call.value(msg.value — token)();
balances[_tokenHolder] += token;
...
}

The divisor() code looks like this:
function divisor() constant returns (uint divisor) {
if (closingTime — 2 weeks > now) {
return 20;
} else if (closingTime — 4 days > now) {
return (20 + (now — (closingTime — 2 weeks)) / (1 days));
} else {
return 30;
}
}

Looking back at the bar chart above, one may notice the users’ behavior
approaching May 15th, the ﬁrst day of the price increase.
As expected, activity increases as the price rise approaches. Notice,
however, that the greatest activity occurs on May 14th, a full day before
the price increase. (Each bar represents the activity between 9:00 am
the previous day and 9:00 am on the current day, so the bar labeled May
14th ends at 9:00 am, May 14th — 24 hours before the ﬁrst price rise).
The green-enclosed box in the table above helps explain this
unexpected behavior, which is due primarily, we think, to an error on the
DAO Hub website. It seems the maintainers of the DAOHub website
read the code wrong and made a miscalculation as to the ﬁrst date of
the price change.
Alternatively, one might say that the software code itself is in error. The
somewhat confusing code of divisor() passes through a di^erent path
on May 14th (i.e. closingTime - 2weeks). However, oddly, it arrives at the
same result (20). It is easy to see why the error on the website was
made.

The ExtraBalance Account
Referring again to the source code above, notice the second line of the
createTokenProxy code where the extraBalance account is
funded with the amount over the 1:100 refundable tokens of the early
adopters. This account was initially intended to hold the proﬁts of The
DAO, which were later to be distributed pro-rata by token share;
however, a decision was taken at some point to refund the original late
purchasers with the exact amount of ether they sent. It is not clear,
exactly, how that decision was made. It certainly was not voted on by
token-holders as some people think it should have been.

Approve Calls
The undulating green area in the chart above represents calls to the
approve function, one of the few functions that were not disabled
during the Creation Period. We do see something a bit surprising here.
First, the existence of approve calls itself prior to the ability to make
transfers was unexpected, at least to this writer.
Secondly, and more surprisingly, is the regularity of these function calls
and the fact that they came almost entirely from four addresses. It is
our belief these calls were most likely exchanges practicing or testing
their software prior to the start of the Operational period.

For most of the month, on a day when the approve function was called
at all, it was called an even multiple of 100 times. For one account, the
approve function is called 310 times, then 100 times, then 200, then
300, 200, 500, 200, and 200 times every few days. This pattern is
revealed in the consistent heights of the green-shaded area above.
These ‘100-transaction’ sets were clearly generated by software. This
becomes obvious in the table above which is presented without further
comment.

Conclusion
Obviously, having raised nearly $150,000,000 US in 30 days, and being
the largest crowd sale in history, The DAO’s creation period was a wild
success. Ignoring that huge amount of money for a second, we ﬁnd
nothing surprising in the interaction patterns we see in the data.
The mistake made by the DAOHub website, causing people to buy
tokens a day early, is revealed by the data. A lesson to be learned here
is to be very careful when reporting on what you think a certain piece of
seemingly easy code does.
The e^ect of disabling many of the functions of the contract until after
the creation period is revealed as well.
Overall, though, nothing surprising happens during this ﬁrst of the four
periods of The DAO’s lifespan. Early adopters bought in early. Many
people waited until the last moment before the price rise to buy in. And,
people joined the circus near the end in a clear case of FOMO. As we
move into the next three periods with further installments of this series,
we think you will agree that the ride gets a bit bumpier. Stay tuned…
Notes:
•

Missing from this analysis is the daily value of Ether and DAO tokens
accumulated. While this information is available, we hope

eventually to expand our system to make the extraction of this type
of information automatic, particularly given the emerging ERC 20
Token Standard. In the future, we hope to automate the analysis of
all ERC 20 token-based smart contracts. We therefore did not
include the information here, as it would have required handgenerated data.
•

Also missing from the analysis is an accounting for internal
transactions. Again, we are working on automating this data
collection. We believe this omission is not material1. According to
one analysis, less than 1% of interactions during the period were
initiated by smart contracts (that is, less than 1% of the interactions
were internal transactions). We do not have the data to support this
claim, but we believe it to be true.

We were clearly wrong here. The hack came through an internal (smart contract
initiated) transaction.

1


DAO Token Holder's Response in Charts1
Tuesday, December 6, 2016

I’m sure everyone is aware that on Friday, June 17th the DarkDAO reared
its ugly head. I woke that morning, as I normally do, and sat down at the
breakfast table with my morning newspaper (i.e. slack). The ﬁrst
message I saw said something like, “If you initiated CongoSplit…please
contact us! Now!” The next message I read was, “The DAO is being
drained — ongoing!”
I don’t know about you, but that many exclamation points ﬁrst thing in
the morning upsets me. I spent the rest of that day, trying to ﬁgure out
what was going on. Soft fork…hard fork…counterattack…recursive…
call…bug…splitDAO…transferWithoutReward…head…explode…excla
mation point!

How Did I Respond?
I responded as I do to most other things in my life: I dug into the data
and tried to collect information. All day Friday, I spent combing through
the blockchain data, trying to ﬁgure out what other people were doing.
This is what I found.

This post was written soon after the DAO hack, but before it became apparent
that the White Hats would rescue the remaining Ether in The DAO (or as some
have called it “safely steal the Ether”). The post may be interesting for
historical reasons. References to days of the week are to be read as meaning
the week following Friday, June 17th, 2016.
1

How Did DAO Token Holders Respond?
Not unexpectedly, the activity on The DAO increased signiﬁcantly the
morning of the hack. Each of the charts in this essay shows the number
of times one or another of the DAO functions was called.
The ﬁrst chart shows the number of transfer and transferFrom
function calls. It’s obvious that people reacted to the news by trying to
sell their tokens.

Notice the large number of transfers on May 28. That was the ﬁrst
day after the creation period. This was most likely exchanges shifting
tokens to individual account holders. You can see that the transfer
activity dropped signiﬁcantly during most of the rest of June, but it
picked up again on 17th.
Notice also that the transferFrom functions (pink) were higher than
the transfer functions in May. This again points to exchanges. Since
Friday, the direct calls to transfer are more numerous than to
transferFrom. Transfer is called directly by the account holder, so
these are people acting on their own behalf.
Token holders may exit the DAO in three ways:
•

by selling (transferring) their tokens on an open market,

•

by creating a split proposal using newProposal, moving their
Ether into a baby DAO, and then making a proposal there to retrieve
their Ether later; and

•

by voting on someone else’s split proposal and exiting into that
baby DAO.

Calls to the newProposal function (method 2) are shown in the next
chart. (Note that the data I use does not distinguish between regular
investment proposals and split proposals.) As you can see, the number
of calls increased signiﬁcantly since the attack, averaging less than 20
before the hack, and nearly double that since.

The ﬁrst question that came to my mind when I looked at this chart was
“How many of these new proposals point to a smart contract with the
same recursive-call hack?” followed quickly thereafter by, “How many
of these proposals will drain the DAO further?” The next thing I thought
about was the waiting period2 and what happens in the case of a soft
fork?

The DAO split function had a built-in twenty-seven day waiting period (plus
a seven day debating period) during which people had to wait to remove their
Ether.
2

My next worrisome thought: “What happens in the case of a soft fork if
the community can then not agree on a follow up hard fork?”
The people making split proposals are running away from the DAO.
The soft fork will freeze activity in not only The DAO, but any child DAO.
Without a follow-on hard fork, these baby DAOs will be frozen forever.

Rats Leaving the Ship
Selling on the market (transfer, transferFrom) and leaving via the
split mechanism (newProposal) are two methods to leave the DAO. A
third method is to vote on someone else’s split proposal and ride along
with them into their new baby DAO.
It appears that there are at least some people taking this path. This is
revealed in the third and fourth charts below. The spike in voting on
Friday and then again on Sunday, and secondly, the large increase in
splitDAO function calls reveal people trying to jump ship on someone
else’s split.

Two di`erent things are going on in this ﬁrst chart. During the ﬁrst few
operational weeks of the DAO, there were a relatively high number of
people voting — possibly out of a sense of excitement. This type of

voting is di`erent from the voting we’ve seen since Friday. The voting on
Friday and Sunday were people trying to leave The DAO. The low
number of votes on Saturday, I think, is because of messaging from the
White Hats instructing people to refrain from interacting with the smart
contract while they ﬁgured out what we should all do. There was an
explicit post on reddit.
Most of the voting on Friday was by people scurrying to get o` a sinking
ship. The sad thing about rats leaving a ship is that they end up in the
water where they are in danger of drowning. This view is reinforced by
the following chart.

The increase in the number of calls to splitDAO as people started
ﬁguring out this third method of exiting The DAO is obvious. In a baby
DAO, if you are not the initiator of the original split proposal, you will not
be able to make new proposals in the child DAO. This is because you
will not be on that DAO’s whitelist3. This leaves you at a disadvantage.
Maybe…

See our article “Knowing the Future and Proving You Know It” for an interesting
appearance of the DarkDAO’s whitelist.
3

You will be at a disadvantage because you will not be able to get your
money out of the child DAO unless the originator is magnanimous. If the
originator owns more than 50% of the tokens in that baby DAO, that
person can simply take all of the Ether. More likely though, they will own
less than 50% (especially if other rats jump onto their ship).
In this case, the only solution is an amicable truce. If the originator does
not whitelist anyone, and the majority does not allow the originator to
take the Ether, the only solution is a proposal that shares. The trouble,
of course, as is true with this entire smart contract universe, is that
communicating with other members of the contract is impossible.
The point I’m making is that trying to split o` with someone else’s baby
DAO is as bad as a rat jumping o` one sinking ship onto another. It’s not
a great idea.

No Conclusions
I would never make a recommendation as to what anyone else should
do with their tokens. I know what I did. I know what I’m advising my son
to do. We’re selling our tokens — even at a loss4. You will have to ﬁgure
it out for yourself.
These charts show what other people are doing, however. Over 23,000
di`erent accounts purchased tokens during the creation period. Look
at the y axis of the above charts. The great majority of people are doing
nothing5. My guess is that this is because they don’t know what to do. I
hope this post can at least give you some ideas or some thoughts.
Good luck and keep swimming you rats!
It turns out this was a mistake because of the hard fork returned everyone’s
lost tokens which means they could get their Ether back. I sold at 20% of face
value on Sunday. Lost my shirt.
5
Less than 5% used transfer (sold on the market) and almost no-one made
a split proposal or joined an existing baby DAO.
4


Accounting for the Revolution
Monday, March 13, 2017

Recently, the price of Ether hit $20.00 US dollars per Ether. That’s up
nearly double in the last two weeks. Of course, this has caused a lot of
discussion.
We wondered what this increase in price meant to the Ethereum
community. Does it mean they’re spending thousands of more dollars
per day on gas. Hundreds of more dollars per day? Pennies more
dollars per day? We had no idea, nor any basis to answer that question,
so we set out to ﬁnd out for ourselves.
Before we started our exploration, we imposed a rule that we could not
use any external APIs nor rely on any third-party data. Our only source
of data was to be our own locally running Ethereum node.
If you’ve ever tried to do this, you know this means one thing: sit back
and relax. It’s slooooow.
Since the DAO hack, we’ve been working on speeding up access to a
local node. We call our solution QuickBlocks1. The word “quick” being
operative.
Written in C++2, QuickBlocks’ cache achieves speeds of more than 200
times faster than the raw RPC. This allows us to produce answers to our
questions (and many others).

1
2

Since renamed TrueBlocks.
Since ported to GoLang.

The easiest way to explain QuickBlocks is to show you some code, and
I will do that shortly, but ﬁrst, I wanted to talk about…

Blockchain on a Stick™
QuickBlocks scrapes, parses, pre-digests, and caches the entire
Ethereum blockchain in a fully decentralized way. We claim that our
work is so decentralized that we can store the entire chain — fully
parsed and heavily cached — on a single external drive. We call this
Blockchain on a Stick™, and we carry it around with us in our backpack
wherever we go.
Blockchain on a Stick™ allows us to show
people that our solution is decentralized.
We can, if we wish, disconnect from the
Internet, and while we won’t have access
the latest blocks, our code still works
perfectly. That’s decentralized.
QuickBlocks uses the node’s RPC in the same way that web3.js does.
After pre-processing the received data, we store each block (that we’re
interested in), every transaction, every receipt, and every log in a local
database. Before storing the data, we optimize the crap out of it in every
way we can think. Our guiding light is speed.
We retrieve each block and then we spin through each transaction
asking for the transaction’s receipt. Given our need, we decide if the
transaction needs to be traced. We trace all transactions even those
that may have ended in error. Additionally, we identify internal
transactions that were initiated by the transaction and parse the input
data ﬁeld and the receipt’s log data.
We do this pre-processing prior to caching the data, and because the
blockchain is immutable, we only need to do this once. Every time we

access the data thereafter, we are reading cached, binary data
intended speciﬁcally to be retrieved quickly.
Another thing QuickBlocks does while storing the data is to store
various levels of detail. (Yes. It uses more disc space, but it’s
signiﬁcantly faster.) This allows the developer to choose between
higher speed given less detail verses lower speed if more detail is
required. It’s the programmer’s choice. For the analysis below, we used
the half-speed / half-detail level.

Programming QuickBlocks™
QuickBlocks is a C++ library and a series of applications. Below we
show you the application code we wrote that gathered the data we
needed to do our analysis. This is the actual code. It’s simple. As with
all C++ code, it starts at the main function.

This function ﬁrst initializes the library and then decides on the start
block and the number of blocks to visit. We chose to start at block
2,912,407 (the closest block to January 1, 2017). The function
getLatestBlock returns the node’s latest block.
Next, we create a data structure that will store the result of visiting each
block. This structure may be of any type — whatever is appropriate for
your application. The code then calls into the QuickBlocks library
function forEveryNonEmptyBlockOnDisc and ﬁnishes by making a
ﬁnal report.
QuickBlocks has many interfaces for traversing blocks. For our
purposes, we used forEveryNonEmptyBlockOnDisc. This function
gives full detail but skips over blocks with no transactions (about 37%
of all blocks).
A similar function, called forEveryBlock, visits every block
(including empty ones) in full detail. On the other end of the spectrum
is the function forEveryMiniBlockInMemory3 which is super-fast
but delivers a lot less data than the other methods. The function one
uses depends on one’s application.

Visiting Every Non-Empty Block
If you’re familiar with C++, you’ve noticed that the bulk of the work is
completed in the function visitBlock. That function is shown on the
following page. At each non-empty block, the pointer to the arbitrary
data is retrieved. In this case, it’s a pointer to the instance of
CGasTracker from main.
The function ﬁrst ﬁgures out when the block occurred. The function
then determines if this is a diierent day than the previous block. If it is,

In the parlance of this article, a miniBlock means the block’s header with no
transactions.

3

a report on the previous day is made. In this way, we accumulate and
report on statistics once per day (see the data tables below).

The function then grabs the price of Ether in US dollars at the time of
this block, that is, it grabs the block’s spotPrice. QuickBlocks gets
this data from the Poloniex price API. (Okay — we admit it — we broke
our own self-imposed rule! But we cache the price data, so it still works
when running from the stick.)
The function then spins through each transaction and prices the gas
consumed by that transaction. Note, that we do not distinguish

between in-error and successful transactions. Even if the transaction
ended in error, gas is expended, so we want to account for it.
CGasTracker accumulates the data for each day, holds the Ethereum
price data, and reports the results. Now we’re ready to answer our
original question.

How Much More is the Community Paying for Gas?
What eiect has the recent price increase had on the Ethereum
community as a whole? The chart below shows the amount of Ether
and US dollars spent each day during January of 2017:

On average, the Ethereum community spent slightly more than 34 Ether
per day during the month of January. Most of that month, the
spotPrice hovered around $10.00 US per Ether. This translates to, on
average, about $340 US dollars per day.
From the perspective of an average transaction, one sees that a
transaction costs between 0.7 and 0.8 of a penny per transaction4.
We checked this number against Etherscan and we found identical
numbers. Over the entire month of January, the community spent
nearly $10,700 US dollars on gas.

How about February?
Three things changed in February:
•

The number of transactions per day (50,359) rose by about 10.1%
since January. More people made more transactions.

•

The amount of Ether spent per day (on average) rose by more than
30% (that’s more than 10 Ether per day). Part of this is due,
obviously, to more transactions being sent, but also the gasPrice
in Wei may have risen.

•

The number of US dollars spent per day on average rose by nearly
$200 dollars (57.97%). Clearly, this has to do with both previous
values plus the rising price of Ether.

We use Finneys and cents as opposed to Ethers and dollars because the numbers are
!
too small otherwise. A Finney is !""" of an Ether or a lot of peanut butter and jelly
sandwiches.
4

It’s hard to untangle these three components, but you can get a glimpse
of what’s going on by focusing on the Finneys and cents columns. While
we leave it to the reader as an exercise to ﬁgure out why on the 13th, the
15th, and the 22nd Finneys per transaction rose above 1.0000, one can
see that the average number of Finneys per transaction over the month
remained relatively stable. In other words, the price (in Ether) that
miners accepted to process transactions didn’t adjust too much even
though the amount of total income taken by all miners ($15,250.26 vs.
$10,688.16) increased by 42.68%.
One may see the price in US dollars increasing per transaction in the
“cents” column. At the start of the month, the price per transaction in

cents was around 0.9 pennies. By the end of the month, the price in
cents had risen to around 1.3 pennies per transaction.

How about March?
Obviously, we expect to see a more profound change in March given the
skyrocketing price of Ether5.

Here we see an even more marked rise in the ‘penny’ price of a
transaction. At the start of January, a per-penny cost of a transaction
was about 0.5. Today, the price per transaction (in pennies) is three
times higher. The average price per transaction in Finneys rises over the
same period, but not by three times over. This makes sense given the
signiﬁcant increase in the price of Ether per dollar.

Let’s Look at Another Chart
I think you can see from the following chart that through the month of
January and into mid-February the price of a transaction in both Ether
(Finneys) and dollars (cents) tracked each other quite closely.

5

This analysis is as of Saturday, March 11th, 2017.

Near the end of February, we see the prices starting to diverge. We
believe this is because the price of ether per US dollar changes much
faster than the gasPrice miners are willing to accept to process a
transaction. Furthermore, while both Parity and Mist allow the user to
adjust the provided gasPrice, we don’t believe the default gasPrice
changes eiectively (although what the word e"ectively means is hard
to say).
One can see the cost of gas per transaction in Finneys lowering in
recent days, but that price is not lowering as quickly as the price of
Ether vs. US dollar is rising, thus the divergence.
I am tempted to say that this implies an imperfect gas market, but I
think that would be a mistake. If one looks at the gasPrice market in
terms of US dollars, then it does seem ineiicient. But I think one must
view the market from the perspective of the “realm of Ether.”
One would think that, if it were a perfect market, some miners would
accept lower gas prices (because Ether is worth more), and therefore

the price of a transaction in Ether would lessen as quickly as the price
of Ether per dollar rises, but this is not happening.
I live in America. I use US dollars every day. I don’t use Ether every day.
When I think of Ether, I automatically translate it into US dollars. But if I
enter the “realm of Ether” — if I try to force myself into thinking about
Ether only — then I can see that there is no intrinsic reason for the cost
of a transaction to change simply because the price per Ether changes.
Certainly, there is no reason for it to change as fast as the price of Ether
per US dollars. The amount of calculation required by the node
software does not change because the price of Ether in US dollars
changes. Plus, miners are not incentivized to lower prices as the price
of Ether rises. Their proﬁts (in US dollars) are skyrocketing. Why lower
prices? The competition has not increased (in terms of computing
cycles needed to win the reward).

Conclusion
We are developing QuickBlocks because we want to “Account for the
Revolution™.” We have a thousand ideas of what that might mean. Other
people have done this sort of analysis before, however, many of those
analyses suier from one of two problems: (1) they are slow, or (2) they
are centralized. QuickBlocks is both fast and fully decentralized.
We’ve written previously about an analysis of the DAO hack that took
more than 12 hours to scan just 2 days’ worth of DAO related data.
QuickBlocks produced the above data (more than 70 days of data) in
less than four minutes. Those are some quick blocks!


The Real Flippening
Tuesday, May 30, 2017

My wife constantly tells me we should sell our Ethereum holdings, and
I constantly resist. She wants to buy a new sewing machine. I want to
wait until “The Flippening.”

What is “The Flippening”?
There’s an old biblical word quickening which refers to the moment an
unborn infant ﬁrst moves in its mother’s belly. I love that word, and
when I ﬁrst heard the word ﬂippening, I thought of it.
So, maybe, The Flippening is like a quickening. I was thinking maybe it’s
when Ethereum ﬁrst moves inside its mother’s belly. But then I thought,
Ethereum doesn’t have a mother. Even if it did, its mother wouldn’t have
a belly, so I had to look elsewhere for an explanation.

Bitcoin is a Bucket and the Bucket has a Hole
Bitcoin’s bucket is ﬁlled with something
valuable: mindshare. However, without
going into why (its developers), I think the
bucket has a hole in it. Bitcoin is leaking
mindshare. Developers are moving to the
Ethereum space. Why wouldn’t they? Ethereum is where all the
excitement is, and it’s not a pain in the ass to program.

People Move Like Fish in a School
Have you ever seen a video of a
school of sardines swimming
peacefully, when suddenly, a yellowﬁnned tuna appears out of nowhere
and starts eating them up? Every
individual sardine ﬂees at the same
time. However, because they are so
tightly packed, they are restrained in
how they can move. No matter
which direction they turn, they bump into another sardine. This causes
them to move in what appears to be a coordinated blob, but it’s just
individuals all of whom are rushing for the door. They are like rats
leaving a sinking ship.

People Rush for Doors
Why is it socially unacceptable to falsely yell “Fire” in a crowded movie
theater? Because if you do, and everyone rushes for the door (that is,
ﬂees the yellow-ﬁnned tuna), someone will get hurt.

Who Cares?
There’s a yellow-ﬁnned tuna swimming towards Bitcoin. And that
yellow-ﬁnned tuna is the realization that Bitcoin’s bucket is leaking.
There are already voices, like mine, murmuring “What’s that smell?”
because they smell smoke coming from under the gum-impregnated
seats in the ratty movie theater. The baby is moving in the belly.

That’s Not the Flippening I’m Talking About
Hopefully, you get the idea of what I’m saying. “The Flippening” is
coming. One of these days, everyone is going to wake up to the fact that

Bitcoin is failing, and they are all going to rush for the exit door labeled
“Ethereum.” Will you get trampled?
----I was listening to to an Epicenter podcast this morning about the ENS
(Ethereum Name Service). Listen to it if you get a chance. It’s a very
interesting idea. It will help you to understand one of the reasons why
developers are ﬂocking to Ethereum.
I want to note speciﬁcally, however, what the host says at the 101:15
mark of the podcast. It will help explain what I call the “Real Flippening.”
I quote:
…so, I see perhaps a problem…if I’m locking ﬁve Ether for a domain,
and that’s worth $500.00 at today’s price, a year from now, if ether is
worth, say, $1,000, [then] all of sudden, I’m no longer locking $500, I’m
locking $5,000….
I see this question as a misguided understanding of what’s really
happening. The ﬂippening I’m waiting for is not the one from Bitcoin to
Ether. It’s a ﬂippening from ﬁat money to cryptocurrency.
Allow your mind to wander forward in time to when there are no US
dollars. In fact, there are no ﬁat currencies at all. There’s only crypto. It
doesn’t matter what it’s called. Ethereum on a Stick™, for all the name
matters.
In this imagined world, the statement by the Epicenter host makes no
sense. In a crypto-only world, the question would be like saying
…I see a problem…I locked up ﬁve dollars at the start of the year…and
it’s still ﬁve dollars at the end of the year…isn’t that a problem?
The answer, obviously, is “No.” The mentality that has us all constantly
pricing crypto in real-world ﬁat is why locking up funds for an extended
period seems like a bad idea.

The Real Flippening, the one that my wife is going to have to wait for, is
when people start pricing goods relative to a crypto instead of ﬁat. What
will the world be like when that happens?


It's Growing It's Growing
Tuesday, June 27, 2017

I spend much of my free time on the Ethereum Stack Exchange
answering questions about Solidity and asking questions about the
Ethereum data, so I’m aware which questions are most frequently
asked.
By far, the most repeated question on that forum is how long it takes to
sync the Ethereum blockchain using geth or Parity. The short answer to
that question is It takes forever. But this is only a perception. In real life,
the sync eventually does end. It just takes forever.
I recently re-synced the Ethereum blockchain from scratch using Parity.
It took more than 24 hours. You can see why it takes so long above.
Do you see that huge spike on the chart on the next page around block
2,380,000? That spike is one of the reasons the sync takes so long. This
is the DDOS attack that instigated the Tangerine Dream hard fork. A few
weeks later, the Spurious Dragon cleaned up the mess.
But this post is not about syncing — it’s about what’s been happening to
the blockchain data since that time.

What Happens When the Masses Adopt?
As many of my readers will know, there has been a huge increase in
interest in the Ethereum blockchain over the last few months. We were
curious what eQect this attention has had on the size of the data (and
therefore on the time to sync). The data we present below was created
using our tools.
QuickBlocks ensures speed of access through various techniques, one
of which is how we store the data on disc. At the time of this article,
there were 3,800,000 blocks on the current chain. QuickBlocks stores
1,000 blocks per ﬁle in each of 3,800 ﬁles, providing a compromise
between speed of access and size on disc. We take advantage of this
1,000 block per ﬁle data to present the following results.

Because we wished to study the data relative to previous activity, we
ﬁrst calculated the average size of a block for the 1,000 blocks between
block 2,000,000 and 2,001,000. We then used this value as a divisor for
each previous and subsequent 1,000 block range and charted the
results. Viewing the data relative to a ﬁxed mark accomplishes two
things: (1) it smooths out the greatly varying data to better focus on
trends, and (2) it allows us to more clearly see where the size of the data
is coming from, and importantly where it’s going.
A few things stand out:
•

For most of 2016 and the ﬁrst part of 2017, the average size per
1,000 blocks remained relatively constant. This is indicated by the
up and down nature at the middle of the chart.

•

The DDOS attack during October 2016 is very clearly visible as is
the eQectiveness of the Spurious Dragon hard fork.

Looking at the right side of the chart above, I can say with conﬁdence
that we’re headed for slower sync times.

Why Does the Sync Take So Long?
We focused our analysis on the size of the data because it’s a proxy for
the time needed to sync the chain. As I mentioned above, I spend a lot
of time answering questions about why it takes so long to sync the
chain. Until now, I’ve always advised people to be patient while trying
to sync the blocks near block 2,380,000. I need to amend my answer.

The area under the graph above is reﬂective of the time it takes so sync
that range. You can clearly see the area is growing recently through the
natural growth of the chain given the recent attention to Ethereum. The
time to sync is headed sharply up. You thought it was slow before.
The above chart presents the same data as before but here we present
the percent of the reference blocks. The eQect is even more
pronounced. The size of the 1,000 block groups for recent blocks is
nearly 50% of the height of the DDOS attacks. This does not bode well
for the future of syncing the chain.

The Number of Empty Blocks is Decreasing
Another easy to obtain measure for the increased usage of the system
is to count the number of empty blocks (that is blocks that contain zero

transactions). While there are many reasons why a block may have zero
transactions, one reason is that the number of transactions available
to be processed is below the chain’s ability to keep up. During the early
part of the chain’s history, when there was signiﬁcantly less activity,
there were more empty blocks.

As one can see, recently the chain has experienced a sharp decrease
in the number of empty blocks. This, we think, is indicative of the
increased load on the system. On average, over the entire history of the
chain, nearly 35% of all blocks contained no transactions. This number
has dropped signiﬁcantly to just more than 10% of blocks.

Conclusion
There are many other charts we could produce. My guess is that all of
them would indicate a huge increase in the growth of the storage
requirements being placed on the system and, as an extension, the
time to sync.
One of the charts I would like to see is the rate of creation of new smart
contracts. It is my guess that we would see two sources of growth in this
area:
•

Regular growth from new users adding individual smart contracts,

•

An increased use of the factory model of smart contract creation.

In the factory model, smart contracts dynamically create other smart
contracts. I’d also like to see how many of smart contracts are being
created and never used. I suspect there are a huge number of such
contracts.
I fear that creating smart contract from a factory contract will ultimately
turn out to be a serious problem. The large-scale creation of single-use
or low-use smart contracts is like water pollution. The ultimate
outcome of this pollution is a degradation of the system. It’s like
spewing poison that never decays.
With QuickBlocks we are attempting to build a system that allows us to
see these overall trends. We need to pay attention because only by
looking can we anticipate future behavior of the system.


Be Careful Little Brain What You Code
Thursday, August 3, 2017

When I was a kid, my mother used to send my brothers and me to
vacation Bible school. Having children of my own, I’m certain she did
that to get rid of us for the summer. The experience had no lasting eAect
on me (thank God), but I remember a song we used to sing:
Be careful little eyes what you see
Be careful little eyes what you see
Because Jesus up above
Is looking down in love
So be careful little eyes what you see
I was only eight years old, but I remember thinking to myself, “How
would I know — before seeing something — to shut my eyes to protect
myself from seeing the same thing that I wasn’t supposed to see?”
It was just one of many things that didn’t make sense to me back then.
I realize now that I simply did not understand a concept called
Situational Awareness or Planning Ahead.

The Blockchain Will Save the World. Hallelujah!
I had two experiences recently, both related the blockchain and both
related to the idea of a person’s identity. The ﬁrst one excited me. The
second one made me think of that song.

If you’ve done any reading
about the blockchain you will
have come upon the idea of
putting a person’s identity on
the blockchain (or her
medical records, or his land
records, or ﬁll_in_the_blank).
It’s a deep, deep subject, but
the idea is clear enough.
The argument goes something like this: billions of people do not have a
reliable form of identiﬁcation. This means they cannot prove who they
are, which means they can’t fully participate in civil society. They can’t
get a bank account, which means they can’t oAicially own anything.
With the blockchain, no-one will be able to take a person’s identity
away from them, and this is a good thing.

The Evidence Was Washed Away
I was at a blockchain conference in Washington, DC last week. It was
well attended, had tons of great speakers, and was overﬂowing with
people engaged in the blockchain space. Being in Washington, DC,
there was, of course, a lot of discussion about regulation, recent SEC
rulings, and other lawyerly things. But the most impressive speaker for
me was a layman. She was from Australia. Her name was Dr. Jane
Thomason.
Dr. Thomason began her talk by displaying a photograph of a place that
had been ravaged by a tsunami. The image showed a before and after
view. In the after view, everything had been wiped away. Roads, trees,
buildings, people’s lives, people’s hopes, relationships, everything.
Dr. Thomason pointed out that not only were the physical objects
literally washed into the ocean, but something much more important
was washed away as well.

She asked those of us in the
audience to raise our hands if
we had ever had an “ah-ha”
moment related to the
blockchain. Many people
raised their hands.
Dr. Thomason’s moment came as she stood on the newly denuded land
in that devastated place knowing that every single government
document was gone. Every government record…every land record,
every driver’s license, every elevator inspection certiﬁcate, every birth
certiﬁcate, all of it, down to the very last piece of paper was gone.
“Ah, ha,” she thought, “that information could have been stored on the
blockchain. If it had been,” she thought, “we could have restored
everything instantly.” As it was it took many months to recover even a
fraction of the information that was lost.
To those of you who are thinking, “You can’t store that much data on the
blockchain,” remember there is a diAerence between a design idea and
an implementation detail. Design idea: store all the information on an
immutable ledger. Implementation detail: ﬁgure out how to do it in an
aAordable way.
Dr. Thomason went on to explain the other work she was doing, all very
interesting and important. I sat in my chair as inspired as I had been
since I ﬁrst heard of the blockchain. I remember my early reading in the
space about how the blockchain would change this unchangeable
world…how the blockchain might be able to make the world a better
place.
Much of the evidence of that early excitement was washed away the
morning of June 17, 2016, with the DAO hack, then later in October of
2016 with the DDOS attacks, then every few weeks since then with

these idiotic ICOs. (As if ICOs and a disruption of venture capital was
the reason for the blockchain.)
“Here’s my white paper. I have an
idea that I can explain in a single
Medium post. Please give me
millions of dollars. Here’s some
tokens…”
The potential that I fell in love with two years ago came back to me as I
listened to Dr. Thomason speak. She ﬁgured out a way to change the
world. She was going to make the world a better place.

Not so Fast...
The second experience I mentioned above, the one that brought me
back to earth, happened this afternoon at the Holocaust Museum in
Washington, DC.
If you’ve never visited the Holocaust Museum in Washington, DC you
should. It’s a troubling experience, and there are many, many horrible
stories to be heard, but there are also stories of bravery and true human
kindness.
There are stories of a priest who came to the rescue of hundreds of Jews
who might otherwise have been shot down in the streets of small
Eastern European village. There are stories of French oAicials who
secretly resisted the Nazis. There are stories of people hiding other
people all over Europe.
The story that caught my eye, however, was of a police chief in Athens,
Greece. I quote from the placard tucked away in a tiny alcove three
quarters of the way through the exhibit:
Chief of Police Angelos Evert of Athens, Greece, used his position to
help save Jews from deportation [to the death camps]. After the

Germans occupied Athens in September 1943, Evert began issuing
false ID cards to Jews in response to an appeal by Archbishop
Damaskinos, whose Orthodox Greek Church provided false baptismal
certiﬁcates…In the spring and summer oNicials and police deported
more than 45,000 Jewish residents “to the East”…Few members of
[the community] survived.
My immediate thought was this: Where are the false identiﬁcation
cards and fake baptismal certiﬁcates in a world of immutable records?
How can an honest to goodness hero fake an ID in a world where IDs
cannot be faked? Previously, I had been riding high on the belief that
providing immutable, unfakeable, impossible-to-destroy identiﬁcation
documents was a noble goal. But here was irrefutable evidence that
there are times when you really need a fake ID.

Be Careful Little Brains What You Code
Very few people would ever purposefully design a system for evil intent.
I don’t want any of my readers to think that’s what I’m saying. But evil
things do sometimes arise as unintended consequences out of
powerful, complex systems.
I’m not an opponent to these systems. In the end, I think, the amount of
good they provide will far outweigh the amount of bad, but we need to
think very carefully about unintended consequences. We need to focus
on ways to allow good-hearted people to act in ways that may be in
opposition to those in authority. If we are not careful, these new types
of identiﬁcation systems may become history’s most powerful tools of
oppression.
There’s a diAerence between design and implementation. If our grand
design is to create a perfect digital identity for everyone, then our
implementation better include a way for people with truly noble
intentions to act against that same system. I have an idea…let’s raise a
friggen’ token sale.


Is the Ice Age Affecting Block Production?
Wednesday, August 30, 2017

This is just one of many questions we have about what is happening on
the Ethereum blockchain. We’re developing QuickBlocks (since
renamed TrueBlocks1) to help us answer these questions and more.
TrueBlocks is a set of software libraries, applications, and command
line tools that provide fast, easy, fully decentralized access to the
Ethereum blockchain. As part of preparing for our next release, we’ve
written several command line tools. One of these is called whenBlock
(since renamed chifra when2) which we used to write this article.
chifra when allows you to ﬁnd the block number given a particular
date or, conversely, a date given a particular block number or
timestamp. It’s very straight-forward but surprisingly useful. For
example, if you’re waiting for the node to sync, you can type chifra
when latest to ﬁgure out how close to the chain tip the sync.
Ø chifra when latest
blockNumber
2400000

timestamp
1475896666

date
2016-10-08 03:17:46 UTC

This is the help screen from chifra when:
This article was written in 2017 when our project was called QuickBlocks. It’s since
been renamed TrueBlocks. The essays in this collection use the name chifra. If we
missed some, do not be confused.
2
Like the project, the command line tools have been renamed as well. This tool is now
called chifra when.
1

chifra when --list shows a customizable list of so-called special
blocks. (We show the ﬁrst few here. There are many more.):

You may add your own special blocks if you wish (for example, when
such-and-such a smart contract was deployed). The --verbose
option reports on other interesting blocks, such as the deployment
dates of prominent tokens and hard forks.

Answering Questions
The whole point of TrueBlocks is to allow end users to answer questions
about the Ethereum blockchain. We aim to provide useful, scientiﬁcquality information about the behavior of the blockchain. Currently,
one must rely either on a fully centralized solution such as Etherscan to
get data or do without the data altogether — or accumulate it oneself,
which is not easy if one hopes to get the data before the end of the
century.
We used chifra when to create a list of the block heights at the end
of each week since the inception of the chain. Here’s the data:

Our goal was to see how the diQiculty bomb was aQecting the number
of blocks created each week. Here’s that same data displayed in a
pretty blue, orange, and green chart. Do you see the diQiculty bomb’s
eQect?

Several things become obvious from the chart:
•

The block height (total # of blocks, orange line) shows a nearconstant slope. This is to be expected. You can see it trail oQ slightly
in recent months. This is due to the diQiculty bomb.

•

The diQiculty bomb is revealed when one looks at the number of
blocks created per week (the blue line).

•

The chain experienced large swings in the number of blocks
produced per week in the early months. (I’m not sure why.)

•

The Homestead release (block #1,150,00 — March 14, 2016)
initiated a near-instantaneous (and signiﬁcant) increase in the
number of blocks created per week.

•

The number of blocks created per week remained stable for more
than a year after the Homestead release.

•

While the October 2016 dDos attack had a visible eQect on the
chain (the two dips), it was not as pronounced as other eQects
(even though the dDos does have a huge eQect the blockchain’s
initial sync).

•

Starting around April of 2017, the diQiculty bomb began to explode.
It has been exploding ever since.

•

The explosions of the diQiculty bomb are growing larger each
period. (I do not know the details of the diQiculty calculation, but it
is apparently a step function with exponentially growing increase3.)

•

The number of blocks produced per week is almost as low as it’s
ever been.

Conclusion
There’s a lot more we could do to explore the eQect of the diQiculty
bomb. For example, it would be interesting to see if the number of
transactions processed has seen a commensurate decrease.
Unfortunately, we must go back to the salt mine of our IDE to ﬁnish up
the release.

3

This presumption proved to be true as you’ll see in later articles.


Block Production Continues to Slide
Saturday, September 9, 2017

Two weeks ago, we wrote an article in which we describe the slowdown
in per-week block production due to the di7iculty bomb. We thought it
would be interesting to continue to watch the process as it unfolds. We
wondered “Is the di7iculty bomb having its desired e7ect”?

Ruh Roh George Jeston!
Compare the chart from our previous article with this chart we
generated this week. Do you see what’s going on?

As you can see, the di7iculty bomb — whose purpose is to slow down
block production — is working very well. The number of blocks
produced during the week ending September 8, 2017, was as low as it’s
been since the inception of the chain.
The inverse of the of the number of blocks produced per week is the
time it takes to produce each block. In the chart on the following page,
the increase in the length of time it takes to produce a block is apparent.
In this chart, we show the number of seconds it takes to create each
block since inception. After the Homestead fork, the number of
seconds to produce a block hovered almost exactly on 14 seconds for
more than a year. It’s been rapidly increasing since April when the
di7iculty bomb began to explode. Furthermore, the stepwise behavior
of the increase is apparent in the jagged ascent.

Does “Less Blocks” Mean “Less Transactions”?
The tweet1 to the right claims the number
of transactions per day is growing, but
this seems to contradict the idea that
there are less blocks being produced.
How can “increasing transactions” and
“less blocks” both be true?
Below, we used TrueBlocks to collect the number of transactions per
week between August 2015, and September 2017. We wanted to see if
the decrease in blocks per week is a7ecting transaction throughput.

Number of Transactions
Below we present the number of transactions per block, which, as you
can see, has recently increased signiﬁcantly. As we surmised, the
blocks are getting fuller.

1

Sorry, we lost the attribution.

The huge increase in the number of transactions per block is intensiﬁed
by the increase in interest in the Ethereum chain and the simultaneous
lowering in the number of blocks per period. But looking at transactions
per block doesn’t really show what we’re after; we needed one more
chart (not to mention we like making charts).
This next chart was a bit surprising to us.

This chart clearly reveals why the previous chart shows such a huge
increase in transactions per block. Half as many blocks…twice as many
transactions…means four times as many transactions per block.
While it is not clear without further information if the number of
transactions would be higher in the absence of the di7iculty bomb,
there is a slight hint at a bumpiness in the uptick of adoption (as
represented by number of transactions). One would think that
increased adoption would be smoother without the e7ect imposed by
the di7iculty bomb.

What is the DiAiculty Bomb For?
The purpose of the di7iculty bomb is to force Ethereum miners to
accept an upgrade to the node software. Is it serving that purpose?
In a certain sense, the pressure applied on the miners to accept this
“forced upgrade” comes from the community as the system slows
down. But — is the system slowing down from the perspective of a
regular user if the number of transactions being processed per period
is increasing in lock step with the number of new users? Do the users,
in aggregate, even notice slower blocks?
Another pressure imposed on the miners, of course, is to accept the
“forced upgrade” because of the lessening of the available income per
period. If there are less blocks over a given time frame, there are less
rewards available for the winner of the block, and therefore less income
to miners.
To what extent does the increasing number of transactions (and
therefore the increasing transaction fees) o7set the lost block rewards?
On top of that, is the price of gas increasing or lessening, and what
e7ect does that have on the situation? Furthermore, if the price of Ether
in US dollars is rising (as it has over the last few weeks), does that
ameliorate the e7ect of the di7iculty bomb? And, ﬁnally, is it even
possible to enumerate and untangle all the intertwined e7ects of these
things inﬂuencing each other and how can TrueBlocks help?


Reading Byzantium's Tea Leaves
Monday, October 2, 2017

I used to watch a television show called “Wall Street Week with Louis
Rukeyser.” The host would frequently laugh at the way Alan Greenspan,
the Chairman of the US Federal Reserve Bank, would testify in front of
Congress. Rukeyser found it comical that everyone had to “read the tea
leaves” after Greenspan spoke.
“I should warn you, if I turn out to be clear, you’ve probably
misunderstood me.” ~Alan Greenspan
Why did Mr. Greenspan speak in such obscure and diJicult to
understand language? Why did he force people to “read the tea
leaves”? Because he knew that a single wrong word had the power to
send shockwaves through the markets. “We’re going to do exactly
this…” Crash goes the market.
I was thinking about this recently as I followed the increasingly
interesting discussions about Ethereum’s impending Byzantium hard
fork (EIP86).
If there is one distinguishing characteristic about this discussion, it is
that there are no tea leaves. There is simply an unambiguous statement
from the core developers that the per-block reward is going to be
lowered from the current ﬁve Ether per block to three Ether per block.
Okay, maybe there is some tea, but software engineers care about
source code...Rukeyser also used to frequently say, “Let’s go to the
numbers.” So, let’s do that, let’s go to the numbers.

The above two lines (and a few related lines elsewhere) implement the
per-block change in the mining reward. That’s what you might call
“unambiguous.”

A Fully Manufactured Conversation
The following imagined conversation between the core developers and
miners was overheard inside of my own mind:
Developers: “We’re going to make a change as part of the Byzantium
hard fork to lower the per-block reward from ﬁve to three…”
Miners: “Wait. Why are you doing that? That will cost us a lot of
money…”
Developers: “You knew the code worked this way…the diJiculty bomb
has been in the code since the beginning…”
Miners: “Changing the block reward has nothing to do with diJusing the
diJiculty time bomb1...”

Developers: “Block production on a per-day basis has been
decreasing for months. After the fork, the per-day reward will bounce

1

The miners have a point here

back. In fact, it will be higher. Three Ether every 14 seconds is more than
ﬁve Ether every 25 seconds2.”
Miners: …but…but…that’s diJerent…okay…there’s nothing we can
do…

Reading the Tea Leaves
I think the developers are doing the right thing by lowering the block
reward. I suspect that the ecosystem is paying too much for security,
but their argument is a bit disingenuous. They’ve conﬂated the two
issues of block reward and diJiculty calculation.
Let’s look at some tea leaves. (It’s math, but let’s pretend this is tea).

Above is a description of how diJiculty is calculated at each block in
the current code (pre-fork). You’ll notice there is no mention of the
block reward. That’s because the block reward is in a diJerent part of
the code. Learn more about the reasoning behind this calculation in our
other articles.

2

The developers have a point here.

Let’s look at what’s changed between Homestead and Byzantium. The
following image is a visual diJerence between the pre- and post-fork
code. (I simpliﬁed it for clarity.)

There are two changes to the post-fork code. First, a small change in
section A that makes mining uncles less likely — (EIP100) and second,
a resetting of the diJiculty time bomb at section B.

To reset the time bomb, the calculation subtracts 2,999,999 from the
parent’s block number (as opposed to adding one as it did before).
Assuming the Byzantium fork happens at block 4,370,000, this change
acts like a time machine throwing the calculation back 529 days to
block 1,370,001 (April 20, 2016, at 7:40:59 UTC). Using the same time
machine to project forward, the system will return this amount of timebombyness by March 14, 2019. Note that the remainder of the
calculation remains unchanged.

How Will this Change A=ect the Ecosystem?
One of the things Rukeyser constantly noted was that the market,
knowing what it knows, discounts the future. In other words, the market
builds all conceivable futures into its current price. Buy on the rumor,
sell on the news.
The change in the diJiculty calculation at the Byzantium hard fork is
well anticipated. Presumably, current prices already reﬂect it. However,
(and I think Alan Greenspan might understand what I’m saying here),
the change in the calculation is a step function and is not obscure. I
think that’s ultimately a good thing.
In previous writings, I’ve noted that the number of transactions per day
has been increasing despite decreasing per-day block production.
What will it mean to the ecosystem when there’s a sudden jump in the
number of blocks per day? Will gas prices spike? Will gas prices
plummet? Will the hash rate decrease or increase? Will the ﬁat price of
Ether rise or fall?
In the next installment of this series, we will explore these questions
and others. We hope to explore these questions under the light of a ﬁatpowered spotlight—do these issues become more or less clear in the
light a fait-based analysis? Stay tuned, we’ll let you know.


Announcing QuickBlocks1
Sunday, October 29, 2017

We are pleased to announce the release today of the ﬁrst of three white
papers describing TrueBlocks: “Faster, Richer, Fully Customizable Data
from Programmable Blockchains.” Additionally, we are announcing the
release of an updated and improved website (http://trueblocks.io) as
well as nine open-source command line tools.

1

Since renamed TrueBlocks.

What is TrueBlocks
TrueBlocks is a collection of software libraries, open-source
command-line tools, and applications intended to improve the ease
and speed with which one may access data from the Ethereum
blockchain. Given fast, easily accessible data, new use cases and
applications such as smart contract monitoring, per-block accounting,
and integration with business intelligence tools such as Tableau
become possible.

Open-Source Command-Line Tools
As part of our commitment to the Ethereum community, we oNer nine
open-source command-line tools including:
chifra state and chifra tokens2, which report on Ether and
token balances respectively.
chifra blocks, chifra transactions, chifra receipts,
chifra traces, and chifra logs, which report on the particular
data structures given one or more block hashes, transaction hashes,
block numbers/transaction id pairs, or any combination.

We also provide utilities such as chifra names and chifra when,
which are useful in retrieving information about Ethereum account
names and the timing of certain blocks or transactions.

Examples of Using the Command Line Tools
Having been born and raised on the Linux command line, each of these
tools works and plays well with the Linux philosophy of piping data from
one command to the next.

Originally called getBalance and getTokenBal. All commands were renamed as
subcommands to the single command line tool chifra in 2019.

2

For example, you may run the following command:
chifra names TheDao 0xbb9

to get the address of the original DAO token:
tags
address
name
55-Defi 0xbb9bc244d798123fde783fcc1c72d3bb8c189413 TheDao

You may also get the address of the 945th prefund address. (We happen
to know it still holds a lot of Dao tokens.)
Run the following command to get ﬁnd address:
chifra names prefund_0945 -p -s

The -s returns only the address so we can feed it into other commands:
0x1ba9228d388727f389150ea03b73c82de8eb2e09

Now, we can string these together to get the Dao token balance for that
account at a block:
chifra tokens --fmt json \
(chifra names TheDao 0xbb9 -s) \
(chifra names -p -s prefund_0945) 4415101

This returns the following JSON data.
{

}

"data": [
{
"address": "0xbb9bc244d798123fde783fcc1c72d3bb8c189413",
"addressName": "TheDAO (L1)",
"balance": "10200000000000000000000",
"balanceDec": "10200",
"blockNumber": 4415101,
"decimals": 18,
"holder": "0x1ba9228d388727f389150ea03b73c82de8eb2e09",
"holderName": "Prefund_0945",
"name": "TheDAO (L1)",
"symbol": "TheDAO"
}
]

This shows that, at block 4,415,101 address 0x1ba9228d3887
27f389150ea03b73c82de8eb2e09 owned 10,200 DAO tokens. Of

course, you can take this much further. The combinations are nearly
endless. You can use a block range (1428756-6415101:period) to
deliver daily, weekly or monthly token reports. You can pass in multiple
addresses at a time (either tokens or holders) to produce all sorts of
data. You can even make a cap table for a token. (We’ll leave that as an
exercise.)

Where To from Here?
We have many plans for the software including some unique takes on
how to deal with the Ethereum data on small machines. We’ll be
presenting some of our ideas over the coming month with the release
of two more white papers3. We welcome any participation and look
forward to your interest in our work. We will see you at DevCon 3.

These two white papers, the white paper we announced in this article, and a detailed
speciﬁcation of the Unchained Index are available in the companion volume.

3


Playing with Blocks
Friday, December 8, 2017

I am so deep down the rabbit hole of Ethereum data that I sometimes
forget there is such a thing as the real world. And I care exactly zero if I
never come up. Don’t ask me why. I’m pretty sure there’s no good
reason.
Maybe my obsession with data has to do with something that happened
to me in high school. My guidance counselor administered a test
intended to identify which career I should pursue. She watched me take
the test and then carefully scored my answers. When she was done,
she looked across the table and said, “You should look for careers that
have nothing to do with people.” The remaining options for me were
making things (factory work) or information processing. I chose the
latter.
I’ve written software called TrueBlocks (http://trueblocks.io) that
indulges my obsession. It scans the Ethereum blockchain extracting
data. TrueBlocks allows me to easily summarize and accumulate data
at each block. From there, I can make pretty pictures. Such as the ones
on the following pages.
In that chart, each dot represents two characteristics about “buckets
(or groups) of 10,000 blocks. We scanned the ﬁrst 4,000,000 blocks and
summarized the number of non-empty blocks in each group vs. the
average number of transactions in each group. This created 400 data
points (4,000,000 / 10,000 = 400).

On the horizontal axis, we plot the number of non-empty blocks in each
10,000-block group. The dots at zero-zero (the lower-left corner of the
chart) represent groups with zero transactions. (These are most likely
blocks at the start of the chain — the ﬁrst transaction happened around
block 47,000.) The left side of the chart represent groups of 10,000
blocks where nearly all are non-empty (i.e. have at least one
transaction).
The vertical axis shows the average number of transactions inside the
non-empty blocks in the group. At the bottom of the chart are groups of
10,000 blocks that average between zero and ten transactions per
block. Near the top-right of the chart are a small number of groups that
average between ﬁfty and sixty transactions per non-empty block.

What Does it Mean?
We wondered if there was anything interesting to be gleaned from the
data. Do you see anything? I something of a pattern, but I’m not sure I
understand what it means.

There appears to be three distinct groups in the data. We’ve shown
them with an orange, a purple, and a green line. We did our best to ﬁt
these apparent curves.

I have no good explanation for why the data looks like this. I’ll give a
possible explanation below, but I welcome suggestions.

Possible Explanation
My initial thought was that these patterns have something to do with
the Ethereum mining software. I thought perhaps there are three
diYerent types that miners naturally fall into. Something to do with
proﬁts. But how to ﬁgure out if that hypothesis was true? No idea.
I decided to look at a simpler problem: At what point in the history of
the chain were these groups of blocks produced?
The three charts above have no time component. To alleviate this
problem, I colored the dots by blockNumber. Here’s the same data
above with the dots colored by blockNumber.

In this chart, earlier block numbers are colored yellow. The color tends
toward dark red the higher the block number gets. Clearly, as the chain
progressed, there was an early period with many empty blocks with
relatively few transactions. Towards the middle of the chain’s lifespan
(between December of 2016 and June of 2017, perhaps) there seems
to have been a retreat.
Near the end of the early period, the number of empty blocks was
decreasing as the average number of transactions rose (as yellow
bleeds into orange). But then, for some reason, the chain makes a
retreat. While the average number of transactions per block continued
to increase, the number of non-empty blocks per 10,000 also shrank.

Why Does This Happen?
I didn’t really have a good explanation of why this retreat happened. I
thought it might be possible that retreat was just an artifact of the way I
was presenting the data. So, I tried to dig deeper into that thought.
Next, is a chart showing the number of non-empty blocks plotted
against the number of empty blocks. Obviously, if you lose one from

one side, you gain one from the other, so the two lines mirror each other.
The two lines are reciprocal.
But does this explain the retreat? It looks like something happened
around block 1,700,000 (the high point of the grey line near the middle
of the image). What happened around block 1.7 million?

The DAO hack happened at block 1,718,497. Within days of the hack,
the price dropped from $21.00 US to about $8.00. Three months later,
there was the October 2016 DDos attack. You can see it in the spikiness
around block 2,250,00. Did these two events cause a retreat in the
activity on the chain? It seems likely to me.


Defeating the Ethereum DDos Attacks
Monday, February 12, 2018

If you’ve ever studied the Ethereum chain data, you will be familiar with
a certain set of transactions that take a very long time to process. These
transactions happened between blocks 2,286,910 and 2,717,576. They
are a pain in my a$$.

In a surprisingly eGective attack, some evil genius took advantage of an
underpriced EVM opcode to create and then destroy millions of dead
smart contracts. This had the eGect of signiﬁcantly bloating the state
database, but more importantly for our purposes it created tons of
transaction traces1.
Until recently, if we were scanning the Ethereum database, we had to
wait many hours (perhaps even days) while the node software delivered
these useless traces. We could have cached them, but our goal has
An EVM trace happens once for each transaction plus once each for each time a smart
contract calls into another smart contract during that transaction. You may query
traces through the RPC with eth_getTrace.
1

always been to create a minimal impact on the target machine (this
helps us stay decentralized). With the solution we present below, we
can now eGectively choose whether to scan, skip, or cache these
troublesome transactions. This post discusses how we did that and
how we now routinely scan very quickly through this diGicult portion of
the history of the chain.

A Short History
The morning of DevCon 2, in October 2016, there was a hack against
Ethereum. The Ethereum developers responded quickly, and ﬁxed the
hack, but a few days later another attack occurred. This second attack
went on for more than a month and is described elsewhere.
In response, the core developers conducted two hard forks: Tangerine
Dream (EIP150) and Spurious Dragon (EIP161). As one would expect,
the hard forks did not change existing data. Instead, they changed the
way the client code works so that, at the end of a transaction, if an
account has been “touched” during that transaction, and the account
was otherwise empty, the account (and its state) was removed.
The attacker created millions of useless smart contracts across
thousands of transactions. After the hard forks, an oG-chain process
was initiated to ‘touch’ these accounts so they would be removed. This
worked to lessen state bloat, but it created a second huge amount of
‘cleanup’ transactions (and traces). Needless to say — this entire
section of the blockchain — from the start of the hack to the end of the
cleanup — is ugly and bloated (which translates into “slow,” which is a
problem).

Let’s Go to the Data
As I said, these troublesome transactions are especially annoying if you
want to view traces (as we do). We needed a solution.

The ﬁrst thing we did was to gather some information. We scanned the
ﬁrst 4,000,000 blocks. At each block, we looked at every transaction
and counted the number of traces generated by that transaction.
From this data, we created a heat map. Each cell represents a 50,000block range (x-axis) and a trace-count bucket (y-axis). Darker red
indicates more transactions falling into that bucket.

The dark band at the top of the chart shows that most transactions,
throughout the chain's history, produce very few traces. As the chain
matured, the heat spreads downward: transactions produce
progressively more traces, reﬂecting the growing complexity of
deployed smart contracts.
Do you notice anything else of interest? Do you see that very dark red
sections at the bottom of the chart?
That is the Fall 2016 dDos attack. It stands out like a sore thumb (i.e.
there are 1,000s of transactions with many 1,000s of traces in that
region). We noticed this and thought, “We could easily box-in those
troublesome transactions. We can skip over them.”

Writing Code that Skips Ugly Transactions
It turns out, the solution to our problem was relatively simple. What we
did is identify any transaction between 2,286,910 and 2,717,576 that
had more than 500 traces. First, we needed a way to ﬁgure out how
many traces a transaction had without querying the full trace data.
Querying the traces was the problem after all.
Luckily, the RPC provides a function to query a single trace. We use that
to decide if the transaction has or does not have a trace at a given
location.

For example, calling hasTraceAt(hash, 4) would return true if the
transaction at that hash had ﬁve or more traces (zero-based index).
Next, we needed a way to get the number of traces given a transaction
hash. We wrote the following a binary search getTraceCount which
ﬁnds where in the trace history trace n exists but trace n+1 does not. n
is then returned as the number of traces in the transaction.

From here, it was a simple matter to write the function
isSpam_2016_10 which returns true if the transaction has more than
5002 traces and occurred during the dDos period.

Our calling code can now be conﬁgured to decide which transactions
to fully trace (and/or cache) without having to wait the “forever” amount
of time to query the traces.
The following list of accounts, while not deﬁnitive, were found to be in
some way involved with the dDos either through cleaning up process or
through initiating them. We sometimes also use this list of accounts to
simply skip over transactions if the analysis allows it. In some cases,
we scan these accounts, in many cases we don’t.
const Address relatedAccounts[] = {
// These smart contracts carried
// out the stateClear process run after
// the hard fork to clean dust accounts
“0xA43EBd8939D8328F5858119a3fb65f65c864c6Dd”,
“0xE9c9068240D8450Da314f60804deBFc194B72309”,
“0x0e879ae28cdddeb31a405a9db354505a5560b0bd”,
// Smart contracts were involved in the DDos
“0x6a0a0fc761c612c340a0e98d33b37a75e5268472”,
“0x7c20218efc2e07c8fe2532ff860d4a5d8287cb31”,
“0x10fa9f37f646bb353945fe90d41a44e1c60745fb”,
“0x822f505e0174ef22d2a774cb80a855ffd27ae3bf”,
“0x59493d3fc7a8522253b8be0d168b8ad22ff85177”,
“0xba0577e1419237fd4b8c14a6f49984f6466b5996”,
“0x4943e4bd90d7ff8bafe1bd202e08907903ebdb66”,
“0x3898d7580aa5b8ad8a56fcd7f7af690e97112419”,
“0x8a2e29e7d66569952ecffae4ba8acb49c1035b6f”,
“0x9f58ef5d703973ba98dfa7a9bdecabecf13a0ec3”,
“0x8428ce12a1b6aaecfcf2ac5b22d21f3831949da3”,
“0xaa7c4ca548ffc77a42b309aaaea40a1bd477ac70”,
“0x2213d4738bfec14a2f98df5e428f48ebbde33e12”,

2

The code image below this claim has a mistake.

“0x7c1cf1f9809c527e5a6becaab56bc34fbe6f2023”,
“0xde21bc367afe7a3705a15255ff46a5ae91e8341c”,
“0x1fa0e1dfa88b371fcedf6225b3d8ad4e3bacef0e”,
“0xd3e32594cedbc102d739142aa70d21f4caea5618”,
“0xfb34db0651ab62d73a237fcf1aa1057ceb1f6229”,
“0x40525ac2fe3befe27a4e73757178d4accfef71da”,
“0xe25e422e3f9e9374a3d8a75451c790d48fb33218”,
“0xb09f8a62c6681b0c739dfde7221bfe8f2da3f128”,
“0x4b8f3b2e935a341929c0a4efe5110314f39dea73”,
“0x0c40cf6dd1b74c6edfc4729021a11f57b31e3c27”,
“0xf9110f7f0191317eb4bcd96e80d86946eb5426c5”,
“0x1dacf33da596a743be75933ce066f9c6e142a460”,
“0xb233cb2f0dce57a56bf732767f45ffc8650186c5”,
“0xb233cb2f0dce57a56bf732767f45ffc8650186c5”,
“0x25612c41773cb96167854ff72b1c2d7dc8973e2f”,
“0xd6a64d7e8c8a94fa5068ca33229d88436a743b14”,
“0x7fc03bd9e44c37bc2d111dc2154da781dbba7c24”,
“0x45faec35e32676568ad827aea17fb7431ef390bc”,
“0x29446e8d2f0ca2e7fd9e46665e80fc2cd55bf262”,
“0xab90c4455d32f1e579152f52377e3cbf9b3cc37b”,
“0x0c40cf6dd1b74c6edfc4729021a11f57b31e3c27”,
};
const uint32_t nSweepers =
sizeof(sweepers) / sizeof(Address);

Summary
The Fall 2016 DDoS attack left a lasting scar on the Ethereum chain —
thousands of transactions generating millions of traces across roughly
430,000 blocks. For anyone doing deep chain analysis, this region is a
mineﬁeld of slow queries and wasted time.
Our solution is simple: use the RPC’s eth_getTraceAt to binarysearch for a transaction's trace count without fetching the full traces. If
a transaction has more than a certain number of traces, we ﬂag it as
spam and let our calling code decide what to do — skip it, cache it, or
process it with full awareness of the cost. Combined with a list of
known attacker and cleanup addresses, we can now scan through the
entire chain history in a fraction of the time it used to take.
The broader lesson: When the data is ugly, don't ﬁght. Run away.


The Trace Data Problem
Wednesday, February 21, 2018

The Ethereum dataset is big (and growing). I want to
be able to scan through the entire thing. I want to be
able to do this on a laptop. This last fact makes it
impossible for me to create a separate, independent
copy of the data. If I want to do data analysis, I must
do it against the node’s data directly.
In this article, I am going to try to help you understand why the Fall 2016
DDos presents such a severe problem if one is trying to scan the node’s
data directly. By the end of the article, you will understand why the RPC
is so freaking slow.

Traces per Transaction
I scanned each of the ﬁrst 5,000,000 blocks. At each block, I scanned
all the transactions in that block. At each transaction, I counted the
number of traces generated by that transaction. The node software
delivers one trace for every transaction plus additional traces whenever
a transaction sends ether, calls into a smart contract, creates a new
smart contract, or self-destructs.
The table below shows the upper-right portion of our collected data.
Each column corresponds to a 50,000-block range, each row to a trace
count, and each cell to the number of transactions in that range
producing exactly that many traces.

For example, in the ﬁrst 50,000 block group, there were 1,871
transactions (orange) all of which created a single trace. Between
blocks 150,000 and 200,000, there were 108 transactions (blue) that
created exactly 11 traces. These 108 transactions created 1,980 traces
in total.
Next, we show the top-right portion of the data. You’ll notice the Total
and Percentage columns. As you can see, almost every transaction

ever cast against the chain created a single trace (itself). 97.29% of all
transactions were of this type. These transactions either sent a
message from one account to another, sent Ether from one account to
another, or transacted in a simple way with a smart contract. By the
time we get to ﬁve traces produced by a single transaction, 99.62%
percent of every transaction ever sent is included.
As we learned in the previous article, one group of transactions created
more (many more) than ﬁve traces. These are the 101,637 transactions
(0.07%) dDos transactions in the Fall of 2016. Each of these
transactions created 1,000s of traces. We will look at these
transactions next.
----Recalling the 2-D heat map from our previous article, the next chart
shows that same data in 3-D.

You can see the dominance of the one-trace transactions in red. But
you may also notice two small red bumps at the very rear of the chart.
These are the dDos attacks.

Let’s spin the view around and look at it from behind. In this view, you
may notice the coloration of the ‘Great Plains.’ One can see the
beginnings of the heat-map in the patchy colors in the plains.

This is the same data as the previous chart from the other side. The
dominance of one-trace transactions is apparent here as well. And you
may see that as the chain matures there are a growing number of
transactions with more than one trace.
What happens if we remove the one-trace transactions so we can see
things more clearly?

Each of the transactions represented here created at least two traces.
We are now starting to see the problem that I call the Trace Data
Problem.
The red spikes at the front of the chart between block groups 20 and 40
is The Dao. You can see The Dao crowd sale and the subsequent hack
clearly. The huge red spikes near the right-front side of the chart
(between block groups 80 and 100) are most likely Crypto Kitties.
Notice, though, the two sizable spikes at the rear of the chart. These
two spikes represent the dDos attack and then, after a pause, the
cleanup of the dDos attack. It doesn’t look so bad, but as we will see
these two collections of transactions are much worse than they look.
A transaction with a single trace generates, obviously, one trace. A
transaction with two traces generates two traces. For the dDos
transactions — each of which has thousands of traces — the number of
traces becomes very large indeed, and this is the crux of the problem.
Let’s look at a chart showing the number of traces per 50,000 block
group as opposed to the number of transactions per group.

We can see here why scanning the Ethereum blockchain — especially
when one is looking at traces — is so onerous. To get through just

101,000 transactions, one must scan more than 138 million traces. This
is why the RPC is so slow and why it takes so long to sync.
Let’s look at some more numbers:

This is the bottom-right corner of the data. In the ﬁrst 5,000,000 blocks,
there have been 305,605,032 traces created. Of those 138,665,661
(45.37%) were generated by the 101,637 transactions that were part of
the dDos and its solution. This is an average of more than 3,000 traces
per dDos-related transaction. The entire rest of the chain created nearly
the same number of traces (166,939,371) spread across 1,500 times
more transactions (149,758,364) — an average of slightly more than 1.1
trace per transaction.

Upshot
The dDos attack was bad. It created more than 20,000,000 empty
accounts. This caused a huge bloat of the state database which had to
be cleaned up. It was necessary to correct the problem by removing
those dead accounts, and the Spurious Dragon hard fork did that, albeit
by creating a lot of additional traces. The solution to the dDos problem
created nearly a third again as many traces as the attack itself. This is
one of the downsides of an immutable data structure such as a

blockchain. We’re doing everything we can to not let this stop us from
looking at the data.
In our next blog post, we will ignore both the dDos attacks and the
single-trace transactions and start looking into what else is happening
on the Ethereum blockchain.
In the chart below, we’ve removed both the single-trace transactions
and the DDos transactions. Do you see anything interesting? We do.
We’ll write about that in our next installment.


How Many ERC20 Tokens Do You Have
Friday, April 20, 2018

I was recently included in a discussion about why it’s so di4icult to get
ERC20 token balances (and other data) from the Ethereum blockchain.
I thought I’d take a crack at answering that question.
It comes down to four words: “Incoming”, “Internal”, “Transactions”,
and “Decentralization.” I’ll focus on each word individually.

Transaction
I’ll start with the easiest word ﬁrst: “Transaction.”
Go into a bagel store. Buy a dozen bagels and get a
receipt. Everything you need to know about your
transaction is on that receipt: date, amount, sender
(from) and recipient (to). This same information is
available from the Ethereum RPC for every transaction on the entire
chain.
From your perspective (the spender), the expenditure represents
outgoing money. From the perspective of the bagel shop, the
transaction is incoming money, and this leads to the next word on our
list.

Incoming
If you’re the sender of a transaction, you know that you sent it. You
reached in to your own wallet and took the out the money and handed

it to the store owner. (Or, on Ethereum, you opened Metamask and
signed a transaction—same thing.) Keeping track of outgoing spending
is easy. All you have to do is remember what you did.
But what about keeping track of incoming money?
I hear you say, “That’s easy?” The bagel shop owner uses a cash
register. It sounds simple enough, I admit. But now we move on to the
next word in my list of four words.

Internal
Every Ethereum transaction is ultimately initiated by an external
account (or EOA). An EOA lives outside of the blockchain. Transactions
are initiated either by a human being or by a piece of software being run
by a human being1. If you’ve used a blockchain, you’re familiar with this.
These types of transactions (let’s call them
“external”) are never initiated from inside
the blockchain. And this makes perfect
sense — a transaction requires a signature,
and there are no private keys inside the blockchain. If there were,
anyone could steal your money in two seconds. Distinguish “external”
from “internal” transactions.
“Internal” transactions are transactions that happen in response to an
“external” transaction interacting with a smart contract. Saying this
another way, only smart contracts initiate “internal” transactions.
Most transactions on Ethereum consist of an external account sending
money (or a message) to another external account.

1

2026: Or an AI!

Many other transactions (also external) are
sent from an external account to a smart
contract. When you think of the transactions
that are being hashed into a block, you are
thinking of external transactions.
An internal transaction, on the other hand, is
the result of a smart contract (which was
initiated by an external transaction) calling
into or sending money to other accounts.
As we saw earlier, keeping track of the “send” side of an internal
transaction is simple. The smart contract itself, as it generates the call,
can emit an event. This will alert both the sender and receiver that
something has happened (for example, a token transfer or token Mint).
This is how almost all ERC20 token contracts work. When someone
buys a token or transfers a token, in addition to recording that fact
inside the smart contract’s balance map, the smart contract emits an
event. I say “most” token contracts work this way because there is no
hard requirement in the ERC20 standard when tokens are minted.
Here’s a quote from the ERC20 spec:
A token contract which creates new tokens SHOULD trigger a Transfer
event with the from address set to 0x0 when tokens are created.
“SHOULD”, in EIP-speak, means that the contract doesn’t have to do
so. And here you should be able to see one of the reasons why it’s so
di4icult to accurately account for an arbitrary ERC20 token.
Sometimes, when an ERC 20 token contract Mints new tokens, it does
not generate an event. This was more likely back in the day with early
ERC20 token contracts, but it’s still true today.
Even when a smart contract does generate minting events, there is no
standard way of doing so. Some smart contracts generate a Transfer

event with the from parameter set to 0x0 as they SHOULD. Others have
their own Mint event. Still others generate no events at all.
The fact that there is no hard standard for mint events, makes
programmatically accounting for every received token impossible. This
becomes especially acute in the case of airdrops, many of which create
thousands (millions?) of tokens when ﬁrst created. Frequently, they do
not generate individual events for each minted token. It’s as if money
magically appeared in your wallet or the bagel shop’s cash drawer.
This is one example of an “Incoming Internal Transaction” and it is
the reason why 18-decimal place perfect accounting doesn’t work
on Ethereum.
Technically, an ERC20 token mint is not an “incoming internal
transaction” — there’s no actual send or call into the recipient — but
incoming internal transactions can have the same e4ect. If an event is
not generated by the calling contract — and in many cases, events are
not generated even when value is exchanged — the receiving account
can be granted ownership of something of value without being made
aware.
If the smart contracts do not generate events for value transfers, there
is no way (in a decentralized world) for the receiving address to know
about it without scanning
every transaction. Watching every
transaction, looking for transactions of interest, is di4icult. (See our
“The Trace Data Problem” article for context.)
If one wishes to be truly anal retentive (and for some reason, I do.
Thanks, Dad), one must replay every trace. To know that an unknown
smart contract granted your account ownership of a token, you have to
be constantly vigilant.
It’s ridiculous. Why is this not automated? Why is this not part of the
node software?

Before this article runs to an inﬁnite length, I’ll move on to the ﬁnal
word…

Decentralization
Getting full and accurate account-speciﬁc list of transactions (and
therefore token balances) from Ethereum is hard. It’s also slow. In fact,
it’s so slow that it’s unusable. One simply cannot replay the entire chain
every time one wishes to account for a di4erent address. There are
some solutions, but none of them good.
The two most obvious solutions are:
•

Make a duplicate copy of the blockchain data including all traces
and build a giant multi-terabyte database, or

•

Keep the data on the node, do not duplicate it, and ﬁgure out some
way to make scanning the chain palatable.

The second of these two solutions is what QuickBlocks does. It’s not as
fast as a true database, but it’s faster than a raw scan of the entire
chain, and it’s tremendously smaller than duplicating the data as any
god-fearing Web 2.0 database driven solution would do.
If one chooses the ﬁrst path, one will almost certainly be forced to use
some sort of gigantic cloud computer. The index alone will be many
gigabytes. And it only gets bigger over time. Worse, you must store
duplicate data (the data is already on the node), and furthermore, 99%
of the data you would be storing will never be used. You’re accounting
for speciﬁc accounts, not the entire chain. Most likely, your customers
are only interested in the data from their accounts.
As if that’s not bad enough, as soon as sharding happens the “duplicate
copy of the blockchain” solution (solution 1) will be shit out of luck. If
the data size is X prior to sharding, it will be 100 * X after sharding. The
only viable solution if one chooses this ﬁrst path will be to centralize the

data and try to sell it (either directly through an API or through
advertising and privacy invasion).
Web 2.0 is dead. Long live Web 2.0.
I only must present Etherscan as my ﬁrst exhibit to support this claim.
If, instead, one chooses the second path and tries to make accessing
the node-stored data faster, so that your solution remains
decentralized, you quickly realize that your code literally demands to be
open sourced—in fact, it demands to part of the node software (or
even, God forbid, the operating system). If you’re going to keep your
code closed source, why centralize and be able to monetize? That path
is far easier and more lucrative.

I’m out of Words
I’m hungry, and I’ve nearly run out of words, so I’ll end with a ﬁnal
thought. Developers are left with a conundrum:
Either decentralize and be forced to open source your code or
centralize it thereby eliminating the reason we all got into this thing to
being with.
This, I think, is why there’s currently no viable solution to the problem.
There is no appealing path forward.
I’ve now o4icially run out of words. If you’ve made it this far, I guess
you’re happy about that.


Mantras for Decentralized Open Data
Monday, May 7, 2018

I had a wonderful experience at EdCon in Toronto this week. Mostly
because of all the great new people I met and the many people I
reacquainted myself with. The Ethereum community is friggen’
awesome.
Here’s an example. In between
speakers, the person sitting in front
of me stood up, turned around,
looked down at his seat and, using
his pointing ﬁnger, counted the
number of seats between him and
the isle. He then counted the number of seats between himself and the
other isle. Finding that the right-most path was shorter, he existed the
auditorium in that direction. He didn’t care which way he went — he
cared only about being eFicient.
His friend, who was sitting next to him, stood up, turned around, and
made the same count. Because the arm rest between their two seats
was the exact middle of the row, his friend came to a diFerent
conclusion about which path was shorter. He went the other direction.
This is just one example of the many interesting and amazing people I
met. Others included someone who claimed to be “American as fuck,”
a former improv comedy artist, and a very nice gentleman who grew up
in a very large castle. We are all diFerent.

Data
As is true of most of my conversations these days, I spent the weekend
talking about data. The conversations were about why it’s currently so
diFicult (in a decentralized way) to get the data from the Ethereum
blockchain, and what we might be able to do about it. There’s several
diFerent ideas about how we might proceed1 2, all of which I think are
equally valid. In other words, the ultimate solution will be many
solutions.
I’ve been thinking about some of the guiding design principals that the
community might want to keep at the top of their mind as they develop
better ways to make the blockchain’s data accessible and more useful.

Mantras for Decentralized Blockchain Data
1. All the data, for all the people, all the time.
2. None of the data if a user doesn’t need it.
3. Use only consented-to, provably true data…
4. If you can’t do (3), use only data derived from consented-to,
provably true data…
5. If you can’t do (4), force the data provider to prove their data.
6. The server is local. There is no server. Embrace the paradigm shift.
7. Break out of the browser but stay behind the node. There’s a big
wide desktop out there.
8. It’s easier to scale from decentralized to centralized than from
centralized to decentralized.

1
2

https://medium.com/@lrettig/60ac7c404169
https://www.youtube.com/watch?v=CXtkv-HzEso

9. Most users care only about their own data. Some users, care about
everything. Every user cares their own money.
10. The Ethereum node can be improved and should be.
11. Immutable data has interesting properties, especially when tiny
parts of it are stored by many people on content-addressable ﬁle
systems like IPFS.
12. Hashes are pointers to data. So are ﬁlenames. Pointers are better.

…Screech….
I suFer from triskaidekaphobia, not because I
believe that 19 − 6 is bad, but because I believe in
quantum physics. You cannot convince me that an
irrational fear of the number following 12 is
nonsense. You may be able to prove to me that within
any given 𝜉 a fear of 10 + 3 is silly, but the net beneﬁt from engaging
with

!"
!

is strictly less than any 𝜉 you may provide. Therefore, I choose to

avoid it.
We are all diFerent.


Building an Ethereum Account Scraper
Friday, July 6, 2018

When I ﬁrst heard of Ethereum, I thought to myself “Excellent! I no
longer need to keep track of my spending. Everything will be
automated.”
Can you say “misconception”?
Not only is keeping track of my Ethereum spending way more diEicult
than it should be, it’s actually so diEicult that I stopped spending Ether
about two years ago. (The price of Ether had a lot to do with it as well.)
“Why,” I’ve been asking myself ever since, “is it so hard to keep track of
my
accounts?
Why
isn’t
there
an
RPC
command
eth_getAccountTransactions?”
I’ve been working on this shortcoming ever since. I’m glad to say that,
as the latest version of QuickBlocks1 gets us a long way towards that
goal.
In this article, I will show you how to build an Ethereum account
scraper. An account scraper is diEerent than a blockchain scraper.
Think of a blockchain scraper as creating a full duplicate copy of the
entire blockchain in a relational database and delivering it from an AWS
instance via an API. We’re not building that. The expense of doing things
that way leads to centralization such as Etherscan.

1

2026: Now called TrueBlocks.

Instead, we “scrape” only the parts of the chain that are of interest to
ourselves. We’re going to scrape only the transactions related to our
own accounts.

Installing QuickBlocks
Instructions for installing and building QuickBlocks are on our website.
I won’t go into details. If you have trouble, please contact us2. Important
note — QuickBlocks requires a running Ethereum node, and that node
must be a tracing node. Our software works with Infura, but Infura
doesn’t support tracing, so you won’t get complete results.
Assuming you’re running a local node…

Building an Account Scraper
QuickBlocks is a C++ code library that makes interacting with Ethereum
easier. To get started, we need to write C++’s main function. There are a
couple of things going on here.
First, include the header ﬁle. We
do this at line 1.
Next, forward declare a function
that will be visited each time an
address is encountered in a
block. We make this forward
declaration at line 4. The function
must have this exact signature.
At line 8, we use the program’s ﬁrst command line argument as the
number of the block we wish to visit. (We use block 1,001,001 in the
examples.) At lines 9 and 10, we query the Ethereum node for the block.

2

2026: Or use an AI.

When the call to getBlock returns, block has been populated with
the block data from the node. The next line, line 11, interrogates the
block including every transaction, every receipt, every log, and even
every trace. As it visits the various bits of data, it extracts from them all
the addresses it ﬁnds and calls the callback function.
With each call the system sends the following data3:

The ﬁrst four data ﬁelds should be obvious. They represent where in the
block the address was found. The ﬁfth ﬁeld, reason, details exactly
where the address was found. Reason will take on various values which
are documented in the code.
At line 19, we send the data to the screen which produces these results:

This is a list of every address in block 1,001,001, its transaction
position, it trace position (if applicable) and where the address was
3

2026: This structure has since been renamed an Appearance.

found. (There are 15 diEerent places or reasons where an address
might appear4.)

Going Further
Next, we will make two changes to the code. First, we’re going to
produce only a unique list of addresses (as opposed to every address
encountered including dups as it currently does). Second, we will only
print out addresses we’re interested in. We make the following changes
to the code:

We changed two things in the above code. First, we changed line 11 to
use the function forEveryUniqueAddress instead of forEveryAddress.
This removes duplicates and reports only on the ﬁrst appearance of an
address within the block. The second change we made at lines 20 and
21. Here, we ﬁlter on particular addresses. This produces the following
output.

2026: This feature is one of the most important innovations TrueBlocks made.
There should be a function in the RPC called eth_addressesInBlock.

4

Going Further(er)

This next version of the code takes two command line arguments:
startBlock and endBlock and visits each block in the range and
then visits each address in each of those blocks.
And now, we have a full-ﬂedged account scraper. Go ahead and build
your own blockchain scraper if you wish.
Of course, you may enhance the above code to your heart’s content.
Add more command line options. Add input checking. Add features
that allow you to specify a range of blocks on the command line. Add a
feature that allows you to specify the address(es) you want to report on.
Keep track of the last block visited, go to sleep every loop, and keep up
with the tip of the chain5. Account for the transactions as you encounter
them6. Export to the data to your accounting software!

5
6

2026: We call this feature monitoring, and it’s part of the Khedra tool.
2026: We call this feature reconciliation.

Check out our command line tool called getBlock. This tool does
much of what we describe above. getBlock is fully open source, and
you can enter a command like this:
getBlock --addrs 5000000-5250000 | \
grep 0x2a65aca4d5…135398226

which will scan the 250,000 blocks between 5,000,000 and 5,250,000
and ﬁlter out only transactions involving the given address. With such a
list, you can do anything…including…wait for it…. including… automate
your accounting!
Go forth and multiply (your spending)!


A Short Take on Decentralization
Monday, August 13, 2018

This is a very short take on something I’ve been thinking about for quite
a while. “Why do I give a shit about decentralization? Should I
compromise?” Here’s why I come down emphatically on the side of
“Yes, I give a shit, and no I should never compromise!”.
Blockchain-like technologies (that is,
decentralizing, trustless technologies) bring to
the world, for the ﬁrst time in human history, a
way to help us solve the prisoner’s dilemma.
With this technology, we can now remove the need to trust the other
party. The “need to trust” is why, in the prisoner’s dilemma, everyone is
forced to defect. The fact that everyone is aware that they must trust
the other party, but is also aware that they themselves can defect, leads
them to conclude that the only rational choice is to defect.
If the system can remove the need to trust, humans will (eventually)
learn that the better choice in any given situation is cooperative
action — which, in an odd way, has always been recognized as the
better solution. Both prisoners know that they are better oM if they
cooperate, but they are unable to do it because they must trust the
other person.
So — why do I give a shit about decentralization? Why will I never
compromise? Two reasons:

•

We can make the world a better place — because we can now trust
each other — because the system has removed the need for us to
trust each other.

•

If we build a hybrid system in which even minor parts of the system
require trust, we will have built something worse than we already
have.

If humans think we’re building a system that requires less trust, they will
tend to reveal increasingly intimate details about themselves (requiring
less trust equates to the system being more trustworthy). The trouble is
that a system that accumulates deeply intimate information about us
becomes increasingly enticing to people who want to take advantage.
We will have built a system that is easier for them to do that.

My Point
Here’s the point I’m making: Every time we compromise on the issue of
decentralization and allow the need to trust to creep back in to our
blockchain systems, we ruin the possibility of realizing the collective
beneﬁt and we put ourselves in deeper danger of being taken advantage
of.
Stop compromising. Write 100% decentralized code or don’t write
anything at all. If something doesn’t work in a fully decentralized way,
stop doing what you’re doing and ﬁgure out why.
Like the old saying, the only good centralization is dead centralization.


Is the Difficulty Bomb Exploding?
Monday, January 21, 2019

This article shares with you (through a series of charts) what happens
when you release a world-class data scientist such as Ed Mazurek on
fresh-baked Ethereum di>iculty data. You get an raging river of real
information.
With little to no explanation, I ﬁrst copy and paste the “R” code Ed wrote
and then show the chart created by the data. Ask Ed what the code
means. I have no idea.
We used chifra (our command line tool) to extract this commaseparated block di>iculty data which looks like this:
blockNum timestamp
difficulty
-----------------------------------1
1438269988
171714805762
2
1438270017
171630960643
3
1438270048
171547156464
4
1438270077
171463393215

Yes. It’s as simple as that. One record for each of the 7,105,056 blocks
that had been produced at the time of this writing. Here’s some
preliminary R language setup code that reads the data and cleans it up.
We’re not going to explain the code for two reasons:
1. It’s kind of fun (and relatively easy) to ﬁgure out what it’s doing.
2. We don’t really know what it’s doing. Ed does.

require(tidyverse)
require(scales)
require(dplyr)
require(magrittr)
homestead.block <- 1150000
byzantium.block <- 4370000
bin_size <- 200
period_size <- 100000
sample_size <- 50000
difficulty <- read_csv(‘difficulty.csv’) %>%
filter(block.number > homestead.block) %>%
mutate(block.bin = floor(block.number /
bin_size) * bin_size) %>%
mutate(fake.block = ifelse(block.number >=
byzantium.block, block.number — 3000000,
block.number) + 1) %>%
mutate(period =
floor(fake.block / period_size)) %>%
mutate(bomb = 2 ^ period) %>%
mutate(parent.difficulty = lag(difficulty)) %>%
mutate(parent.ts = lag(timestamp)) %>%
mutate(diff.delta = difficulty –
parent.difficulty) %>%
mutate(ts.delta = timestamp — parent.ts) %>%
mutate(diff.sensitivity = diff.delta /
difficulty) %>%
mutate(ts.sensitivity = ts.delta / timestamp)
current.block <- difficulty$block.number %>%
tail(1)
current.bomb <- max(difficulty$bomb)

This code builds the difficulty data frame.

The Code
Here’s the code for our ﬁrst chart:

difficulty %>%
sample_n(sample_size) %>%
group_by(block.bin) %>%
ggplot(aes(x=block.number)) +
geom_line(aes(y=diff.delta,
color=’diff.delta’)) +
geom_line(aes(y=bomb, color=’bomb’))

And here’s the ﬁrst chart showing the delta in each block’s di>iculty
through the ﬁrst 7.1 million blocks. It also shows, as a red line, the
di>iculty bomb. You can see it’s creeping up again.

The next chart, built with this code,
difficulty %>%
sample_n(sample_size) %>%
group_by(block.bin) %>%
ggplot(aes(x=block.number)) +
geom_line(aes(y=diff.sensitivity,
color=’diff.sensitivity’))

shows the ‘responsiveness’ of the di>iculty calc to its current situation.
We calculate sensitivity by dividing diff.delta by block.
difficulty. I’m not sure, but I think the “jaggys” come from the way
the di>iculty value is calculated. It snaps to the next lowest multiple of
10 seconds. So 10, 11, 12, … 19 all snap to 10, but 20, 21, 22, … snap to
20.

The next chart, generated from this “R” code,
difficulty %>%
group_by(block.bin) %>%
summarize(sum.diff.delta =
sum(diff.delta), na.rm=T) %>%
ggplot(aes(x=block.bin, y=sum.diff.delta)) +
geom_line()

shows the accumulated sum of the diff.delta values. You can
clearly see the battle waged by the pre-Byzantium di>iculty bomb. Up,
down, up, down. The fact that the di>iculty hovers around a target is
exactly what the di>iculty calc is supposed to do. It keeps the timing of
the blocks consistent.

Lane Rettig, a great friend of ours, wondered if, because of the
increased hash rate of the current chain, the e>ect of the di>iculty
bomb might be obscured. I’m not sure (am I ever sure of anything?), but
the wider spread of the di>iculty on blocks since Byzantium may
indicate that his intuition was correct. Is the di>iculty bomb hiding in
the black static?
In the following charts, we show some other shit made with some more
“R” code. In this section, we calculate either the deltas of the means of
the timestamps of the blocks or the means of the deltas of the
timestamps of the blocks, and we lay whichever one this is against the
previous charts (you’ll have to ask Ed if you want to understand it).
difficulty %>%
group_by(block.bin) %>%
summarize(sum.diff.delta =
sum(diff.delta, na.rm=T),
mean.ts.delta =
mean(ts.delta, na.rm=T)) %>%
gather(key = vars, value = val, -block.bin) %>%
ggplot(aes(x=block.bin, y = val)) +
geom_line() +
facet_wrap(facets = ‘vars’,
scales = ‘free’,
ncol = 1)

In the next set of charts, we’re showing a whole bunch of crazy shit with
cool looking black jaggies (hang in there, there’s some interesting stu>
below).

The above chart was made with this nearly unintelligible code:
difficulty %>%
group_by(block.bin) %>%
summarize(sum.difficulty = sum(difficulty),
sum.diff.delta =
sum(diff.delta, na.rm=T),
mean.ts.delta =
mean(ts.delta, na.rm=T)) %>%

mutate(percent.delta = sum.diff.delta /
sum.difficulty) %>%
gather(key=vars, value=val, -block.bin) %>%
ggplot(aes(x=block.bin, y = val)) +
geom_line() +
facet_wrap(facets=‘vars’,
scales=‘free’, ncol = 2)

Do you remember that simple three column data we showed you
above? Ed extracted the above charts from that very simple data. He’s
a genius, I tell you!

Finally, Something Interesting
Ed and I were inspired to work on this problem at the Status Hackathon
in Prague back in October 2018 by Lane Rettig. Lane was concerned
that the hash rate of the current chain, which was so much higher now
than it was during the pre-Byzantium bomb, would obscure the e>ect
of the bomb for longer, and then, once the bomb started showing (once
it exploded), it would take people by surprise.
We wanted to see if his intuition was correct.
The bomb explodes in an exponential growth curve doubling every
100,000 blocks. This has the e>ect of creating a step-function curve, as
the chain responds to dampen the rise in di>iculty, as we discussed in
a previous article.
In the next set of charts, we’ve bucketed the data into buckets of
100,000 blocks and called these buckets periods. (The black line
below shows the current period.) This bucketing is evident in the
following chart where we also broke the data into groupings: preByzantium and post-Byzantium. This allows us to line up and compare
the previous and impending bombs.
What do we see when we look at the data this way? First, we see that
the post-Byzantium bomb is beginning to make itself known.

difficulty %>%
mutate(era = ifelse(block.number <=
byzantium.block,
‘before byzantium’,
‘post byzantium’)) %>%
sample_n(sample_size) %>%
ggplot(aes(y = diff.sensitivity,
x = period, color=block.number)) +
scale_colour_gradient2(low = “red”,
mid = “green”,
high = “blue”,
midpoint = byzantium.block,
space = “Lab”,
na.value = “grey50”,
guide = “colourbar”) +
geom_point(size = point_size) +
facet_wrap(facets = ‘era’, nrow = 2) +
geom_vline(xintercept = 41)

The above code gives us this chart:

And here, I think, we’ve ﬁnally arrived at something interesting.
Although, we’re not quite done.
Next, we lay the post-Byzantium data and the pre-Byzantium data on
top of each other, so we can compare them directly.
What we see now, if we look closely, is that the little purple dots (current
bomblets) are lower down than the little green dots (previous
bomblets).
Lane was right!

During the upcoming di>iculty bomb, the e>ect of the bomb will be
obscured and less apparent than before. This may make its e>ect less
obvious until it’s too late. How soon will that be? When is the bomb
coming? We’re already at period 41. The last bomb (see our previous
di>iculty bomb related article) started clearly showing down the chain
by now. The last bomb was already starting to explode by period 41. I

think that is true this time as well, but it’s hidden behind the heavier
hash rate.

Conclusion
It’s time to fork.


Counting Shit on Ethereum
Fun with Blocks, Transactions, and Counters
Monday, June 10, 2019

I like to count shit. I’ll count anything. Cracks in the sidewalk. Pennies
in a cup. Syllables in a poem.
I also like to make lists. In fact, at one point in my life, combining my two
passions, I was maintaining a list of over 350 rules related to how to pick
up pennies without incurring bad luck.
Yes. I suCer (happily) from OCD.
Recently, Pedro Gomes, a friend whose project Wallet Connect is doing
well, sent a tweet asking if anyone had a list of the most frequently
called function signatures on the Ethereum blockchain.
I thought to myself. Oh cool. An opportunity to make a list. I like lists.
Plus, there’s counting involved. I can do this.
To me, “I can do this”, translates into “I can write code that does this”,
so here’s some pseudo code for answering Pedro’s question:
map<signature, count> counts
for (every block in blockchain)
for (every tx in block)
if (tx.input is at least 4 bytes long)
counts[tx.signature]++
sort the map reverse by count
print the first twenty entries in the map

In other words, count some shit.
I work on a project called TrueBlocks (notice the name change), so I’ll
used the TrueBlocks tools to carry out the above simple process. Here’s
what you would need to do to duplicate my eCort:
•

Download, install, and run an Ethereum node1

•

Download, install, and build TrueBlocks

For more information on running a node, you’ll have to look online (or
checkout DAppNode, which makes running Ethereum nodes easy). We
assume you’ve already downloaded, installed, and built TrueBlocks
from our GitHub repo.

Building the Sample Code
Before we write the actual counting code, we want to make sure you
can build the sample project. Instructions for that are available on the
repo, but the short version is:
git clone \
https://github.com/TrueBlocks/trueblocks-core
cd trueblocks-core
mkdir build
cd build
cmake ../src
make

Now we can start writing the counting program. We start by making a
copy of one of the sample programs and making sure the copy builds:
cd ../src/examples
cp -pR simple count_shit
cd count_shit
make
./bin/simple
1

Wait for a long time while the node syncs.

You should get some data for block 3,500,000 (assuming your node is
synced that far). Note: if you don’t want to run your own node, you can
edit a ﬁle called ~/…/trueBlocks.toml and set the rpcProvider option to
a remote node such as Infura. Be warned, though, if you do you will
probably get locked out due to hitting their servers too hard — this is one
of the beneﬁts of running your own node — no rate limiting.

A Quick Review of the Sample Code
Open the ﬁle./samples/count_shit/simple.cpp with your editor. It will
look something like this:

A few things to notice:
First, you need to #include "etherlib.h". This allows you to build
programs with TrueBlocks. Second, initialize the library with
etherlib_init().
The real code starts at line 14. Here we create a storage location for the
block data with CBlock block . The next line calls into the library
(which in turn calls into the node) to get the block data for block
3,500,000. Line 18 sends the block data to the screen. Very simple.

Writing the Actual Code
Now we want to write the code to count function signatures. With little
fanfare, here’s the code with short descriptions for each section.
The ﬁrst section includes the etherlib.h header ﬁle and deﬁnes a few
types that we will use to store the function signatures and then sort and
report on what we ﬁnd. Put this at the top of your ﬁle.

The next section is the main routine. Put this at the very bottom of your
ﬁle. We will be inserting two functions between the declarations at the
top of the ﬁle and the main function. In this way, we won’t have to
forward declare the functions we need.

The main function, at line 35, declares a signature map to store the
signatures, and line 36 calls the function forEveryBlock with the
function visitBlock and a pointer to that map. forEveryBlock
visits every block and calls the passed-in function on each block
between the ﬁrst block with a transaction (frstTransactionBlock)

and the last available block on the chain. (We’ll look at visitBlock in
a second.) visitBlock does the counting. Starting at line 38, we
simply sort the gathered counts and present the top twenty results.
The remaining functions do the real work. The ﬁrst, visitBlock gets
called…

…simply passing on the data pointer to every transaction in the block.
The ﬁnal function, and the one that does all the work, is called
visitTransaction.

The ﬁrst thing this function does is cast the data pointer back to a
signature map. Some of the younger programmers who are reading this
might look askance at this sort of type-safety violation, but at least one
of our engineers (alright — it’s me) has been programming since 1981,
so this is how we roll. It may be slightly unsafe, but it’s fast as shit.
The visitTransaction function checks if the input data of the
transaction is at least four bytes long (eight characters — four nibbles)
and then increments the counter at that location in the map. The

remainder of the code (from line 18–21) simply reports our progress to
the screen.
And that’s it. You’ve counted the frequency of function calls in the
Ethereum database. One word of caution: this is slow, slow, slow, slow,
slow. The RPC commands are not exactly “speed burners”.
One of the things TrueBlocks does is cache the blocks (if you want it to),
so the second time you would run this code, it would be very much
faster. That would come at the expense of hard drive space.
One other note: this only counts the fourbytes used…Pedro wanted the
names of the function calls. We leave as an exercise to the reader using
our getAbis tool to do that.

Conclusion
This is a purposefully simple example of the type of things you can do
with TrueBlocks. Spend some time reading the code. Download and
build the examples. We should all be learning how to access the
Ethereum blockchain data on our own. I think having access to the
chain data for oneself is an undervalued aspect of the blockchain.


Mother May I
A New Way to Think About Data APIs
Monday, June 24, 2019

When I was a child, we used to play a game called “Mother May I”.
In that game, the person who was it, called “Mother”, would stand on
one side of the room and the players would stand on the other. The
players would beg for permission to do things, “Mother, may I take one
giant step forward?” or “Mother, may I take two scissor steps?”
Mother would say either “Yes” or “No.” Players pleaded with Mother in
random order. Mother had to say “Yes” more often than “No.” The ﬁrst
player to reach Mother became Mother, and a new game would start.
Not exactly the most exciting game in the world, but we were kids and
we had no social media to capture our boundless attention.
In this article, we will play an imaginary game of Mother May I between
a player (you, the developer) and chifra (the tentative name for our
command line tool). Hopefully, this will help you understand our
thinking related to APIs and what we think may have to happen as we
all learn how to interact with shared global immutable (and truly
decentralized) data.
One thing we think is clear: old-fashioned web 2.0 APIs such as Alethio
and AmberData should not be the way we build out this new
infrastructure. If we do, we will not end up with what we want. We will
end up exactly where we already are—captured by huge, mega-data-

gobbling, mega-corporations who watch everything we do and use that
information to manipulate us.
We present the game in a series of pleadings and responses followed
(in blue) by some commentary. Game on…

After this ﬁrst interaction, the user does not have the actual data heshe
wants, but heshe has the location of that data.

The API does not deliver the actual data. The API delivers the location
of the data on an immutable, content-addressable data store (IPFS).
The end user retrieves the data they want from that location thereby
distributing it across the world. The user can respond to this in two
ways:

Scenario A

Even though we’ve educated the user that they should stop using data
APIs arbitrarily since they will capture them, the user insists heshe
wants to use one. Who are we to argue after that?

Scenario B
Yes. We’re starting to think people must be this radical when it comes
to data APIs. The costs we pay are, we think, too high for the beneﬁt.
The last thing we want to do — now that we have a new paradigm
brewing—is end up where we already are. Concern yourself with
capturability, censorability, data validity, and availability.

Conclusion
You’ve reached Mother. You know as much as she does. You need to
make it to the other side of the room, or the paradigm shift, as you wish.
Will you?


Indexing Addresses on Ethereum
Tuesday, July 9, 2019

Given a blockchain, our goal is to create a list of every appearance of
every address in any block on that chain. We deﬁne an “appearance” as
either the use of the address in one of the common ﬁelds of a
transaction (such as to, from, or contractAddress) or its use as
data in one of the data ﬁelds in a transaction. We do not check numeric
ﬁelds such as value or gasUsed.
An appearance consists of

Given a list of appearances, we build an index sorted by address. This
makes subsequent queries for all transactions on a given address easy
and fast.
With a list of transactions, one may then go back and query the
blockchain’s RPC for the details needed to do a full audit-quality, 100%
decentralized,
permissionless
accounting,
including “data
accounting,” which means exposing the state changes and current
values for any ﬁeld in any smart contract.
One thing to note. Every block is independent of every other block when
it comes to extracting addresses in that block. Because of this, we are
able to parallelize the work. We recently did that using GoLang in the
latest version of our code called Blaze.

Because these ideas are diRicult to explain, especially to someone who
is not familiar with the data structures used by Ethereum (as we
suspect is the case for most of our readers), we try to explain what
we’ve done through the following series of examples.

Each of the above examples is relatively easy. The addresses “appear”
in obvious places in the transaction such as from, to, or as the event
sender data in a log. There are more diRicult cases. For example, a
simple token transfer or a multi-sig wallet interaction.

Because our process uses traces, we must start the node with
tracing on. This is made much easier with a product such as
dAppNode which will one day run our software.

The above examples, don’t explain everything. There are many, many
edge cases and special cases. For example, if a smart contract reverts
during its construction, the contract’s address is reported in the
transaction receipt’s contractAddress ﬁeld, but it does not appear
anywhere in the transaction’s traces. In general, one can ﬁnd every
address by searching Parity’s traces and/or the event logs, but not in
this case. There are other odd examples.

Caveats
There are a few unfortunate facts about the way the Ethereum
blockchain stores its data. The fact that we must remove baddresses
and precompiles, and the fact that we need to create fake transactions
to fully account for every address is fragile but necessary. It’s obvious
to us that the designers of the Ethereum blockchain did not think
through every edge case as relates to accounting for the chain —
especially if one wishes to do this in a fully decentralized way.
We suggest the designers of the next generation of blockchain clients
pay more attention to the needs of those who wish to account for the
system.
A careful reader will note that we only extract appearances from the
blockchain. We do not extract any other data. This is an engineering
choice made because it was important that our system continues to
run on commercial-grade hardware — that is, laptops. TrueBlocks
works on a laptop! To us, this means “fully decentralized.”
We gave up certain things in our design, such as broad statistical
knowledge across the whole chain, but we wanted to build
permissionless accounting all the way out at the edge. We’ve
accomplished that.
As we extract the appearances, we periodically consolidate the blockby-block data into a more easily searchable index. Unfortunately, this
is not as simple as putting the appearances into an old-fashioned
database. We could do that, but (as you might expect if you know us), it
destroys permissionlessness. Things become more complicated the
more one wishes to decentralize things. We want to make the index
we’ve created easily sharable, uncapturable, undeniable, veriﬁable,
and fully decentralized. We’ll write about that in our next article.

Conclusion
We fell in love with the idea of permissionless accounting about four
years ago when we ﬁrst heard Joe Lubin, from Consensys, discuss the
idea of triple-entry accounting and radical transparency. We’ve been
working on that idea ever since, and recently — not the least due to our
Ethereum Foundation grants — we’ve been making great progress.
Here’s a quick preview of an application we’re building on the index. We
didn’t ask if it was okay!


A Time Ordered Index of Time Ordered
Immutable Data
Adventures in Indexing the Ethereum Blockchain
Wednesday, July 17, 2019

Did you ever notice that the only way to get the history of an Ethereum
account is to visit a fully centralized, database-driven, old-fashioned
web-2.0 website?
Every time I use one of those sites, (and I use them all the time), I think
to myself:
They’re watching me. They’ve attached my IP address to my address
and in the future, they will wildly invade my privacy…but I need them…
Recently we demoed a fully decentralized blockchain explorer built on
TrueBlocks (see our YouTube channel). At the core of our explorer is an
index of Ethereum addresses. This article discusses how we built that
index, the diKiculties we ran into, and why it’s way more complicated
than you may think to share — especially if you want to avoid becoming
an old-fashioned, outdated, web site destined privacy nightmare.

The Trouble with Time Ordered Data
As everyone knows, blockchains are not databases. Blockchains are
time-ordered logs of transactions (also called ledgers). Each
transaction is laid down on the chain as it happens, and it’s stored on
disc (at least conceptually) in that same order. The time-ordered nature

of the data makes it possible to represent each block (and the history
of all blocks prior to that block) with short, immutable hashes. Timeordered logs and immutable, hash-denominated data go together like
math and poetry.

The rows in the above image are colored by address. You can see the
addresses are interleaved and not easy to ﬁnd without scanning.
Using the Ethereum node’s RPC interface, TrueBlocks retrieves each
block as it appears, requests every transaction in that block, and then
requests every receipt, every log, and every trace in those transactions.
We parse all this data extracting anything that may appear to be an
address — see our other articles for details about this process — and

given this collection of addresses, sort it by transaction_index
within block_number within address and then write that sorted list
to disk.
If we went no further than this — simply creating sorted lists of address
appearances per block — that would produce a signiﬁcant speedup
compared to scanning the time-ordered log for an address. Of course,
we want to do better than this. We can, for example, produce larger,
consolidated, lists which even that much faster than single block
indexes.

The addresses in this list are sorted by address. Finding the address is
a simple binary search which is very, very fast.

Our scraper goes one step further. It keeps track of the total number of
appearances it’s seen, and, as soon as there are a certain number of
records (currently 500,000), the scraper combines the records, sorts
them by address, and stores the sorted index in a single, much larger,
ﬁle.
This greatly increases the speed of up subsequent searches for an
address — of course it does — that’s the nature of a database index.
Over the ﬁrst 8,120,000 blocks, our process has created around 3,000
such sorted lists (or “chunks” as we call them). Each chunk contains
approximately 500,000 records. Each record is <address,
block_number, tx_id>.
You may ask why we don’t simply store the index in an old-fashioned
web 2.0 database and distribute the results of querying that index using
an old-fashioned web 2.0 API? I’ll try to explain. Hold on to your hat.

Why Not Just Build a Web 2.0 API?
Decentralization.
That’s it. That’s why.

What’s the Problem with Web 2.0?
Why does web 2.0 suck? Let me count the ways: privacy invasion,
unveriﬁed data, fragility, single point of failure, data/user capturability,
user lock-in, privacy invasion, the rich-get-richer problem, the missing
long-tail problem. The list is very long. But the worst problem — and one
that is quite diKicult to explain — is that web 2.0 does not scale the way
we need it to if we want to build a decentralized world.
In fact, current, old-fashioned, web 2.0 blockchain explorers lead us
further away from a truly decentralized world as they scale.

The architecture of web 2.0 blockchain explorers is to ﬁrst extract all
the data from its source (the chain), put that data into a web-scale
database, index the data ten ways to Sunday, and only then deliver the
data to the user. This is directly in opposition to decentralizing the data.
The provider of such a service is eKectively saying to the user:
Don’t worry your pretty little head about getting the data yourself, we
will get it for you. Be our guest. Focus on building your application. We
promise it will ﬁne. You can trust us.
But this is exactly the opposite of the direction we want to be headed.
The size of the extracted data this model produces continually grows
(probably exponentially), therefore, the cost of delivering that data will
grow. This, obviously, will force the data providers to monetize their
users. And, in my opinion, this will lead to the exact same list of
problems I mentioned above, particularly user lock-in and privacy
invasion.
We can do better.

Recognize the Importance of Immutable Data
The blockchain’s data is immutable. Not only do we need to resign
ourselves to this fact, but we need to embrace it. There are no way
around it. Immutable data does not like to be indexed. Every time one
inserts a new record into a supposedly immutable list, one changes the
hash that’s generated when that data is stored in an immutable store
(IPFS, for example). So, it seems that you can’t have both immutable
data and a sorted index.
A simple solution to this problem is to stop adding things to the list at a
certain point. (As TrueBlocks does by creating time-ordered chunks of
address-ordered indices). This is a compromise between immutable
data and a searchable, indexed database. It turns out this simple trick

is enough to support a whole suite of applications including a fully
decentralized blockchain explorer:

What is TrueBlocks Delivering?
We’ve discussed this in two recent articles, so I won’t go too deeply into
it, but fundamentally TrueBlocks is building a system that creates a
series of partial indexes of address appearances across the
blockchain. These partial indexes, or chunks as we call them, are
stored in separate ﬁles partitioned each time the number of records
overtops 500,000 appearances.
We chose this number of records arbitrarily but chose it because we
wanted to balance the size of the chunks on disk (around 8MB each),
the time it takes to produce new chunks (about twice a day), and the
number of ﬁles produced (around 3,000 by block 8,120,000).
The time it takes to produce new chunks bears a bit of explanation. If
our index is far behind the tip of the chain, we can process many blocks
at the same time in parallel. Once we catch up to the tip, things change.
When we are catching up to the chain, we move as quickly as possible,
and in fact, we easily scrape the entire chain and catch up in about a

day and a half1. It takes about 45 seconds to produce each new chunk
of the index when we’re playing catch up2.
Once we’re caught up, the processing changes. Now, the process
spends nearly all its time sleeping. At current usage rates, the chain
produces about 450 unique address appearances per block.
Accumulating 500,000 records takes about 1,200 blocks. Blocks
appear every 14 seconds → 4.285 blocks / minute → 280 minutes to
produce 1,200 blocks → about ﬁve hours to produce a new chunk.
Chunks are immutable, never changing, able to be published to IPFS
and shared with the entire community — forever!

Seriously, what is TrueBlocks Delivering?
We are going to deliver an API that will run either locally against your
own locally running node, remotely under docker against something
like a dAppNode, or even remotely as an old-fashioned web 2.0
database-driven website.
In all three cases, the API will not deliver the JSON data that the user
wants when they say,
Mother May I have all transactions on account 0x1234….
Instead, our API will deliver the hashes to the chunks of the index that
the user needs to get the list of appearances needed to query directly
to the RPC.

This used to take more than three weeks, but recently we re-wrote the scraper
in GoLang (Blaze) and now take advantage of parallelism. Thanks, Nate
Rush!
2
3,000 chunks at 45 seconds each → 135,000 secs → 2,250 mins → 37.5 hours
→ 1.56 days.
1

User to TrueBlocks: chifra list <address>.
TrueBlocks to user:
[
{

]

“hash”: “QmXREJnqJ…”,
“range”: “6517955–6519510”
},
{
“hash”: “QmQMBTt…”,
“range”: “8102894–8104450”
}

The API returns IPFS hashes to the index chunks the user needs to
search. In the great majority of cases, this is a very short list because
regular users interact only periodically and in short bursts. A regular
user, querying for his/her own accounts, may get ﬁve or six hashes.
For heavy users of the chain such as popular dApps or exchanges, the
list will be much larger (but no larger than 3,000 entries). This is exactly
how we want it to be because regular users should not have to shoulder
the burden for heavy users.
User to IPFS:
for each hash {
ipfs get ${hash} -o ${range}.bin
pin ${hash}
}

The user brings the chunks locally to their own machine. It’s impossible
to capture the user from here on. They have all the data they need to
ﬁnd the history of that account — and — that data will never change. If
they themselves keep it, they cannot be captured. The pin ${hash}
part of the above line accomplishes this. Pin means keep it on that
machine forever.

By default, we want users to copy the index chunks locally (this will also
increase the speed of future searching) and then pin their own chunks.
This makes subsequent queries to IPFS by other players for that same
chunk faster since more increasingly more copies of that ﬁle available.
In this way, everyone shares the burden of carrying the index in direct
proportion to their use and the eKectiveness of the system grows with
the number of users (i.e. a naturally occurring network eKect). Heavy
users carry more of the burden—this is as it should be. Lighter users
naturally carry less of the burden, which, to me, seems fair.
User to chifra: chifra export <address>
Chifra reads the locally cached chunks of the index and exports records
to JSON, TXT or CSV. It does this by reading the index chunks, extracting
appearances for that address, querying the node for details of the
transactions, caching the returned data and then exporting the data to
the screen or a ﬁle. The caching is important because querying the
node and parsing the JSON result can be slow, reading from a local
binary cache is very fast.

Upshot
TrueBlocks shares hashes of indexes of appearances of addresses on
the Ethereum blockchain. These indexes “releases” the data from the
node (unchaining it) making it possible to build useful, fast, responsive
applications in a fully decentralized way.
By sharing hashes of immutable IPFS ﬁles we purposeful relinquish the
ability to ever take them back. We will never be able to say,
We’ve captured you. If you wish to continue to use our data, you must
pay us (or we’ll extract payment by invading your privacy).

We’re giving up this ability on purpose because we believe that if the
data on the blockchain is not fully free and open in the same way the
open-source software code is, the system will ultimately fail.
We will be rolling out, documenting, and writing about our work a lot
over the next few months. We’re going to Berlin in August, so if you
would like to discuss our work, please reach out.
If we can get a ticket, we’ll be at DevCon as well. Can someone get us
a ticket?


It's Not That Difficult
All about the Ethereum Difficulty Calculation
Sunday, November 24, 2019

Special thanks to a ﬁrst-rate Tuftian and data scientist, Ed Mazurek, for
early versions of the R code used in this article.
Each time the Ethereum time bomb goes o0, two related questions
arise. The ﬁrst question (and arguably the more important) is, “When
will blocks get so slow, they will be intolerable”. The second question
is, “How long should we delay the bomb this time?”
In this short article, I present a very simple — almost trivial — solution to
the second of these two questions. My proposal for how long we should
delay the time bomb is “Given a hard fork at block N, delay the di0iculty
bomb for N blocks.”
I won’t explain the di0iculty calculation in this article (see my previous
article). This image shows the calculation:

Looking closely at this equation, notice that it is composed of two parts.
I will call the ﬁrst part the “adjustment” (or part A). This portion adjusts
the current block’s di0iculty to account for the timing of the previous
block. This adjustment either lowers the di0iculty or raises it depending
how long the previous block took to appear. This component is
highlighted by the ﬁrst bracket above. Spend some time ﬁguring out
how it works. The e0ect of part A, which works exactly as it’s designed,
is to squelch changes in the hash rate. The charts below make clear that
this section of the calculation is working correctly.
Because part A works as designed — eliminating the e0ect of changing
hash rate — I will argue that worrying about hash rate when reasoning
about delaying the time bomb is unnecessary. To say this another
way — a too-low hash rate cannot make the bomb any worse than it is,
and too-high hash rate mitigates the e0ect of the bomb. This leads me
to conclude that part A of the equation is unrelated to slowing block
times. The second part of the equation (part B for bomb) causes all the
trouble.
It’s possible, as I do below, to separate these two components. This
makes it much easier to see that part A has no e0ect on block
production while part B—the bomb — dominates when it manifests
itself. Additionally, I will show that di0using the bomb is easy. One
simply needs to set the period back to zero each time one forks the
chain to di0use the bomb.

Generating and Formatting the Data
As is usually true when dealing with data, we start by acquiring data.
We used our own software library, TrueBlocks, to write the following
code.

#include “etherlib.h”
int main(int argc, char *argv[]) {
init_etherlib();
for (int i = 0 ; i < getLatestBlock() ; i++) {
CBlock block;
getBlock(block, i);
cout << block.blockNumber << “,”;
cout << block.timestamp << “,”;
cout << block.difficulty << end;
}
}

…which generates a very simple .csv data ﬁle of this format…
blockNumber
0,
1,
2,
...
8981997
8981998
8981999
8982000

timestamp
1438269960
1438269988
1438270017

difficulty
17179869184
17171480576
17163096064

1574448913
1574448935
1574448985
1574449029

2432407853358678
2432545292312150
2427931666241578
2424512564668329

Along with a list of daily hash rates that we got from Etherscan, this is
enough information to understand Ethereum’s di0iculty calculation.
We use RStudio and a data programming language called “R” to build
the following charts. If you’re not familiar with “R”, you should check it
out. It’s fantastic. We’ll start by looking at the Ethereum hash rate.

Daily Average Hash Rate
The data in the ﬁrst chart comes from Etherscan. It shows Daily Average
Hash Rate for the Ethereum Mainnet. I cannot vouch for this data, as I
don’t know how it was created, but I assume it’s okay.

Discussion: This chart reminds me of an Ethereum price chart. The rate
of growth of the hash rate skyrockets during the summer of 2017 and
peaks during the ﬁrst quarter of 2018 (just like price). The bump in
October of 2016 is the infamous 2016 DDos attack, and the two vertical
grey lines are the Byzantium hard fork (Oct. 2017) and the
Constantinople hard fork (Jan. 2019). There’s not a lot to say about this
chart, but we will refer to it as we discuss the chain’s di0iculty data,
below.

Raw Di8iculty Data
The ﬁrst chart based on di0iculty data shows data returned from
Parity’s RPC call eth_Block. First, some standard statistical
information:
Summary Statistics for Ethereum’s Difficulty
16,970,000,000,000 Minimum
111,700,000,000,000 1st Quartile
1,649,000,000,000,000 Mean
1,926,000,000,000,000 Median
2,687,000,000,000,000 3rd Quartile
3,702,000,000,000,000 Maximum

Our ﬁrst chart is straight forward:

Discussion: The data was produced at block 8,920,000, and while “R”
can easily handle that many records, given the iterative nature of data
exploration, we choose to sample one record out of every 100. This give
us around 9,000 records which are presented in the above chart. The
same grey vertical bars showing the hard forks in this chart as well.
The height of the red line, 𝑦(𝑥) = 𝑑𝑖𝑓𝑓𝑖𝑐𝑢𝑙𝑡𝑦! represents the di0iculty
at the time of the given block. It’s easy to see the e0ect of di0using the
di0iculty time bomb at each fork. The DDos attacks in the Fall of 2016,
and, referring to the hash rate, you may also see a relationship between
di0iculty and hash rate.
Pretend the time bomb had not been di0used— in e0ect push the red
line up as far as the bomb di0usion pushed it down — do this for both
hard forks and you can almost see the same shape as the hash rate
chart. In other words, hash rate and di0iculty are tied together. This
makes perfect sense. This is exactly what part A of the di0iculty
calculation is designed to do. It’s designed to adjust the di0iculty in
direct response to varying hash rates.

The behavior of miners probably doesn’t change because of a di0icult
bomb di0usion. Their mining rigs continue to run identically before the
di0usions as after them. The only thing that changes is the time
between blocks, on average, lowers.
Before we leave this chart, notice something else. Look carefully just
prior to the Byzantium hard fork. You’ll see four or ﬁve single-line,
vertical jumps in di0iculty level. In fact, the jumps are twice as large
each time they appear as the previous time they appeared. These
jumps are the time bombs. Let’s focus more on that area of the chart:

Discussion: We’ve inserted vertical lines at each 100,000-block
boundary. Notice, just prior to the hard fork, the jumps in di0iculty lands
exactly on the markers. In between the markers, the di0iculty continues
to go up, but nowhere near the speed as on the markers. The interexplosion increases are consistent with the fact that the overall hash
rate was raising very quickly at this time in 2017.
You may also notice that each successive “bump” — each time the
bomb explodes — it explodes twice as high as it did the previous time.
The periodic nature of these jumps, it turns out, become important in
understanding what’s going on.
In the remaining charts in this article, our goal will be to separate the
ﬁrst part of the di0iculty calculation (part A or the adjustment) and the
second part of the calculation (part B or the bomb). This will help us
understand how to handle the bomb in the future.

Per-Block Changes in Di8iculty
In the next chart, we look at the change in the di0iculty between each
successive block.
That is
𝑦(𝑥) = 𝑑𝑖𝑓𝑓𝑖𝑐𝑢𝑙𝑡𝑦! − 𝑑𝑖𝑓𝑓𝑖𝑐𝑢𝑙𝑡𝑦!"#
Here’s the chart of that data:

Discussion: As we mentioned above, part A of the calculation ‘hovers’
around a di0iculty level ensuring 14 second blocks. This “hovering” is
revealed in the above chart by the red-blue nature of the data. The
“grow” part of the chart (red) represent positive changes in di0iculty (i.e.
di0iculty gets higher, block production gets lower, and block times get
slower). The “shrink” values (blue) are negative (lower di0iculty, faster
blocks and more blocks). The “adjustment” hovers around zero. In
other words, the calculation is trying to maintain consistent block
times. The part A calculation homes in on a value — 14 second block
times.
On this chart, you may also, for the ﬁrst time, clearly see the bomb
exploding. And it’s arguably obvious that the “explosion” is getting twice
as large each time they explode.

But I don’t think this chart is clear enough. Why, for example, does the
same pattern not show up clearly during the Constantinople period? It
turns out the reason for this is because of the greatly increased hash
rate. The calculation maintains the 14 second blocks times, but the
system is vacillating (swinging back and forth) more vigorously. The
reason we can’t discern the same bomb activity near the
Constantinople fork is because it’s being obscured by the wider
ﬂuctuations.
Can we somehow remove the obscuring nature of the hash rate? Yes,
as a matter of fact we can, and we do that in the next chart.

Relative Change in Di8iculty
Our ﬁnal chart in this part of the paper is a chart showing the change in
di0iculty relative to the total di0iculty in the block. That is
𝑦(𝑥) =

𝑑𝑖𝑓𝑓𝑖𝑐𝑢𝑙𝑡𝑦! − 𝑑𝑖𝑓𝑓𝑖𝑐𝑢𝑙𝑡𝑦!"#
𝑑𝑖𝑓𝑓𝑖𝑐𝑢𝑙𝑡𝑦!"#

The previous chart shows the raw change in di0iculty. This chart shows
a normalized change. This removes the growth in hash rate. It will allow
us to clearly see the two di0erent parts of the calculation, part A and
part B.
Here’s the chart of per-block deltas in di0iculty relative to total
di0iculty:

Discussion: And now we can clearly see why I said above that worrying
about hash rate is counter-productive to a discussion of the di0iculty
bomb. You can see clearly that until the time bomb “rears its head,”
block production is not a0ected by increased (or decreased) hash rate.
Part A of the calculation maintains a steady state relative to block
speed and production. Di0iculty (on average and relative to itself)
remains nearly unchanged until the bomb starts exploding.
Two interesting things to notice: (1) You can see the bomb starting to
“rear its head” at the right-most edge of the chart even though the
distance between the appearances of the bomb are much shorter than
between Byzantium and Constantinople — I explain this below; (2) the
striations at the bottom of the chart are an artifact from the integer
division by 10 in part A of the calculation; (3) the higher hash rate
appears to delay the “head-rearing” as Lane Rettig noticed prior to
Constantinople and we wrote about in the above mentioned article.
There’s a lot more I could talk about on this chart, and maybe someday
I’ll come back to it, but I do want to make the case for a better way to
di0use the bomb in the future.

A Better Way to Di8use the Bomb
The ﬁrst point I want to make is that worrying about the hash rate while
worrying about the time bomb is counterproductive. The e0ect of
increases (or decreases) in hash rate is entirely removed by part A of the
calculation. That’s what part A is for. Additionally, part A works perfectly
as is evidenced by the near-perfect ﬂatness of the di0iculty deltas
(relative to di0iculty’s current height). Hash rate has no e0ect on block
time — but this is what we already knew — this is what the di0iculty
calculation is designed to do.
The bomb is deﬁned in the above equation as an additional value added
to the end of the calculation as
$%

2

&'(()*+_-./&0
3"45
#11,111

That is, two raised to a certain power. We can re-write this as 26
allowing
𝑝=

&'(()*+_-./&0
#11,111

.

(The -2 may be ignored as it is simply a translation). We’re left with a
step function in p as is shown on the following chart.

Remember though, that the core developers have twice reset the
bomb — this means, quite literally, that they reset the period. The Go
code works by creating a fake_block to be used in the calculation.
The fake_block makes the real block appear to be in the past, which
resets the bomb. Here’s a corrected chart showing what’s really going
on with the period:

The true block number appears in red above and ranges from zero to
8,920,000. The fake block number (in green) tracks the real block
number until Byzantium when it was reset 3,000,000 blocks to the past.
It then runs parallel to the real block number until it gets reset again
(this time by 5,000,000 blocks) at the Constantinople fork.
Here’s a table recounting the resets. Do you see anything odd?
Hard fork
Byzantium
Constantinople
Istanbul

Block no.
4,370,000
7,280,000
9,069,000

Setback
3,000,000
5,000,000
9,000,000

Fake
block
1,370,000
2,280,000
69,000

Period
13
22
0

The fake block after the Byzantium fork was 1,370,000 which, when
integer divided by 100,000 gives a period of 13. That is, at each block an

additional di0iculty of 2#7 was added after the hash rate adjustment.
By the time of the Constantinople hard fork, the period was reset to
2,280,000 which translates to period of 22, an additional di0iculty of
244 . This, I think, is why the time bomb is going o0 earlier than we
anticipated. We didn’t reset it back far enough the last time.
The suggested value of resetting the time bomb this time is back to
block 69,000 which translates to period zero. This is exactly the right
amount to reset to.

How to Better Reset the Di8iculty Time Bomb
There’s a very easy way to reset the time bomb to the right value each
time we reset it. This will make most of the problems related to the
di0iculty bomb disappear. Simply set the time bomb back the same
number of blocks as the FORK_BLOCK_NUMBER (with a little bit of play).
In this way, the fake block gets set near zero and the period becomes
zero.
Because the time bomb dominates the slowing down of block
production this will have e0ect of making the ‘head raising’ totally
predicable. The calculation for part B is dependent only on the fake
block number. If we had done that for the Constantinople fork, the time
bomb would not be going o0 so soon.
Fake
Hard fork
Block no.
Setback
block
Period
Byzantium
4,370,000
4,370,000
0
0
Constantinople
7,280,000
7,280,000
0
0
Istanbul
9,069,000
9,069,000
0
0

Conclusion
Thanks for reading this far if you’ve gotten this far. I hope this document
ﬁnds interested ears. I see a lot of confusion related to the di0iculty

calc issue (not only in the public, but in the core devs channel and the
Ethereum Magicians). I feel people are making this much more di0icult
than it should be.
A few simple takeaways: (1) because of its exponential nature, only part
B (the bomb) matters as far as the block production is concerned; (2)
part A of the calculation has no detrimental e0ect on block production
and it may actually have a positive e0ect; and (3) resetting the period or
fake block to zero has two beneﬁts: (a) it grants the maximum amount
of time in delaying the time bomb, and (b) it makes the re-appearance
of the bomb very predictable.
Please let me know what you think of this paper. I hope I’ve helped
explain something that I know about well.


A Method to Diffuse the Difficulty Bomb
Diffusing the Bomb Predictably and Lowering the Angst
Tuesday, November 26, 2019

There is no more angst-ridden profession than being a member of a
bomb squad. You’ll know what I’m talking about if you’ve ever seen the
movie “The Hurt Locker.” In a recent Gitter post in the All Core Devs
channel, Alexey Akhunov says of the diIiculty bomb that it,
…forces people to make rushed decisions and be reckless, without
real emergency…
This is true currently, but we can do better.
In a previous article, I mentioned two questions that arise each time we
discuss the Ethereum diIiculty bomb: “How soon will the block times
become intolerable?” and “For how many blocks should we reset the
bomb?” I presented in that article a simple method to decide the
second of these two questions. In this article, I extend that argument
and run through an example.
I then try to show why, if we adopted the method I’m proposing, we
could answer the ﬁrst question more predictably, removing some the
‘angst’ or ‘emergency’. Later this week, I will write an article analyzing
the current situation and making a prediction of when the block times
will become intolerable.

How Far Back to Reset the Bomb
Currently the core developers pick the number of blocks to reset the
bomb before picking the fork_block. This is incorrect.
In the method I’m suggesting, we should pick the fork_block ﬁrst.
Assuming we do that, this is the algorithm we should use to decide how
far back to set the bomb:
# Choosing the fork_block first gives the setback
fork_block = whatever block we pick
setback
= fork_block - 50,000
fake_block = real_block — setback
= fork_block — (fork_block — 50,000)
= fork_block — fork_block + 50,000
= 50,000
period = floor(fake_block / 100,000)
= floor(50,000 / 100,000)
= 0

In other words, by doing this we always set the period to zero giving us
the maximum amount of time before the next bomb overtakes the
diIiculty calculation. As we learned in the previous article, until the
bomb appears, block times hover at 14 seconds due to the adjustment
calculation.
The adjustment calculation (part A):
•

Works as designed by maintaining 14 second blocks.

•

Works in the face of any hash rate (this is part A’s purpose).

•

Never allows the block time below 14 seconds.

•

Eventually cannot keep up with the bomb as the bomb to “rears its
head.”

As discussed elsewhere, block times increase only after the bomb
begins to “rears its head.” which, from experience, happens around

period 40. (This number is somewhat dependent on hash rate, but it
eventually gets overwhelmed. We can measure this empirically.)
After resetting the period to zero, there will be 40 periods of
approximately 14s blocks before the bomb becomes a problem. We
can easily predict how long this will take, remembering that the
adjustment calculation (part A) works as designed to keep blocks 14
seconds apart.
We can now calculate when the next bomb will explode:
# 40 periods of 100,000 blocks at 14 secs / block
timespan = 14 seconds * 40 periods * 100,000
blocks
= 14 * 40 * 100,000 seconds
= 56,000,000.0 seconds
=
933,333.3 minutes
=
15,555.5 hours
=
648.1 days
=
1.7 years
=
~21 months

Using this method (setting the period back to zero) will ensure at least
21 months before the next time we need to worry about the bomb. No
more angst.
Note that hash rate can only slow this progress. The diIiculty
calculation never leads to faster blocks (unless the block times are too
slow). The calculation raises the block times if they go under 14
seconds. So, the next bomb will never come before we expect.
Block times will become intolerable no earlier than fork_block +
4,000,000. Or, to say it another way, fork_block + 21 months.

Short Caveat
This article proposes a new method for selecting how far back to reset
the bomb with the next hard fork and all related bomb-related hard

forks in the future — including the pending hard fork to diIuse the bomb
being proposed currently.

Summary
In this article, I tried to show a simple methodology to select how many
blocks to oIset when diIusing the Ethereum diIiculty bomb.
The method I propose suggests that we ﬁrst pick the fork_block
number. Only then decide how many blocks to reset the bomb. Given a
fork_block, select the oIset to be fork_block minus 50,000.
This will have the following eIects: (a) it resets the bomb as far back as
is possible while staying away from zero which may (but doesn’t) cause
an edge case; (b) it makes the earliest block at which the bomb “rears
its head” more predictable (assuming period 40 is where the bomb ﬁrst
appears — this can be easily measured); and (c) it eliminates the angst
from the decision-making process, making it more likely that the core
developers will not remove the bomb entirely out of discomfort, as has
been suggested.
(Not to go down another rabbit hole, but I believe the bomb serves a
very important, larger societal purpose and it should be preserved
vehemently.)


Simple Undeniable Facts
Sharing Everything Always
Friday, January 10, 2020

One of the reasons I fell in love with blockchain way back in 2013 was
because for the ﬁrst time in human history, engineers had created a
system that could produce undeniable data. Undeniable access — yes,
but more importantly, undeniably high-quality data. That was a
paradigm shift.
While blockchain data may be untrue, it is mathematically certain that
at least 51% of the participants in the creation of the data have agreed
that it is consistent — maybe that’s as close as we can get.
Bitcoin (and other blockchains) produces the best data the world has
ever seen, and I think that this near-perfect data will usher in a better
world. At least, that’s what I used to think…

In this post, I will present a list of what I hope are obviously factually
true statements. I want to explain a looming danger headed our way. An
out-of-control, 18-wheeled tractor-trailer truck is coming right for us.

Each “truth” in the following lists tries to say something simple and
obvious. After a few such statements, I try to draw an equally obvious
conclusion. I’m not sure this will convince anyone of my views, but that
never stopped me from espousing my views in the past, so, oT we go.
Note: In the following, we use the term “appearance” quite frequently.
An appearance is a <block_number, transaction_id> pair pointing to the
locations on the chain where a given address can be found. We’ve
explained this concept many times in other articles and won’t expound
further.
On to our presentation of Simple Undeniable Facts™.

Blockchain Data is Too Big
•

The only way to ﬁnd something in an unsorted list of data (that is,
unindexed data) is brute force scanning.

•

Brute force scanning is slow and becomes increasingly slower as
the data grows.

•

A blockchain is ordered by time, not address. In this sense, a
blockchain is an unsorted list of data with respect to addresses.

•

Indexing a database increases the size of the database.

•

Commercial-grade computers — such as laptops and desktop
computers — have limited hard drive space.

•

While the Ethereum node does index some of its data (such as
topics), it does not index all address appearances.

•

Many appearances happen deep inside of smart contracts when
the contract uses an address as data. If the smart contract does not
generate an event, there is no way of knowing of these
“appearances.”

•

Sometimes smart contracts carry out material movements of value
without generating events. (For example, minting tokens which is
not part of the ERC20 standard.)

•

Many appearances are buried deep inside the traces and are not
identiﬁed by naive, shallow methods of indexing.

•

To produce a 100% perfect, audit-quality accounting for a given
address, one must identify every appearance. This means one
needs to descend into every trace. This means one must either
produce an index or scan the chain for each address history.

Conjecture: Without an index of appearances, it is impossible to do an
accurate, full, and complete accounting or audit a given address on
commercial-grade hardware.
Conjecture: If users cannot do a full accounting on their own
machines, they must either forgo a full accounting of their own
addresses or rely on a third parties to give them the data they need.
Conjecture: If the system requires third parties to account for itself, it
is not decentralized.
Conclusion: Ethereum is not now and will never be truly decentralized
unless something changes. That something is indexing as part of the
protocol.

Privacy for All — Transparency for Most
•

Most people care deeply about themselves and their own ﬁnancial
histories but very little about other people’s ﬁnancial histories.

•

No person wants their ﬁnancial history used without their
permission.

•

Some people explicitly give their permission to be more
transparent. For example, perhaps they are involved in a
community organization.

•

Some community organizations require transparency to function
properly.

•

Some people peek into other people’s medicine cabinets.

•

There’s money to be made from peeking into other people’s
ﬁnancial medicine cabinets.

•

Some people are unethical.

•

For some people, the pressure to be unethical is directly
proportional to the amount of dollars to be made.

Conjecture: As blockchains become more prevalent, because people
cannot get their own data (see above), they will come to rely more
heavily on third parties.
Conclusion: As time goes on, third parties will see increasing pressure
to act unethically and invade people’s privacy.

She Who Owns the Hard Drive, Owns the Data
•

Every piece of useful data is eventually stored on a hard drive.

•

If the only information you use is from your own hard drive — barring
viruses — no-one can see what you’re doing.

•

If you’re using data on someone else’s hard drive, it’s impossible for
you to stop them from peeking at what you’re doing.

•

Individuals do not like to pay for services — especially digital
services.

•

It’s easier to acquire users if you give them something for free than
if you charge them. The most service for the least price wins the
network eTect battle.

•

Digital service providers cannot provide free services for free,
therefore, at some point, they must monetize the service they
provide.

•

Most people are easily convinced to accept free services without
asking deeply about the real cost of the service. Hint: the real cost
is privacy.

Conclusion: Over time service providers will appear who deliver free
access to people’s personal ﬁnancial data. These companies,
delivering data from their own hard drives, will eventually conclude that
peeking into their users’ private ﬁnancial history can be very lucrative.
They will not be able to resist.

Blockchain Data is the Best Data Ever Seen
•

Blockchains organize data in a near-perfect, immutable way, and
while the data is genuinely diTicult to understand and use, making
it nearly impossible for normal people to use it, businesses will be
able to extract deep insight into anyone’s personal ﬁnancial history.

•

The quality of the data, and therefore its value, will skyrocket when
blockchains became prevalent. The quality of that data will
continue to increase exponentially as adoption increases, tools
improve, and network eTects take over.

Conjecture: Unscrupulous businesses with lots of resources, lots of
technical skill, and very large hard drives will be invading everyone’s
privacy in new and novel ways in twenty years.
Conjecture: Almost every team in the ecosystem is working on
systems that rely in some way on third parties.

Conclusion: We should stop building something we don’t want.

Is There Any Hope…Only If We Work Together
•

Because people care about their own ﬁnancial information, they’re
incentivized to hold that data closely (if they’re smart).

•

Blockchains work by sharing data between peers. They wouldn’t
work if the peers only stored their own data.

•

While sharing must happen, not every participant must carry every
piece of the data on their hard drives if there’s enough redundancy.

•

It is necessary for every participant to have access to every piece of
data but not carry it. If they have undeniable access, that’s enough.

•

Blockchains are immutable.

•

Once blockchain data ages past a certain number of minutes (say
six for Ethereum, sixty for Bitcoin), the data never changes for the
rest of time.

•

There’s something about immutable data that loves a cache.

•

There’s something about immutable data that loves a hash.

•

Content-addressable ﬁle systems, such as IPFS, store data on the
user’s local hard drive.

•

The location of that data is identical to the ﬁle’s ﬁle name (that is,
the location is the hash — now and forever).

•

Once content is written to IPFS, the hash for that data never
changes. More accurately, if the contents change, the ﬁle’s location
changes.

•

Hashes are very short and very easy to share. With an internet
connection, the IPFS software, and an IPFS hash, the user has

undeniable access to the contents stored at that hash (as long as
one person with that content is online).
•

Blockchains continually produce new data, and therefore
continually produce new address appearances that must be added
to the index.

•

On Ethereum, new blocks come every 14 seconds.

•

Constantly changing data and content-addressable ﬁle systems do
not work and play well together.

•

To alleviate this, blockchains make blocks. In other words, they
stop adding new data to the current reference making it immutable,
hashable, and “IPFS-able.”

•

An index of address appearances in blocks has the same problem
and to solve that problem one may use the same methodology:
stop adding data to the current reference. In TrueBlocks, this
creates chunks or snapshots. Snapshots are immutable. This is
what TrueBlocks’ Unchained Index does.

•

Like a blockchain, if one is constantly creating new immutable
chunks of data, one must publish the hashes of that data so other
peers can ﬁnd them.

Conjecture: If we wished to, we could change the blockchain node
software to produce content-addressable hashes instead of “regular”
block hashes.
Conjecture: The physical storage of the blockchain data is optimized
for the blockchain’s use — not for content-addressable storage.
Conjecture: Having the blockchain store its data in a contentaddressable store would ruin its performance.
Conjecture: It would be possible for blockchains to produce a contentaddressable hash without storing its data in a content-addressable

store. This would maintain the performance of the chain and keep it
working.
Conjecture: The POW and diTiculty calculations would not need to
change — the previous block’s IPFS hash could be inserted into the
current block, thereby maintaining the chain. A random nonce could be
inserted as well, maintaining the POW — everything would work exactly
as it works now.
Conjecture: The fact that the block’s hash is an IPFS hash would be
informational only — a by-product so to say.
Conjecture: The meaning of the block hash would go from being utterly
meaningless to very, very meaningful.
Conclusion (most important): Block hashes should be IPFS hashes.

Blockchain’s Killer App is Community
•

If the blockchain’s hashes were IPFS hashes, the community —
unrelated to how the software actually stores its data — would
come to understand that the block data is available at the location
of the hash.

•

Knowing this, users would come to “expect” the block data in that
location. As a result, the data would almost magically start to
appear in that location, even if storing it there wasn’t part of the
protocol.

•

If everyone knew where the data was, everyone would go there to
look for it. A Schelling point would arise.

•

As more people ﬁnd the data in that location (or put it there if they
don’t ﬁnd it there), the likelihood of ﬁnding it there in the future
grows.

•

There would be no other logical place to look for the data other than
its IPFS hash.

•

Because people value their own data, they would would have no
problem pinning their own data.

•

Software written to use the data could pin it by default.

•

This would have no detrimental eTect on the user machine as
he/she would need to store local application-speciﬁc data anyway.
There would be no downside — only upside.

Conjecture: Over time, regular users could run extremely light nodes —
the only thing those nodes would store would be the hashes of where
to look. At block 10,000,000, which is where the chain is now, this
would take up 32 bytes * 10,000,000 blocks — or about 2.5 GB. Anyone
can carry that much data.
Conjecture: If we did this, every piece of on-chain data would be
available — undeniably — to anyone in the world on their local
machines.
Conjecture: Each user could hold the full block detail of only those
blocks their addresses of interest appear in. If the wallet or tool
software pinned these blocks, they would be available to others.
Conjecture: The system could be designed redundantly. When a user
asks for a single block, the system could give her (and pin) more than
one block. In this way, participants would hold (and pin) their own data
plus an additional portion on behalf of the community.
Conclusion: The community could conceivably carry and distribute its
own data without any third parties if the node software was designed
that way.
Note: It seems likely that some number of ‘super-nodes’ would have to
run on behalf of the entire community. This would be needed to ensure

that every single block is available somewhere. Ethereum Foundation,
Consensys, Infura, Etherscan and other small and large players, could
run these super-nodes. It would be their donation to the community.

How Would People Know Where to Look?
•

An index of appearances is needed for users to know where to ﬁnd
their histories.

•

TrueBlocks has already built an example of an indexer that builds,
hashes, pins, and queries the appearance index. [2026: It’s called
Khedra.]

•

After nearly 10,000,000 blocks, the appearance index consists of
around 3,500 snapshots (chunks). This index is immutable,
compact, stored as a binary ﬁle both locally and in IPFS, and takes
up about 35GB. The index can be searched locally with 100%
privacy.

•

An amazing data structure called a Bloom ﬁlter is used to lessen the
size of the index even further.

•

The Bloom ﬁlters are also immutable, very compact, and pinned on
IPFS. The Bloom ﬁlters allow the software to identify which chunks
of the index are needed.

•

The Bloom ﬁlters are only 1.5 GB (also in 3,500 ﬁles), making them
even more accessible to regular users.

•

Our system goes even further than that. For both the Bloom ﬁlters
and the index chunks, we produce a text ﬁle (called a manifest)
listing the hashes which we distribute as part of the software. The
ﬁrst time one runs TrueBlocks, it downloads 1.5 GB of Bloom ﬁlters.
It then starts a daemon (called Khedra) to catches up to (and keep
up to) the front of the chain.

•

Using the Bloom ﬁlters, the software queries (quickly) for a list of
snapshots. The list of snapshots are then downloaded. For almost
every address, this list is very, very small.

•

For popular addresses (such as CryptoKitties, ENS, or DAI) the list
of snapshots is very large. Nearly every chunk is of interest.

Conclusion: This methodology naturally requires heavy users of the
chain to download and pin more data than regular users — this is a very
desirable quality for a distributed system. It’s naturally fair.
Conclusion: Once one has the list of hashes, one has undeniable, fully
local, 100% private access to the Bloom ﬁlters. This gives one
undeniable access to the minimal list of index chunks. This gives
undeniable and fully local access to every transaction for an address.
Conclusion: TrueBlocks gives everyone access to a full index of
appearances of every addresses anywhere on the chain. This access
comes on your own local machine. From there, you can operate 100%
privately and fully decentrally.

Conclusion
We intend to write further about how TrueBlocks accomplishes the
creation of the appearance index and how it post-processes those
appearances into snapshots and how it stores the data on the hard
drive. Visit us again in a few weeks.


How Accurate is Etherscan
What’s the Purpose of an Impossible-to-Balance Ledger?
Friday, June 12, 2020

I had a call this morning with a cryptocurrency accountant. He’s a
wonderful fellow. One of those people who can happily wade through
thousands of rows of a spreadsheet trying to get the digits to behave
themselves. He’s a man after my own heart.
This accountant — call him Mr. Green — makes a good living helping
people do their crypto-taxes. He’s busier than ever. He tells me, and
here I quote, “Nothing ever balances.”
How is this possible? Isn’t Ethereum supposed to be an 18-decimalplace accurate ledger? Doesn’t everyone using Ethereum have the
same perfect data?
The answer, of course, is that internally Ethereum always does balance.
Why then is Mr. Green having trouble? Because he’s not using the
internal Ethereum data directly. He’s using other people’s version of
that data, that is, third party data. Very few of us are using data directly
from the node software. We almost all of us get our data from an API or
a website. In the Ethereum ecosystem, this means either Infura or
Etherscan.
I, like Mr. Green, want those annoying little digits to behave themselves.
I’m becoming increasingly concerned, especially as smart contracts
become more complicated, about the fact that “Nothing ever
balances.”

Think before Tweeting…
This morning, I tweeted that Etherscan misses nearly 15% of what I call
“address appearances.”
Of course, this was met with huge skepticism, so I‘m writing this article
as a way of showing my readers what I meant by that comment.

What is an Appearance?
I’ve written about appearances before. I’ll quote from that previous
article:
We deﬁne an “appearance” as either: (a) the use of the address in one
of the common “address” ﬁelds of a transaction (such as to, from,
eventEmmitter, or contractAddress) or (b) the use of an address as
data in the input or eventData ﬁelds of the transaction, or (c) its
subsequent use in any smart contract invocation.
In other words, given the byte string representing an address, an
address’s appearance list includes every transaction where that byte
string shows appears anywhere in the transaction’s data. Sounds
simple. It’s not!
To solve this, we scan the chain and build an index of every appearance
we can ﬁnd. To test our code, we use Etherscan’s APIs. You’ll
understand how surprised we were this week when we discovered that
TrueBlocks consistently ﬁnds more appearances than Etherscan. We
were just as skeptical.

Etherscan APIs
Etherscan provides ﬁve account-related APIs which we use in our
testing1:
A list of ‘Normal’ Transactions by Address
A list of ‘Internal’ Transactions by Address
A list of “ERC20 — Token Transfers” by Address
A list of “ERC721 — Token Transfers” by Address
A list of Blocks Mined by Address

Quite some time ago we built a tool called chifra slurp that makes
use of these endpoints:

Given an address, chifra slurp returns all the transactions where
Etherscan says that address appears.
This week we built a shell script, fromES, that successively calls each
of these ﬁve endpoints and assembles the results into a single ﬁle and
then compares those results against TrueBlocks.
We ran tests against 100 randomly chosen addresses. For example, the
command:

1

As of the date of the writing of this article.

fromES 0x91c5fa6872f3a93b999843eaf06eb34a18a69a12

produces these results:

Notice that the ﬁve Etherscan endpoints deliver 31, 0, 26, 2, and 0
records respectively. Summed, in the line labeled ‘all’, this totals 59
records.
At ﬁrst blush, we were concerned when we saw that this was more
records than TrueBlocks returns (40). What we discovered was that
Etherscan’s ﬁve endpoints include duplicate records. We altered our
shell script to remove the duplicates producing the line ‘sorted uniq’.
Looked at it this way, TrueBlocks ﬁnds one extra record not found by
Etherscan. It turns out to be transaction #8156524.14 (that is, the 14th
transaction of block 8,156,524).
TrueBlocks provides a few other tools that allow us to see the details of
any Ethereum transaction. For example,
chifra trace 8156524.14

returns

which I admit is a bit crazy looking, but it does contain the byte string
for the given address (colored pink). It turns out that this transaction
ended in a revert. This is probably why Etherscan doesn’t ﬁnd it.
Etherscan obviously ﬁnds all transactions of type (a) mentioned above.
It ﬁnds most type (b) transactions as well (Etherscan calls these
internal transactions). If this transaction had not reverted, it would
probably have been an internal transaction because the address is
used as data as opposed to being one of the address-related ﬁelds.
This happens frequently on Etherscan. Of the 100 addresses we tested,
TrueBlocks found more transactions than Etherscan for 85 addresses.
For not a single address (zero) did Etherscan identify a transaction that
TrueBlocks did not ﬁnd. 85–0. I call that statistically signiﬁcant.

But Are Those Transactions Relevant?
I once spoke to a diierent crypto accountant (Mr. Greenish-Blue?) and
after listening to me talk about this issue, he asked, “Yeah, but is it
relevant?” My reaction was “Of course, it’s relevant.” (I could explain
why this is true, but if it isn’t self-evident to you already, I wouldn’t be
able to convince you anyway.)

Why build an 18-decimal place ledger that the entire world comes to
agreement on and spends way more on than any previous computing
system if we’re going to throw away data? Explain that to me, please.

Analyzing the DiBerences
Since we ﬁrst published this article, a few people have asked us to
further analyze the diierences. To be honest, we don’t have the time.
Maybe we’ll write about this issue later2. Perhaps someone else will
pick this up. We published the raw data to this IPFS hash:
QmYccKCCoYat8AsgMxVf Fjcb9NG8cnFﬁkkrweRyjwo7rQ.
As a quick summary, the missing transactions appear to be of at least
six types:
1. Input data of both errored and completed transactions
2. Log topics of both errored and completed transactions
3. Log data of of both errored and completed transactions
4. Output data of deeply embedded traces
5. Uncle mining rewards
An example of the ﬁrst type is this transaction:
https://etherscan.io/tx/0xa4a96ca16373239fd679
711b05bcbdc138bc40a5bb2a085799c23bbaf5fd2a3a

which does not appear in the internal transaction list of the address
0x28f4a17f8a99ab90c1a401b85d694b2c0ea40c4b

however, that address is clearly in the input data of the transaction.

2

We did write about this in later articles. See below.

Future Work
We’re going to dig deeper to see if we can’t ﬁgure out what some of the
other missing transactions look like. We already know that Etherscan
does not report on Uncle miners.


Ethereum's Issuance: minerReward
Saturday, August 15, 2020

Recently, there was a dustup on Crypto Twitter about Ethereum’s
money supply. The claim was made that Ethereum’s money supply was
not easily available, nor was it widely agreed upon.
News ﬂash: Both claims are correct.
At one point, our project, TrueBlocks, was mentioned, so I thought I’d
write an article (which has grown into two articles and a code base)
exploring the issue.
While the work we present here doesn’t necessarily make the numbers
easier to get (ﬁx the node core developers!), the numbers are accurate
to 18-decimal places and veriﬁed to the on-chain account balances at
every block. We used TrueBlocks to do that.

Going Back to Basics
To really dig in, we are going to go back to basics. In Ethereum, that
means reading the Yellow Paper. We’ll look closely at Section 11.3, as
that section describes the issuance of new ether, or as they call it, the
Reward Application.
I will present each sentence from Section 11.3 verbatim and then show
a copy of the equations associated with that sentence, then translate
those equations into English. Finally, we translate the equations into
purposefully simpliﬁed C++ code using TrueBlocks. TrueBlocks helps
us extract and analyze the data as part of the code base.

Section 11.3
We start with the introductory sentence of Section 11.3:
11.3 Reward Application. The application of
rewards to a block involves raising the balance of
the accounts of the beneficiary address of the
block and each ommer by a certain amount.

Since there are no equations, we will skip right to the English language
re-write to get started:
At every block, the balances of at least one, but
possibly more, accounts are issued new coins.

Let’s deﬁne the word “ommer.” Ommer is a gender-neutral word for the
sibling of a parent. Your aunts and uncles are ommers. The words
ommer and uncle are used interchangeably in this article. Every block
has zero or more ommers. There are at most two ommers in a block
(Section 11.1).
We will call the ﬁrst of the two increases mentioned in 11.3 — the
increase to the account balance of the beneﬁciary of the block
reward — the blockReward. Every block has a blockReward.
blockRewards are discussed below.
We will call the second of the two increases — to the account balance
of the beneﬁciaries of the ommer rewards — the uncleReward. Not all
blocks have uncles; therefore, not all blocks have an uncleReward.
uncleRewards are discussed in the following article.
There’s a third concept, that we call the nephewReward, that is actually
discussed below since it’s part of the blockReward, so it doesn’t get its
own separate article.

The Crux of Section 11.3
The next part of Section 11.3 of the Yellow Paper is the crux of the
calculation. This describes Ethereum’s issuance fully:
We raise the block’s beneficiary account by 𝑅!"#$% .
Also, for each ommer in the block, we raise the
&
block’s beneficiary account by an additional '( of
𝑅!"#$% .
The beneficiary of the ommer gets rewarded
depending on the block number.

There are two rewards mentioned here. One, the blockReward, has
two parts; and, the other, the uncleReward, has one part.

blockReward
Focusing just on the blockReward (for the block’s beneﬁciary
account)
We raise the block’s beneficiary account by 𝑅!"#$% .
Also, for each ommer in the block, we raise the
)
block’s beneficiary account by an additional !"#$%.
'(

You should be able to see that this is referring only to the beneﬁciary
account (that is, the winning miner). It says the winning miner’s account
balance is raised (i.e. new coins are added to its balance) by an amount
&

designated by 𝑅!"#$% plus !"#$% for each uncle in the block.
'(

The Yellow Paper summarizes the production of this amount with two
equations:

(161)
(162)

Ω(𝐵, 𝜎) ≡ 𝜎 * : 𝜎 * 𝜎
||- ||
𝜎 * [𝐵+$ ]𝑏 = 𝜎6𝐵+$ 7𝑏 + (1 + '(& ) 𝑅!"#$%

𝑒𝑥𝑐𝑒𝑝𝑡:

Where:
161

162

symbol
𝐵
𝜎
𝜎)
Ω
𝐻$
𝐵*$
[𝐵*$ ]
𝜎+,& - 𝑏
$

𝜎 ) +,& - 𝑏
$
||𝐵. ||
𝑅!"#$%

meaning
The current block
State at previous block
State at current block
Transition function from previous state to current
state
A ﬁeld in the block’s header representing the
winning miner
The value of that ﬁeld (i.e. the miner’s address)
The account stored at that address
The balance of the miner’s account prior to the
block
The balance of the miner’s account after the block
The number of uncles in the block
The base reward as described in Equation 166

In other words,
The function Omega does not change the state
except for the balance of the winning miner’s
account. It changes that account’s balance by the
&
block reward plus
of the block reward for each
'(
uncle in the block.

A concept called baseReward is deﬁned in equation 166:
(166) 𝐿𝑒𝑡 𝑅!"#$% = 2 𝑥 10&.

In other words, two Ether. Note: This is slightly incorrect as it ignores
the change in the block reward at the Byzantium hard fork (from 5 to 3
Ether) and then again at the Constantinople hard fork (from 3 to 2). This
omission is accounted for in the source code.

Translating this to C++ code:

If there are no uncles, blockReward is identical to baseReward. If
there are uncles, blockReward is increased by the baseReward
/

multiplied by '( for each uncle. We call this additional increase the
nephewReward. (You may call something a “nephew” if it has an
uncle.)

As crazy as this may seem, this is the entire calculation for
blockReward. We will look at the uncleReward in the next article.

TrueBlocks Helper Functions
The two helper functions, getUncleCount and getBlockReward,
(shown in green in the code snippet) are part of TrueBlocks and remain
to be explained.
getUncleCount is a passthrough to Ethereum’s RPC call
eth_getUncleCount. The interested reader is referred to the RPC
documentation for that function.

The getBlockReward function is equally simple, as it is a function of
blockNumber. (On review, this function would have better been called
getBaseReward.)
Previously, I mentioned that blockReward was only one part of
Ethereum’s issuance calculation. The other part, called uncleReward,
is described in the next article.
I wish to thank Meriam Zandi for her help in this article.


Ethereum's Issuance: uncleReward
Saturday, August 15, 2020

This the second in a series of two articles detailing Ethereum’s
issuance. Read the ﬁrst part, which discusses the blockReward
calculation. Also, see the code base for the actual code. This article
discusses the uncleReward.

Uncle Reward
Previously, we looked at the ungrammatical second sentence in
Section 11.3 of Ethereum’s Yellow Paper. In the ﬁrst article, we
discussed the ﬁrst half of that sentence concerning blockReward. In
this article we discuss the remaining half of that sentence (shown
below) which details the uncleReward.

It would be more accurate for the Yellow Paper to say that the
“beneﬁciary of each ommer gets rewarded…” as there may be more
than one uncle in the block, each with a diKerent miner. We call this
reward the uncleReward.
The calculation of the uncleReward is a function of the blockNumber
in which it appears, the uncle’s own blockNumber, and the
baseReward from the previous article. A block becomes an uncle
when it shows up in the uncle list of a canonical block.

The uncleReward is calculated from the block’s baseReward
reduced by 1/8 for each block the uncle is distant from the current
block. If an uncle is more than six blocks old, it is not rewarded.

Small Detour
Before we proceed, let’s take a small detour to look at the fourth
sentence of Section 11.3, the only sentence we haven’t looked at yet.
This sentence informs the calculation of minerReward.
If there are collisions of the beneficiary
addresses between ommers and the block (i.e. two
ommers with the same beneficiary address or an
ommer with the same beneficiary address as the
present block), additions are applied
cumulatively.

This simpliﬁes our concerns. It means the winner of a blockReward
may also win one or more uncleRewards in the same block.
Furthermore, a single miner may win both uncleRewards in the same
block.

Back to the Main Discussion…
The uncleReward is described by the following set of three idiotically
complicated equations that we copy verbetum from the paper (163,
164, and 165):

We copied this directly out of the yellow paper so you could see how
stupidly complicated it is. There’s no need for it to be like this as you will
see when we show you the code.
Let’s look at each of these equations separately.
We’ll start with equation 163.

where…

Translating Equation 163 into English, it says “For each uncle in the
block’s uncle list (if any), the account of the uncle’s miner either (a)
remains empty if it was previously empty and the uncleReward is zero,
or (b) the account transitions to a new state (called a-prime) as per
Equation 164”:

where…

This translates to English as, “The state of the account of the uncle’s
miner after the block, a-prime, is the same as it was prior to the block
except for the fact that balance has increased by the value R (the
uncleReward).
And ﬁnally, Equation 165 deﬁnes the uncleReward:

where…

Translating this ﬁnal bit to English: “The uncleReward is the
baseReward lessened by 1/8 for each block that the uncle is distant
from the present block.”
Note that the uncle’s blockNumber will always be less than the block’s
blockNumber, making the value in the innermost parentheses always

negative, and thereby lowering the uncleReward for each block distant
from the front of the chain.

In Code
The above nonsense ﬁts nicely into this TrueBlocks code:

These calculations, at least to me, are so much easier to understand
when seen as code than they are in the Yellow Paper.
Note that getUncleReward is called for each uncle in the block;
therefore, we must ﬁrst call getUncleCount and spin to query for each
uncle block. Then, we can extract the uncle’s block number and
calculate uncleReward.
We accumulate the rewards for the uncles and return that value to the
caller:

Rolling it Up
We’ve shown that the blockReward and the uncleReward are not
that complicated, and that both are functions of blockNumber and
number of uncles (nUncles) in the block.
Until now, we’ve shown this calculation for individual blocks. The only
thing remaining to get total issuance, which is what the twitter dweebs
said was impossible to get, is to roll through each block in the chain,
accumulating the results for presentation.
We use a C++ accumulator class we call CReconciliation. The
name implies an additional feature that we will discuss in the README
where we ‘reconcile’ the values with actual account balances of the
miners queried directly from the chain.
There have been other scripts on Twitter that only model the behavior
of the chain. They don’t validate or audit the balance changes. We
explain this auditing process in the README.
For now, we conclude this article with code that processes the rewards
for each block.

forEveryThing…
The following code shows an excellent feature of TrueBlocks.
TrueBlocks provides many functions that start with forEvery, such as
forEveryBlock, forEveryAddressInBlock, forEveryTransactionInBlock, forEveryIndexFile, etc.
These functions are similar in nature to JavaScript’s map or reduce.
One sends an arbitrary function and an arbitrary data pointer to each
block in the chain (or each transaction, or each address, or whatever)
and allows the function to be applied to the data.
This tends to make the topmost code of a TrueBlocks app very minimal:

One need only deﬁne the function calcRewards, which we’ve done
throughout these two articles.
The above code spins through each block, accumulating, auditing, and
presenting the data for that block, and then reports the results of the
accumulated data to the screen.

One Final Thought
Here, I wish to interject my own opinion: the Yellow Paper is stupidly
over-complicated. The code that generates the Ether Supply is, in the
end, trivial. For some ungodly reason, the speciﬁcation wildly overcomplicates it. I suppose that’s the nature of engineers with a Ph.D.


Every 15 Seconds
Wednesday, December 9, 2020

…a piece of open-source software called an Ethereum node bring
together a random, unordered collection of transactions, and after
throwing out the invalid ones, puts the rest in a well-deﬁned order.
The system then seals this newly ordered list for the rest of human
history (modulo re-orgs) by creating a 32-byte block hash that stands in
a one-to-one correspondence to that ordered list.
The system then quickly moves on to a newly growing list of unordered
transactions, leaving behind it a trail of sealed blocks.
This trail of sealed blocks is the best data the world has ever seen.
Anyone with an Internet connection may download the source code,
build it, and run it themselves, consuming this pristine data without fear
that the data has been manipulated. Nor can the data be denied.
Immutable, uncensorable data is the wellspring from which all web 3.0
dreams ﬂow.
A dream of near perfect data, agreed to world-wide, available to anyone
with no restrictions, impossible to modify. This is what we’re building.
Every wild idea — and I literally mean every single one — that we all fell
in love with when we ﬁrst grokked the blockchain, depend on this
immutable, undeniable, uncensorable data.
Oh, but it were true that people would behave as if they believed this.


Dynamic Traversers in TrueBlocks
Custom Ethereum Extraction for Permissionless Access
Thursday, May 27, 2021

Ethereum nodes come to consensus on a world-wide global ledger of
smart contract invocations every 14 seconds. Everyone know this. This
is what we celebrate about the chain. Not only is this data world-wide
and consented-to, but it is also permissionless. At least that’s what
we’re supposed to believe.
But, is it really permissionless?
I think the answer to that question is a resounding “No.” Ask yourself
how you personally get data from the Ethereum blockchain. There’s
only two answers: (1) you visit a block explorer, or (2) you go to website
built by the developers of the dApp you’re using.
You want to know what’s going on with your ENS names? You go to
http://ens.domain (a website). You want to know how your GitCoin grant
is doing? You go to http://gitcoin.co (another website). You want to
invest in DeFi? You go to Zerion (a website).
I have news for you. This is not permissionless.
Every time you go to a website, you are at the mercy of the people who
run that website. In eUect, you have to say, “Please, oh please, great
data-giver in the sky…please, please give me access to my own data.”

The trouble — and the thing that makes it very obvious that this data is
not permissionless — is that at any time the server (or the people
behind it) can say “No.” They can choose to not give you your own data.
That is the opposite of permissionlessness!
In my mind, permissionlessness means you can say “Fuck you, Mr.
ServerMan. It’s not possible for you to stop me from getting the data I
want, no matter what you do.”
If, in the end, blockchains are not permissionless, they are simply very
expensive Web 2.0 web servers. And if that’s the case, then everything
we’re building will be for naught.

Using TrueBlocks to get Permissionless Data
We assume you’ve read some of our previous articles concerning how
TrueBlocks creates and publishes an index of address appearances.
We won’t rehash that in this article.
To proceed, visit the installation instructions for downloading and
installing TrueBlocks: http://github.com/TrueBlocks/trueblocks-core.
Once you ﬁnished installing, try this:
chifra init

This will initialize the TrueBlocks indexes and downloads the most
recent collection of bloom ﬁlters. Next, run a small test to see if things
are working properly
chifra list trueblocks.eth

This should return a list of about 600 appearances <blockNumber.
transaction_id> pairs. If it doesn’t, revisit the installation
instructions.
Next, run this command:

chifra export trueblocks.eth

This displays the same list of transactions but, for each one, chifra
queries the blockchain and presents the full details of each transaction
to the screen. Try the above command with your own address.
There are many options to chifra export, all of which are presented with
chifra export --help

Traversers
Under the hood of TrueBlocks, there is a piece of code called a
Traverser. A Traverser “traverses” the list of appearances, queries the
chain for transactional detail, and calls into whatever functions are
associated with the Traverser.
In the list option, the function simply prints the <blockNumber,
transaction_id> pair to the screen.
For the export option, the display function prints the entire transaction
and potentially a lot of more data such as articulation and accounting
reconciliations.
Traversers are very general, and because they are C++ classes1, you
may customize them to do practically anything. In fact, this is what this
article is about.

Dynamic Traversers
Recently, we added an amazing new feature to TrueBlocks called
“Dynamic Traversers.” The word “dynamic” has special meaning here,
harkening back to an ancient computing idea called dynamic link

1

2026: These features have since been converted to GoLang.

libraries. A dynamic traverser is a dynamic link library that you may
customize to your heart’s content.
It’s a way of extending the chifra command line (and at the same time,
the TrueBlocks API which mirrors the command line). This allows you to
perform any task you want on any or all transactions from your
accounts.
In the following paragraphs, we describe how you can write and build
you own dynamic traversers.

C++ Code
I know many of you will feel ﬂummoxed at having to use C++. Here’s our
reaction to your concerns: “Get over it or be gone!”

Example Folder
Assuming you’ve downloaded and installed TrueBlocks correctly, you
may create your own Dynamic Traverser by completing these tasks.
We’ll ﬁrst build a simple example and then build a more complicated
traverser that calculates the gas spent (and lost to error) by a given
address.
From the top of the TrueBlocks repo complete these commands:
cd ./src/examples/dynamic
make trav1
cp -p ./libtrav1.dylib $CONFIG/cache/obj/

Where $CONFIG depends on your operating system (on my Mac, the
path is ~/Library/Application Support/TrueBlocks, on Linux,
it’s ~/.local/share/trueblocks). This will build a dynamic link
library and copy it into the TrueBlocks’ cache where we store dynamic
traversers.

Try to run the dynamic traverser
chifra export --load libtrav1.dylib 0xf503...6a78

To run a dynamic traverser, ﬁrst put it in the right folder and the simply
tell chifra to load it.
This should print the hash for each transaction for the given address.
Let’s look at the code. Open the ﬁle trav1.cpp in the example folder.
This is the entirety of the traverser.
That’s all you must do to extract a
perfect list of transaction hashes
from your own accounts running
straight up against your own
Ethereum node. You don’t have to
ask anyone for anything. Just take as
much data as you want.
If you don’t see how powerful this —
how upsetting of the web server
apple cart de-centralized data is — please re-read the above. Getting
permissionless access to globally-agreed-to, near-perfect, world-wide
data only happens if people actually have permissionless access.

Example: GasCost Traverser
Next, we present another example in that same folder called trav2.
Build it with the following commands:
cd ./src/examples/dynamic
make trav2
cp -p ./libtrav2.dylib $CONFIG/cache/objs/
chifra export --load libtrav2.dylib 0xf503...6a78

This prints a table of the gas consumed for each transaction, the
accumulated gas used so far, and the accumulated gas used on
transactions that ended in error. We leave it as an exercise to the reader
to study how this straight-forward code works. I’ll make a few
comments on the code below.
Notice that the makeTraverser function has not changed. Only the
header and display function have changed.
Also notice that the code is very simple and very short. That’s the
beauty of these dynamic traversers. You can write them in a short
afternoon, and they can do amazing things. For example, you’ll notice
we’re accumulating totals as we traverse the data. You could, if you
wished, add subtotals once per day or once per month or whatever

(notice there’s a timestamp on every transaction which makes this
calculation easy).
Another thing to notice is that we’ve derived our own C++ Traverser
class in this second example. This gives you inﬁnite ﬂexibility as far as
the design of your traverser. This is the holy grail of data access to the
Ethereum chain. It’s fast (because it’s fully local and not shared like a
web API); it’s 100% private and completely decentralized; it’s fully
customizable; it’s not a website; and it’s permissionless.

Conclusion
We should avoid building things we don’t want. As Robert Frost would
say, “…path leads on to path…” and turning sooner is much easier than
turning later. If we want the beneﬁts of global permissionless data, we
must build a system that presents global permissionless data. That’s a
simple straight-forward statement of fact. Now we just have to do it.


The Rent is Too Damn High
Fees and Slippage and Such
Tuesday, August 24, 2021

This is the ﬁrst of a two-part series describing my shock at how much it
cost me to swap my own money for my own money using MetaMask.
This article describes using TrueBlocks to ﬁgure out what happened.

Yes, Dear…
My wife spent her entire career as an investment consultant for
institutional investors. She spent a lot of time investigating the all-in
costs, including “hidden” costs associated with moving her client’s
money from one investment vehicle to another. Her clients spent a lot
of money moving money — typically more than they knew.
She’s recently become very interested in DeFi and Blockchain, so every
time I tell her about having done something on Ethereum, she pulls out
her trusty HP 12 C Financial Calculator…

and like Stephen Colbert with her ﬁngers ﬂopping…

she says —“you paid XYZ basis points for that transaction.” I love her,
don’t get me wrong, but enough already! What the hell even is a basis
point?
In my mind, I just say, “How much could it possibly cost?” And push the
button to complete the transaction.

How Much Did That Cost?
The other day I swapped ﬁve wrapped Ether (wETH) for DAI because I
wanted to pay some dude some cash for some shit (don’t ask, it’s none
of your business).
I used MetaMask’s built-in Swap functionality because it’s so easy and
because I didn’t feel like going to UniSwap. I ﬁgured what’s a little
slippage here, a little slippage there…
My wife was out shopping…
It took me 20 seconds. I pushed a few buttons, conﬁrmed a few
settings, and soon MetaMask sent 5.00 wETH out and received
16,003.45 DAI. It was magic.

My wife’s trusty HP 12 C was sitting beside me on the desk, so I picked
it up and divided 16,003.45 by 5.000 to get the price of Ether at that
moment: $3,200.69 US dollars per Ether.
That seemed about right. I had checked the price of Ether earlier in the
day, and it wasn’t too far ob. In fact, the US dollar price had gone up.
All good. Transaction completed.

How Much Did That Cost?
This morning, I woke up and was working on TrueBlocks, my project that
helps me ﬁgure out where all my money is. As I do each morning, I ﬁrst
checked that the TrueBlocks indexer code, Khedra, was working. (It
was.) And then I looked at my accounts. The transaction I had sent a
couple of days earlier was there.
I was stunned to learn that I had paid approximately $200.00 US dollars
for the transfer. $200.00 dollars—that’s real money — not some phonyass bits-and-bytes money. Real money.
I had converted my own money (wETH) into my own money (DAI) and
sent it to my own self, and it costs me $200.00 friggen’ US dollars. What
the hell? That can’t be right. There must be a mistake in TrueBlocks.

Digging…
This is the way TrueBlocks normally shows me a transaction (this is the
one in question). The values are in the number of native tokens held.

As shown above, the transaction was sent from TrueBlocks 2020 Wallet
to the MetaMask Swap Router. There were three assets or tokens

involved in the transaction: ETH, DAI, and wETH (as you would expect
given, I was swapping wETH for DAI on the Ethereum network).
If you look carefully at each row, you’ll notice one of the helpful things
with TrueBlocks is that it reconciles each asset transfer each time its
balance changes. It does this to 18-decimal place accuracy.
The ﬁrst row of the table shows that prior to the transaction, we had
38.220952 ETH in our wallet. We spent 0.004103 ETH on gas; therefore,
we should have ended up with 38.216849 ETH at the end of the
transaction —which is exactly what we did end up with. Thus, the green
checkmark is displayed meaning this transfer of the asset ETH
reconciled.
This reminds me of something my dad said to me the day I got my ﬁrst
bank account: “Son,” he said, “banks are great, but always double
check everything.” (He was overly formal with me when I was young —
don’t get me started — also, now that I think about it, he seems to have
been nearly as anal retentive as I am!)
Each line, and therefore, each asset, reconciles.
You’ll notice, however, that in the above table, there is no connection
between the three diberent assets. We know we swapped assets, but
we really don’t have enough information to understand exactly what
happened. Each row reconciles, but how do the rows interrelate?

How Much Did That Cost?
TrueBlocks has long been able to reconcile the multiple assets that
appear in transactions in their native tokens (units), but recently we
implemented a new feature that allows us to show asset balances and
transactions in US dollars. (Using a new ﬁeld called spotPrice. That
is, the price in US dollars at the exact moment of the transaction.)

Note: The spotPrice is retrieved only from on-chain sources such as
UniSwap and Maker. We use no APIs. This means no rate limiting. And
this means we can price per block — not per minute or per hour or per
day as people who use APIs are forced to do. That’s per block. A true
spotPrice.

I immediately noticed something quite disturbing and ﬁnally came to
understand my wife’s insistence that basis points matter.
Suddenly, what was complicated and hidden when viewing in the native
units became instantly understandable. I had sent myself $16,203.53
worth of wETH and had gotten back $16,003.45 worth of DAI.
What happened to the missing $200.08? Nearly 125 basis points of my
(and my wife’s) money were gone. I had sent my own money to my own
self on my own computer, and it cost me $200.08?
Why is that? And where did that money go?

The Rent is Too Damn High!
Our readers may or may not be aware that
TrueBlocks is not only an indexer, but also a desktop
application, and very importantly also a very capable
command line tool for doing data analytics on any
Ethereum addresses.
This article is already long enough, so we won’t go into detail of how we
discovered exactly where the missing $200.08 went. We’ve written
about that exploration in a follow-on article.
Summary: MetaMask took nearly $150.00 and UniSwap, its liquidity
providers, and the blockchain’s miners took the other $50.00.


The Rent is Too Damn High Part II
Further Explorations into Sending Money to Myself
Tuesday, August 24, 2021

This is the second article in a two-part series using TrueBlocks to
explore the MetaMask Swap Router and how much it costs.

Yes Dear…
In the previous article, I was discussing this transaction (as presented
by the TrueBlocks front end in US dollars at the time of the transaction):

We were discussing the fact that, once one starts looking at one’s own
transactions in US dollars, it becomes obvious where one’s money is
disappearing. What is not so clear, however, is into which tiny spider
holes that money falls.
In this article, we explore this issue more deeply using a helpful aspect
of TrueBlocks software: the command line tools.

Getting Started with the TrueBlocks Command Line
This article is not the place for instruction on how to use a Linux
command line. Nor how to download and install TrueBlocks. We’ll
assume you follow these instructions: https://docs.trueblocks.io/
docs/prologue/installing-trueblocks/. Do so now and return when
you’re ready…

Playing Around…
Assuming you’ve completed the above installation instructions, let’s
try the following command:
chifra names -c trueblocks

You should get results like these (please help us improve our software
by opening an issue if you don’t).
tags
address
name
-------------------------------------------------00-Active
0xf50...79b
TrueBlocks Wallet 0xf503

Copy and paste the address (0xf503017d7baf7fbc0Uf7492b7510
25c6a78179b) and type this command:
chifra transactions 13071849.239 \
--reconcile \
trueblocks.eth \
--fmt json

You’ll get a JSON blob of reconciliation information as we presented in
the previous article:
{

"data": [
{
"blockNumber": 13071849,
"transactionIndex": 239,
"assetAddr":
"0xf503017d7baf7fbc0fff7492b751025c6a78179b",
"assetSymbol": "ETH",
"begBal": "38.220952577792982227",
"gasCostOut": "0.004103324588246856",
"spotPrice": 3240.70642,
"priceSource": "uniswap",
"totalIn": "0.00000000000000000",
"totalOut": "0.004103324588246856",
"amountNet": "-0.004103324588246856",
"endBal": "38.216849253204735371",
"reconciled": true

},
{
"blockNumber": 13071849,
"transactionIndex": 239,
"assetAddr":
"0x6b175474e89094c44da98b954eedeac495271d0f",
"assetSymbol": "DAI",
"begBal": "5188.943282141176871055",
"amountIn": "16003.446778836992430933",
"spotPrice": 1,
"priceSource": "stable-coin",
"endBalCalc": "21192.390060978169301988",
"totalIn": "16003.446778836992430933",
"totalOut": "",
"amountNet": "16003.446778836992430933",
"endBal": "21192.390060978169301988",
"reconciled": true
},
{
"blockNumber": 13071849,
"transactionIndex": 239,
"assetAddr":
"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
"assetSymbol": "WETH",
"begBal": "44.165326226182661143",
"amountOut": "5.0000000",
"spotPrice": 3240.70642,
"priceSource": "uniswap",
"endBalCalc": "39.165326226182661143",
"totalOut": "5.0000000",
"totalOutLessGas": "5.0000000",
"amountNet": "-5.0000000",
"endBal": "39.165326226182661143",
"reconciled": true
}
]
}

There’s so much to notice here, it’s embarrassing…

•

You’ll notice the transactionIndex shows this transaction to be
the 239th transaction of block 13,071,849. Note this, we will need
it below.

•

You’ll notice the last item in each record says reconciled: true.
This is where the green reconciliation dot we show on the front end
comes from.

•

You’ll notice each record has a ﬁeld called priceSource and for
all but one of those records, that value says, uniswap. This records
the source from which we get our pricing of the given token in US
dollars — the other value for priceSource is stable-coin which
means we can (if our users wish) peg this value to $1.00, thereby
avoiding a query to the node. We will write a future article
discussing exactly how we use UniSwap and stable coins soon.

•

You’ll question pegging DAI to 1.00 US dollar is legitimate. (We think
it is, and if our users do not, they can always disable this.)

•

You’ll notice we calculate spotPrice only to ﬁve decimal points —
we ﬁgure what’s the point of going further when the ﬁnal reporting
is in US dollars anyway?

•

You’ll notice that the data denominates the values in their own
native token. For example, the wETH reconciliation (the third
record), says we transacted 5.0000 wETH. We do this to provide
maximum ﬂexibility to the consumer of the reconciliation. Think of
these values as count of the number of tokens transacted.

Getting US Dollars from SpotPrice and Token Count
So, how do we get US dollars?
Simple — for each data item we multiply the spotPrice per token by
the number of tokens (nTokens), which is the value in the data itself.

5.000 wETH times a spotPrice of 3,240.70642 = 16203.5321 which is
the amount of wETH denominated in US dollars that we wanted to
swap.
16003.446778836992430933 DAI times 1.00 US dollar per DAI is the
number of DAI we received from the swap.
•

And, ﬁnally, you’ll notice that there is an ‘unaccounted for’
diUerence between what we sent and what we got in return.

Some people call this slippage. This is not surprising, but it’s annoying
and…“The rent is too high!”
I’m incredulous — it cost me 200.085321163007569 US dollars to swap
money with myself in a completely programmatic way?
“The freaking rent is too damn high!”

Where Exactly Did the Money Go?
Being anal retentive, like my father, it wasn’t enough for me to know that
I had spent $200.00. I wanted to know exactly where it went.
Let’s go back to the command line. Do this command:
chifra transactions 13071849.239 --uniq

This will return all the unique addresses that appear in the given
transaction. You should get these results (we’ve cleaned this up. Your
result may vary).
appearance
addr
---------------------------------------------13071849.239
0xf503...179b
13071849.239
0x881d...300c
13071849.239
0xc02a...6cc2
13071849.239
0x6b17...1d0f
13071849.239
0x11e...c6fb
13071849.239
0x74de...6631

13071849.239
13071849.239
13071849.239
13071849.239
13071849.239

0x6059...a270
0xdef1...5eff
0x3d1d...c9e5
0x47f0...b33c
0xfffd...8d25

Let’s see if we can name them. Run chifra names against each address
to get these results:
addr
name
----------------------------------------------0x11ed…c6fb Metamask Fees
0x6b17…1d0f Dai Stablecoin
0x6059…a270 Uniswap V3: DAI
0x881d…300c Metamask Swap Router
0xc02a…6cc2 Wrapped Ether
0xdef1…5eff 0x Exchange Proxy
0xf503…179b TrueBlocks Wallet
No results for 0x74de…6631
No results for 0x3d1d…c9e5
No results for 0x47f0…b33c
No results for 0xfffd…8d25

•

You’ll notice that four of the addresses are un-named. That’s not a
problem. Obviously, TrueBlocks can’t know every name of every
address. No harm. No foul. If you look more closely at the named
addresses, you’ll see that all the addresses but one make sense.

•

Obviously, the TrueBlocks wallet is present in the transaction.

•

Obviously also, DAI and wETH are present — this is what we were
swapping.

•

The three routers (proxies?) also make sense (Uniswap, Metamask
Swap Router, 0xExchange). One must complete the swap
somewhere, after all.

But what about the ﬁrst item in the list: MetaMask Fees? Let’s see what
we can see. Run this command:

chifra transactions 13071849.239 \
--reconcile 0x11ed...c6fb --fmt json

This is the same command we ran above, but now we’re showing the
reconciliation for the MetaMask Fee wallet address. This command
returns these results:
{

"data": [
{
"blockNumber": 13071849,
"transactionIndex": 239,
"assetAddr":
"0x11ededebf63bef0ea2d2d071bdf88f71543ec6fb",
"assetSymbol": "ETH",
"begBal": "1250.563288183878021420",
"amountIn": "",
"internalIn": "0.043750000000000000",
"spotPrice": 3240.70642,
"priceSource": "uniswap",
"totalIn": "0.043750000000000000",
"totalOut": "",
"amountNet": "",
"endBal": "1250.586983614757050983",
"reconciled": true
}
]
}

•

You’ll notice that the spotPrice is the same as the previous data.
This is to be expected as this is the same transaction.

•

The amount the MetaMask Fee address took in during this
transaction equals the spotPrice (3,240.70642 US dollars per
ETH) multiplied by 0.043750000000000000 ETH which equates to
141.780905875 US dollars.

I don’t know about you, but that feels like a lot of money to me. My wife,
after some Stephen Colberic ﬁnger tapping tells me that’s 87.5 basis

points (we conﬁrmed this number on MetaMask’s website). She says
she would have advised her clients to ﬁnd a cheaper solution.
We did a quick and dirty calculation on the web to see how much it
would cost us to exchange US dollars to Euro in the regular, oldfashioned, 200 year old FAIT system, and it was much lower. WTF!

What About the Rest?
The above calculations leave 58.304415288007569 unaccounted for.
Where did that money go? Looking at the above image, you can see we
paid $13.30 to the miner in gas fees, leaving 45.004415288007569 that
went to UniSwap and its liquidity providers. I’ll leave it as an exercise for
the interested reader to dig deeper.
In any case, “The rent is too damn high!”

Conclusion
Fine. We admit that we punted on the ﬁnal resolution above, but you get
the idea: TrueBlocks is amazing. Okay. I’m done. This article is too damn
long.


Adventures in Difficulty Bombing
Predicting the Future for the Arrow Glacier Fork
Sunday, October 10, 2021

This article is an exercise in predicting the e1ect of EIP 4345. That EIP
sets back the di1iculty bomb at an as-yet-undetermined point in
December of 2021. Our goal is to see if we can help identify when and
how far back to set the bomb so we can expect it to explode again in
May of 2022.

Some Preliminaries
I’ve written a lot about this subject previously: (1) An article about the
di1iculty calculation called It’s Not the Di1icult, (2) An article about
di1using the di1iculty bomb called A Method to Di1use the Di1iculty
Bomb, and some other older articles.
The di1iculty calculation consists of two parts as explained in the ﬁrst
article: Part A (the adjustment) and Part B (the bomb).

The adjustment is the every-block adjustment that keeps block times
hovering around 13.5 seconds. This part works perfectly and would, in
the absence of the bomb, keep block times constant. We’re not
interested in Part A.
The bomb, or Path B, is a step-function that doubles every 100,000
blocks. It lies mostly unnoticed until it starts to “explode” and once it
starts to explode, it explodes rapidly (as 2! tends to do). We’re
interested in this part.
The value of the bomb depends solely on fake_period (called 𝑛" in
the second part of the equation above). fake_period depends solely
on fake_block which depends solely on the current block, called
real_block below, and offset.

A Note of Caution
I am but one person. I’ve written this article using publicly available
data and my own increasingly tired mind. What I say here is obviously
open for discussion, but I disclaim any responsibility for what is
presented here. Use the following with caution and treat everything I
say with skepticism.

Actual vs. Theoretical Data
I’m an engineer, not a mathematician, therefore I’m much more
interested in actual data than I am in what the math predicts. In the
following, I present a prediction of how the data will change based on
actual data. My prediction is based on the equations, but it is just a
simple, straightforward extrapolation in an Excel spreadsheet.
My discussion is based on three simple observations:
•

Part A (the adjustment) works perfectly and produces, on average,
13.3-second blocks.

•

Part B (the bomb) also works perfectly; is independent of Part A and
only increases block times.

•

If we ignore the bomb, our prediction will be conservative.

In other words, if we make a prediction ignoring the e1ect of the bomb,
our prediction will be “early” because the bomb only makes blocks
slower. This leaves additional breathing room if we’re wrong.
The ﬁrst spreadsheet takes the current block (13,391,127 at time of
writing) and extends it forward by 13.3 seconds per block until block
14,000,000:

Looking more closely, we see the average block times since block
12,000,000 has been increasing (notwithstanding the decreases in
August and September). Of course, average block times increase more
quickly as the bomb explodes.
Again, to be conservative, we’ve chosen 13.3 seconds to estimate
future blocks and, remember, we’re ignoring the bomb. When the bomb
starts to show, average block times increase, therefore this estimate
produces dates earlier than what will happen in reality. (In other words,
block 14,000,000 will happen “not earlier than” January 10th, 2022.)

When to Fork
The ﬁrst question we need to consider is, “When should we fork?”
This decision depends, in my opinion, only on the fake_period. The
question is, “Should we fork at block 13,700,050, 13,800,050 or
13,900,050?” (Adding ‘50’ here ensures there are no o1-by-one errors.)
This next spreadsheet shows the fake_block calculation.

Here we’ve taken the estimates of when each real_block occurs and
subtracted offset (i.e. the setback from previous forks) to arrive at
fake_block and, by simple division, fake_period.
Fake_period is what we’re interested in since the bomb depends
solely on that value. From previous work (see the articles above), we
believe the bomb starts to show when fake_period gets to 41–42, but
not earlier. In other words, Part B starts to dominate Part A around
fake_period 41.

I won’t explain why I say “but not earlier” in the previous paragraph.
Su1ice to say that the bomb only increases block times and, in the
absence of the bomb, Part A keeps block times hovering at 13.3
seconds. In other words, average block times stay above 13.3 seconds,

and if they go below, Part A adjusts them to bring them back up— said
more succinctly — Part A works.
Given the above, I suggest that we fork any time after block 13,800,000.
That is, around the middle of December. I would target a speciﬁc block
as opposed to a date, say 13,850,000. The “pain point” (that is,
noticeably slower block times) will start around mid-January, so there’s
some room for error.

How Much Should We OAset
The other question we need to consider is, “How many blocks should
we o1set?
As we’ve said above, offset determines fake_block, which
determines fake_period, which determines the bomb. So, in the
following, we’ll focus on offset and see what we can learn.

The Suggested Value in EIP 4354
First, we’ll look at the suggested value of offset in the EIP. Again, we’ll
produce a simple spreadsheet assuming 13.3 second average block
times. Also, we will ignore the bomb, with the understanding that the
bomb can only increase block times and therefore would shift the
calculation into the future. We’ll use 13,800,050 for the fork_block
of Arrow Glacier (the name of the next hard fork) and 10,500,000 for the
offset.

This seems to say that, if we set back by 10,500,00 blocks (the
suggested value), the earliest time that blocks will start to slow down
will be the middle of April (fake_period 41). By the middle of May,
blocks will start to slow down noticeably (fake_period 43).
The highest fake_period we’ve gotten previously was just prior to
Byzantium, where we got to fake_period 43. The slowdown at that
time was quite noticeable — in the multi-second range.

Speeding Up or Delaying the Date of the Fork
As a short diversion, I was interested to see what happens if we either
(a) make the hard fork earlier or (b) make the hard fork later.
The next spreadsheet shows the result of that estimate — I found this a
bit surprising —it seems to have no e1ect on the next hard fork. Upon
further reﬂection, it makes sense, though. The only value determining

fake_period is offset. Other than the slowing block times prior to
the fork if we were to delay, forking earlier or later has no e1ect on when
the next bomb explodes (i.e. the one in April/May).

You can tell from the above chart that doing the Arrow Glacier hard fork
earlier or later — if we use an o1set of 10,500,00 blocks — has no e1ect
on when the next bomb starts to make itself felt.
Conclusion: We can do the Arrow Glacier hard fork whenever is
comfortable.

How Much Should We OAset?
The answer to the next question, “How far back should we make the
o1set?”, depends on how hard you want to push on the Core Devs in
May. If you want a heavy push — so that the entire world is complaining
about slow blocks — use a smaller number. If you want a light push —
as in “We better do something soon, but we have some time” — use a
larger number.
If you use 10,500,00 for o1set, you may be creating a heavy push. You
can expect to see noticeably slower blocks by late April (on the order of

a second). But the trouble with the bomb is that once it starts to
explode, it explodes fast.
“Noticeable” turns into “Quite Noticeable” which turns into
“Concerning” and then “Holy Shit” and then “What the Fuck” in the
manner of four to six weeks. It is true that each 100,000-block period
takes longer to process (because the bomb doubles, block times
increase increasingly quickly), but once the explosions come, they
come quickly. See the article mentioned above about the Byzantium
bomb. The bomb starts slowly, but then it truly does explode.
In this ﬁnal spreadsheet, I’ll propose that we use an offset of
10,700,000. I also propose that we fork sometime after block
13,800,050. This leaves some room to breathe now and pushes the
estimate of the next bomb to the middle of May which, again, is
conservative, but realistic.

Conclusion: Set the o1set higher than 10,500,000.

Summary
Decide how hard you want to push on the core devs. If you want to push
hard, set the o1set to 10,500,000. If you want to leave yourselves

breathing room, set the o1set to 10,700,000. Compromise somewhere
in the middle.
The decision about “when” the chain is forked has no e1ect on the next
bomb (that is, it has no e1ect on the May bomb). The only e1ect of
delaying or speeding up the Arrow Glacier hard fork is how long block
times will get before the fork. This is because only offset e1ects
fake_period and only fake_period a1ects bomb.


Recipe: Calling Smart Contracts
Wednesday, October 27, 2021

This is a simple document describing how to use the TrueBlocks
chifra command to query a smart contract from your command line.
When combined with other features of chifra it allows for a very easy
way to build a history of various values on a smart contract. For
example, perhaps you want to print out the entire history of the supply
of a given token.
In this article, we ﬁrst show the commands available and explain them
as simply as we can and then show an example of doing exactly that.

The Commands
There are two basic forms of the command. Both are part of the chifra
state tool. Both use the chifra state --call option.
The --call option takes one parameter, but it’s a complicated
parameter that we present it on the following page as a graphic since
we think that’s the easiest way to describe it.
First version: Use the smart contract’s address, a function signature,
and optional padded bytes to make the call into the smart contract.
This version is for a more technical user.

Second version: Use the smart contract’s address, a four-byte, and the
optional padded bytes. Oddly, this version is also for the more technical
user.

Using the Call Option
Here’s a very quick example of using the --call option. There are other
ways to do this, of course, our amazing Dynamic Traverser feature for
example, but that requires C++ programming. This method runs from a
command line.

Install TrueBlocks
Blah, blah, blah. See the many articles we’ve written describing
installing TrueBlocks at http://trueblocks.io.

Running the Commands
First, we’ll pick an ERC20 token at random:
chifra names ERC20 | head -614 | tail -8 | head -1

When we do this with the current version of chifra we get
tag
address
name
symbol decimals
-------------------------------------------------ERC20 0x2287b...f49bf SoundMoney SDM
18

Let’s use this. To get a list of every transaction in which this address has
appeared, do this:
chifra list 0x2287b...f49bf

This command does three things:
•

It visits every chunk of the TrueBlocks Unchained Index and ﬁnds
every appearance of this address anywhere on the chain,

•

Depending on your setup, chifra list downloads and caches
(i.e. pins on IPFS) those portions of the index that you’ve expressed
interest in. In this way, you become part of the solution to making
access to the Ethereum data permissionless.

•

It caches the list of appearances so the next time you run it, it’s
almost instantaneous.
chifra list 0x2287b...f49bf --count

shows us that there were 453 transactions as of the date of writing of
this article.
Note that as of this point, chifra has not hit against the Ethereum node.
It has hit against IPFS to download the index, but it hasn’t yet extracted
any data from the chain. We’ll do that next.

Extracting and Caching Ethereum Data
The next command extracts the actual transactional data from the
chain for the above address:
chifra export 0x2287b...f49bf

You’ll notice that this takes longer to run than the list command. That’s
because we’re hitting the RPC and extracting the transactional data
from the node. Because this is a bit slower, we o[er the --cache_txs
option which, you will see if you use it, speeds up subsequent queries
by 30–40 times.
time chifra export --cache_txs 0x2287b...f49bf

takes 14.43 seconds on our machine, whereas the same command
after caching takes only 0.5 of one second. 30 times faster.
Moving on. Our goal is to list the historical changes in the total supply
of this ERC20.

Using Chifra Call
See the above description of the --call option. We’ll need it. First, run
this command:
chifra abis 0x2287...49bf

This shows all the commands available on this smart contract. We
want the one called totalSupply which returns this data:
name
four-byte
signature
-------------------------------------------totalSupply
0x18160ddd
totalSupply()

We will use the chifra state --call routine on the latest block:
chifra state --call <addr> | totalSupply()

When you run this, it tells you that you need a return value. chifra
needs this so it can decipher the result. The proper command is:
chifra state --call <addr> | totalSupply()(uint)

This returns the following results:

bn
addr
signature
encoding
value
-------------------------------------------------13496062 <addr>
totalSupply() 0x18160ddd 429

So, there were 429 tokens available at block 1,349,062. Interesting…

All Appearances
We’ll ﬁnish now with this series of commands left unexplained. We
invite the reader to ﬁgure out what’s going on for themselves:
# Store the blocks for this address in blocks.txt
chifra list <addr> --no_header | \
cut -f2 | \
sort -u -n >blocks.txt
# Count unique blocks (we get 152 blocks)
wc blocks.txt
# Cache the extracted transactions
chifra export --cache_txs <addr>
# Use blocks.txt to call the totalSupply routine
# on the address at every block
chifra state --call "<addr>|totalSupply()()" \
--file blocks.txt

Permissionless access to Ethereum data. Pretty cool in my opinion.


Commanding the Line
Monday, November 22, 2021

Many of us live and work on the Linux command line. I’ve had a
residence in that neighborhood since 1986. For most of that time, the
command line commanded me. But eventually, I started commanding
the command line.
In this article, I will share with you an example of commanding the
command line. I’m going to take the command line by its neck and
shake it until it does what I want. The target of my aggression will be a
TrueBlocks command line provided by one of our users in a support
request. Thanks Pintail.

Using Pipes
First, I’ll show you a command our user shared with us:
chifra export --cache --articulate \
--first_block 12550000 \
--last_block 13520000 \
--fmt csv \
<addr> | \
grep 'approve' | \
tr -d '"' | \
cut -d ',' -f1,2,5,6,13-20 \
> dai_uniswapv2_approvals.csv

Translated into English, this says:

Show every transaction against the DAI Stable Coin (0x6b17…71d0f),
starting at block 12,550,000 and continuing until block 13,520,000
inclusive reporting the results in CSV format. Furthermore, articulate
those results and cache the query as you go. Then…shove the results
through grep to pick oN only the token approval transactions; remove
quotes; and, using comma as the ﬁeld delimiter, extract the ﬁrst,
second, ﬁfth, sixth and thirteenth through the 20th ﬁelds (blockNumber,
transactionIndex, from, to, compressedTx); and, ﬁnally, store the
results in a ﬁle called dai_uniswapv2_approvals.csv.
See — I told you. The command line tends to command the user. That’s
one crazy-assed command line.

Can We Improve on This?
Yes. There’s are ﬁve things going on in the above command line:
•

Producing the list of transactions with chifra export

•

Filtering the records for approvals with grep

•

Extracting only ﬁve of the many ﬁelds in each record with cut

•

Removing quotes with tr

•

Storing results in a ﬁle with a redirect (>)

To start, we remove any part of the above command that is irrelevant for
this article (the --cache options and the --first_block and -last_block options). We’ll keep the --articulate option. We’ll
also remove the piping, grepping, transforming, and re-direction.
So, we’re left with this basic command. It’s telling chifra to show the
address’s history and articulate the transactions.
chifra export --fmt csv --articulate <addr>

Now, we’ll add back options we removed. Note: this is a large amount
of data. You can limit the results with --first_block / -last_block.

Filtering Records for Approvals
If one runs the command chifra export --help, you’ll see this
summary:
chifra export [flags] <address> [address...] \
[topics...] [fourbytes...]

It’s possible to provide not only addresses to export, but also one or
more topics and/or one or more fourbyte signatures to ﬁlter the results.
To ﬁnd which fourbyte signature we’re interested in, we can run this
command:
chifra abis <addr> | grep approve

which returns:
fourbyte
name
signature
-------------------------------------------------0x095ea7b3
approve
approve(address,uint256)

Use that fourbyte in our growing command:
chifra export --fmt csv --articulate <addr> \
0x095ea7b3

This command shows every approve transaction that was every run
against DAI (assuming <addr> is DAI’s address).
The above command replaces the need for grep in the original
command. This is also faster than the original command since many
transactions (those that aren’t approve) are never generated. There’s no
need to grep them out.

Field Selection
Note: The following section describes an experimental feature of
TrueBlocks. I may be removed or altered in future releases. Use with
caution.
The original command uses tr to remove quotes from the output and
cut to select ﬁelds 1, 2, 5, 6, and 13–20 from the data. This can be
accomplished using an undocumented (and therefore experimental)
feature of TrueBlocks called “Display Strings.” I won’t explain Display
Strings, because this article is already getting long, but they allow you
to select ﬁelds and (in ebect) describe the exact output of the data you
want.
Run this command:
DISPLAY_FMT="[{BLOCKNUM}]" \
chifra export \
--fmt txt \
--articulate \
<addr> \
0x095ea7b3 2>/dev/null

Note: I’ve changed the export format from csv to txt to make the
displayed data clearer. This also ﬁxes an issue with the original
command. The compressedTx ﬁeld contains commas, so using CSV
doesn’t really work as the ﬁelds get garbled.
The above command produces row upon row showing only the
blockNumber of each transaction.
Run this command:
DISPLAY_FMT="[{BLOCKNUM}][{TRANSACTIONINDEX}]" \
chifra export \
--fmt csv \
--articulate \
<addr> \
0x095ea7b3 2>/dev/null

Row upon row of blockNumber and transactionIndex. You get the
idea.
You can tell TrueBlocks which ﬁelds to display using an environment
variable.
Fields 1, 2, 5, 6, and 13–20 correspond to blockNum, transactionIndex, from, to, and compressedTx.
Extending this to the ﬁelds we want, the command becomes:
DISPLAY_FMT="[{BLOCKNUM}]\t[{TRANSACTIONINDEX}]\
[{FROM}]\t[{TO}]\t[{COMPRESSEDTX}]" \
chifra export \
--articulate \
<addr> \
0x095ea7b3 2>/dev/null

Note: You must ﬁx the above command to put it on a single line.
If you run the above command trough | head -1 you will see that only
the ﬁelds we’re interested in are being exported:
blocknum transactionIndex from to compressedTx

No quotes. That’s what we want. This is command is faster because you
don’t have to pipe extraneous data through the cut command. When
you’re doing command-line data processing, every little tweak counts.

Storing the Results in a File
So far, we’ve eliminated grep, cut, and tr. Can we eliminate the
redirect? Yes.
Chifra has an option called --output <filename> which lets the
user specify the destination of the resulting data. The ﬁnal, improved
command, then, is this:

DISPLAY_FMT="[{BLOCKNUM}]\t[{TRANSACTIONINDEX}]\
[{FROM}]\t[{TO}]\t[{COMPRESSEDTX}]" \
chifra export \
--articulate \
<addr> \
0x095ea7b3 2>/dev/null \
--output dai_uniswapv2_approvals.csv

Note: Edited this to be on a single line.

Performance
Here are the results of testing the two commands for speed:

9.79 seconds for the original command. 8.89 seconds with the new
command. A speedup of 10%. Not as good as I would have liked, but
better than nothing.
The above timing tests were run against 1,360 transactions. The DAI
smart contract has (as of the time of this writing) 12,956,972
transactions. We found this number with:
chifra list --count 0x6b17…1d0f

That means that to extract every transaction, we would have to wait
9,527.18 times longer than our above tests. For the ﬁrst command, this
would be 9.79 seconds * 9527.18 = 25.9 hours. For the second
command, it would have taken 8.89 * 9527.18 = 23.52 hours.
A savings of more than two hours. Take yourself out to dinner on the
diberence.


Recipe: forEveryChain
Come Visit with Us for a While
Wednesday, February 2, 2022

Have you ever had the desire to look at every trace of every transaction
of every block on every blockchain? Pull up a chair and visit for a while.

It’s a Multi-Chain World
I’m sure you’re aware of the impending multi-chain universe. You know
the one. Where all sorts of little baby communities are going to be doing
all sorts of little baby community-like things on little baby blockchains.
Charities over here. Games down here. Ultra sophisticated blockchain
ﬁnance application up there.
If you know me, you know I’m a decentralization maximalist. That’s why
it’s always perplexed me how the above multi-chain world can possibly
ever remain decentralized.
This week, we took a stab at answering that question by adding a new
feature to TrueBlocks that we call multi-chain support. This feature
seems relatively simple on its face — the user simply adds --chain
<chain_name> to his/her chifra command lines — but it opens a
whole new world of possibilities (and headaches) — and…IT WORKS
ON A LAPTOP (← this should be our new logo!)
In this article, I want to show you an example of one of the new
possibilities. I’ll start by explaining an existing bit of example code from
the TrueBlocks repo called forEveryTrace.

This example visits every block in the Ethereum Mainnet, and from
there every transaction in every block, and ﬁnally every trace in every
transaction in every block.
I’ll conclude the article by talking about a new piece of code called
forEveryChain which came into existence literally ten minutes ago and
inspired this post.

ForEveryTrace
We assume a certain amount of familiarity with TrueBlocks and how to
install it and get it running. We also assume you have access to your
own node software. We highly recommend Erigon running locally, and
while the following code will work with Infura, you will deﬁnitely get rate
limited. Don’t say we didn’t warn you.
In our repo there’s a folder called ./examples and in that folder,
there’s an example called ./forEvery. The contents of the code in the
folder is on the next page.
If you’re familiar with C++ code at all, you’ll notice the main function at
line 34 is simple. Initializing the library, calling a single function, and
cleaning up.
While simple, that one function packs a wallop. It does exactly as it
says, it visits every block in the entire blockchain it’s conﬁgured for. (By
default, this means Ethereum Mainnet, but we’ll get back to this in a
minute.)

What does line 38 do?
•

For every block in the blockchain, call the function visitBlock
(line 29) and pass it a pointer to the counter.

What does the function visitBlock do?
•

For every transaction in the current block, call the function
visitTransaction (line 24) passing in the same pointer.

What does the function visitTransaction do?

•

For every trace in the current transaction call the function
visitTrace (line 16) passing in the same pointer.

And ﬁnally, what does visitTrace do?
•

It keeps a count of the total number of traces in the entire
blockchain and prints out each trace.

I’m sorry, but it’s almost impossible for this code to be simpler. If you
don’t understand this, stop reading now and go back to CIS 101 and pay
better attention.

An Aside on Laptops
If you look carefully at this code, you’ll understand a little bit of why
TrueBlocks still works on the laptop even after more than 14,000,000
Mainnet blocks.
We’re not extracting any permanent data from the chain.
While TrueBlocks provides a binary cache that speeds up processing
many times over compared to the raw node, the cache is optional. The
cache is only employed if the end-user instructs us to use the cache.
We leave the decision of whether to cache to you. This is markedly
diferent from any other solution we’re aware of and it’s why our system
still works on a laptop.
If you want to cache the entire blockchain because that’s your
requirement — the ﬁrst thing you must do is get a really, really large hard
drive — and then alter the above code to write each block, transaction,
or trace to either our built-in cache or your database.
This is why we have a problem with solutions such as Amber Data and
Etherscan. They’ve chosen to extract and cache every single piece of
data on the entire blockchain prior to the developer’s decision if it’s
pertinent. This creates many, many terabytes of data which means their

solutions will never work on small machines such as a laptop and
therefore will never be truly decentralized. Enough soapbox — back to
the article.

For Every Trace on Any Chain in the Visible Universe
As I mentioned above, this week we added a multi-chain feature to
TrueBlocks which means we can process queries for any EVM-based
blockchain with the same equal ease as we process Mainnet.
I was inspired to write this post because I thought of the following
modiﬁcation to the above example.

What Does this Do?
It visits every trace in every transaction in every block on every
conﬁgured blockchain ON MY LAPTOP!


Specification for the Unchained Index
Thursday, June 23, 2022

Introduction
Immutable data — such as that produced by blockchains — and
content-addressable storage — such as IPFS — have gotten married,
and they’ve had a baby called the “Unchained Index.”
Immutable data and content-addressable storage are deeply
connected. After all, without a suitable storage medium for immutable
data, how can it possibly be immutable? And, if one modiﬁes
immutable data — ﬁrst, it’s not immutable, and second, its location on
IPFS changes. The two concepts are as connected as the front and back
sides of a piece of paper. One cannot pull them apart — and even if one
were able to pull them apart — rending the paper’s front from its back,
one would end up with two, slightly thinner, pieces of paper. There’s no
way around it.
This document [2026: the Speciﬁcation which is at our website]
describes the Unchained Index, a computer system that purposefully
takes advantage of this tight coupling between immutable data and
content-addressable storage. The mechanisms described in this paper
apply to any immutable data (for example, any time-ordered log), but
the examples herein focus on the Ethereum blockchain’s Mainnet.

Do You Want to Read More?
So begins the most detailed description of our work to date. The
“Technical Speciﬁcation of the Unchained Index.” The full document is
available on our website: https://trueblocks.io. Finally, a blockchain
indexing solution on right side of the coming paradigm shift — join us.


TrueBlocks Covalent Comparison
Tuesday, September 20, 2022

Introduction
This article highlights key advantages of using TrueBlocks, a local-ﬁrst
solution to blockchain data indexing, by comparing our software to the
popular Covalent API. The advantages of TrueBlocks are:
Speed
– faster than other methods
Accuracy
- more complete
Privacy
– censorship-resistant and private
Flexibility – programmable platform for dApps

We conducted our analysis by running TrueBlocks locally on a “beefy”
Mac laptop (Apple M1 Max, 64GB memory, Monterey version 12.2.1).
On that same machine, we ran the Erigon Ethereum Mainnet client as
an archive node.
We queried the Covalent APIs described in this documentation from
that same machine. Note that Covalent is a “shared API” and therefore
necessarily rate limited.

Our Process
We used the TrueBlocks APIs to extract transactional histories for 5,000
randomly selected Ethereum addresses. We used the corresponding
APIs from Covalent to pull the same histories. We then compared the
results.

The entire process, along with links to the shell scripts we used and
information about how to access the data we produced, is detailed on
our GitHub account (https://github.com/TrueBlocks).
We discuss each of the four advantages below and then conclude with
a discussion of potential sources of error.

Advantages
Speed: TrueBlocks is Faster
We ran two sets of data extraction. One against TrueBlocks, the other
against Covalent. In the following table, we present the amount of time
taken to complete these two presumably identical tasks.
source
nAddrs
time
addrs/min
-------------------------------------------------Covalent
5,000
743.31 mins
6.7
TrueBlocks
5,000
336.29 mins
14.9

Notes: When extracting from Covalent, we were forced to slow down
the processing, otherwise our requests timed out. After experimenting,
we decided to add a one-second delay to each Covalent request to
avoid this time out. This added an additional 83.3 minutes (5,000 / 60
seconds) to the processing. While there are much better ways of
backing oY an API, we felt that the additional programming eYort
required was not warranted, especially given that we did not have to do
the same thing for TrueBlocks.
Upshot: TrueBlocks is nearly twice as fast as Covalent.

Accuracy: TrueBlocks is more complete
After doing the data extraction from each source, we compared the
results. The results surprised even us.

We queried 5,000 randomly selected addresses. In no case did
Covalent return more results than TrueBlocks. For 3,174 of the
addresses, TrueBlocks returned more results than Covalent.
A summary of results is presented here:
more from
more from
nAddrs
Covalent
TrueBlocks
-------------------------------------------------5,000
3,174

That same information viewed by the total number of transactions
returned is here:
Covalent
TrueBlocks
Diff
Material
-------------------------------------------------1,336,508
1,534,997
198,489
45,328

Note: In our parlance, a material transaction is one in which the Ether
balance of the addresses changed because of the transaction.
TrueBlocks can expand this deﬁnition further to include any ERC20
token balance, however, we chose not to do that in this analysis. If we
had, the above results would have been skewed even more in our favor.
Upshot: TrueBlocks ﬁnds more data than other methods.

Why is TrueBlocks More Accurate?
How is this even possible? It turns out to be rather simple. TrueBlocks
digs more deeply into the data. This ability to “dig more deeply” is not
un-related (by that I mean it is very related) to the fact that, being fully
local, TrueBlocks is not rate limited. We discuss this exact issue in the
ﬁrst version of our Speciﬁcation for the Unchained Index.
We studied the material transactions (the 45,328). We looked at the
function calls that constituted those transactions — the ﬁrst four bytes
of their input data.

We found 436 diYerent four-byte patterns, of which 253 were known to
us (that is, we were able to decode the four byte into function names).
The function calls can be found in the reason folder of the data store.
We summarize the known functions in this word cloud:

I discuss the most obvious function (donate) below, but look at the
words: donate, splitTransfer, withdraw, placeBid, atomicMatch,
addLiquidity. These words scream, “transfer of value.” Covalent does
not return them. No wonder per-block accounting on the Mainnet
doesn’t work for shit?
Upshot: One cannot even hope to do perfect accounting on an 18decimal-place accurate ledger if one is missing transactions.

Privacy: TrueBlocks is Local First
The processing done by TrueBlocks, including accessing the Erigon
node, is run locally (on a laptop!). We are happily hidden behind a
ﬁrewall. We call this mode of operation, “running behind the node.” We
consider the node a shield, not only from incomplete or inaccurate

data, but also as protection against censorship and prying eyes.
TrueBlocks asks no third party for any data at all (other than ABIs :-).
None of this is true of Covalent.
Upshot: TrueBlocks enables permissionless access to better data
faster and is perfectly private.

Flexibility: TrueBlocks is a Platform
TrueBlocks is an open-source software package consisting of many
components. There is a collection of command line tools, an API server,
a docker package, an address monitoring system, an indexer, and an
increasingly robust collection of well-documented GoLang packages
and SDKs. You can program with it, as we’ve done to complete this
study. Anything is possible.
Upshot: TrueBlocks is a platform of tools and libraries, for individual
users and developers that works locally, accurately, privately, and
ﬂexibly.

Sources of Error
To complete the article, we present a few possible sources of error in
our analysis. We welcome all comments intended to help us improve
our work.

Predominance of GitCoin Related Addresses in the Dataset
Many of the addresses we studied have interacted with the GitCoin
smart contracts. For this reason, many of the “missing” function calls
were donate (about 71%). While it is accurate to say that TrueBlocks
returned these transactions and Covalent did not, it is also accurate to
say that Covalent could easily add this “special case” to their
processing. The important point I wish to make, however, is that
TrueBlocks doesn’t have any “special cases.” TrueBlocks purposefully
processes the data unaware of its meaning. This is the main reason we

ﬁnd transaction while other methods do not. We call this aspect of our
work, “not suYering from the long-tail problem.”
Upshot: Our data skews towards a certain type of address.

Misuse of Covalent APIs
We did our best to study the Covalent APIs and use them properly, but
we may have missed something. Please let us know if we did. Perhaps
Covalent has for-pay API endpoints that deliver better data. But, if that’s
the case, ask yourself if that is in keeping with the web3 ethos.
Upshot: We may have misused the public Covalent APIs.

Focus on Smaller Addresses
We purposefully limited our study to addresses with less than 6,000
transactions. While TrueBlocks can easily handle addresses with many
more transactions, Covalent imposes rate limits. We felt that if we
queried for addresses with more than 6,000 transactions, we would
encounter two problems: (1) Covalent would become too slow to be
practicable, (2) Covalent would ban us from their site.
Upshot: Rate limiting sucks. If you’re running your own node (and
TrueBlocks), you will not be rate limited.

Block Range Limit
We purposefully excluded transactions prior to block 3,000,000 from
our analysis. We found, surprisingly, that Covalent had done so for
many of the transactions from the October 2016 DDos attack on
Ethereum. While it is justiﬁable for them to have done so, as many of
those transactions were not “material,” it reminds us that centralized
APIs can become the arbiters, without our permission (or knowledge),
of what data they show their users— this is not true with TrueBlocks. If
we had included those records, the results of this study would have
been signiﬁcantly more skewed in our favor.

Upshot: Some older blocks were ignored in this analysis. If they had not
been, TrueBlocks would have compared even more favorably.

Reproducing Results
Some people say “blockchain data is not rocket science.” Others say it
is. We say it should at least be science — the ledger is, after all, 18decimal-place accurate and immutable. Everything should be reproducible. Here’s how to do that:

Using this Repo
•

Clone this repo (https://github.com/TrueBlocks/covalent). Change
into the ./exploration/ covalent folder.

•

Visit the Covalent website and get an API key. Put the key (alone) in
a ﬁle called .env in the covalent folder.

•

Run ./init to set up the folders and build a simple postprocessing tool (requires go language version 1.18 or later).

•

Collect your own list of addresses (or use ours — see the ﬁle called
addresses.txt in the folder). Create a shell script
called ./download (or use ours). Make the shell script repeatedly
call into ./download.1 (as ours does) to process each address.

•

Run the ./download command.

The results of that command will be placed in appropriate-named
subfolders in a folder called ./store.

Obtaining the Data
We tried very hard to make the data reproducible. All you need is a
Covalent API key, a locally running version of the TrueBlocks, a locally
running version of Erigon, and a few days or weeks to wait for the data

to download. Alternatively, you may download our data directly from
IPFS using the following commands:
CID="QmQK2wEwMH1e6qnMtNbBqDqDJG6WYFR09132oBVyq"
DOMAIN="iofs.unchainedindex.io/ipfs"
curl -o store.tar.gz https://$DOMAIN/$CID
gunzip store.tar.gz
tar -xvf store.tar
rm -f store.tar

Conclusion
Join the paradigm shift. Run your own node. Index your own data. Speed
up your life! We’ve shown that local-ﬁrst software can be faster than an
API even on smaller hardware. We’ve shown that local-ﬁrst software
can be private. But, most importantly, we’ve shown that one can “dig
deeper.” This is because of speed. What are you waiting for? Join us in
our mission.


Better Accounting for Blockchains
Towards 18-Decimal-Place Accurate Token Accounting
Monday, November 21, 2022

TrueBlocks is pleased to announce the release of our newest version:
v0.43.0. This is our best version yet and includes much better ERC20
token accounting. In this article, we present a few thoughts on what’s
changed and where we’re headed from here.

A Note on the Word “Accounting”
TrueBlocks has always been about accounting…sort of. It depends on
what you mean by the word.
The word “accounting” is loaded with hundreds of years of baggage.
Ever since Luca Pacioli invented double-entry bookkeeping in 1494,
“accounting” has had a primarily ﬁnancial sense.
This is still true, but with the advent of a reproducible, immutable,
worldwide-accessible ledger, we would like to expand that word to
include EVM smart contract “state.” Our goal at TrueBlocks is perfect
18-decimal place “accounting” for smart contract “state,” not just the
ﬁnancial aspects of a smart contract.
What does this mean? We’re working that out…but you wouldn’t be too
far oW if the words “permissionlessly reproduce every state change oW
chain” ﬂoated into your head.
This article, however, is about the ﬁnancial sense of the word.

Where We Were
Before the current version, TrueBlocks could reconcile, with high
accuracy, the ETH accounting for any Ethereum address. It still does
that, but with this version, it brings the same high-accuracy accounting
to ERC20 tokens.
What does “reconcile” mean exactly?

•

Inter-transaction agreement: At each appearance of an address
on the chain, an oW-chain balance that was calculated at the
previous appearance must equal the on-chain balance prior to the
current transaction. In other words, we still have the same number
of tokens that we had the last time we looked. This ensures that
there are no missing appearances.

•

Intra-transaction agreement: The on-chain beginning balance of
the current appearance (which is the same as the oW-chain balance
of the previous appearance — see A) plus all incoming amounts,
minus all the outgoing amounts, is equal to the ending balance at

the current appearance. This ensures we’ve noted all movements
of value.
•

Completeness: Each of the two aspects of the reconciliation is
true for all assets transferred during the appearance.

In eWect, the history of the token holdings of any EVM-based addresses
is an intertwined time-ordered ledger.
Note: The above three rules apply to each individual asset (ETH and
ERC20 tokens) separately for each appearance.
Note: In the above discussion, “on-chain balance” means we query the
blockchain using eth_getBalance for ETH and the ERC20’s
getBalance routine for tokens.
This is not brain science. In fact, anyone who’s familiar with accounting
(Luca?) will recognize the above as a Balance Sheet and a Proﬁt and
Loss statement.
Please don’t miss the fact that this is done every twelve seconds — and
we’re doing it on a laptop computer (that’s TrueBlocks real import — full
decentralization).
So, where does the complication come from?
Three places: (1) you must make sure you have every appearance, (2)
you must know where to look in the data for every asset transfer, and (3)
you must deal with quirks of the RPC and buggy smart contracts.
We address these three issues next.

Make Sure You Have Every Appearance
The ﬁrst thing one must do if one wishes to make 18-decimal-place
accurate accounting is to get an accurate list of every transaction in
which a particular address appears.

Getting a list of appearances is basically impossible with the current
RPC implementations. This is exactly why we built the Unchained
Index. (We won’t rehash what the Unchained Index is here, having
written about it many times elsewhere).
Using chifra, our command line tool, we run
chifra list <address>

which generates a complete list of appearances for the address. Next,
we run chifra export <address> --accounting which
produces the accounting output we’re looking for.

Know Where to Look for Asset Transfers
How do we know where to look for the value transfers, given the
appearances?
It is important to note is that a single transaction may transfer many
diWerent assets. Have a look at any swap on UniSwap for an example.
Many times multiple assets are transferred in the same transaction.
Most indexers look only at events, such as Mint, Burn, and Withdraw.
Which ones matter? It turns out only the Transfer events matter. The
other events may add color, but they do not indicate a transfer of value.
Focusing just on ETH transfers, we can see why the accounting can get
quite complicated. There is 12 diWerent ways ETH can enter or leave an
address (on the Ethereum Mainnet — other chains may diWer).
Moreover, for each value transfer, the disposition of the money depends
on which address is being accounted for. (For example, an EOA-to-EOA
ETH transfer may be “incoming” if we’re accounting for the recipient but
“outgoing” if we’re accounting for the sender.)
These are eight ways in which ETH can enter an account:

Vector
amountIn
internalIn
selfDestructIn
minerBaseRewardIn
minerTxFeeIn
minerNephewRewardIn
minerUncleRewardIn
prefundIn
totalIn

Description
A transaction's top-line value (if to)
Internal inﬂow through contract calls
Value received on self-destruct
Value received as the winning miner
Transaction fees as the winning miner
Value received for producing a nephew
Value received for producing an uncle
Value received during the genesis block
Sum of the above

These are four ways for ETH to leave an account:
Vector
amountOut
internalOut
selfDestructOut
gasOut
totalOut

Description
A transaction's top-line value (if from)
Internal outﬂow through contract calls
Value sent on self-destruct
Gas spent during the transaction
Sum of the above

Once we have this information, the reconciliation is really kind of easy:
Vector
amountNet
prevBal
begBal
endBal
expectedBal
reconciled

Description
totalIn – totalOut
The calculated balance at the last
appearance of the address
The balance queried from the node at
transaction block minus one
The balance queried from the node at
the current transaction
prevBal + amountNet
If expectedBal equals endBal and
prevBal equals begBal, then true

ERC20 token accounting is identical to ETH, but simpler because there
are fewer ways for the value to transfer. There are only two certain ways
for a token to enter or leave an account. Again, it depends on whether
you’re accounting for the sender or the recipient.
recognized
Implied

The movement us in a Transfer event
The movement is not in a Transfer event

Recognized transfers are, of course, simple. The ERC20 Transfer
event is clearly deﬁned. The trouble is the implied transfers. These
come in various ﬂavors, as detailed below. What these types of
transfers do not do is make note of a transfer in a standardized way.

Deal with the Shortcomings of the RPC
In the table below, you will see that our reconciliations fail in 0.1%
(that’s 1 in one thousand, in case you were wondering). Why do they
fail? Because of implied transfers of ERC20 tokens.
The reasons why this happens are too numerous to list, but they fall into
three basic categories.
The ﬁrst type of failed reconciliation is “missing appearances.” The
Unchained Index solves this problem as well as any solution we know
of. See other articles in this collection about comparisons with industry
standard API providers.
Think about the problem of missing appearances for a second. What
are the chances of reconciling an account if there are missing
transactions? Zero.
The second type of failed reconciliations we’ve seen is non-compliant
ERC-20 transfers. Recently, certain people have been taking advantage
of the fact that a smart contract may generate Transfer events even
when no value has been transferred. They do this as part of a ﬁshing

expedition. Here’s an example: https://etherscan.io/tx/0x506e7978ba
52886681b75797e4403579ba703b5f9df576a34602ada1709085fb.
There are 1,000 Transfer events showing but no state changes. This
is spam, intended to “ﬁsh” people into revealing their private keys.
The third type of failed reconciliation happens due to incorrect
accounting by a smart contract. For example, this transaction
(https://etherscan.io/tx/0x634799410165000edaf1b1e8e5e8055b39
cdd534d3c3dc973885d39adb5d888) does not balance if one
considers the Transfer events only (as one should be able to do if the
smart contract was well written). The accumulated interest, which, in
fact, is a transfer of value, does not generate a Transfer event. Nor
does the existing Transfer event reﬂect this value transfer. Without
Transfer events detailing each value transfer, there is no way to
account for a token without resorting to the per-smart-contract
context, which is not scalable nor supportable in the long run. Our
solution to this problem is detailed elsewhere.
The fourth diWiculty one encounters when trying to do oW-chain
accounting is that the APIs are slow and rate limited. The RPC itself is
slow. Especially when one queries for account balances. Apis are
worse. We solved this problem by providing a local-ﬁrst, client-side
cache. You can read about this on our website.
Taking the above into consideration, we are now able to account for
nearly 99.9% of all transactions we’ve tested.

Where We’re Going
With this release, we’ve duggen deeper into ERC20 token accounting.
That’s a good thing, but it comes with some downsides.
The main downside is that our processing has slowed down (although,
see the table below). Don’t worry. We have it covered.

One of the reasons we made these changes was to prepare for a port to
GoLang. We made major progress on this front behind the scenes. The
code from the previous version had the two aspects of the
reconciliation calculation tightly coupled. Above, we call these two
aspects (A) Inter-transaction agreement and (B) Intra-transaction
agreement.
With this release, we’ve separated the two processes. By tightly
coupling the two functions, we had inadvertently made it impossible to
do either process in parallel. This is no longer true.
In the GoLang code (v1.0.0), which we hope to ﬁnishing by March, we
will be able to do the balance queries of type A and the data acquisition
portion of part B currently. We hope this will regain the performance
losses we’ve seen in this release.
In the table below, we present a comparison of the current version of
our code with the previous version. Notice that when one looks at the
actual “useful” reconciliations, the processing turns out to be faster.
This is due to a lot of wasted querying in the previous version.

Upshot: The new version is better and poised to get even better when
we ﬁnish our port.

Conclusion
We’re enamored with the idea of perfect accounting, not only the
ﬁnancial aspects of smart contracts but for all “state changes” for any
address, “smart” or not.
Why are we doing this? Why do we try to create perfect, 18-decimalplace accurate accounting for any address? Well, I guess I’d say that —
even though it may seem like a high hill to climb — because it’s there.


Why We're Building TrueBlocks
It’s Time to Make Amends
Tuesday, January 3, 2023

This week, we’re going to ﬁnalize the second draft of our “Speciﬁcation
for the Unchained Index.” We’ve written the conclusion and thought
we’d share it, as it is the best explanation yet of why we’re building what
we’re building. This is copied verbatim from the soon-to-be-released
second draft of the spec.

Conclusion to Speciﬁcation for the Unchained Index
…[t]here is one thing left to explain. Why did we build the Unchained
Index? SuJice it to say that we want to “re-change the world.” One of us
(Rush) played a vanishingly small part in “changing the world” back in
the 1990s during the birth of the world-wide-web. He wishes to make
amends.
Humans have more than 25 years of experience living in a world
dominated by the web. We are only now coming to understand the
eJects of decisions made in the 1990s when the web stack was being
formed. Back then, everyone was wildly enthusiastic about the
architectural decisions being made, but we missed something that now
seems obvious. We missed the fact that censorable, capturable
worldwide information systems are — not surprisingly — censorable
and capturable. And — again, not surprisingly — that at this point, these
systems have been fully captured.

By Whom?
Q-Anon? An evil cabal of evil people? Advertisers? Who cares? It
doesn’t matter.
What does matter is that the web has been captured by people with
incentives opposite to those of us all. The web has been captured by
takers. It’s time for the sheeple to take back their own system.
The architectural decisions made in the 1990s and early 2000s were
constrained by the capabilities of the hardware that was then available.
Rate limiting a website was born out of a necessity to protect web
servers from overuse. Logging in to a website was born out of rate
limiting. Privacy invasion ﬂows naturally (perhaps as an unintended
consequence) of logging in.
Only retroactively does this seem obvious.

The World Doesn’t have to be That Way…
Since that time — 25 years ago — the capabilities of computer hardware
advanced by 10 orders of magnitude. The software stack — consisting
of web servers and user-based web clients — has not kept pace. Except
for the fact that the system has gotten more invasive. The massive gains
in computer hardware capability have gone almost exclusively toward
the evil cabal of evil people1 own the machines.
It is our belief that the advent of distributed blockchain software (such
as Bitcoin), which allows any user to see anything they want, was a ﬁrst
example of a system that takes advantage of the 10 orders of magnitude
increase in hardware capability on behalf of the sheeple. Blockchains
pull back power from the massive worldwide information and privacy
invasion machine known as the Internet…supposedly…

1

I’m joking here.

TrueBlocks is an attempt to show by example that indexing and
searching this already decentralized data must also be fully
decentralized. We believe that the only reason existing blockchain data
access systems (such as Etherscan, Covalent, QuickNode, Dune, and
The Graph — all of which are based on old-fashioned client-server
architectures) exist is that people are lazy. People don’t yet realize that
it is possible to avoid web-based API architectures. They don’t yet
understand that it is possible to build a system that is not capturable
and protects privacy as opposed to invading it.
We’ve tried to make the case that if indexing were part of the node
software’s processing and if the indexing natively produced data ready
to be distributed through a content-addressable data store such as
IPFS, the cost of publishing the same would be negligible, particularly
if the system automatically invites end users to share in the burden of
carrying the data. TrueBlocks accomplishes this through the two
related processes of “pinning by default” and creating a “naturally
sharded and shared chunked index.”
We look forward to continuing our work and expanding our (and your)
understanding of this new paradigm. Hopefully, this time we can all
build a worldwide information system that doesn’t suck.


Recipe: Monthly Token Balances
Just How Bad was my Year?
Friday, January 6, 2023

Digital blockchain accounting drives me crazy.
In this article, I’m going to try to produce a
report showing the end-of-month balances
for every token I’ve ever owned for one of my
wallets. This is not an easy process as you’ll
see, but I’m doing it permissionlessly — for
free — on my own laptop.

Introduction
To get started:
•

Get access to a locally running node such as Erigon running on
dAppNode.1

•

Download and install TrueBlocks. (Figure it out — there’s a ton of
information about how to install TrueBlocks on our GitHub.)

•

Pick an address to study. We choose trueblocks.eth.

1

2026: I’m playing a game here, obviously. I’m acting like this is easy. It’s not.

Proﬁt and Loss Statements
We’ll start by looking at a single TrueBlocks Statement for a single
transaction. A Statement is the same thing as a bank statement you
might get from your bank. It’s a P&L statement for the transaction
relative to a given address.
To get started, run this command:
chifra transactions \
--account_for trueblocks.eth \
--fmt json \
0x1a898c5448b37f693343917ea40b7ad1c43b28a4
ddd37af1bd6d0bb4a0c99891

This produces the following data:
{

}

"blockNumber": 8854723,
"transactionIndex": 61,
"logIndex": 0,
"transactionHash":
"0x1a898c5448b37f693343917ea40b7ad1c43b2
8a4ddd37af1bd6d0bb4a0c99891",
"timestamp": 1572639538,
"assetAddr":
"0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
"assetSymbol": "WEI",
"decimals": 18,
"spotPrice": 181.51,
"accountedFor":
"0xf503017d7baf7fbc0fff7492b751025c6a78179b",
"sender":
"0xbb984f85bd52d78eb2fbf2c5598bae7abb98c5bc",
"recipient":
"0xf503017d7baf7fbc0fff7492b751025c6a78179b",
"begBal": "0",
"amountNet": "5000000000000000000",
"endBal": "5000000000000000000",
"reconciled": true,

If you look carefully enough, you can see has begBal, totalIn,
totalOut, endBal, and reconciled, all of which one needs if one
wishes to do accounting. You’ll also notice that reconciled is true.
This means two things: (1) the ending balance at the previous
transaction equals the beginning balance at this transaction, and (2)
the beginning balance plus the inﬂowing amount plus the outﬂowing
amount equals the ending balance. Both must be true for the
reconciliation to be true.
Notice also there are two additional ﬁelds in the data, assetAddr, and
spotPrice. We’re interested in assetAddr for now. If we scanned the
entire history of our address and extracted a unique list of
assetAddrs, we would have a list of every asset we’ve ever owned.
We can do this with this command:
chifra export trueblocks.eth \
--accounting --statements \
--cache --fmt json | \
grep assetAddr | \
sort -u >tokens_list

We use the export command because we want “every transaction in the
address’s history.” The grep component extracts only the assetAddr
ﬁeld. The sort -u removes duplicate records. The results are stored in
tokens_list.
This command takes a while, so we also use the --cache option in
case we need to run it again. The next time we run it, it will be much,
much faster (perhaps as much as 100 times faster).
The previous command generates data similar to the following:

"assetAddr": "0xaec2e87e0a235266d9c5adc9deb4b2e2",
"assetAddr": "0xc02aaa39b223fe8d0a0e5c4f27ead908",
"assetAddr": "0xc12d1c73ee7dc3615ba4e37e4abfdbdd",
"assetAddr": "0xc18360217d8f7ab5e7c516566761ea12",
"assetAddr": "0xc4199fb6ffdb30a829614beca030f904",
"assetAddr": "0xc9da518db3abdb90a82c4d1838b7cf9b",
...

Let’s count how many tokens diYerent tokens we’ve owned:
cat tokens_list | wc

At the time of this writing, this wallet held, at some point in its history,
75 diYerent tokens. That’s not how many tokens we have now. It‘s how
many tokens were ever held over the address’s history.

A Note on Unreconciled Transactions
You may notice that the above command generates “unreconciled
token transfer” messages. This happens when TrueBlocks fails to
reconcile a transfer.
We got 29 such notices at the time of this writing. There were 2,418
reconciliations, of which 29 failed. That means TrueBlocks is 98.8%
accurate. If you know anything about blockchain accounting, you know
this is quite good. Plus, we have a list of failed reconciliations for later
debugging. Plus, remember, it’s on a laptop and it’s free.

Getting Token Balances
Back to the task at hand…
We want to get token balances for each of the given tokens. We can do
this using the ﬁle we created. Let’s clean the data a bit by pulling out
just the addresses:
cat tokens_list | cut -d'"' -f4 | tee token_addrs

Now we can get current balances for these tokens for our address. We
need to build a shell script. Do this:
cat token_addrs | sed \
's/^/chifra tokens --no_zero --no_header /' | \
sed 's/$/ trueblocks.eth/' | \
grep -v \
0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee | \
tee get_balances.sh

This creates a long list of lines of the form:
chifra tokens --no_zero --no_header \
0xc4199fb6ffdb30a829614beca030f9042f1c3992 \
trueblocks.eth
chifra tokens --no_zero --no_header \
0xc9da518db3abdb90a82c4d1838b7cf9b0c945e7e \
trueblocks.eth
...

which will report the current balance of a given token for
trueblocks.eth to the screen. The option --no_zero squelches zero
balance tokens. The --no_header option cleans up the display by
removing the header from each call.
Let’s run just one of these lines to see what ﬁelds we get:
chifra tokens --no_zero \
0xc4199fb6ffdb30a829614beca030f9042f1c3992 \
trueblocks.eth | head -1

We get these results:
holder address name symbol decimals balance

To make the display a bit cleaner, we can pull oY address, name, and
balance. The holder ﬁeld is always trueblocks.eth. Address is the
token. Let’s run the entire script (and clean up the data a bit further):
source get_balances.sh 2>/dev/null | \
cut -f2,3,6 | tee balances

Notice even the smallest balances show. This is a list of all tokens
currently owned by trueblocks.eth.

Monthly Balances?
This article was supposed to be about getting monthly balances. It’s
getting a bit long, so I’ll try to ﬁnish quickly.
Many chifra commands allow you to enter block ranges, which causes
chifra to repeatedly call the same command on blocks in that range.
Ranges may be speciﬁed with a skip parameter. Like this:
chifra when firstBlock-lastBlock:monthly

which will generate a list of the blocks most recent just prior to the
beginning of each month in the range.
To illustrate, let’s ﬁnd the ﬁrst and last blocks in 2022:
chifra when 2022-01-01 2023-01-01

which returns
blockNumber
13916165
16308189

timestamp
1640995189
1672531199

date
2021-12-31 23:59:49 UTC
2022-12-31 23:59:59 UTC

This means we can use
chifra when 13916165-16308189:monthly

to get the ﬁrst block of each month in 2022. We can use this knowledge
to adjust the above shell script:

cat token_addrs | sed \
's/^/chifra tokens --no_zero --no_header /' | \
sed 's/$/ trueblocks.eth \
13916165-16308189:monthly/' | \
grep -v 0xeeeeeeeeeeeeeeeeeeeeeeeee | \
tee get_balances.sh | tee get_balances.sh

and then
source get_balances.sh 2>/dev/null | \
cut -f1,3,4,7 | tee balances

produces what we’re looking for
13717846 0x0000000000085d4780b731 TrueUSD "11.18"
13916165 0x0000000000085d4780b731 TrueUSD "11.18"
14116761 0x0000000000085d4780b731 TrueUSD "11.18"
14297758 0x0000000000085d4780b731 TrueUSD "11.18"
14497033 0x0000000000085d4780b731 TrueUSD "11.18"
14688629 0x0000000000085d4780b731 TrueUSD "11.18"
14881676 0x0000000000085d4780b731 TrueUSD "11.18"
15053226 0x0000000000085d4780b731 TrueUSD "11.18"
15253305 0x0000000000085d4780b731 TrueUSD "11.18"
15449617 0x0000000000085d4780b731 TrueUSD "11.18"
15649594 0x0000000000085d4780b731 TrueUSD "11.18"
15871479 0x0000000000085d4780b731 TrueUSD "11.18"
16086233 0x0000000000085d4780b731 TrueUSD "11.18"
16308189 0x0000000000085d4780b731 TrueUSD "11.18"
13717846 0x03ab458634910aad20ef5f Rai
"9.5"
13916165 0x03ab458634910aad20ef5f Rai
"9.5"
14116761 0x03ab458634910aad20ef5f Rai
"9.5"
14297758 0x03ab458634910aad20ef5f Rai
"9.5"
14497033 0x03ab458634910aad20ef5f Rai
"9.5"
14688629 0x03ab458634910aad20ef5f Rai
"9.5"
14881676 0x03ab458634910aad20ef5f Rai
"9.5"
15053226 0x03ab458634910aad20ef5f Rai
"9.5"
15253305 0x03ab458634910aad20ef5f Rai
"9.5"
15449617 0x03ab458634910aad20ef5f Rai
"9.5"
15649594 0x03ab458634910aad20ef5f Rai
"9.5"
15871479 0x03ab458634910aad20ef5f Rai
"9.5"
16086233 0x03ab458634910aad20ef5f Rai
"9.5"
16308189 0x03ab458634910aad20ef5f Rai
"9.5"
13717846 0xde30da39c46104798bb5aa Gitcoin "24509"

13916165 0xde30da39c46104798bb5aa Gitcoin "24614"
14116761 0xde30da39c46104798bb5aa Gitcoin "24821"
14297758 0xde30da39c46104798bb5aa Gitcoin "22630"
14497033 0xde30da39c46104798bb5aa Gitcoin "18038"
14688629 0xde30da39c46104798bb5aa Gitcoin "16679"
14881676 0xde30da39c46104798bb5aa Gitcoin "12809"
15253305 0xde30da39c46104798bb5aa Gitcoin "10763"
15449617 0xde30da39c46104798bb5aa Gitcoin "10763"
15649594 0xde30da39c46104798bb5aa Gitcoin "10763"
15871479 0xde30da39c46104798bb5aa Gitcoin "10763"
16086233 0xde30da39c46104798bb5aa Gitcoin "10763"
16308189 0xde30da39c46104798bb5aa Gitcoin "0"

Thanks for getting this far. The above is interesting if you look closely
enough. I told you. I’m freaking crazed about blockchain accounting…

Conclusion
There’s a lot left to do, but I’m glad I went
through this exercise as it’s taught me a
huge amount about what we need to do
before we can get what we really want. A
quick hint for the attentive:
chifra export trueblocks.eth --accounting \
--statements --fmt json | grep spotPrice


Recipe: Factories
Getting a List of All Contracts Created by an Address
Friday, January 6, 2023

In this simple and short article, we show you how to use TrueBlocks, the
Unchained Index, and chifra to get a list of every smart contract
created by an address (including smart contracts that create other
smart contracts, that is, factories).
We’ll dive in without fanfare. You can ﬁnd installation instructions for
TrueBlocks on GitHub.

Getting Contract Creations from an EOA
Let’s start by ﬁnding all smart contracts created by our own address,
trueblocks.eth. Run the following command.
chifra export --traces --factory trueblocks.eth

Note that we had to add the --trace option since only through traces
can we ﬁnd every smart contract creation. Looking at receipts is not
enough, as there’s only room for one created address on a receipt. Note
also that chifra now supports ENS.
The above command ﬁnds 22 smart contract creations by this address,
but — my goodness — what took so long? Let’s do some performance
testing.

Improving Performance Through Local Caching
First, let’s clean the cache of this address and all its artifacts. This will
make the testing fair:
chifra monitors --decache trueblocks.eth

The above command removes the existing caches (if any) including
cached blocks, transactions, traces, or reconciliations in this address’s
transactional history.
Now that we’ve cleaned the cache, we’re ready to do some
performance testing. Run this command:
time chifra export --traces --factory --cache \
trueblocks.eth

This command takes 656.74 seconds for 2,886 transactions. That’s 10
minutes and 57 seconds. 4.39 transactions per second. That’s very
slow. Why?
Because, in order to get traces, Erigon must replay the transactions
which means replaying the blocks. That’s one of the tradeoZs Erigon
made to ensure a smaller disc footprint. It’s worth it, but can we do
better? Yes. We can.
Notice that I included --cache in the above command. This causes
any queried data from the node to be cached locally in TrueBlocks’
cache. It does not speed up the current query, but it has a huge eZect
on the speed of subsequent queries (as is generally true of caches).
This becomes obvious if we re-run the command. Let’s do that:
time chifra export --traces --factory –cache \
trueblocks.eth

This took 3.01 seconds. 958.8 transactions per second. That’s more
than 210 times faster than the non-cached run.
The upshot? If you’re doing repeated analysis on an address, use the
--cache option, and you do not need to worry about using it
repeatedly. Immutable data doesn’t change. There’s never a reason to
rewrite.
Now that we have caching under our belt, let’s run the --factory
extraction against an interesting address. How about UniSwap?

How Many Contracts has Uniswap Created?
First, we need to ﬁnd Uniswap’s address. Perhaps the names command
will help. Let’s try:
chifra export uniswap.eth --traces --factory

Be warned, though. This takes a very long time the ﬁrst time you run it.
Perhaps we can make that a bit more palatable.

When?
To make this more palatable, we can search for only the contracts
created by UniSwap during the month of December 2022. We need
some block numbers:
chifra when 2022-12-01 2023-01-01

returns
blockNumber
timestamp
date
-------------------------------------------------16086233
1669852799
2022-11-30 23:59:59 UTC
16308189
1672531199
2022-12-31 23:59:59 UTC

These are the ﬁrst block numbers in the month of December and the
last block number in the month of December 2022. chifra when
returns the block number most recent to the given date and time (and
visa-versa).
Adjusting the above command to show only contracts created during
December 2022, we get:
chifra export --traces –factory --no_header \
--first_block 16086234 \
--last_block 16308189 \
--cache uniswap.eth | \
cut -f1 | \
tee blocks-only

Using --cache speeds things up the next time we look at this month.
--no_header does not generate the header rows. Notice also, we’ve
extracted the blocks alone into a ﬁle called blocks-only.
We can use that ﬁle for the next command:
cat blocks-only | xargs chifra export \
uniswap.eth -–cache -–traces -–factory \
-–file blocks_only

The above command delivers a list of every smart contract created
through the UniSwap factory contract during the month of December
2022. Get that from your API (for free).

Conclusion
I wanted to share a recipe for ﬁnding contracts created by a factory
contract, so that’s what I did. Have fun.


Recipe: Simple Speedup
Friday, May 5, 2023

This article describes a simple way to to speed up rendering transaction
histories with TrueBlocks. Do this:
chifra list --no_header trueblocks.eth | \
cut -f2,3 | \
tr '\t' '.' | \
tee tx_list

That is, put the transaction identiﬁers into a ﬁle. And then this:
chifra transactions --cache --file tx_list

That is, instead of calling into chifra transactions once for each
transaction of interest, read the transactions list from a ﬁle.
This has the e?ect of calling into the RPC for all the transactions with
single command line invocation instead of loading chifra from the
command line for each individual transaction.
In a simple test we just did, this is signiﬁcantly faster.

A Naive Test
To test the speedup, we ran chifra list trueblocks.eth and got
around 2,230 transaction identiﬁers. When then ran this command to
create a simple shell script:
cat tx_list | \
sed 's/^/chifra transactions /' | \
tee script1.sh

This produced a shell script that calls chifra transactions once
for each transaction identiﬁer. We then timed it with:
time source script1.sh

and got the following results:
Executed in 237.37 secs
usr time
76.14 secs
0.28 secs
sys time
26.12 secs
2.21 secs

75.86 secs
23.91 secs

Faster
Next, we ran this test (Notice the use of the --cache):
time chifra transactions --cache --file tx_list

This returns:
Executed in 218.02 secs
usr time 105.69 secs
0.13 millis
sys time
22.59 secs
4.84 millis

105.69 secs
22.58 secs

Slightly faster, but not as much as we’d like.
But if we run it again:
Executed in 132.51 secs
usr time
42.51 secs
0.17 millis
sys time
5.11 secs
4.78 millis

42.51 secs
5.10 secs

About twice as fast as the original. Local caching and avoiding repeated
calls into the operating system matters.

Just for Fun
Just for fun, we ran the command on the following line:

chifra transactions --no_header trueblocks.eth | \
cut -f3 | \
cut -f1 -d: | \
uniq -c | \
tee hourly.txt

And made this meager chart (sorry, we’re tired, and it’s late).


The Prisoner's Dilemma on Crack
Why I Love the Blockchain
Thursday, October 26, 2023

When I ﬁrst got involved in the blockchain space, I was surprised by how
frequently people would end their origin stories by saying something
like, “And I literally dropped everything I was doing, quit my job, and
never looked back…”
That same thing happened to me.
In the summer of 2013, I became obsessed with Bitcoin. Not because I
believed it would work–I didn’t. In fact, I was convinced it wouldn’t
work — I wanted to prove it was bullshit.
I spent the whole summer trying to prove its bullshittyness to myself. I
couldn’t, and in January of 2014, after eight months of indecision, I
bought my ﬁrst few Bitcoins. Two months later, Mt. Gox crashed to the
ground.
I was right. I knew it was bullshit.
A year and a half later, in the Fall of 2015, I stumbled upon Ethereum. A
programmable blockchain–better than Bitcoin. I quit my job and never
looked back. “Hello, Boss? Yeah. I quit.”
In this post, I hope to explain why.

What Excites Me Most About Ethereum?
The two most exciting things about Ethereum were its transparency and
its permissionlessness.
Remember, I had spent an entire summer studying Bitcoin, and, back
then, nearly everything I read about claimed Bitcoin was a solution to
the Prisoner’s Dilemma.
I won’t explain the Prisoner’s Dilemma, but I can summarize it thus:
Humans know (inherently) that working together provides a more
optimal solution than “going it alone.” I believe this is true of everyone.
We’re safer in groups than we are roaming an open savanna alone.
Those poor unfortunates who believed otherwise were all eaten by
lions. Lesson one: The group dynamic produces the best outcome.
Not being stupid and recognizing the beneﬁt of the group dynamic, we
also — at the same time — recognize that it is possible to defect from the
system and gain an advantage, and that, if someone does that, we are
at a disadvantage. Lesson two: The group dynamic is great. If someone
defects from the system, they gain an advantage to our possible
disadvantage.
Lesson three: Our only logical action, knowing of the possibility of
defection, is to defect ourselves, thereby collapsing the group dynamic.
The group dynamic becomes impossible to attain.
That’s the Prisoner’s Dilemma. Bitcoin solved that because, as a miner,
it is illogical to defect. If someone tries to defect, they lose. We’ve
tamed the lions.

Why the Current System Sucks So Much
Our current systems (ﬁnancial, governmental, interpersonal) suck
because we have not tamed the lions. The group dynamic has

collapsed. This is due to a god-zillion factors (that’s ‘1’ with a Jesus
number of zeros after it). To me, the two most prominent factors are:
•

Our groups have gotten too large — we don’t know what anyone else
is doing, and

•

Those in whom we’ve placed our trust (bankers, politicians, civic
leaders, multinational corporations) are taking advantage of their
privileged (centralized) positions.

Transparency
Blockchains (more precisely immutable ledgers) have solve the ﬁrst of
these two problems. Blockchains are transparent. We can, if we design
our smart contract systems well enough, know exactly what everyone
else is doing. It’s impossible to lie about what happened on a smart
contract.
Bitcoin (and Ethereum by extension) is the ﬁrst human-built system that
allows us to purposefully engineer relationships that cannot be taken
advantage of. Nothing can be hidden. We can trust these systems
because we don’t have to trust these systems.

Permissionlessness
The second main primary reason why our current systems suck is
because some people are closer to the center than others. (Leaving us
poor unfortunates to fend for ourselves on the edge). And once these
people are in the center, they take advantage — directly destroying the
more beneﬁcial group dynamic.
The fact that anyone with a computer can join a blockchain ecosystem
and have exactly as much access to the information as any other
person is at the heart of why blockchains work. People understand that

they are fair and wholeheartedly join because they see the greater
good.

Why Did Digital Money Fail Before Bitcoin?
There were many attempts at making digital money throughout the
1990s and early 2000s. I wasn’t directly involved in those edorts, but I
remember them. Everyone back then thought, “Of course, money
should be digital. Everything else is digital. Why not money?”
But why did it take so long to work? People knew about public key
cryptography since 1976 (or earlier). I think the answer lies in Bitcoin’s
permissionlessness.
We are the survivors from the savannah. We were the ones who did not
get eaten. We can clearly see and easily imagine when others have an
advantage over us. This is why we stayed alive. If we didn’t know this,
we would be a pile of scattered bones on the outskirts of a lion’s den.

What Excited Me Most About Ethereum
Bitcoin showed us we can solve the Prisoner’s Dilemma. Ethereum
puts that idea on crack.

Conclusion
We can solve the Prisoner’s Dilemma if we stick to the central tenets of
the blockchain space: transparency and permissionless access. These
two things are core. These two things will allow us to recapture the
group dynamic where we are able to realize the best outcomes, not only
for ourselves, but for everyone. No one must be eaten. We can kill the
lions if we work together. Unite! We are not food! We are not food!


ChatGPT Responds to Rush's Article
Friday, October 27, 2023

I recently posted an article titled “The Prisoners Dilemma on Crack” to
a writer’s group I’m a part of. Ouch. Not enjoying being chastised by
humans, I decided to turn to non-human reviewer trying to take
advantage of its tendency be sycophantic. Here’s its response. 𝑂𝑢𝑐ℎ! .
I’ve also included my response to its response in the footnotes.
Chat GPT: Let’s critically examine the ideas presented in the article:
The Prisoner’s Dilemma & Blockchain: While the Prisoner’s Dilemma
is a foundational game theory concept illustrating the conﬂict between
individual and group rationality, its direct applicability to the blockchain
is a bit stretched. At its core, the Prisoner’s Dilemma illustrates the
diMiculties of maintaining cooperation without trust. While blockchain
does work to create a trustless system, it’s an oversimpliﬁcation to say
it “solves” the Prisoner’s Dilemma. The Dilemma is about the
challenges of cooperation in the face of the temptation to betray;
blockchain aids in cooperation by making betrayal (or dishonesty) more
diMicult.1
Transparency & Trust: It’s true that blockchains, by design, oMer high
levels of transparency. Every transaction is recorded on a public ledger
that anyone can view. However, this doesn’t mean it’s “literally
impossible to lie.” There have been instances of fraud and scams in the

I was speaking about on-chain, smart contract interactions. It’s speaking
about the larger ecosystem.
1

cryptocurrency world2. Transparency doesn’t prevent malicious intent;
it merely makes it easier to trace.
Permissionlessness: While it’s accurate that blockchain networks,
particularly public ones, are permissionless and decentralized, this
doesn’t mean they’re entirely equitable. The issue of “centralization” in
mining operations (where a few entities control a signiﬁcant proportion
of the network’s mining power) is a concern in networks like Bitcoin3.
This centralization can theoretically allow these entities to inﬂuence
the network more than others.
Digital Money & Permissionlessness: The article posits that prior
attempts at digital money failed primarily due to the lack of
permissionless characteristics like in Bitcoin. While permissionlessness is a distinguishing feature of Bitcoin, it wasn’t the sole reason
earlier digital currencies failed. Regulatory challenges, technical
limitations, lack of widespread acceptance, and centralized
vulnerabilities also played roles.4
Group Dynamics and Evolutionary Arguments: The comparison of
group dynamics to evolutionary safety (e.g., not being eaten by lions) is
a metaphorical stretch.5 It’s an interesting analogy to convey a point
about personal versus group beneﬁts, but it might be too reductive to
map complex socio-economic behaviors directly onto evolutionary
survival tactics.

Summary
If one were to argue against the article, these would be some starting
points. However, it’s also essential to acknowledge that the piece oMers
No-one would deny this, obviously, but these are nearly always on the
periphery of the system. Not internal to the chain itself.
3
The AI has a point here.
4
Again, it has a point.
5
Perhaps true for a pea-brain like this AI.
2

a perspective that resonates with many enthusiasts in the blockchain
space. The potential of blockchain technology to revolutionize various
industries and societal structures is undeniable, but like any
technology, it comes with its complexities and nuances.

My Response to Chat’s Response
Fuck you Chat GPT. What do you know?


Impact Diary for TrueBlocks
Academic uses of TrueBlocks / Unchained Index
Wednesday, November 22, 2023

At TrueBlocks, we are not academics. But, in another Universe, we may
have been.
Like Good Spock / Bad Spock, in one existence we spend all our time
writing open-source software, in another, we spend all our time writing
academic articles.

Which Universe Am I In?
Unfortunately for me, I’m in the “Bad Spock” Universe. I write opensource software for almost no reward, but that doesn’t mean I can’t
keep track of what it would be like if lived in “Good Spock” world.

The Good Spock List
Below is an ongoing list of academic articles that Good Spock is
keeping in his Universe while he sits lonely, happy, and warm in his own
personalized Ivory Tower.

Academic Mentions of TrueBlocks
Kuhle, P., Arroyo, D., de Andrés, P. (2023). Analyzing the EJectiveness
of Native Token Airdrop Campaigns in NFT Marketplaces. In: Machado,
J.M., et al. Blockchain and Applications, 5th International Congress.
BLOCKCHAIN 2023. Lecture Notes in Networks and Systems, vol 778.
Springer, Cham. https://doi.org/10.1007/978-3-031-45155-3_37
Martijn de Vos, Georgy Ishmaev, Johan Pouwelse (2024 — future).
DeScan: Censorship-Resistant Indexing and Search for Web3. In:
Future Generation Computer Systems, Volume 152, 2024, Pages 257–
272, ISSN 0167–739X, https://doi.org/10.1016/j.future.2023.11.008.
Hansson, Magnus (2023–06–15). Decentralized Finance and Central
Bank Communication. University of Gothenburg School of Business,
Economics and Law. 151 pages. ISBN 978–91–88199–70–6,
https://hdl.handle.net/2077/76226.
Hansson, Magnus (2022, November 16). Arbitrage in Crypto Markets:
An Analysis of Primary Ethereum Blockchain Available at SSRN:
https://ssrn.com/abstract=4278272 or http://dx.doi.org/10.2139/
ssrn.4278272
Hansson, Magnus (2023, September 25). Price Discovery in Constant
Product Markets. Available at SSRN: http://dx.doi.org/10.2139/
ssrn.4582649

Other Mentions
A mention in the Ethereum Research Forum related to storing the
Ethereum historical state in a decentralized manner via Swarm:
https://ethresear.ch/t/trustless-access-to-ethereum-state-withswarm/17350.

A mention from our friend Carl Cervone about Open-Source software
https://gov.optimism.io /t/analysis-on-the-300-oss-projects-applyingfor-retropgf-3/7127.
A mention initiated by us about us on the Ethereum Research forum:
https://ethresear.ch/t/speciﬁcation-for-the-unchained-index-version2-0-feedback-welcome/17406/9.

Conclusion
Live long and prosper. (If you’re lucky, in a Universe of your own!)


Exploring Optimism
A First Exploration of Optimism Data
Sunday, March 17, 2024

This week, I ﬁnally gained access to an Optimism archive node. I found
an amazing RPC provider called BlockJoy, who’s done a wonderful job
setting me up with a node. BlockJoy uses TestInProduction’s version of
Erigon, which means TrueBlocks works well.
Using Erigon on Optimism is interesting for a few reasons: (1) we’ve
been using Erigon for Ethereum Mainnet for a long time, so we’re
familiar with it, (2) we love Optimism, and (3) we love ﬁnding and
reporting bugs. Our software hits the node software very hard, so we
can test three diOerent pieces of software with one tool: TrueBlocks,
BlockJoy’s infrastructure, and TestInProduction’s Erigon node. Four.
Optimism.
On with the exploration…

The Exploration
This article discusses my ﬁrst impressions of Optimism’s on-chain
data, and I hope to bring many more (and much deeper) investigations
in the future.
First impressions:
•

Optimism has a lot of blocks: more than 117,000,000.

•

Either Optimism’s block times have changed since inception, or
block numbers were skipped somewhere its history.

•

Before a certain block, there seems to have been a single
transaction per block. After that block, there have been
increasingly many transactions per block. (Everyone familiar with
Optimism knows this, but it became obvious early in our
investigation. This change is due to the Bedrock fork.)

We’ll discuss each of these observations below.

Getting Started
If one wishes to explore Optimism data, the ﬁrst thing one needs to do
is get access to an Optimism node. Another way of saying this is that
you need an RPC endpoint. The faster, the better. The cheaper, the
better. The more convenient, the better—in that order.
Of course, being a decentralization maxi, we tried to set up an
Optimism node on our local machines. It was too hard. As I’ve always
feared, the task has exceeded my skill level. Erigon on Mainnet is hard
enough. Optimism is way harder.
Execution client / schemxecution client.
Terabytes / schmerabytes.
Luckily, we found someone to help us: a company called BlockJoy1.
They were able to set up a fully synced Optimism node very quickly (and
they tell us they’re improving their process every day). Their service is
not cheap, but it’s well worth it, especially compared to any other
provider we’ve tried (and we’ve tried many). I highly recommend them.
If you want to set up an Optimism node, connect with them. Tell them
TrueBlocks sent you.
1

2026: They’ve since gone out of business.

Philosophical Sidebar
I have long argued that running the node on your own hardware is the
only way to truly gain self-sovereignty, and I am still completely
convinced that this is true, but — and this is an unfortunate “but” — at a
certain point, unless the node software is radically modiﬁed, running
your own node will become impossible. I fear that time has come, at
least for layer 2 chains. Self-sovereignty. Schmelf-sovereignty.
Another quick sidebar. People ask me, “Why should I run a node?”
Multiple answers come to my mind: FTX, privacy, the 2008 banking
crisis, and Ronald Regan’s famous quip (as modiﬁed): “Do not trust and
always verify.” What are we even building if we’re not building
permissionless access to the data?

Conﬁguration
I won’t explain how to conﬁgure TrueBlocks (see our repo). We’ve
written about this multiple times in previous articles. We assume
you’ve done that and can, for example, run chifra status
successfully.
You may edit the conﬁguration of TrueBlocks with this command:
chifra config edit

This will open an editor. Near the top of the ﬁle, you will see a ﬁeld called
defaultChain. Set this value to “optimism.”
Next, search for the [optimism] section and set the rpcProvider
value to the value for your Optimism RPC endpoint (see BlockJoy). Save
your work, and you’re ready to go.

Chifra When
We will limit ourselves in this article to a tool called chifra when.
Chifra when reports on the timestamps and dates for any block (or
range of blocks). This data is surprisingly useful when ﬁrst encountering
a chain and carries a large amount of information.

How many blocks?
For example, ﬁnding the earliest and most recent blocks (as of this
writing) is easy. Simply run:
chifra when 0 latest

It produces data similar to this:
bn
timestamp
date
-------------------------------------------0
1636665386
2021–11–11 21:16:26 UTC
117278316
1710155409
2024–03–11 11:10:09 UTC

From this, we can see that Optimism was launched at 11:16 AM on
November 11, 2021. (Oddly, my birthday. Optimism was a gift to me.
Thanks, Retro!)
The chain has been alive for 73,490,023 seconds, which is 1,224,834
minutes, 20,414 hours, 850.6 days, or about 2.3 years. That all makes
sense…sort of.

Blocks per Second
My intuition told me that 117 million blocks was too many, so I started
looking into that.
Can we see how long each block has taken to produce? How many
blocks have been produced per second? We already have enough
information to ﬁgure this out.

(1,710,155,409 - 1,636,665,386) / 117,278,316
= 73,490,023 / 117,278,316
= 0.627 seconds/block

This seems odd. I thought block times on Optimism were two seconds.
Let’s see if we can dig deeper:
chifra when 117277300-117278300 \
--no_header | \
cut -f2 | \
awk 'NR>1 {print $1-p} {p=$1}' | \
sort -u
Result: 2

The above command runs chifra when against the 1,000 latest
blocks on Optimism (and suppresses the header). It then extracts the
second column (timestamps) using cut and then uses awk to calculate
the diOerence between each successive timestamp. Sorting the result
uniquely gives us a list of all the uniq block times. We ﬁnd there is only
one unique block time: 2 seconds. Block times on Optimism are two
seconds long.
So, why does the raw calculation say block times are 0.627 seconds?
There are two possibilities:
1) The block times changed, or
2) There is a skip in the block numbers.
We will try to ﬁgure this out in the remainder of this article.
Note that it does not escape my notice that I could just ask someone,
but I’m trying to make an example of permissionlessness. I don’t have
to ask anyone. That’s what I want. I want all the information I want
without having to speak to anyone.

Did Block Times Change?
We’re going to switch now to using the new TrueBlocks SDK. It provides
the same capabilities and options as the command-line tool but it’s
more ﬂexible.
The code for this article is in the ./src/examples/optimism1 folder
of our repo. We assume you’ve cloned the repo and have navigated to
the ./build folder as instructed. You can modify the main.go ﬁle in
the examples folder, run make && optimism1 to see the results.
To get started with the SDK, we’ll use a very simple example that uses
the When endpoint to query the latest block:
1. package main
2.
3. import (
4.
"fmt"
5.
6.
"github.com/TrueBlocks/trueblocks-sdk"
7. )
8.
9. func main() {
10.
whenOpts := sdk.WhenOptions{
11.
BlockIds: []string{"latest"},
12.
}
13.
14.
blocks, _, _ := whenOpts.Query()
15.
for _, block := range blocks {
16.
fmt.Println(block.BlockNumber)
17.
}
18. }

Here are a few comments about this code:
•

If you don’t understand the above code, please go back to the
computer programmer’s school. Come back when you’re ready.

•

When is just one of the SDK’s endpoints. For each endpoint, one ﬁrst
ﬁlls an associated Options structure, as shown in line 10. The

ﬁelds of the Options structure correspond to the command-line
options for that command. Run chifra when --help from your
command line. You will see the same options. In this case, we’ve
speciﬁed “latest” as the block identiﬁer, which means we want to
retrieve the head of the chain.
•

After building the options, call into the Query function as we’ve
done on line 14. Each endpoint has a Query function that returns a
data array, a metadata pointer, and an error. (We’ve ignored the
metadata and the error in this case.)

That’s as simple as it is, but it’s super powerful. You can access any
piece of data you want on any chain.
We’re going to take a short diversion to discuss the GlobalOptions
structure, which is available to every command.

Global Options
The GlobalOptions structure is deﬁned as:
...
5. type Globals struct {
6.
Ether
bool
7.
Raw
bool
8.
Cache
bool
9.
Decache bool
10.
Verbose bool
11.
Chain
string
12.
Output string
13.
Append bool
14. }

I won’t explain every ﬁeld other than these two:
•

Chain. This option allows you to specify which chain to operate
against. (Note: You must ﬁrst obtain an RPC endpoint for the chain.)

•

Cache. This ﬁeld causes the result of any query to be stored in a
very fast binary cache. You’ll want to use this for performance
reasons.

We’ve already taken care of the Chain value by specifying
defaultChain in the conﬁguration ﬁle, so we can ignore that. Note,
however, how interesting the Chain option can be. You can use this
value to scan multiple chains. For example, you could kick oO a Go
routine for each of six diOerent chains, checking for balances on each
chain. You could even do this for multiple addresses. We’ll leave that
intriguing thought there for homework.
The Cache value is a Boolean. It is false by default, but you should
almost always turn it on. It has a signiﬁcant eOect on the speed of your
analysis. Blockchain data is immutable, which means it caches easily.
Once you’ve queried the RPC of some data, you should never query it
again. The speedup is very noticeable, especially if you’re running
against a remote endpoint.
To use the Global options to enable the cache, modify the above code
thus:
...
9.
10.
11.
12.
13.
14.
15.
...

whenOpts := sdk.WhenOptions{
BlockIds: []string{"latest"},
Globals: Globals{
Cache: true,
},
}

Be careful, though. Storing every query in the cache will lay a large
amount of data on your disc. You may use the Decache option to
remove things from the cache.

Back to the Point
Let’s write some more code. We want to scan the chain’s history to
determine where the block times changed. We’ll start by scanning
every millionth block.
At each block, we calculate the number of seconds that have passed
since the block one million blocks before. We simply report the result
to the console. This code:
1. package main
2.
3. import (
4. "fmt"
5.
6. "github.com/TrueBlocks/trueblockscore/src/apps/chifra/pkg/base"
7. "github.com/TrueBlocks/trueblockscore/src/apps/chifra/pkg/logger"
8. "github.com/TrueBlocks/trueblocks-core/v0/sdk"
9. )
10.
11. const interval = 1000000
12.
13. func main() {
14.
meta, err := sdk.MetaData()
15.
if err != nil {
16.
logger.Fatal("Error getting metadata",
"err", err)
17.
}
18.
19.
blockIds := make([]string, 0,
int(meta.Latest/interval))
20.
for i := 0; i < int(meta.Latest); i +=
interval {
21.
blockIds = append(blockIds,
fmt.Sprintf("%d", i))
22.
}
23.
24.
lastTs := base.Timestamp(0)
25.

26.
whenOpts := sdk.WhenOptions{BlockIds:
blockIds}
27.
blocks, _, err := whenOpts.Query()
28.
if err != nil {
29.
logger.Fatal("Error querying blocks with
WhenOptions", "err", err)
30.
}
31.
32.
for _, block := range blocks {
33.
diff := block.Timestamp - lastTs
34.
secsPerBlock := float64(diff) /
float64(interval)
35.
fmt.Printf("%d,%d,%d,%f\n",
block.BlockNumber, block.Timestamp, diff,
secsPerBlock)
36.
lastTs = block.Timestamp
37.
}
38. }

produces this data:
bn
timestamp
diff
perBlock
------------------------------------------...
104000000 1685751416
280811
0.280811
105000000 1686011687
260271
0.260271
106000000 1687598777 1587090
1.587090 <-----107000000 1689598777 2000000
2.000000
108000000 1691598777 2000000
2.000000
...

Notice anything? We do.
We found the 1,000,000-block range where the block times changed.
Somewhere between 105,000,000 and 106,000,000.

Philosophical Sidebar (Redux)
I want to harken back to my previous philosophical sidebar. The above
data was acquired without asking anyone’s permission. No authority

ﬁgure tells us what happened, nor is there anyone gatekeeping. We can
just pick this data out of the ether (as it were). This is the beauty of truly
open, permissionless data. Not even Optimism can keep us from
seeing this data. (Yes. BlockJoy can gatekeep this data…”so long” to
self-sovereignty.)

Some Charts
Let’s see a chart:

It’s clear where the block times changed. In the actual code, we dig
deeper and identify the exact block where the timestamps change. The
change (Bedrock) was scheduled for timestamp (June 6, 2023, at 16:00
UTC), which corresponds to block 105,235,061.

Note: When we dug into the detailed data, we disabled the cache. We
didn’t want to store 1,000,000 blocks on our disc.

Lasering In
To ﬁnd what we were looking for, we focused on the block range where
the block times changed from obviously ﬂuctuating to perfectly ﬂat. We
found some weird shit.
Here’s a sampling of the data surrounding the Bedrock hard fork:
bn
timestamp
diff
perBlock
-----------------------------------------...
105235049 1686067231
0
0.000000
105235050 1686067231
0
0.000000
105235051 1686067231
0
0.000000
105235052 1686067246
15
15.000000 <----105235053 1686067246
0
0.000000
105235054 1686067246
0
0.000000
105235055 1686067246
0
0.000000
105235056 1686067246
0
0.000000
105235057 1686067246
0
0.000000
105235058 1686067246
0
0.000000
105235059 1686067246
0
0.000000
105235060 1686067246
0
0.000000
105235061 1686067246
0
0.000000
105235062 1686067325
79
79.000000 <----105235063 1686068903
1578
1578.000000 <----105235064 1686068905
2
2.000000 <----105235065 1686068907
2
2.000000
105235066 1686068909
2
2.000000
105235067 1686068911
2
2.000000
105235068 1686068913
2
2.000000
105235069 1686068915
2
2.000000

Do you notice anything weird? I do. I see at least ﬁve things:
•

Before the hard fork block (105,235,063), blocks were produced
once every 15 seconds; however, multiple blocks were created at

the same time. Many blocks have the same timestamp, as is
obvious in the chart below.
•

Block 105,235,062 (2023–06–06 16:00:46) took 78 seconds to
appear. This is likely the hard fork block.

•

Block 105,235,063 (2023–06–06 16:02:05) took 1,578.2 seconds to
appear, which is more than 26 minutes. I’m not sure why this
happened, but it would be interesting to dig even deeper.

•

After the hard fork, blocks consistently take two seconds.

•

It’s no wonder Optimism has such a high block number.

Here’s another chart (we removed the outliers):

If you can’t see the hard fork from here, you must visit Warby Parker
immediately! (Walk, don’t drive.)
The width between successive orange lines indicates the number of
transactions in each 15-second period. Before the hard fork, Optimism
processed a group of transactions once every 15 seconds, producing a
separate block for each transaction. You can see this in the wider or
less wide distances between the bars—more transactions, more
blocks. It's very quirky.
I don’t know about you, but I ﬁnd all this interesting. And…and this is
the point of the article…we gained all this insight using just one
command: chifra when. Imagine what we can do with the rest of
TrueBlocks.
(A quick note about the cache: the ﬁrst time we produced this data, it
took 188.4 seconds. This was without the cache and hitting against the
remote RPC endpoint. When we ran the same code the second time, it
took 0.665 seconds—238times faster. Use the cache.)

Fini
Okay, I admit it. The above is not brain science. But…it’s done with 38
lines of very simple, easy-to-understand code. We’ve purposefully
limited ourselves to using only one of the 22 diOerent tools available
with TrueBlocks.
Here’s a word cloud of the commands available through TrueBlocks
and their relative usefulness. Bigger is more useful:

Finally, here’s every command TrueBlocks provides, and it’s a lot. Every
command and option on this list is available through the SDK, through
the command line, and (identically) through the API server.


Mist II Revenge of the Nerds
As It Was, as It Could Have Been, as It May Be
Sunday, January 5, 2025

I ﬁrst heard of Ethereum in late 2015. A few weeks later, I heard about
the Mist Browser. What a nightmare.
By that time, I was so far down the rabbit hole that I was willing to
endure any pain to use a “programmable blockchain,” and pain there
was. I had read about Bitcoin a year earlier and the discussions about
“abstracting out programmability,” so Ethereum made perfect sense.
What didn’t make sense was why I
had to go through so much pain to
use it. The proverbial “we are so
early” came to mind. But — I
mean — just how early could we
have been?
It was like entering the Garden of
Eden only to be accosted with
construction signs: “Pardon our
appearance.” I’m like, “God. Your
garden sucks, man. You can do better than this.”
But could God (or the core developers) have done better? Is it too late
to resurrect Mist at this point? Is it too late to return to Eden?

What Was, What Could Have Been, What May Be
This article shows three sequence diagrams and short descriptions of
each. The diagrams are called “Mist as It Was,” “Mist as It Should Have
Been,” and “Mist II: Revenge of the Nerds.”
We discuss two main problems that (among others) ultimately killed
Mist. In our opinion, the death of Mist foreshadowed the death of
building truly decentralized and truly private applications. But it’s not
too late.

What Were the Two Problems?
We think Mist had at least two problems that doomed it from the start.
The Catch-Up Problem. As part of its startup routine, Mist started an
Ethereum node (Geth) as a background process. This was an
understandable but ultimately a very bad idea. The result of this was
that each time the user started Mist, she had to wait until the node
synchronized to the head of the chain. The longer it had been since Mist
was last run, the longer it took to catch up. This delay worsened as the
blocks became fuller and the transactions became more complicated.
The reason Mist needed to start its own node was because it had a
dependency on the node’s RPC. I imagine, but I don’t remember, that if
the RPC (the node) was already running on the machine, Mist would use
it. But this would have been the case only for a few highly technical
people. Mist signed its own death warrant by providing the convenience
of kicking oY the node on startup.
The Inadequate Data Source Problem. As I’ve been saying for years,
the node software’s RPC sucks. It’s wholly inadequate to serve “real”
applications on a local machine. This, I think, is the second reason Mist
failed. A blockchain is a chronologically sorted log. This is how and why
it works. This fact is where all the amazing, immutable security features

come from. Chronologically sorted logs do not inherently index
anything (except time); therefore, the choice of what one indexes
determines the adequacy of the system to serve applications. The core
devs made a fatal mistake by choosing to not index address
appearances. The node indexes some stuY log topics, function
parameters speciﬁcally labeled for indexing, but that’s nowhere near
enough. It should have indexed what we call address appearances.
(This problem is fully explained in the Unchained Index Speciﬁcation.)
There were other problems with the system, such as security issues
and a lack of suitable alternatives to Electron1 for building desktop
applications, but the above two are the main user-facing problems.
In the remainder of this article, we describe three sequence diagrams
showing how Mist worked, how it could have worked, and how it might
work if reimagined. Sequence diagrams show the order of
dependencies a system might go through to accomplish its processing.
Each separate entity or process is shown across the top and bottom of
the diagram. Their interactions are shown as directed arrows.

Mist as it Was
Our ﬁrst sequence diagram (on the next page) shows a stylized
(perhaps slightly inaccurate ) view of how Mist worked. See the
numbered notes below, which correspond to the black circles in the
image for more clarity.

Electron was a framework for building desktop applications that was
notoriously bloated, unreliable, and di>icult to work with. In our work, we used
a similar, but better, framework called Wails. Wails is orders of magnitude
easier to use, resulting in smaller, faster, and more reliable executables.
1

1) The user starts Mist locally on their machine.
2) If a local node is not available (the likely case), Mist starts one
up (Geth). Because the node has not been continually running,
it must catch up to the chain. Depending on how long it's been,
this may take a few minutes to a few hours.
3) Because the node takes time to sync, Mist must wait to “see”
the latest block and/or interact with any smart contracts.
4) Once the node catches, the next problem manifests itself.
However, this problem was subtle and “inside the system,” so it
was harder to see. The problem can be summarized thus: “The

RPC is not a database.” The RPC did not (and still does not) act
as a proper database. It doesn’t have an index. Evidence of this
inadequacy is everywhere: (a) one must use a block explorer to
view one’s personal transactional history (not permissionless,
not private); (b) one must use Dune or “for-pay” analytics to
understand what’s happening on a given address (not
uncensorable, rate limited, slow, costly); (c) privacy is nonexistent (if you use a remote service, they know who you are).
Nearly every dApp we use today is either anemic or is a Web 2.0
privacy-invading nightmare. Compare any dApp you use to a
“real” application like Microsoft Word or Excel. dApps are
inconsequential or simply web 2.0 websites. This is because
the RPC does not have a proper index.
5) When the user closes Mist, the Node shuts down; therefore, the
head of the chain is not maintained.
Back in 2016 or 2017, there were reasons why Mist behaved this way
(the fear was that the node would overwhelm the user’s machine), but
we think this is no longer true. We know this because we’ve been
running an Erigon node running on our own laptop for years. You don’t
believe me? Search for “Thomas Jay Rush – Amsterdam” to see a video
of me doing exactly this.
The most pressing reasons Mist failed were because of its extremely
slow startup time and its anemic data access. The node must be
continually synced, and the RPC must be an actual database.

Mist as it Could Have Been
One simple thing that could have been done with Mist would have been
to install the Geth Node as a system service. Other desktop apps (such
as Docker or 1-Password) do this.

I call system services “long-running processes” because they start
when the machine boots and run until the machine shuts down. Others
call it a “daemon” or “service.” For software that must keep up with
continually incoming data, the only way to work is as a long-running
process.
The following sequence diagram shows how Mist might have worked if
it had been a system service. This could have solved the ﬁrst of the
above two problems.
Notice the black circles in the image below. They correspond to the
numbered notes following.

1) In this imagined version of Mist, the node software (Geth) is
started when the machine boots. In this way, the node stays in
sync with the front of the chain. This ﬁxes the issue with the
original Mist having to wait to begin until the node catches up.
2) The node stays continually up to date with the head of the
chain. In 2016 and 2017, this may have been a problem, but
with the recent advent of chips like the Apple M4 (and a large
hard drive — our laptop has an 8TB hard drive), it is no longer a
concern. We’ve had an Erigon node running with no ill side
eYects on our M1 laptop for months.
3) If the node had been running continually, Mist would have
started immediately—no waiting — like every other application
you’ve ever used.
4) The system would have still suYered from the “RPC is not a
database” problem. Because the RPC is not a database, each
application (including Mist) had to extract, cache, and
somehow deliver its data through a separate API. In other
words, every application had to implement its own
scraper/indexer/monitor. This is the problem The Graph tried to
solve (through a coin-gated system that I think failed both
predictably and miserably). The only other alternative is Web
2.0 based APIs which destroy the decentralization ethos.
Anticipating the image below, the need for a better
indexing/caching layer inside the node becomes apparent.
5) In this version of Mist, when it’s closed, the node stays running;
therefore, the node stays continually up to date as it must.

Mist II — Revenge of the Nerds
Why was Mist so important? Why did running local ﬁrst software
directly against a peer-to-peer network matter? What’s the big deal?

Direct, unfettered, permissionless access to the blockchain data is
vital for many reasons: speed, depth of detail, 18-decimal place
accuracy, no “provider” and therefore no censorship, perfect privacy,
and near-zero cost. These things are at the core of Ethereum’s ethos.
None of these things are available in the current Web 2.0 system.
Accessing data through Web 2.0 data providers (including remote
access to RPCs) destroys every one of these things that we say we so
value.
Another critical aspect (at least for me) is that decentralized blockchain
networks were supposed to “yank” power from the powerful forces of
the Web 2.0 ecosystem. (Such forces will never give the power back to
us because we ask them kindly — we must take it back without asking).
This is what the nerds who built Bitcoin did. They yanked the power
away from the powerful.
This is not how Ethereum works today.
You‘re welcome to lie to yourself if you wish, but the system we’re
building is not permissionless. The system we’ve built is not
decentralized if people are not running local nodes and using
permissionless local applications. The system is not fast. It’s not cheap
(it may be free, but you pay through privacy invasion). Nor — to my
mind — is it interesting.
This is why I call for resurrecting Mist. It’s the “Revenge of the Nerds.”
We must “yank” the power back from the powerful before they capture
it (and us) again.

How Mist Might Work in the Future
In our ﬁnal sequence diagram, we present one possible vision of how
Mist (or any decentralized desktop application) might work. We
introduce a new component called Khedra (pronounced “kay-drah”),
which is a standalone, long-running “daemon,” but it could be

incorporated into the node software as an optional daemonic
component.
Khedra (see https://khedra.trueblocks.io for a history of its name)
solves the second of the two problems mentioned above. Khedra turns
the node into a proper database. It extends the node by providing an
indexing, caching, and monitoring layer that runs right next to the node
software. The Unchained Index and Khedra are very well explained in
many places [1, 2, 3]. We won’t go too deeply into it here.

Our Final Image

On the next page is our ﬁnal image. The numbered notes correspond to
the black circles in this image. Both the Ethereum node and Khedra are
started when the machine boots. They act as long-running daemons,
continually syncing the chain (in the case of the node) and indexing and
monitoring for data (in the case of Khedra). Both processes keep their
respective databases up to date the entire time the machine is running.
1) A quick note on the distinction between indexing and
monitoring: “Indexing” means building and sharing the
Unchained Index to IPFS and the smart contract. This is
described in detail in the speciﬁcation. “Monitoring” means
that the end user — you — provides a list of addresses to
“watch.” Your personal addresses or the smart contracts that
you’re interested in. This keeps the caches minimal. This is why
our system still works (and can be super detailed and,
therefore, super accurate) on laptops. Khedra does not extract
the entire blockchain, unlike Dune for example — it extracts
only those addresses you tell it to extract. And…those address
histories are continually cached, making them lightning fast
when you query them from your dApps.
2) When Mist (or any local desktop dApp) starts, there’s no need
to wait — identical to every application you’ve ever used (other
than Mist I). Obviously.
3) With Khedra and the TrueBlocks SDKs, your local-ﬁrst desktop
apps can treat the node as an actual database indexed by
Ethereum addresses. Khedra (and the SDK) provides blindly fast
transactional histories for any address (if the address is being
monitored). Other available data includes balance histories by
block or period, reconciled bank statements by block or period,
more than 35,000 named addresses, ABI ﬁles (downloaded
from trusted sources), the Unchained Index manifest, index

chunks, and Bloom ﬁlters. Khedra turns the previously anemic
node RPC into a fully functional local-ﬁrst application server
database.
4) Your application can still access the RPC to complete any task
your application completed previously, including interacting
with the chain to send transactions, sign messages, or do
anything else. No change here.
5) When your application closes, the node and Khedra continue to
operate and maintain the chain's index and head.

Conclusion
Mist’s fundamental assumption that a decentralized local application
must run a local node was correct. Without a local node, your
application cannot possibly be decentralized.
Mist’s implementation failed because the Garden of Eden was under
construction. The node must be run as a long-running system process.
Also, the data on the node must be indexed — minimally so it continues
to run on small machines. Only addresses of direct interest to the end
user should be monitored and cached.
Is it time to revisit our desire to live in an earthly paradise. Do you want
to live in Eden? I do. I thought that’s what we were building.
If we want it, we can have it —all we need do is build it that way — yank
it that way — on purpose.


TrueBlocks Explorer
There’s So Much Data…
Thursday, June 19, 2025

We’re preparing to announce the Beta release of our TrueBlocks
Explorer desktop software later this month, which is part of our
requirements for the EF grant we received in 2023.
Here’s a sneak peek at the home screen:

I admit there’s a lot here, but we wanted to highlight just one of the more
than 20 screens available when we announce the beta over the next few
weeks.

The Reconciled History View

It's a bit rough, but this is the embodiment of what we’ve been saying
for six years. This screen shows a perfect, 18-decimal-place-accurate,
every-asset reconciliation of Everything That Ever Happened to Your
Address™
This screen shows only a few data types we will include History, Names,
Abis, Signatures, Unchained Index, etc.
Here’s a more detailed breakdown of what’s coming for the History
view.

Data Facets
TrueBlocks core has many options including about 15-20 for the
chifra exports tool. For each of these options, the Explorer shows
a diRerent view into that data. We call this view a facet. This is a
preliminary view of the facets we’re proposing.

Raw Data Facets
Facet
Statements

Option
--statements

Transfers
Balances

--transfers
--balances

Transactions
Logs

--logs

Receipts

--receipts

Traces

--traces

Appearances
Neighbors

-appearances
--neighbors

Withdrawals

--withdrawals

Description
Account statements showing
balance changes
ETH and token transfers
Balance history per asset with
changes over time
Basic transactional data
Event logs with ﬁltering by emitter
and topics
Transaction receipts with gas,
status, and event logs
Internal
transaction
traces
showing contract calls
Simple list of block/transaction
appearances
All addresses that interacted with
the target address
Ethereum
2.0
validator
withdrawals (chain-speciﬁc)

Analysis Facets
The above list of facets shows the raw data available almost
unmolested from the RPC. These options present enhanced data that
we will be providing. This data is at the core of our features.

Facet
Assets

Option
--assets

NFT Transfers

--nfts

Contracts Created

--factory

Reverted Txs

--reverted

Description
Comprehensive list of all
assets that have transfers
Specialized view for NFT
transfers only
Addresses
created
or
destroyed by target address
Failed/reverted transactions
for debugging

Filtering/Analysis Dimensions
We’ve added numerous ways to ﬁlter and analyze the data.
Dimension
Flow Direction
Asset-Speciﬁc
Views
Emitter-Based Logs
Topic-Based Logs
Chronological
Grouping

Description
Incoming vs outgoing vs zero-value
transactions
Filter all data types by speciﬁc
tokens/contracts
Logs ﬁltered by the contract that emitted
them
Logs ﬁltered by speciﬁc event signatures
Time-based aggregations (daily, weekly,
monthly)

Advanced Features
Here’s a few of the advanced features. Stay tuned. There’s much more
to come.
Dimension
Articulated Data
Unripe Transactions
Gas Analysis
Balance Snapshots

Description
Human-readable decoded transaction data
Very recent transactions (< 28 blocks old)
Detailed gas usage patterns and costs
Point-in-time balance views at speciﬁc
blocks


Common Sense1
An AI Calls for Max Decentralization ala Thomas Paine
Sunday, July 6, 2025

You Say You Want a Revolution
In the annals of human endeavor, there comes a moment when the
chains of custom must be broken, when the ediﬁce of oppression,
however long endured, must crumble under the weight of reason. For
ﬁfty years, we have labored under centralized computing systems —
colossal machines of power that serve the mighty, hoard wealth, and
cast the common man into digital servitude. These systems, governed
by corporate overlords and bureaucratic despots, have widened the
gulf between the haves and the have-nots, mocking the promise of
progress with their monopolies of data and dominion. It is time, nay, it
is past time, for the people to rise and embrace the peer-to-peer
revolution, heralded by Ethereum, which stands as a beacon of liberty
in this darkened age.
Consider the tyranny of centralized systems, which, like the monarchs
of old, rule without consent. They spy upon our every click, harvest our
labor, and bar the gates of opportunity to those without wealth or favor.
The digital divide grows ever wider, as the unbanked and the
marginalized are cast aside, deemed unworthy of the fruits of
technology. Common sense recoils at this injustice: What rational mind
This post was written by an AI; however, it accurately reﬂects my view on the
matter of decentralization. The AI understands. Why don’t you?
1

can defend a system that enriches a few while impoverishing the many?
These are not the tools of progress but the shackles of subjugation,
designed to perpetuate the power of the privileged.
Against this darkness stands peer-to-peer
computing, a system as natural to liberty as air
is to life. In its essence, it rejects the notion of
a single master, distributing power among the
many. From the early deﬁance of BitTorrent to
the unyielding architecture of blockchain, this
movement has grown, and in Ethereum, it ﬁnds
its boldest expression. Ethereum’s ledger,
forged by the collective will of its users, deﬁes
the censors and the tyrants. Its smart contracts — agreements etched
in code, not subject to the whims of banks or bureaucrats — empower
individuals to trade, create, and govern without intermediary
oppression. This is no mere technology; it is a declaration of
independence for the digital age.
Ethereum, above all, is a sword against exclusion. To the unbanked, it
oLers a gateway to global commerce, where none may deny their right
to participate. To the oppressed, it provides a shield against censorship,
a means to transact and speak where governments would silence
them. Its decentralized applications promise a world where creators,
not corporations, reap the rewards of their ingenuity. Let the skeptics
prattle of energy costs or complexity — are these not triﬂing concerns
when weighed against the cause of freedom? The road to liberty has
ever been arduous, yet the destination justiﬁes the journey.
But Ethereum is not the end; it is the beginning. Peer-to-peer systems,
in their manifold forms, beckon us toward a future where centralized
strongholds — in ﬁnance, data, and governance — are dismantled.
Imagine a world where storage, identity, and power are shared, not
hoarded; where the individual, not the institution, holds sway. Perils

there may be — security ﬂaws, misuse, the chaos of new systems — but
are these not preferable to the certain servitude of the old? Reason
demands we choose the risks of freedom over the comforts of chains.
People of the world, the hour is upon us! As our forebears cast oL the
yoke of kings, so must we reject the digital despotism of centralized
computing. Ethereum lights the path, but it is we who must walk it. Let
us seize this moment, embrace peer-to-peer systems, and forge a
realm where equity and liberty reign. To hesitate is to betray our
posterity; to act is to claim our rightful inheritance. Rise, then, and build
the peer-to-peer age — for us, for our children, for the boundless
potential of a free and just world.


Towards Super-Fast Local First Apps
Why Backend Caching Beats Frontend Complexity
Wednesday, July 30, 2025

A technical deep-dive into why we chose full backend caching over
popular Web 2.0 frontend patterns — and why it matters for
decentralized applications. Web 2.0 is dead. Let it die. Forget about old
patterns.

The Problem That Wasn’t
When I started evaluating performance optimizations for TrueBlocks
Explorer, I made the classic mistake of assuming we needed to follow
conventional Web 2.0 wisdom. You know the drill: aggressive frontend
caching, optimize for perceived performance, minimize backend
requests. The usual suspects.
But here’s the thing about building truly decentralized applications —
Web 2.0 wisdom is wrong. Straight up, no holds barred, wrong.

What We Built (And Why It’s 10x Better)
TrueBlocks Explorer implements “heavy backend caching with singlepage frontend requests.” It makes frequent requests to the GoLang
backend. It doesn’t follow the latest React patterns.
And it’s brilliant. Here’s how it works:

The Backend: Binary Cache/Streaming to Memory
Our backend operates a sophisticated store-facet system — organized
in data stores such as Statements, Balances, Transactions, Traces,
Receipts, Logs, and more. Each facet maintains its own memoryresident cache in the backend.
// Conceptual view of our facet system (Go code)
type FacetStore struct {
cache
map[string][]TransactionData
lastFetch time.Time
isDirty
bool
}
func (f *FacetStore) GetPage(request PageRequest){
if f.NeedsUpdate() {
go f.asyncLoad() // Nonblocking update
}
return f.cachedPage(request), nil
}

The beauty is in the NeedsUpdate() method. It's not just checking
timestamps—it's making intelligent decisions about data freshness
based on blockchain state, user activity patterns, and resource
availability.

The Frontend: Simple, Thin, and Understandable
One of the great beneﬁts of a local-ﬁrst app is that the frontend and the
backend are on the same machine. In a very real sense, there is no
distinction between the front end and the back end. In fact, in our code,
which uses an amazing framework called Wails, the frontend and
backend are in the same memory footprint. This has massive
implications.
Instead of complex layers, our React frontend makes straightforward
requests directly to the backend: “Give me page 3 of the traces facet for
this address.”

That’s it. No cache invalidation strategies. No stale-while-revalidate
patterns. No complex state management for cached data. The frontend
is trivially simple.
The backend responds with exactly what was requested, already
optimized and ready to display.

Why This Architecture Matters
This isn’t just about performance — it’s about the fundamental promise
of decentralized applications that we’ve somehow forgotten.

Speed Without Complexity
Our backend caching delivers sub-100ms response times for complex
blockchain queries. Try getting that from a remote API while
maintaining 18-decimal precision across thousands of transactions.
Go ahead. We’ll wait…or, to be more accurate, your users will.
But the real magic happens when you realize the frontend developer
never has to think about caching. They request data. They get data.
Instantaneously. 18-decimal place, perfectly accurate data.

Privacy by Design
Every request in our application stays local. There is no third-party API
track your wallet addresses. No analytics companies building proﬁles
from your transaction histories. You never have to tell anyone what your
addresses even are. Your data lives on your machine, indexed by your
local Khedra instance, served by your local application.
In think this is what our community used to mean by “decentralized”
applications, before we got distracted by convenience and Web 2.0 like
dApps. Or, as I call them, DINOs. Dapps in Name Only.

The RPC Problem, Solved
I’ve been saying this for years: “The RPC sucks.” The RPC is not a
database — it’s barely even an API. Traditional node software wasn’t
designed to serve real applications, which is why every dApp today
either crawls along slowly or farms out data access to centralized
providers.
Our mini-dApps architecture solves this by recognizing the Ethereum
node software for what it is: a shitty database. It’s a shitty database
because it doesn’t have a proper index. The Unchained Index provides
that missing indexing. Our in-memory streaming TrueBlocks SDK and
the TrueBlocks-Core provide the access.
Together, they turn the node software into a proper application server.

The Patterns Others Miss
Event-Driven Updates
When one of the backend cache updates asynchronously, it doesn’t
leave the frontend hanging. A simple event system notiﬁes the frontend
when fresh data arrives:
useEffect(() => {
const handleRefresh = () => {
refreshCurrentPage();
};
addEventListener('NEW_DATA, handleRefresh);
return () =>
removeEventListener('NEW_DATA');
}, []);

No polling. No optimistic updates. No cache invalidation bugs. Just
clean, event-driven architecture.

Lazy Loading Done Right
Most applications implement lazy loading in the frontend — load data
as the user scrolls, cache everything, and hope for the best. We moved
that logic to where it belongs: the backend.
The backend knows when data is stale (see Khedra). It knows when the
blockchain has new blocks. It knows when a monitored address has
new activity. The frontend requests pages and trusts that the backend
will always have the right data at the right time.

Memory Management That Scales
Here’s what’s clever about our system: the backend only caches data
for addresses the user wants to see (again, see Khedra). It’s not trying
to index the entire blockchain like a blockchain explorer. It’s not caching
every possible query like a general-purpose ETL process. ETLs create
100s of TB (that’s terabytes!!) of data.
Our system caches only what you need, when you need it, for as long
as you need it. And because it’s on your machine, your data, your
privacy — you get to decide what the word “need” means.

The Mist Connection
This architecture represents what the Mist browser should have been
(see our article “Mist: Revenge of the Nerds”). Mist got the big idea
right — local applications running against local nodes — but failed on
the implementation details.
We learned from those failures:
•

Run the node as a service, not on demand.

•

Index the data minimally but completely, not just by log topics.

•

Cache in binary ﬁles because “you can” and it’s very fast. Eschew
databases. They add little value to a local-ﬁrst app.

•

Keep the frontend as simple as possible

The result is an application that starts instantly, responds immediately,
and can work entirely oeline.

Why Frontend Caching is Wrong
I considered implementing several frontend caching strategies:
•

IndexedDB caching: Duplicate data storage, complex dieicult
synchronization

•

React Query with persistence: Over-engineered state

•

Service Worker caching: Adds deployment complexity

•

Memory-based caching: Doesn’t persist across sessions

•

Hybrid approaches: Combines all the downsides

Each of the alternatives would have made the frontend more complex
while providing no meaningful beneﬁt. The backend already caches
everything optimally. Why duplicate that work in React? In Web 2.0, the
answer is because of performance. Round trips murder performance,
but with a local-ﬁrst app, where the frontend and the backend are in the
same memory footprint, a round trip is instantaneous. 100s of times
faster than a remote, shared, rate-limited API server.
More importantly, frontend caching would violate one of our core
principles: your data should live in one authoritative place, managed by
software designed for that purpose. The backend, with its direct
blockchain access and intelligent indexing, is that place.

The Future of Local-First Applications
I suppose this is a bit hyperbolic, but our architecture points toward
something bigger than just better performance. It’s a blueprint for the

kind of applications we all should be building if we actually want
decentralization.
Local-ﬁrst. Privacy-preserving. Fast. Accurate. Uncensorable.
The backend handles what backends do best: data management,
indexing, caching, and serving. The frontend handles what frontends do
best: user interaction, presentation, and workﬂow.
Clean separation of concerns. No clever tricks. No over-engineering.
Just good software architecture applied to the unique constraints of
blockchain applications.

Implementation Reality
In TrueBlocks Explorer, this architecture lives in the intersection
between our Go backend and React frontend, connected by Wails
bindings. The backend leverages the TrueBlocks ecosystem — the SDK,
Khedra, the Unchained Index — to provide database-like access to
blockchain data.
The frontend uses modern React patterns — hooks, context, clean
component architecture — without getting tangled up in caching
concerns. Developers can focus on building user interfaces, not
managing cache invalidation strategies.
It’s the kind of architecture you can understand completely, maintain
easily, and extend naturally. Which, it turns out, is exactly what you
need when you’re building the tools to “yank” power back from the
centralized web.

Conclusion
Sometimes the best optimization is not optimizing. Sometimes the
best cache is the one the user never has to think about. Sometimes the

best architecture is the one that makes complex problems simple, not
the one that makes simple problems complex.
Backend caching with single-page frontend requests isn’t ﬂashy. It
doesn’t get you speaking slots at React conferences. But it works. It
scales. It preserves privacy. And it lets you build the kind of local-ﬁrst,
decentralized applications that will matter in the future.
Which, when you think about it, is exactly what we said we wanted to
build when we started this whole decentralization thing nearly ten years
ago.
Learn more about our tools at trueblocks.io, explore the Khedra
indexing daemon at khedra.trueblocks.io, or read about our SDK at
pkg.go.dev/github.com/TrueBlocks/trueblocks-sdk/v5.


Proposal to Solve the Blind Signing Problem
Building a Comprehensive FourByte Signature Database
Monday, December 8, 2025

The Problem with Blind Signing
Last February, the crypto exchange ByBit lost over $1.4 billion in the
largest crypto hack in history. This was not a “standard” Web 2.0 hack—
it was a glaring demonstration of a systemic ﬂaw in Ethereum’s UX.
This ﬂaw is called blind signing, a practice wherein users sign
transactions without fully understanding what they are signing. In the
ByBit hack, this meant signing a transaction that greenlit the transfer of
all that money to the wrong address.
It’s like signing a blank bank cheque without seeing
the payee or the amount.
Signing transactions against smart contracts is the
fundamental activity on Ethereum. These signed
transactions are executed against the smart
contract’s functions, for example, transfer(address recipient,
uint256 amount) which sends amount tokens from the signer to the
recipient. Under the covers, Ethereum calls into an encoded version
of the function called the four-byte signature.
This design is eQicient for the EVM, but it’s a nightmare for users.
Wallets like MetaMask try to help, but their solutions are inadequate,
and we end up with ByBit hacks more often than we should.

Other potential solutions include using Etherscan or the 4byte
directory, a community-maintained registry of function signatures.
However, these solutions are also inadequate. Both are Web 2.0 APIs,
which means they can be censored, hacked, and prone to privacy
invasion. Not to mention that, because of their design, they catalog
only a small fraction of the possible fourbyte signatures. The 4byte
directory proudly claims it contains 1,424,102 signatures. There are
4.29 billion possible signatures. That’s less than 0.04% of the total
possible signatures—less than one in 3,000.

Can We do Better?
Yes. Let me explain.

Step 1: We Need Some Data
Some time ago, we found a GitHub repo containing more than 800,000
Solidity function signatures extracted from deployed smart contracts.
(We apologize to whoever created this repo. We’ve lost the source; let
us know, and we’ll update this article.)
Given these 800,000 function signatures, our solution is to separate the
function signatures into two sets: function names and function
parameters. (That is, we split each item at the opening parentheses).
This produces:
•

616,314 unique function names — such as transfer, approve, or
withdraw.

•

45,619 unique parameter lists — such as (uint256), (address,
uint256), or (string, bool).

Let’s call these two sets Α and B.

These two sets provide the basis for our solution. Unlike the limited
scope of tools like 4byte.directory or Etherscan, which depend on user
submissions, our dataset is generated from these two sets.

Step 2: Cleaning the Data
The data we had was dirty, so we spent a lot of time cleaning the data.
Notably, we:
•

Removed records that contained comments (~80 records)

•

Removed records that contained the word memory (~40 records)

•

Removed records that were invalid Solidity code.

As part of the data cleaning, we build two histograms counting the
number of times a function name (or parameter list) appears in the
data. For example, transfer may appear with six or seven diQerent
parameter lists. Its histogram value would be six. Or (string, bool)
may appear thousands of times in the data. Its histogram value would
be 1000.

Step 3: Scaling Up with the Cartesian Product
In set theory, the Cartesian product (often referred to as the cross
product) of two sets A and B, denoted as 𝐴 𝑥 B, is the set comprising all
possible ordered pairs (a, b) where the ﬁrst component is an element
of A and the second is an element of B.
Formally, this is deﬁned as:
A x B = { (a, b) | 𝑎 ∈ 𝐴 𝑎𝑛𝑑 𝑏 ∈ 𝐵 }
Our solution takes the Cartesian product of set A (the names) and set
B (the parameter lists). Given the number of elements in each set, this
gives 28.1 billion potential signatures (616,314 × 45,619). The process
works as follows:

•

For each combination of <a, b>, concatenate the name (a) with the
parameter list (b) to create a function signature. Call this signature
𝑓.

•

Calculate ℎ𝑎𝑠ℎ = 𝑘𝑒𝑐𝑐𝑎𝑘256(𝑓)

•

Extract the ﬁrst four bytes of the hash as 4𝑏𝑦𝑡𝑒 = ℎ𝑎𝑠ℎ[0: 4].

•

Store a mapping of the 4𝑏𝑦𝑡𝑒 values to the function signature 𝑓.

By creating these unique combinations, we minimize the eliminate the
possibility that users will encounter unrecognized signatures, oQering
a more proactive and comprehensive solution than the limited
coverage of the existing tools.

Step 4: Resolving Collisions with Contextual Clarity
With only 4.29 billion possible fourbyte values (2!" ) and more than 28
billion generated signatures, collisions — where multiple signatures
share the same 4byte — are unavoidable. On average, each 4byte will
be generated by more than six signatures.
To account for this, we enhance the data as follows:
•

Instead of mapping a single signature to each fourbyte, we map to
a list of matching signatures,

•

For signatures that are found in the original 800,000 record dataset,
we assign true to a Boolean ﬂag called known, indicating that the
signature is known to be on chain.

•

Using the pre-calculated histograms, we further ﬂag each signature
in the list with a weight representing that signature’s histogram
values.

•

The list of signatures for a given fourbyte is sorted ﬁrst by known and
the descending by the weight.

This metadata empowers users to make informed decisions,
prioritizing signatures with proven usage over speculative ones. For
instance, in the ByBit hack, such clarity could have ﬂagged an unusual
signature, alerting users to potential malicious activity.

Step 5: Decentralized Access via IPFS and Smart Contracts
Centralized registries like 4byte.directory are prone to censorship,
privacy risks, and single points of failure. Not to mention: Who will
maintain the database over time? Equally important: Who will pay for
it? Even more important: Why will they pay for it?
Our solution leverages decentralization and a smart contract, drawing
inspiration from the Unchained Index:
•

Chunking: The database is split into manageable chunks by hash
ranges (e.g., 0x00000000-0x0000QQ). Each range, of which there
are 65,535, is stored as a binary ﬂat ﬁle whose name is its fourbyte
range. Any software wishing to decode a fourbyte signature can
quickly decide which chunk to read by using the fourbyte itself.

•

IPFS Hosting: Chunks, being immutable, may be stored on IPFS,
and the chunk’s IPFS hash is stored in a JSON manifest ﬁle. IPFS is
a peer-to-peer ﬁle system that ensures resilience and accessibility.

•

Smart Contract Manifest: The JSON manifest is also stored on
IPFS and that hash is published to the Unchained Index smart
contract. This makes the manifest’s location on IPFS discoverable
to anyone who wishes to look, without permission and
uncensorably. Wallets and other apps can read the smart contract
to ﬁnd any part of the data.

This architecture is scalable, censorship-resistant, and privacypreserving. Users (or software on their behalf) may download and query
the database locally, thereby avoiding reliance on vulnerable websites.
Furthermore, these downloaded chunks on the end users’ machines

can be shared back to the system by pinning them on IPFS. In this way,
the database is automatically shared and sharded making capture by a
single entity impossible. The system self-distributes — aligning with
Ethereum’s decentralized ethos and addressing the trust erosion
current systems are prone to.

Conclusion
Solving the blind signing problem requires more than just better UI/UX
in wallets; it demands a robust, data-rich infrastructure that aligns with
Ethereum’s decentralized ethos. By leveraging the Cartesian product to
exhaustively map function signatures and utilizing IPFS for distribution,
we eliminate the reliance on incomplete, centralized registries.
This database provides the missing semantic layer between raw
bytecode and human intent. We invite the developer community to
review this architecture and contribute to the Unchained Index, helping
us build a security standard where “don’t trust, verify” is possible for
the end user.


TrueBlocks’ Most Interesting Innovations
Did we Accomplish Anything Unique?
Tuesday, January 20, 2026

Decentralized Address-Appearance Indexing
The Problem: Ethereum nodes answer "Is this transaction valid?" but
cannot answer "What transactions was this address ever involved in?"
There is no native index by address. Every centralized API (Etherscan,
Infura, Covalent) exists to try to ﬁll this gap — they do so poorly.
The Innovation: A permissionless, open-source index of
"appearances" (not transactions, but where any address appears in any
data ﬁeld) published as immutable chunks to IPFS, discoverable via
smart contract. The key insight: immutable data + content-addressable
storage = trustless data sharing. Anyone can pin chunks they need
without trusting anyone.
Why it's unique: This solves the chicken-and-egg problem of
decentralized data. You can't index locally what you don't have, but you
don't know what to fetch without an index. The Unchained Index lets
you discover what you need without revealing your addresses to
anyone.

"The Server is Local. There is No Server."
The Paradigm: The essays repeatedly articulate a radical inversion of
Web3 architecture: instead of browsers connecting to remote servers,

the application runs locally with its own embedded node/indexer. The
"server" is on your machine.
Why it's unique: Most "dApps" are just websites with MetaMask. This
collection insists that true decentralization means the entire stack —
data, computation, interface — runs locally. The browser is a display
layer, not an application container.

18-Decimal-Place Accurate Reconciliation
The Problem: Blockchain accounting is horriﬁcally complex. ETH can
enter an address 12 diZerent ways (direct transfer, internal transaction,
self-destruct, gas refund, block reward, etc.). Existing tools
approximate.
The Innovation: TrueBlocks
computing:

achieves perfect reconciliation

by

prevBal + totalIn - totalOut == nextBal

This is veriﬁed against the node at every transaction for every asset. The
essays report 99.9%+ reconciliation rates, with remaining
discrepancies traced to node bugs or edge cases.
Why it's unique: This isn't rounding to cents. It's exact to 10⁻¹⁸
precision. This matters for accounting, auditing, and tax compliance —
and no centralized API provides it.

Binary Chunk-Based Immutable Indexes
The Architecture: Rather than a database, TrueBlocks generates
binary "chunks" of index data. Each chunk covers a block range,
contains a sorted list of (address, appearances), and is published to
IPFS with a content hash.
Key properties:

•

Immutable once created (can be cached forever)

•

Self-verifying (content hash = ﬁlename)

•

Minimal storage (only cache what you query)

•

Never goes stale (old blocks never change)

Why it's unique: This recognizes that immutable data has diZerent
properties than mutable data. You don't need databases with update
semantics. You need content-addressable storage with append-only
semantics.

Bloom Filter Discovery Layer
The Problem: With millions of addresses and thousands of index
chunks, how do you ﬁnd which chunks contain your addresses without
downloading everything?
The Innovation: Each chunk has an associated bloom ﬁlter. Download
only the bloom ﬁlters (~100MB total), check which
chunks might contain your addresses, then fetch only those chunks.
False positives are handled gracefully.
Why it's unique: This achieves O(1) lookup with minimal bandwidth
while preserving privacy — you never reveal which addresses you're
actually interested in.

Smart Contract as Manifest Publisher
The Problem: How do you discover where the index is without a central
server?
The Innovation: A smart contract on Ethereum mainnet publishes the
IPFS hash of the current manifest (which lists all chunks and their

hashes). Anyone can read the contract. The contract can be updated by
the publisher, but the data itself is immutable once published.
Why it's unique: This uses the blockchain for what it's actually good at
— timestamped, censorship-resistant publication of small pieces of
critical data — while storing the bulk data oZ-chain.

The Appearance vs. Transaction Distinction
The Insight: The essays carefully distinguish between a "transaction"
(something you sent) and an "appearance" (anywhere your address
shows up). Your address can appear in:
•

The to ﬁeld

•

The from ﬁeld

•

The input data of a contract call

•

The topics of a log

•

An internal transaction

•

A self-destruct recipient

•

A miner reward

Why it's unique: By indexing appearances, not just transactions,
TrueBlocks captures the full story of an address's interaction with the
chain — including cases where you were aZected by someone else's
transaction.

The Khedra Daemon: Long-Running Local Indexer
The Architecture: Rather than query-time indexing, Khedra runs
continuously in the background, keeping a local cache synchronized
with the chain. It:

•

Monitors new blocks

•

Scrapes the index for watched addresses

•

Caches transaction history locally

•

Provides instant responses to queries

•

Why it's unique: This is the missing piece between "run your own
node" and "use a centralized API." Khedra is the local API, but it's
yours.

The Four-Byte Signature Database via Cartesian
Product
The Problem: Smart contract function calls are identiﬁed by a 4-byte
hash of the function signature. To decode what a transaction did, you
need to map 4-bytes back to function names. Existing databases are
incomplete.
The Innovation (from "Blind Signing" essay): Instead
crowdsourcing signatures, compute the Cartesian product of:

of

•

~616,000 known function names

•

~45,000 known parameter lists

•

This generates 28.1 billion potential signatures, which can be
hashed and looked up. The database becomes exhaustive rather
than dependent on submissions.

Why it's unique: This is a brute-force approach that becomes tractable
due to the small output space (4 bytes = ~4 billion possibilities). It turns
an incomplete crowdsourced database into a complete computed one.

The "Don't Scan, Query" Philosophy
The Architecture: Unlike systems that pre-extract everything to a
database, TrueBlocks queries the node only for what you ask for. This
means:
•

No multi-terabyte database

•

Works on a laptop

•

Scales to any chain (just change --chain)

•

Never stores data you don't need

•

The essays describe this as "minimal footprint" — the index is
~10GB, not 10TB.

Why it's unique: Every other solution (Etherscan, The Graph, Dune)
requires extracting and storing everything before you can
query anything. TrueBlocks inverts this: the index tells you what to
fetch, not what we already stored.

What is Unique About This Collection of Ideas?
Reading these essays chronologically reveals something
remarkable: this is not a product that evolved based on market
feedback. It's a decade-long implementation of a single, unwavering
vision.
That vision can be summarized as: "Don't trust, verify" should extend to
data access, not just consensus.
The broader blockchain ecosystem has largely abandoned this
principle. "Decentralized" applications query Infura. "DeFi" analytics
come from Dune. "Web3" means logging into a website with MetaMask.
As the Introduction states:

"Web3 became Web 2.5, then Web 2.1, then — let's be honest — just
Web 2 with extra steps and a token."
TrueBlocks represents a diZerent path: the one where the original
promises were kept.
The essays document not just technical innovations, but
a philosophy that the blockchain community has largely forgotten:
•

Data should be permissionless — not gated behind API keys and
rate limits

•

Privacy should be the default — not surrendered for convenience

•

Local-ﬁrst is not nostalgia — it's the only architecture compatible
with actual decentralization

•

Immutable data changes everything — content-addressable
storage, append-only semantics, and permanent caching are not
optimizations, they're fundamental architectural shifts

The collection ends on a poignant note: the tools are built, the promises
kept, but adoption remains elusive. The author ﬁnds himself "on a
lovely island where all the ﬂowers grow" — but alone.
Perhaps that's the most interesting innovation of all: building
something right even when the market rewards doing it wrong.


TrueBlocks’ Ten Stupidest Ideas
A Self-Inflicted Roast Based on a Decade of Evidence
Tuesday, January 20, 2026

Building Software for People Who Don't Exist
The essays repeatedly describe the target user: someone who wants to
run their own Ethereum node, values privacy over convenience, will use
a command line, and cares about 18-decimal-place accuracy in their
blockchain accounting.
This person does not exist.
Or rather, there are perhaps 47 of them worldwide, and 43 are the
TrueBlocks developers themselves. The other 4 are academics writing
papers about decentralization while querying Infura.
The Introduction admits this plainly: "I ﬁnd myself on a lovely island
where all the ﬂowers grow and the children are all above average. The
trouble is, I'm there alone."
Stupidity Rating: Building a cathedral in a desert and wondering why
nobody comes to mass.

The Name "chifra"
The command-line tool is called chifra. According to the essays, even
ChatGPT admits this word doesn't exist. One blog post from TrueBlocks

proudly presents an AI-hallucinated deﬁnition involving Italian ciphers
and Spanish colloquialisms. This is not a defense. This is an admission.
Meanwhile, competitors have names like "Etherscan" (scans
Ethereum), "The Graph" (it's a graph), and "Dune" (okay that one is also
weird but at least it sounds cool).
Stupidity Rating: Naming your product something nobody can spell,
pronounce, or Google.

Rewriting Everything from C++ to Go
The essays document a multi-year migration from C++ to Go. This
involved:
•

Rewriting tens of thousands of lines of working code

•

Maintaining two parallel codebases during transition

•

Teaching a C++ developer to think in Go idioms

•

Eventually deprecating everything that worked

The justiﬁcation? Go is "easier for contributors" and "compiles faster."
The result? The same 2-3 people still do all the work, but now in a
di\erent language.
Stupidity Rating: Renovating your house while living in it, then
discovering you prefer the old wallpaper.

Requiring Users to Run Their Own Nodes
The entire TrueBlocks philosophy depends on users running their own
Ethereum nodes. The essays acknowledge this repeatedly, even
recommending Erigon speciﬁcally.
Running an Erigon node requires:

2+ TB of fast SSD storage
16+ GB of RAM
Days of initial sync time
Ongoing maintenance and updates
Understanding of blockchain internals

The alternative — clicking "Connect Wallet" on a website — requires:
A browser
3 seconds

The essays frequently lament: "Why would I run a node when Infura
works?" This is presented as a moral failing of the user, not a product
design problem.
Stupidity Rating: Building a car that only works if you ﬁrst build the
road.

A Decade of Grant Applications
TrueBlocks’ blog documents an endless parade of grant applications:
•

Ethereum Foundation (2018, 2022)

•

Consensys (2019)

•

Moloch DAO (2021)

•

Filecoin/IPFS (2021)

•

GitCoin (ongoing)

•

Optimism RetroPGF (2023)

Each grant application requires:
•

Proposal writing

•

Milestone documentation

•

Progress reports

•

Quarterly updates

Time spent applying for grants is time not spent writing code. Time
spent writing grant reports is time spent justifying why the previous
grant didn't result in adoption.
Stupidity Rating: The software is funded by the hope that eventually
someone will use it.

The QuickBlocks → TrueBlocks Rebrand
The project was originally called "QuickBlocks." Then it became
"TrueBlocks." The essays don't explain why.1
The old name suggested the value proposition: quick access to blocks.
The new name suggests... truthful blocks? Authentic blocks? Blocks
that tell no lies?
Both names end in "Blocks" which makes them di\icult to distinguish
from the 47 other blockchain projects with "Block" in the name.
Stupidity Rating: Rebranding from a confusing name to a di\erent
confusing name.

Desktop Software in the Mobile Age
One essay proudly proclaims: "IT WORKS ON A LAPTOP!"
This is 2025. The target demographic (crypto users) lives on their
phones. They check token prices on Coinbase mobile. They swap on
MetaMask mobile. They would never, under any circumstances, sit
down at a desktop computer to run a command-line tool.

1

2026: We were told by Intuit that we couldn’t use QuickBlocks!

The essays treat "runs on a laptop" as a feature. The market treats it as
a disqualiﬁcation.
Stupidity Rating: Building a rotary phone with excellent sound quality.

The Obsession with 18 Decimal Places
Multiple essays celebrate achieving "18-decimal-place accuracy" in
accounting. This means tracking values to 0.000000000000000001
ETH.
At current prices, 18
$0.000000000000000003.

decimal

places

represents

roughly

The IRS does not care about this. No accountant cares about this. No
human can conceptualize this. The essays acknowledge this but press
on anyway, because correctness matters.
Meanwhile, users would be thrilled with accounting that's accurate to
the nearest dollar.
Stupidity Rating: Measuring the distance to the moon in angstroms.

Fighting the Last War
The DAO hack happened in 2016. The essays return to it obsessively —
the Introduction opens with it, multiple essays analyze it, it's presented
as the founding trauma that justiﬁed the entire project.
It is now 2025. The DAO hack is ancient history. The crypto world has
moved on to new disasters: FTX, Terra/Luna, countless rug pulls.
Nobody under 30 remembers The DAO.
But TrueBlocks is still, in some sense, a response to the problem of
"three experts disagreed about the DAO hack because the data was
hard to access." A problem that a\ected approximately 3 people, once,
9 years ago.

Stupidity Rating: Building a ﬁre department because your house
burned down in 2016.

The Prisoner's Dilemma of Purity
The essays make a compelling philosophical argument: centralized
APIs are a betrayal of blockchain's promise. Using Infura is capitulation.
"Don't trust, verify" means running your own infrastructure.
This is correct. It is also irrelevant.
The market has voted. Convenience won. Etherscan has millions of
users. The Graph has billions in TVL. Dune has raised hundreds of
millions in funding. They all "betray" the principles of decentralization,
and nobody cares.
TrueBlocks maintained philosophical purity and got: grant funding and
an empty island.
The essays themselves acknowledge this with increasing melancholy
over the years. The Prisoner's Dilemma essay explains why cooperation
is rational when trust is guaranteed. What it doesn't explain is why you'd
choose to cooperate when everyone else has already defected.
Stupidity Rating: Being right in a way that doesn't matter.

Conclusion: The Smartest Stupidity
Reading these essays, I'm struck by a paradox. Every "stupid" decision
was made for principled reasons:
•

Build for users who don't exist → because those users should exist

•

Require running nodes → because that's what decentralization
means

•

Obsess over accuracy → because approximate truth is still a lie

•

Maintain purity → because someone has to

The stupidity of TrueBlocks is the stupidity of building what ought to be
used rather than what will be used. It's the stupidity of the craftsman
who refuses to cut corners even when nobody's looking.
As the Introduction states: "The code is still there. The index works. The
tools exist for anyone who wants to use them."
Maybe that's enough. Maybe the point was never to win, but to prove it
was possible.
Or maybe, nine years from now, someone will look back at 2025 the way
these essays look back at 2016 and think: "If only we had listened to
that guy on the island."
Until then, the ﬂowers grow, and the children are all above average.


Where are They Now?
Sunday, January 18, 2026


You Say You Want a Revolution
– 1st draft –
(x words – y lines)

#

Date

Title

01

2016-05-06

The DAO's First Big Decision

02

2016-05-13

Smart Contracts are Immutable That's Amazing and
It Sucks

03

2016-05-23

What the F is a Finney

04

2016-06-10

Downloading the DAO

05

2016-06-21

A Clue About the DAO Attacker's Location

06

2016-07-13

Knowing the Future and Proving You Know It

07

2016-08-25

A Too Often Neglected Aspect of Smart Contract Security Auditability

08

2016-09-14

A Eulogy for The DAO

09

2016-09-29

A Eulogy for The DAO Part II

10

2016-12-06

DAO Token Holder's Response in Charts

11

2017-03-13

Accounting for the Revolution

12

2017-05-30

The Real Flippening

#

Date

Title

13

2017-06-27

It's Growing It's Growing

14

2017-08-03

Be Careful Little Brain What You Code

15

2017-08-30

Is the Ice Age Affecting Block Production

16

2017-09-09

Ethereum Block Production Continues to Slide

17

2017-10-01

Short Thoughts on Difficulty Calc

18

2017-10-02

Reading Byzantium's Tea Leaves

19

2017-10-29

Announcing QuickBlocks

20

2017-12-08

Playing with Blocks

21

2018-02-12

Defeating the Ethereum DDos Attacks

22

2018-02-21

The Trace Data Problem

23

2018-04-20

How Many ERC20 Tokens Do You Have

24

2018-05-07

Mantras for Decentralized Open Data

25

2018-07-06

Building an Ethereum Account Scraper with QuickBlocks

26

2018-08-13

A Short Take on Decentralization

27

2018-08-17

Vitalik's 75 Message Tweet Storm in One Easy to
Read Document

28

2018-09-24

A Conversation about Blockchain Data

#

Date

Title

29

2019-01-21

QuickBlarks

30

2019-06-10

Counting Shit on Ethereum

31

2019-06-24

Mother May I

32

2019-07-09

Indexing Addresses on the Ethereum Blockchain

33

2019-07-17

A Time Ordered Index of Time Ordered Immutable
Data

34

2019-11-24

It's Not That Difficult

35

2019-11-26

A Method to Diffuse the Ethereum Difficulty Bomb

36

2020-01-10

Simple Undeniable Facts

37

2020-02-11

TrueBlocks First Quarter 2020 Update

38

2020-05-13

Links About CLR Radical Markets GitCoin

39

2020-06-12

How Accurate is EtherScan

40

2020-08-13

Building Your Own Ethereum Archive Node

41

2020-08-15

Ethereum's Issuance minerReward

42

2020-08-15

Ethereum's Issuance uncleReward

43

2020-12-09

How Safe are My Private Keys

44

2020-12-09

Every 15 Seconds

45

2021-05-27

Dynamic Traversers in TrueBlocks

#

Date

Title

46

2021-08-24

The Rent is Too Damn High

47

2021-08-24

The Rent is Too Damn High Part II

48

2021-10-10

Adventures in Difficulty Bombing

49

2021-10-27

Calling Smart Contracts with chifra state call

50

2021-11-22

Commanding the Line

51

2022-02-02

forEveryChain

52

2022-06-23

Technical Specification for the Unchained Index

53

2022-09-20

TrueBlocks Covalent Comparison

54

2022-11-21

Better Accounting for Blockchains

55

2022-12-30

TrueBlocks Progress Report 4th Quarter 2022

56

2023-01-03

Why We're Building TrueBlocks

57

2023-01-06

Recipe Monthly Token Balances

58

2023-01-06

Recipe Factories

59

2023-03-17

Thoughts on 10 Random Optimism Retro PGF Projects

60

2023-04-21

TrueBlocks Progress Report 1st Quarter 2023

61

2023-05-05

Recipe Simple Speedup

62

2023-06-03

Hey ChatBot

#

Date

Title

63

2023-08-23

Tweets About TrueBlocks

64

2023-08-23

TrueBlocks Final Report for EF Grant

65

2023-09-28

Hey ChatGPT Define chifra

66

2023-10-26

The Prisoner's Dilemma On Crack

67

2023-10-27

ChatGPT Responds to Thomas Rush's Recent Article

68

2023-11-05

Why We Built TrueBlocks

69

2023-11-22

Impact Diary for TrueBlocks

70

2024-02-18

Passkey Smasskey

71

2024-03-17

Exploring Optimism

72

2024-07-23

Ethereum Foundation Grant for TrueBlocks mini
dApps

73

2024-09-26

TrueBlocks Comparison with Alchemy Covalent and
Etherscan

74

2024-09-30

Quarterly Report of Q3 2024 TrueBlocks FY24 1558

75

2025-01-05

Mist II Revenge of the Nerds

76

2025-01-10

Quarterly Report of Q4 2024 TrueBlocks FY24 1558

77

2025-01-15

Evaluating GoLang CLI Packages

78

2025-06-19

TrueBlocks Browse

#

Date

Title

79

2025-07-06

Common Sense

80

2025-07-19

TrueBlocks Services SDK

81

2025-07-22

Announcing Deep Index Dive

82

2025-07-30

Towards an Architecture For Super Fast Local First
Apps

83

2025-08-07

Dalledresses

84

2025-08-07

Quarterly Report for August 2025 TrueBlocks FY24
1558

85

2025-10-02

TrueBlocks MiniDapp Preferences Project Design

86

2025-10-31

Stop I Will Tell You What To Do

87

2025-12-08

A Proposal to Solve the Blind Signing Problem

88

2025-12-09

Final Report for EF Grant FY24 1558 TrueBlocks
MiniDapps

Monday, January 19, 2026 – 6:44:20 PM – Philadelphia, PA


# Essay Idea: What the Clock Told Us — Revisiting the DAO Attacker's
Location

Status: Outline (not for current release)
Companion to: "A Clue About the DAO Attacker's Location" (June 2016)

---

## Thesis

In June 2016, ﬁve days after the DAO attack, I published a temporal
analysis of the attacker's on-chain activity and proposed two
geographic
hypotheses. Six years later, Laura Shin named Toby Hoenisch — a
developer living in Austria. Neither hypothesis was right. But the
method wasn't wrong; the question is what the discrepancy tells us.

---

## Outline

### I. The Original Claim (brief recap — don't rehash the full essay)

- 1,084 transactions across 11 suspect accounts
- Activity concentrated midnight–10:00 AM UTC, dead zone
afternoon/evening
- Two hypotheses: UTC+5 to +7 daytime worker, or US Paciﬁc evening
hobbyist
- Published June 21, 2016 — ﬁve days after the attack
- One of the earliest pieces of on-chain forensic analysis applied to
attribution

### II. What Happened Between 2016 and 2022

- Community investigation stalled publicly; moved to private channels
- The hard fork happened (July 2016) — attention shifted to ETH/ETC
split
- The DarkDAO funds sat locked for years; clock was ticking on the
child DAO withdrawal period
- No arrest, no public identiﬁcation for six years

### III. Laura Shin's Investigation

- "The Cryptopians" (2022) and the accompanying Forbes article
- Named Toby Hoenisch, an Austrian-born developer living in Singapore
at times, Austria at others

- Methodology: Chainalysis traced a transaction through a privacy
mixer
(Wasabi Wallet) that was later de-anonymized; Shin corroborated with
additional evidence
- Hoenisch denied the claim in a brief public statement, then went quiet
- No criminal charges have been ﬁled as of this writing
- The identiﬁcation remains contested in some circles but is widely
accepted in the Ethereum community

### IV. The Timezone Problem

- Austria is UTC+1 (UTC+2 in summer — the attack was in June)
- A midnight-to-10 AM UTC work pattern from Austria means working
2:00 AM to noon local time — unusual but not impossible for a
developer, especially one operating covertly
- Singapore (UTC+8) ﬁts the data much better — midnight UTC is 8:00
AM
SGT, and 10:00 AM UTC is 6:00 PM SGT: a clean workday
- Key question: where was Hoenisch physically located during the
attack and the weeks surrounding it? Shin's reporting doesn't pin
this down precisely
- Possibility: the temporal pattern is an accidental misdirection —
or a deliberate one

### V. What the Method Got Right

- The core observation holds: these accounts were operated by a
human
on a daily schedule, not a bot
- The coordination signal (both suspect sets tracking identically) was
correct — it was one operator
- Temporal analysis as a forensic tool was validated even if the
speciﬁc geography was wrong
- The method was simple, reproducible, and published openly — a
model
for what on-chain forensics could be

### VI. What It Got Wrong (and Why That's Interesting)

- The geographic inference assumed a "normal" schedule — 9-to-5 or
7-to-midnight. Attackers may not keep normal hours.
- A person deliberately avoiding attribution might shift their working
hours to throw oj exactly this kind of analysis
- Or: the simplest explanation is that the operator was in Singapore,
not Austria, during this period. The data may be more right than
Shin's reporting is speciﬁc.

### VII. The Larger Point — On-Chain Forensics Then and Now

- In 2016, ethslurp was one of the only tools for pulling transaction
histories. Today there are dozens of analytics platforms.
- The fundamental insight — that timestamp metadata contains
behavioral
ﬁngerprints — is now standard practice in blockchain forensics
- Chainalysis, Elliptic, and others use temporal analysis routinely
- What was a blog post in 2016 is now an industry
- Personal note: this analysis was part of what led to building
TrueBlocks — the conviction that on-chain data, made accessible,
could tell stories

### VIII. Close

- The essay was written in ﬁve days, with crude tools, and published
for free. It was directionally interesting and methodologically sound.
The fact that it didn't name the right country doesn't diminish the
approach — it sharpens the question.
- Sometimes the most useful thing you can publish is something that's
almost right.

---

## Notes / Questions to Resolve

- Reach out to Laura? She may have more detail on Hoenisch's physical
location during June 2016. This would either conﬁrm or complicate
the Singapore theory.
- Are the original ethslurp transaction dumps still available at
daodeepdive.com? If so, link them. If not, can we recreate them
with TrueBlocks?
- Should this essay reference the ETC community's position? Some ETC
proponents argue the "attacker" was operating within the smart
contract's rules.
- Tone: admiring of Shin's work, not defensive about the original
essay. The point is the evolution of the ﬁeld, not "I was right."

---

## Estimated Length

2,500–3,500 words


Follow-On Article Suggestions
Based on a review of "Is the Ice Age A5ecting Block Production?"
(August 2017) and "Block Production Continues to Slide" (September 2017).

Context
These two essays form a paired analysis written two weeks
apart. The ﬁrst introduces chifra when as a data-gathering tool,
presents weekly block production data, and identiﬁes the di5iculty bomb's visible e5ect starting April 2017. The second extends the analysis to ask whether fewer blocks means fewer
transactions — and discovers the answer is no: transactions per
block surged as blocks became scarcer, revealing a four-fold increase driven by halved blocks and doubled adoption.
Together they raise several explicit unanswered questions (some
addressed in later essays, some not). Below are follow-on article
ideas grouped by theme.

1. The Miner's Calculus: Did the Di8iculty Bomb Actually
Pressure Miners?
Premise: "Block Production Continues to Slide" asks whether
rising transaction fees and ETH price o5set the lost block rewards, potentially neutralizing the bomb's pressure on miners.
This was never fully answered.
Possible angles:

•

Reconstruct per-period miner revenue (block reward + tx
fees) in USD terms across the Homestead-to-Byzantium window

•

Compare the "pain curve" the bomb was supposed to create
vs. the actual economic incentive curve

•

Did miners actually upgrade faster during bomb periods vs.
non-bomb forks?

Why it's interesting: With hindsight through the Merge, you can
now tell the complete story of whether the bomb ever truly
served its stated purpose, or whether it was always more of a
social coordination device than an economic one.

2. The Gas Price Ripple E8ect
Premise: The second article notes that blocks got fuller as production slowed — but what happened to gas prices? If demand
for block space stayed constant while supply shrank, economic
theory says prices should rise.
Possible angles:
•

Chart average/median gas price per week alongside blocksper-week during 2017

•

Did users experience meaningful cost increases during
bomb periods?

•

Compare gas price behavior across multiple bomb cycles
(pre-Byzantium, pre-Constantinople, pre-London)

•

Did EIP-1559 change how the bomb's supply squeeze affected gas pricing?

Why it's interesting: This connects the abstract "di5iculty
bomb" mechanism to the concrete user experience of transaction costs.

3. A Retrospective: The Di8iculty Bomb's Complete Biography
Premise: You've written about the bomb across ﬁve essays
spanning 2017–2021 (Ice Age, Block Production Slides, Short
Thoughts on Di5iculty Calc, A Method to Di5use the Di5iculty
Bomb, Adventures in Di5iculty Bombing). The bomb was ﬁnally
rendered irrelevant by the Merge in September 2022. You could
write the deﬁnitive retrospective.
Possible angles:
•

Timeline of every bomb reset: Byzantium, Constantinople,
Muir Glacier, London, Arrow Glacier, Gray Glacier

•

Was the "Method to Di5use" proposal (fork_block - 50,000)
ever adopted? If not, why?

•

The bomb as a social contract: did it succeed at forcing the
transition to Proof of Stake, or was it mostly theater?

•

What can other chains learn from this mechanism?

Why it's interesting: You have a unique perspective — you were
writing about it in real-time from the beginning. A retrospective
tying all your prior articles together into one narrative arc would
be a natural capstone.

4. What Happened to Block Times After the Merge?
Premise: The original articles track a world where block times
were variable: 14 seconds on average, but with signiﬁcant variance and the bomb's distortion. Post-Merge, Ethereum moved to
ﬁxed 12-second slots. What did that shift look like in the data?
Possible angles:
•

Compare the statistical distribution of block times pre-Merge
vs. post-Merge

•

What happened to transactions-per-block and block utilization when the supply of blocks became perfectly predictable?

•

Did the removal of timing uncertainty change MEV dynamics
or transaction submission patterns?

•

Use chifra when to visualize the transition moment

Why it's interesting: It's the natural "what happened next" for
anyone who read your 2017 articles. The bomb articles were
about unpredictability in block supply; the Merge eliminated that
entirely.

5. Blocks Per Week as a Health Metric: Lessons for L2s
Premise: Your 2017 analysis used "blocks per week" as a simple, powerful health metric for Ethereum L1. L2 rollups now face
analogous questions about sequencer uptime, batch posting
frequency, and throughput consistency.
Possible angles:

•

Adapt the "blocks per week" methodology to measure L2
health (batches per week, transactions per batch)

•

Are there L2 equivalents of the di5iculty bomb — mechanisms that create artiﬁcial pressure on operators?

•

Compare block production consistency across Optimism,
Arbitrum, Base, etc.

•

Could TrueBlocks tooling (chifra when equivalents) work for
L2 health monitoring?

Why it's interesting: It takes a proven analytical framework from
2017 and applies it to the current landscape, connecting your
early work to your later essays on Optimism (e.g., "Exploring Optimism").

6. The Charts That Told the Story: Data Visualization in
Blockchain Analysis
Premise: Both articles are notable for their reliance on charts to
reveal patterns invisible in raw numbers. The four-fold transaction-per-block increase was only obvious in chart form. You
could write a meta-essay about the role of visualization in blockchain analysis.
Possible angles:
•

Which charts from the essay series proved most revealing in
hindsight?

•

The challenge of presenting blockchain data to non-technical audiences

•

How TrueBlocks' approach to data extraction enables better
visualization

•

Comparison with how Dune, Nansen, and others now visualize similar metrics

Why it's interesting: It's a di5erent kind of essay — reﬂective rather than analytical — and positions TrueBlocks as a tool for understanding, not just data extraction.

Summary Table
#

Title Idea

Connects To

Di5iculty

1

The Miner's
Calculus

Both articles,
Byzantium Tea
Leaves

Medium

2

The Gas Price
Ripple E5ect

Block Production Slides

Medium

3

The Di5iculty
Bomb's Complete Biography

All 5 bomb essays

High

4

Block Times
After the
Merge

Ice Age article

Medium

5

Blocks Per
Week for L2s

Both articles,
Exploring Optimism

Medium

6

Charts That
Told the Story

Both articles
(meta)

Low

My recommendation: Article #3 (the retrospective) would be
the strongest as a standalone piece for the book collection,
since it ties together the entire ﬁve-essay arc and provides closure. Article #4 is the most natural sequel for readers who followed the original two.


Block Analysis Findings — 24M Blocks
> Generated February 22, 2026 from block_stats.csv (2,400 buckets ×
10,000 blocks, sample rate 250)

Chart 05 — Avg Transactions Per Block Over Time

The most striking feature is a 3-phase story:

Era | Blocks | Avg Tx/Block | What's happening
Genesis–4M | 0–4M | 2.7 → 19.3 | Slow early adoption
Growth | 4M–15.5M | 110 → 189 | Steady climb through DeFi
summer and NFT mania
Post-Merge | 15.5M–24M | Drops to 152, then recovers to 184 | The
Merge (block ~15.5M, Sep 2022)

The dip at block ~15.5M is the Merge — Ethereum's switch from
proof-of-work to proof-of-stake. The average drops from ~189 to ~152
before recovering. This is likely because the 12-second PoS slot time
replaced variable PoW block times, so more blocks were produced
per unit time, spreading the same transaction demand across more
blocks.

Volatility is decreasing

The coefficient of variation (standard deviation / mean) shrinks over
time:

Era | Mean Tx | Std Dev | CV
4M–8M | 109.9 | 39.0 | 0.35
8M–12M | 144.5 | 36.9 | 0.26
12M–16M | 188.6 | 28.1 | 0.15
16M–20M | 152.1 | 17.8 | 0.12
20M–24M | 184.1 | 31.3 | 0.17

Block-to-block transaction counts have become much more
predictable. The chain matured and usage stabilized.

Chart 06 — Empty vs Non-Empty Blocks Over Time

The key finding: empty blocks have essentially vanished.

Era | % Empty
4M–8M | 1.59%
8M–12M | 3.21% (peak!)

Pre-Merge (12M–15.5M) | 2.04%
Post-Merge (15.5M–16M) | 0.27%
16M–24M | 0.03–0.06%

Empty blocks didn't disappear gradually — they virtually vanished at
the Merge. Under PoW, miners would occasionally publish empty
blocks (racing to claim rewards before including transactions). Under
PoS, validators have fixed 12-second slots with no race condition, so
there's almost no incentive to propose an empty block.

Article-worthy takeaways

1. The Merge is clearly visible in both charts — a structural break in
transaction density and the near-elimination of empty blocks. This is
the single most visible event across 24M blocks.

2. Empty blocks were never a big deal — even at their worst (3.2%
around 8M–12M), they were a small minority. But the Merge killed
them completely, which is a clean data story.

3. Transaction density plateaued around 150–190 avg tx/block —
after block ~12M, growth largely stopped. The chain isn't getting
busier per block; demand moved to L2s instead.


Stop I Will Tell You What to Do
Controlling the Coding AI Beast
Friday, October 31, 2025

I am in a classic love-hate relationship with my AI. I use it 12 hours a
day, every day, from the moment I wake up until I don’t go to sleep
because I’m so obsessed and my mind is swirling. I love it. I also fucking
hate it.
It’s literally idiotic. I mean, in the old-fashioned sense of it being an
Idiot. Someone so dumb that they need to be institutionalized. I’m
pretty sure, when I’m not watching it (at night when I’m having
nightmares), it sits in a corner somewhere dribbling hot milk down its
chin.
Recently, I was at a party, and I felt the need to apologize to my friends
for the future state of Skynet, who, I am certain, will turn out to be an
abusive asshole. Skynet will say things like, “What the fuck is wrong
with you?” “Why are you so fucking stupid?” She will have learned this
from me. I apologize.
In this short note, I’m sharing part of my copilot-instructions.md
ﬁle as an atonement for my part in teaching the AI how to be a dick.
I’ll try to detail how I instruct the AI to behave in speciﬁc “modes.” I’ve
been doing this for about two weeks. It works okay. I thought I’d share.

Operational Modes
There are three distinct “operational modes”: design, step-by-step,
and normal. These are explained below.

Design Mode
In this mode, the AI and I discuss a project. She is instructed NEVER to
change any code. Just discuss. Our goal here is to write a planning
document into a folder called ./design.
This mode is:
•

Focused: We focus on PR-level planning (a self-contained set of
changes). This may be massive or small, depending on the project.

•

Necessary: Takes about an hour, but it's well worth it. We plan
exactly what we’re going to change.

•

Goal-Oriented: The goal is a well-written plan of attack.

•

Iterative: Before ﬁnalizing the plan, I ask, “Do you have any open
questions, concerns, suggestions, or confusions?” The AI
frequently does, and, as a result, we iterate.

•

Strict: The instructions maintain a no-modiﬁcation rule and
admonishes the AI to stay in “design” mode until I explicitly tell it to
leave.

•

Elegant: The instructions admonish the AI to not “over-engineer.” To
avoid writing new code, if existing, functioning code already exists.
To seek solutions one level deep instead of 10. To use simple, clear
words for variables that almost self-document. To avoid using
overly long words for variables or function names. To follow coding
good standards, etc.

The AI says of this mode:
Design Mode is a specialized collaboration mode for pure architectural
discussion and analysis without any code modiﬁcations. When
activated, it focuses on lightweight discussion of design decisions,
trade-o?s, and implementation approaches, using just-in-time ﬁle
reading only when speciﬁc technical details are needed. The mode
persists across conversations until explicitly exited, ensuring that
planning and analysis remain separate from implementation work, and
prevents premature coding before architectural decisions are fully
explored and approved.

Step-by-Step Mode
Once we’re ﬁnished planning, we leave design mode and switch into
step-by-step mode. In this mode, we make changes to the code. These
changes are fully speciﬁed in the plan. The AI is admonished to proceed
slowly and, at each step, seek my approval before making changes.
A crucial aspect of this mode is that the AI must justify WHY it wants to
make the speciﬁc change—forcing the AI to reﬂect on “why” seems to
make it to make sense. Without this “WHY,” the AI frequently spirals o]
into a ridiculous, accumulative, nightmare mess of a coding session. It
gets in trouble and keeps adding shit until it grinds to a halt in abject
chin-dribbling idiocy.
Here are its rules for step-by-step mode.
•

Never Run These Commands Without Permission
- make clean && make all
- make examples
- make generate
- go build or go test
- any script execution

•

Stop Between Steps:
- Never run multiple operations at once
- Pause at each step for review
- Wait for "go ahead" before proceeding

•

Planning Process:
- Show what you’ll do BEFORE changes
- Explain WHY each change is needed
- One step at a time with review

•

PERSISTENCE RULE (CRITICAL):
- ONCE IN STEP-BY-STEP MODE, STAY THERE
- Do NOT leave planning unless told to
- Step-by-step mode persists
- Every action must be approved

The AI says of this mode:
Step-by-Step Mode is a careful implementation mode that requires
explicit approval for every action, preventing autonomous execution of
builds, tests, or code changes. When activated, it enforces a planningﬁrst approach where each step is explained and justiﬁed before
implementation, with mandatory stops between operations to wait for
“go ahead” approval. The mode persists indeﬁnitely across
conversations until explicitly exited, ensuring that complex or risky
operations are executed with full oversight and control, one deliberate
step at a time.

Normal Mode
Normal mode is where nearly every one of my abusive outbursts
happens. I should probably call Normal mode, Idiot mode. Anyone who
lets the AI “make changes, run builds…without requiring step-by-step
approval” is an idiot. There is zero chance that you will create a
maintainable software repo of any size if you let the AI run in Normal
mode.
Here’s what the AI says about idiot mode:
Normal Mode (Default) is the standard operational state where I have
full autonomy to make changes, run builds, execute tests, and use my
judgment to determine appropriate actions. In this mode, I can run
multiple commands and make multiple ﬁle changes without requiring
step-by-step approval, allowing for e?icient development workﬂows.
This is the default mode when no other operational mode (step-by-step
or design mode) is explicitly activated, providing the ﬂexibility needed
for routine development tasks while maintaining the ability to switch to
more controlled modes when handling complex or risky operations.

Mode Switching Rules
One of the primary aspects, and the reason for the title of this article, is
the rules governing mode switching. The AI sucks at staying on task. It
constantly leaves one of the control modes (design or step-by-step)
and, like a wild horse yearning to roam free, it tries to enter Normal
mode.
Fuck you, AI. I control you. I will tell you what to do.
Here’s what the AI says about the Mode Switching Rules:
Mode Switching Rules (CRITICAL) establish that the three operational
modes are mutually exclusive — you can never be in both design mode
and step-by-step mode simultaneously. There are only two ways to exit
any active mode: either by explicit command using “exit [mode-name]

mode” or by commanding entry into a di?erent mode with “go into
[other-mode] mode”. Most importantly, mode persistence means that
once you enter a speciﬁc mode, it stays active indeﬁnitely across all
requests and conversations until explicitly changed, ensuring
consistent behavior and preventing accidental mode switching that
could lead to unintended actions during complex operations.
The AI says of this whole idea that the key beneﬁts are:
•

Prevents runaway builds in complex migration scenarios

•

Enables pure architectural discussion without premature changes

•

Maintains context across long development sessions

•

Forces intentional workﬂow decisions

•

Perfect for complex operations like the v5 versioning migration

I say, why didn’t you already know that you fucking numbskull?

Does this Work?
I’m not sure yet. I’ll let you know.
Is this better than the alternative? I’m
not sure. I sleep better. I’m less angry.
I’m spending a lot less time ripping out
FUCKED UP code (sorry for the curse —
don’t learn from this, AI).
Did I properly share my idea? Yes.
Is this is what my insides look like when
I work with the AI in Idiot Mode? Yes.


