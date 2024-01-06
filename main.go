package main

import (
	"fmt"
	"math/rand"
)

type Memes = []string

var memes = Memes{
	"During his own Google interview, Jeff Dean was asked the implications if P=NP were true. He said, \"P = 0 or N = 1\". Then, before the interviewer had even finished laughing, Jeff examined Google's public certificate and wrote the private key on the whiteboard.",
	"Compilers don't warn Jeff Dean. Jeff Dean warns compilers.",
	"The rate at which Jeff Dean produces code jumped by a factor of 40 in late 2000 when he upgraded his keyboard to USB 2.0.",
	"Jeff Dean builds his code before committing it, but only to check for compiler and linker bugs.",
	"When Jeff Dean has an ergonomic evaluation, it is for the protection of his keyboard.",
	"gcc -O4 emails your code to Jeff Dean for a rewrite.",
	"Jeff Dean once failed a Turing test when he correctly identified the 203rd Fibonacci number in less than a second.",
	"The speed of light in a vacuum used to be about 35 mph. Then Jeff Dean spent a weekend optimizing physics.",
	"Jeff Dean was born on December 31, 1969 at 11:48 PM. It took him twelve minutes to implement his first time counter.",
	"Jeff Dean eschews both Emacs and VI. He types his code into zcat, because it's faster that way.",
	"When Jeff Dean sends an ethernet frame there are no collisions because the competing frames retreat back up into the buffer memory on their source nic.",
	"Unsatisfied with constant time, Jeff Dean created the world's first O(1/N) algorithm.",
	"When Jeff Dean goes on vacation, production services across Google mysteriously stop working within a few days.",
	"Jeff Dean was forced to invent asynchronous APIs one day when he optimized a function so that it returned before it was invoked.",
	"When Jeff Dean designs software, he first codes the binary and then writes the source as documentation.",
	"Jeff Dean wrote an O(N2) algorithm once. It was for the Traveling Salesman Problem.",
	"Jeff Dean once implemented a web server in a single printf() call. Other engineers added thousands of lines of explanatory comments but still don't understand exactly how it works. Today that program is the front-end to Google Search.",
	"Jeff once simultaneously reduced all binary sizes by 3% and raised the severity of a previously known low-priority python bug to critical-priority in a single change that contained no python code.",
	"Jeff Dean can beat you at connect four. In three moves.",
	"When your code has undefined behavior, you get a seg fault and corrupted data. When Jeff Dean's code has undefined behavior, a unicorn rides in on a rainbow and gives everybody free ice cream.",
	"When Jeff Dean fires up the profiler, loops unroll themselves in fear.",
	"Jeff Dean is still waiting for mathematicians to discover the joke he hid in the digits of PI.",
	"Jeff Dean's keyboard has two keys: 1 and 0.",
	"When Jeff has trouble sleeping, he Mapreduces sheep.",
	"When Jeff Dean listens to mp3s, he just cats them to /dev/dsp and does the decoding in his head.",
	"When Graham Bell invented the telephone, he saw a missed call from Jeff Dean.",
	"Jeff Dean's watch displays seconds since January 1st, 1970. He is never late.",
	"Jeff starts his programming sessions with cat > /dev/mem.",
	"One day Jeff Dean grabbed his Etch-a-Sketch instead of his laptop on his way out the door. On his way back home to get his real laptop, he programmed the Etch-a-Sketch to play Tetris.",
	"Google search went down for a few hours in 2002, and Jeff Dean started handling queries by hand. Search Quality doubled.",
	"Jeff Dean puts his pants on one leg at a time, but if he had more legs, you would see that his approach is O(log(N)).",
	"The x86-64 spec includes several undocumented instructions marked private use. They are actually for Jeff Dean's use.",
	"Knuth mailed a copy of TAOCP to Google. Jeff Dean autographed it and mailed it back.",
	"When he heard that Jeff Dean's autobiography would be exclusive to the platform, Richard Stallman bought a Kindle.",
	"Jeff Dean once shifted a bit so hard, it ended up on another computer.",
	"Jeff Dean can losslessly compress random data.",
	"When asked if the facts about him are true, Jeff Dean responded \"111111\". While the interviewer was still trying to figure out what he means, he clarified \"every single bit of it is true.\"",
	"Jeff Dean mines bitcoins. In his head.",
	"Jeff Dean knows the last digit of Pi.",
}

func sampleMeme(memes Memes) string {
	return memes[rand.Intn(len(memes))]
}

func main() {
	fmt.Printf("%s\n", sampleMeme(memes))
}
