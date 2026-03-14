# Derivative Chains in Everyday Life

## A Catalog of Familiar Quantities Linked by Differentiation

### Overview

The chain linking position, velocity, acceleration, jerk, snap, crackle, and pop is one
instance of a general pattern: any measurable quantity that changes over time (or space) has an
infinite chain of derivatives, each describing a successively subtler aspect of how that quantity
evolves. This document catalogs thirty such chains drawn from physics, engineering, economics,
biology, and everyday experience.

In each chain below, each arrow represents differentiation with respect to time unless otherwise
noted.

---

## Mechanical / Motion Chains

### 1. Linear Motion (the Distance-Based Derivative Chain)

$$x \to v \to a \to j \to s \to c \to p$$

Position → Velocity → Acceleration → Jerk → Snap → Crackle → Pop.

This is the chain explored in the companion essay *Snap, Crackle, and Pop*. It is exceptional
in having names through the 6th derivative because human bodies and precision machinery *can*
detect those differences.

### 2. Angular Motion (Rotation)

$$\theta \to \omega \to \alpha \to \zeta$$

Angle → Angular velocity → Angular acceleration → Angular jerk.

Every rotating thing you encounter — wheels, gears, hard drives, figure skaters pulling in
their arms. The angular version of the linear chain.

### 3. Force Chain

$$\text{Impulse} \xleftarrow{\;\int\;} F \xrightarrow{\;\frac{d}{dt}\;} \text{Yank} \xrightarrow{\;\frac{d}{dt}\;} \text{Tug} \xrightarrow{\;\frac{d}{dt}\;} \text{Snatch}$$

Yank ($dF/dt$) matters in crash testing — it's not just the force that injures, it's how fast
force onset is.

### 4. The Integral Side of Position (Absement and Beyond)

$$\cdots \leftarrow \text{Absity} \leftarrow \text{Abseleration} \leftarrow \text{Absement} \leftarrow x \to v \to a \to \cdots$$

Absement ($\int x\, dt$, in meter·seconds) is real and used in engineering: it measures *how
far from zero, for how long*. If you hold a door open 30 cm for 10 seconds, that's 3 m·s of
absement. Fluid valves are controlled by absement — how open, for how long, determines total
flow.

---

## Electrical Chains

### 5. Charge Chain

$$q \to I \to \frac{dI}{dt}$$

Charge → Current → Slew rate.

Every circuit designer cares about this. Slew rate (how fast current changes) determines
inductor voltage ($V = L\,\frac{dI}{dt}$) and is the reason transformers work.

### 6. Magnetic Flux Chain

$$\Phi \to \mathcal{E} \to \frac{d\mathcal{E}}{dt}$$

Magnetic flux → EMF (voltage) → Rate of voltage change.

This *is* Faraday's law. Every electric generator, every induction cooktop.

### 7. Energy Chain

$$E \to P \to \frac{dP}{dt}$$

Energy → Power → "Surge" or "ramp rate."

How fast your power consumption changes matters to the electrical grid. Utilities penalize
industrial customers for high $dP/dt$ ("demand charges").

---

## Thermodynamic Chains

### 8. Temperature Chain

$$T \to \dot{T} \to \ddot{T}$$

Temperature → Heating/cooling rate → Rate of change of heating rate.

Your oven preheating: temperature rises, but *how fast* it rises also changes (it slows as you
approach the set point). That slowdown is the second derivative.

### 9. Entropy Chain

$$S \to \dot{S} \to \ddot{S}$$

Entropy → Entropy production rate → Acceleration of entropy production.

Relevant in irreversible thermodynamics and climate modeling.

---

## Fluid Chains

### 10. Volume / Flow Chain

$$V \to Q \to \dot{Q}$$

Volume → Volumetric flow rate → Flow acceleration.

Turning a faucet: volume fills the sink, flow rate is how fast, and the twist of the handle
controls flow acceleration.

### 11. Pressure Chain

$$p \to \dot{p} \to \ddot{p}$$

Pressure → Pressurization rate → Rate of change of pressurization.

Aircraft cabins are pressurized during climb. The *rate* of pressure change is controlled to
protect eardrums. The *rate of that rate* is controlled for comfort.

---

## Economic / Financial Chains

### 12. Wealth Chain

$$\text{Wealth} \to \text{Income} \to \text{Income growth} \to \text{Growth acceleration}$$

"I'm getting richer" = positive income. "I'm getting richer *faster*" = positive income growth.
"The rate at which I'm getting richer faster is itself speeding up" = growth acceleration
(compound interest does this).

### 13. Price / Inflation Chain

$$\text{Price level} \to \text{Inflation} \to \text{Change in inflation}$$

Inflation is the first derivative of prices. When economists say inflation is "decelerating,"
that's negative second derivative — prices still rising, but the *rate of rise* is slowing.

### 14. Debt / Deficit Chain

$$\text{National debt} \to \text{Deficit} \to \text{Change in deficit}$$

"The deficit is shrinking" means $d(\text{Debt})/dt$ is positive (debt still growing) but
$d^2(\text{Debt})/dt^2$ is negative (it's growing more slowly). Politicians frequently confuse
these levels.

### 15. GDP Chain

$$\text{GDP} \to \text{GDP growth rate} \to \text{Growth acceleration}$$

A "recession" is typically negative first derivative. A "slowdown" is negative second derivative
(still growing, but decelerating).

---

## Demographic Chains

### 16. Population Chain

$$N \to \dot{N} \to \ddot{N}$$

Population → Growth rate → Population acceleration.

The global population growth *rate* peaked around 1968. Since then, the second derivative has
been negative — population is still growing, but the growth is decelerating.

---

## Biological / Medical Chains

### 17. Blood Glucose Chain

$$\text{Glucose level} \to \text{Rate of change} \to \text{Acceleration}$$

Continuous glucose monitors (CGMs) show diabetics not just their glucose level but the *trend
arrow* — the first derivative. Doctors also look at how fast the trend is changing (second
derivative) to predict dangerous spikes.

### 18. Cell Growth / Tumor Chain

$$N_{\text{cells}} \to \text{Growth rate} \to \text{Growth acceleration}$$

A tumor whose growth rate is accelerating is far more dangerous than one growing at a constant
rate.

### 19. Drug Concentration Chain (Pharmacokinetics)

$$C \to \dot{C} \to \ddot{C}$$

Concentration → Absorption/elimination rate → Rate of change of that rate.

The "half-life" of a drug relates to the first derivative. Extended-release formulations are
engineered to control the second derivative.

---

## Geometric / Spatial Chains

These chains are differentiated with respect to arc length or horizontal distance rather than
time.

### 20. Curve Geometry Chain

$$\text{Position on curve} \to \text{Tangent} \to \kappa\;\text{(curvature)} \to \tau\;\text{(torsion)}$$

A straight road has zero curvature. A circular curve has constant curvature. A spiral (highway
on-ramp) has *changing* curvature. Torsion adds the third dimension — a helix.

### 21. Terrain / Elevation Chain

$$\text{Elevation} \to \text{Slope (grade)} \to \text{Curvature of terrain}$$

Hiking: slope is how steep. Curvature is whether it's getting steeper or leveling off. Road
engineers design vertical curves to bound the second derivative for driver sight lines.

---

## Acoustic / Sound Chains

### 22. Sound Displacement Chain

$$\text{Displacement} \to \text{Particle velocity} \to \text{Pressure variation}$$

A speaker cone's displacement, velocity, and acceleration produce the sound wave.

### 23. Loudness Chain

$$\text{Sound level} \to \text{Rate of loudness change} \to \text{Loudness acceleration}$$

A sudden loud noise (gunshot) has extreme loudness acceleration. A gentle crescendo has bounded
first derivative. This is why audio engineers use compressors — to limit the first derivative of
loudness.

---

## Information / Signal Chains

### 24. Data / Bandwidth Chain

$$\text{Data (bits)} \to \text{Throughput (bit rate)} \to \text{Rate of change of throughput}$$

Downloading a file: data is position, bit rate is velocity, and buffering/throttling is
acceleration.

### 25. Signal Chain (Any Measured Signal)

$$f(t) \to f'(t) \to f''(t)$$

In signal processing, the first derivative detects edges. The second derivative detects
*corners* (changes in edge direction). This is how image sharpening and edge detection work.

---

## Psychological / Perceptual Chains

### 26. Hedonic Adaptation Chain

$$\text{Happiness} \to \text{Rate of change of happiness} \to \text{Adaptation rate}$$

Winning the lottery: huge positive first derivative. Within months, the first derivative returns
to zero (hedonic treadmill). The *speed* of that return is the second derivative.

### 27. Pain Perception Chain

$$\text{Pain level} \to \text{Rate of onset} \to \text{Onset acceleration}$$

A slowly building headache vs. a sudden stab. The difference is primarily in the first and
second derivatives of pain intensity.

---

## Chemical Chains

### 28. Concentration / Reaction Chain

$$[\text{Reactant}] \to \text{Reaction rate} \to \text{Rate acceleration}$$

A catalyzed reaction accelerates — its rate increases. An inhibited reaction decelerates. Enzyme
kinetics is largely the study of these derivatives.

---

## Astronomical Chains

### 29. Orbital Position Chain

$$\text{Orbital position} \to \text{Orbital velocity} \to \text{Perturbation acceleration}$$

Planetary perturbations: Jupiter's gravity doesn't just change Mars's velocity — it changes the
*rate of change* of that velocity over orbital timescales.

### 30. Cosmic Scale Factor Chain

$$a(t) \to \dot{a} \to \ddot{a}$$

Scale factor → Hubble expansion rate → Cosmic acceleration.

The discovery that $\ddot{a} > 0$ (the expansion of the universe is *accelerating*) won the
2011 Nobel Prize in Physics. The mystery of *why* is dark energy.

---

## Summary

| Domain | Chains Listed |
|--------|:---:|
| Mechanical / Motion | 4 |
| Electrical | 3 |
| Thermodynamic | 2 |
| Fluid | 2 |
| Economic / Financial | 4 |
| Demographic | 1 |
| Biological / Medical | 3 |
| Geometric / Spatial | 2 |
| Acoustic / Sound | 2 |
| Information / Signal | 2 |
| Psychological / Perceptual | 2 |
| Chemical | 1 |
| Astronomical | 2 |
| **Total** | **30** |

These are the named, well-known chains. In principle, *any* measurable quantity that changes
over time (or space) has an infinite derivative chain. The reason we don't have hundreds of
*named* chains is that most peter out after the 2nd or 3rd derivative — the higher derivatives
either aren't perceptible or aren't useful. The position chain is exceptional in having names
through the 6th derivative precisely because human bodies and precision machinery can detect
those differences.
