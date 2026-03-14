# Plan for Book II — *The Hidden Architecture*

> **Variation attributes** (arc, ending, structure, entry, register, setting, math\_visibility)
> are maintained in the companion file `Plan for Math Book II — attributes.yaml`.
> The tables below carry only content columns; the YAML is the pipeline's source of truth.

## Thesis

The structures you can't see that shape everything: probability, shape, information,
optimization, curves. Mathematics as the invisible scaffolding beneath familiar
experience. Start with leading digits. Expand through probability, topology, signals,
strategy, and curves. End with a thread, a spiral, and the quiet mathematics of
living systems.

Where Book I said "look, there's math in your cereal bowl," Book II says "the thing
you can't see — the pattern beneath the pattern — is doing most of the work." The
reader who finished Book I noticing math everywhere now discovers the *structures*
that organize what they were noticing.

**Character:** More intellectual, more "wait, really?" Some essays resolve, some
provoke. The reader finishes sensing patterns that were always there.

## The Storyline in One Sentence

> *Behind every familiar experience — the lottery, the weather forecast, your
> stock portfolio, a coiled seashell — there is an invisible architecture you
> were never taught to see, and it is breathtakingly elegant.*

---

## Book Introduction

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | introduction | intro-book-2 | The Architecture You Can't See | How invisible structures organize the world you already know | —  |

---

## Part 1: The Numbers Behind the Numbers

*Before you look at the world, look at the numbers themselves. They have their own
secret life. Leading digits follow a law. Wealth obeys a distribution. Growth fools
your intuition. Shuffling has a sharp phase transition. And metabolic rate scales
with a power law that gives every mammal the same number of heartbeats. Five essays
that say: the numbers are not random — they have architecture.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-1 | The Numbers Behind the Numbers | - | -  |
| 1 | essay | benfords-law | Benford's Law | 30% of real-world numbers start with 1 | Scale invariance forces a logarithmic distribution; detects fraud  |
| 2 | essay | pareto-distribution | The Pareto Distribution | 80/20 in wealth, words, cities, earthquakes | Power-law tails: most distributions are not Gaussian  |
| 3 | essay | exponential-growth | Exponential Growth and Your Intuition | The pond is half-empty one day before it's full | Humans linearize growth; exponentials ambush us every time  |
| 4 | essay | shuffling-seven | Why Shuffling Seven Times Is Enough | Card players argue about how much to shuffle | Riffle shuffle has a sharp phase transition at 3/2 log₂(n) — below it, order; above, chaos  |
| 5 | essay | kleibers-law | Kleiber's Law: Why Elephants Live Longer | An elephant and a shrew have the same heartbeat count | Metabolic rate ∝ M^(3/4), not M^(2/3) — fractal vascular networks explain why  |

---

## Part 2: Probability Against Intuition

*Your intuition about probability is wrong. Not a little wrong — systematically,
reliably, spectacularly wrong. Five essays that recalibrate: the Monty Hall door
you should switch to, the false positive that isn't what you think, the regression
that fools coaches and doctors, the 37% rule for commitment, and the birthday room
that defies counting.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-2 | Probability Against Intuition | - | -  |
| 1 | essay | monty-hall | The Monty Hall Problem | Should you switch doors? | Conditional probability: switching wins 2/3 of the time  |
| 2 | essay | bayes-doctor | Bayes' Theorem at the Doctor's Office | Your test came back positive. How worried should you be? | False positive paradox: P(disease|positive) is shockingly low for rare conditions  |
| 3 | essay | regression-mean | Regression to the Mean | Supplements "work," hot streaks "cool off" | Not a force — it's a statistical artifact of imperfect correlation  |
| 4 | essay | secretary-problem | The Secretary Problem | Interview candidates one by one — when do you commit? | Optimal stopping: reject the first 1/e ≈ 37%, then take the next best  |
| 5 | essay | birthday-paradox | The Birthday Paradox | 23 people, >50% chance of a shared birthday | Pairwise combinations grow quadratically; our brains think linearly  |

---

## Part 3: Shape and Space

*Now look at shape itself. A coffee mug is a donut. Stacking oranges is a 400-year
hard problem. The nearest-neighbor map is everywhere. A walk across bridges invented
a field. And a strip of paper has only one side. Five essays where geometry stops
being about measurement and starts being about identity.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-3 | Shape and Space | - | -  |
| 1 | essay | topology-coffee-mug | The Topology of Your Coffee Mug | A mug and a donut are the same shape | Genus classification: topology cares about holes, not angles  |
| 2 | essay | packing-oranges | Packing Oranges | The grocer's pyramid — is it optimal? | Kepler's conjecture proved (Hales 2017): face-centered cubic is densest  |
| 3 | essay | voronoi-diagrams | Voronoi Diagrams | School districts, cell towers, giraffe spots | Nearest-neighbor partition: Voronoi appears in nature, logistics, and Dirichlet's mathematics  |
| 4 | essay | eulers-bridges | Euler's Bridges and the Birth of Topology | Can you cross every bridge in Königsberg exactly once? | First theorem of graph theory — Euler proved it impossible and invented a field. Enriched with Euler biography.  |
| 5 | essay | mobius-strip | The Möbius Strip | One side, one edge — cut it and it doesn't split | Non-orientability, Euler characteristic, why it matters for conveyor belts and recycling symbols  |

---

## Part 4: Signals and Information

*Information has a physics. Sound has a sampling rate. Power has imaginary numbers.
Weather has memory (but only one day). And the most beautiful equation in mathematics
connects five constants. Five essays on the hidden information layer of daily life.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-4 | Signals and Information | - | -  |
| 1 | essay | shannon-entropy | Shannon Entropy | Information is surprise | H = −Σ p log p: the less likely a message, the more information it carries  |
| 2 | essay | nyquist-sampling | Nyquist Sampling in Digital Music | Why CDs use 44,100 samples per second | Sample at 2× the highest frequency or you get aliasing — Shannon-Nyquist theorem  |
| 3 | essay | euler-formula-ac | Euler's Formula in AC Power | Imaginary numbers are in your walls | e^(iθ) rotates a vector — power engineers use complex exponentials daily  |
| 4 | essay | markov-chains-weather | Markov Chains in the Weather Forecast | Tomorrow's weather depends on today — not yesterday | Memoryless chains, transition matrices, stationary distributions  |
| 5 | essay | eulers-identity | Euler's Identity | $e^{i\pi} + 1 = 0$ — five constants, one equation | Why it works: rotation in the complex plane links analysis, algebra, and geometry. Enriched with Euler biography.  |

---

## Part 5: Strategy and Tradeoffs

*When you optimize, you discover that the universe fights back. The traveling
salesman can't be solved fast. The optimal bet is smaller than you'd think. And
more processors won't help if the problem has a serial bottleneck. Three sharp
essays about hitting walls.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-5 | Strategy and Tradeoffs | - | -  |
| 1 | essay | traveling-salesman | The Traveling Salesman Problem | Planning the shortest errand route | NP-hard: the number of routes grows factorially, no shortcut known  |
| 2 | essay | kelly-criterion | The Kelly Criterion | How much should you bet? | Optimal growth rate = edge / odds — bet too much and you go broke  |
| 3 | essay | amdahls-law | Amdahl's Law | Trying to get ready faster by parallelizing | Speedup ≤ 1/(s + (1−s)/N) — the serial fraction dominates; too many cooks  |

---

## Part 6: The Gentle Curves

*Curves are mathematics made visible. A seashell grows without changing shape.
Kittens chasing each other trace a logarithmic spiral. A pendulum draws Lissajous
figures. A Spirograph recapitulates Ptolemy. Archimedes found beauty in nested
semicircles. The ocean adds sine waves. Six essays where the math is the picture.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-6 | The Gentle Curves | - | -  |
| 1 | essay | logarithmic-spiral | The Logarithmic Spiral of Seashells | Growth without change of shape | Self-similar curve: r = ae^(bθ) — the nautilus, the hurricane, the galaxy arm  |
| 2 | essay | curves-of-pursuit | Curves of Pursuit: Kittens Chasing | Four kittens at corners, each chasing the next | Pursuit curves converge in logarithmic spirals — the geometry of chasing  |
| 3 | essay | lissajous-curves | Lissajous Curves in Pendulum Art | Two pendulums, one pen, one picture | Frequency ratios create closed curves; irrational ratios fill the plane  |
| 4 | essay | spirograph-epicycloids | Spirograph and Epicycloids | The toy that teaches Fourier series | Epicycloids = Ptolemy's deferent-and-epicycle model — the math is identical  |
| 5 | essay | arbelos | The Arbelos of Archimedes | Three semicircles, a world of surprises | Nested semicircles yield Pappus chains, twin circles, and area equalities — pure geometric beauty  |
| 6 | essay | sine-waves-tides | Sine Waves in Ocean Tides | Why tide tables work at all | Tidal harmonics: superposition of ~37 lunar-solar sine components  |

---

## Part 7: Thread and Fabric

*A short, textured part. Crochet makes hyperbolic space physical. Braids have
an algebra. Stock prices wander randomly. Three essays about what you can hold
in your hands — or can't.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-7 | Thread and Fabric | - | -  |
| 1 | essay | hyperbolic-crochet | Hyperbolic Crochet | A physical model of a geometry Euclid couldn't imagine | Constant negative curvature: the crochet coral reef project makes it tangible  |
| 2 | essay | math-of-braiding | The Mathematics of Braiding | Three strands, infinite complexity | Artin braid group: braids form an algebraic group with deep connections to knot theory  |
| 3 | essay | random-walks | Random Walks on Your Stock Portfolio | Prices wander like a drunk on a number line | Brownian motion, Wiener process, why you can't predict tomorrow's close  |

---

## Part 8: Nature's Other Ledger

*The book closes with a return to nature — but now the reader sees the structures,
not just the surfaces. An aspen trembles because of eigenvalues. A vine knows
its chirality. Spider silk solves a multi-objective optimization problem. Perfume
obeys the diffusion equation. And a wildflower meadow is a bell curve. Five
essays that bring the hidden architecture home.*

| # | Type | Slug | Title | Hook | Hidden Math |
|---|------|------|-------|------|-------------|
| — | section | sec-2-8 | Nature's Other Ledger | - | -  |
| 1 | essay | aspens-tremble | Why Aspens Tremble and Oaks Don't | Every leaf on an aspen shivers in the lightest breeze | Eigenvalue flutter: the aspen petiole is flat, placing the resonant frequency in the wind band  |
| 2 | essay | vines-twist | How Vines Know to Twist | Morning glories always coil counterclockwise | Helical geometry, chirality, differential growth on left vs. right  |
| 3 | essay | spider-silk | Why Spider Silk Outperforms Steel | Stronger than steel, tougher than Kevlar — spun at room temperature | Multi-objective optimization: evolution found a point on the strength/elasticity Pareto frontier  |
| 4 | essay | diffusion-perfume | The Diffusion Equation in Perfume | Spray perfume; it fills the room without wind | Fick's laws: concentration gradient drives flow, Gaussian bell spreading over time  |
| 5 | essay | bell-curve-wildflowers | The Bell Curve in a Wildflower Meadow | Height, bloom time, petal count — all normally distributed | The CLT in nature: many small independent forces produce a Gaussian  |

---

## Summary

| Part | Title | Essays | Sections | Introductions |
|:----:|-------|:------:|:--------:|:-------------:|
| — | Book Introduction | — | — | 1 |
| 1 | The Numbers Behind the Numbers | 5 | 1 | — |
| 2 | Probability Against Intuition | 5 | 1 | — |
| 3 | Shape and Space | 5 | 1 | — |
| 4 | Signals and Information | 5 | 1 | — |
| 5 | Strategy and Tradeoffs | 3 | 1 | — |
| 6 | The Gentle Curves | 6 | 1 | — |
| 7 | Thread and Fabric | 3 | 1 | — |
| 8 | Nature's Other Ledger | 5 | 1 | — |
| | **Total** | **37** | **8** | **1** |

## Arc Distribution

| Arc | Count |
|-----|:-----:|
| Slow Build | 7 |
| Cold Open Mystery | 4 |
| Paradox | 3 |
| Demolition | 2 |
| Walk | 4 |
| Zoom | 2 |
| Countdown | 2 |
| Recipe | 2 |
| Catalog of Wonders | 2 |
| Dual Timeline | 2 |
| Confession | 2 |
| Argument | 1 |
| Letter | 1 |
| Inheritance | 2 |
| Debate | 0 |

## Ending Distribution

| Ending | Count |
|--------|:-----:|
| Resolution | 11 |
| Awe | 10 |
| Communion | 7 |
| Provocation | 5 |
| Honesty | 2 |
| Handoff | 3 |

*Note: Book II shifts toward more Provocation and Handoff endings than Book I,
reflecting its role as the structural middle — the reader has confidence now
and can handle more open threads.*
