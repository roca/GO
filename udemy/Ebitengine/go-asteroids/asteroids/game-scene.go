package asteroids

import (
	"fmt"
	"go-asteroids/assets"
	"log"
	"math"
	"math/rand"
	"time"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/solarlune/resolv"
)

const (
	baseMeteorVelocity   = 0.25
	meteorSpawnTime      = 100 * time.Millisecond
	meteorSpeedUpAmount  = 0.1
	meteorSpeedUpTime    = 1000 * time.Millisecond
	cleanUpExplosionTime = 200 * time.Millisecond
	baseBeatWaitTime     = 1600
	numberOfStars        = 1000
	alienAttackTime      = 3 * time.Second
	alienSpawnTime       = 12 * time.Second
	baseAlienVelocity    = 0.5
	maxNumberOfAliens    = 1
)

type GameScene struct {
	player               *Player         // The player
	baseVelocity         float64         // The base velocity for items in the game
	meteorCount          int             // The counter for meteors
	meteorSpawnTimer     *Timer          // The timer for spawning meteors
	meteors              map[int]*Meteor // A map of meteors
	meteorsForLevel      int             // # of meteors for a level
	velocityTimer        *Timer          // Thw timer used for speeding up meteors
	space                *resolv.Space   // The space for all collision objects
	lasers               map[int]*Laser  // A map of lasers
	laserCount           int             // The counter for lasers
	score                int
	explosionSmallSprite *ebiten.Image       // The sprite for the small explosion
	explosionSprite      *ebiten.Image       // The sprite for the explosion
	explosionFrames      []*ebiten.Image     // The frames for the explosion
	cleanupTimer         *Timer              // The timer for cleaning up the explosion
	playerIsDead         bool                // Is the player dead?
	audioContext         *audio.Context      // The audio context for the game
	thrustPlayer         *audio.Player       // The audio player for the thrust sound
	exhaust              *Exhaust            // The exhaust for the player
	laserOnePlayer       *audio.Player       // The audio player for the laser sound one
	laserTwoPlayer       *audio.Player       // The audio player for the laser sound two
	laserThreePlayer     *audio.Player       // The audio player for the laser sound three
	explosionPlayer      *audio.Player       // The audio player for the explosion sound
	beatOnePlayer        *audio.Player       // The audio player for the beat one sound
	beatTwoPlayer        *audio.Player       // The audio player for the beat two sound
	beatTimer            *Timer              // The timer for the beat sound
	beatWaitTime         int                 // The wait time for the beat sound
	playBeatOne          bool                // Is the beat one sound playing?
	stars                []*Star             // The stars in the game
	currentLevel         int                 // The current level the player is on
	shield               *Shield             // The shield for the player
	shieldsUpPlayer      *audio.Player       // The audio player for the shield sound
	alienAttackTimer     *Timer              // The timer for the alien attack
	alienCount           int                 // The counter for the aliens
	alienLaserCount      int                 // The counter for the alien lasers
	alineLaserPlayer     *audio.Player       // The audio player for the alien laser sound
	alienLasers          map[int]*AlienLaser // A map of alien lasers
	alienSoundPlayer     *audio.Player       // The audio player for the alien sound
	alienSpawnTimer      *Timer              // The timer for the alien spawn
	aliens               map[int]*Alien      // A map of aliens
}

func NewGameScene() *GameScene {
	g := &GameScene{
		meteorSpawnTimer:     NewTimer(meteorSpawnTime),
		baseVelocity:         baseMeteorVelocity,
		velocityTimer:        NewTimer(meteorSpeedUpTime),
		meteors:              make(map[int]*Meteor),
		meteorCount:          0,
		meteorsForLevel:      2,
		space:                resolv.NewSpace(ScreenWidth, ScreenHeight, 16, 16),
		lasers:               make(map[int]*Laser),
		laserCount:           0,
		explosionSprite:      assets.ExplosionSprite,
		explosionSmallSprite: assets.ExplosionSmallSprite,
		cleanupTimer:         NewTimer(cleanUpExplosionTime),
		beatTimer:            NewTimer(2 * time.Second),
		beatWaitTime:         baseBeatWaitTime,
		stars:                GenerateStars(numberOfStars),
		currentLevel:         1,
		aliens:               make(map[int]*Alien),
		alienCount:           0,
		alienLasers:          make(map[int]*AlienLaser),
		alienLaserCount:      0,
		alienSpawnTimer:      NewTimer(alienSpawnTime),
		alienAttackTimer:     NewTimer(alienAttackTime),
	}
	g.player = NewPlayer(g)
	g.space.Add(g.player.playerObj)

	g.explosionFrames = assets.Explosion

	// Load Audio
	g.audioContext = audio.NewContext(48000)

	thrustPlayer, _ := g.audioContext.NewPlayer(assets.ThrustSound)
	g.thrustPlayer = thrustPlayer

	laserOnePlayer, _ := g.audioContext.NewPlayer(assets.LaserOneSound)
	g.laserOnePlayer = laserOnePlayer

	laserTwoPlayer, _ := g.audioContext.NewPlayer(assets.LaserTwoSound)
	g.laserTwoPlayer = laserTwoPlayer

	laserThreePlayer, _ := g.audioContext.NewPlayer(assets.LaserThreeSound)
	g.laserThreePlayer = laserThreePlayer

	explosionPlayer, _ := g.audioContext.NewPlayer(assets.ExplosionSound)
	g.explosionPlayer = explosionPlayer

	beatOnePlayer, _ := g.audioContext.NewPlayer(assets.BeatOneSound)
	g.beatOnePlayer = beatOnePlayer

	beatTwoPlayer, _ := g.audioContext.NewPlayer(assets.BeatTwoSound)
	g.beatTwoPlayer = beatTwoPlayer

	shieldsUpPlayer, _ := g.audioContext.NewPlayer(assets.ShieldSound)
	g.shieldsUpPlayer = shieldsUpPlayer

	alienLaserPlayer, _ := g.audioContext.NewPlayer(assets.AlienLaserSound)
	g.alineLaserPlayer = alienLaserPlayer

	alienSoundPlayer, _ := g.audioContext.NewPlayer(assets.AlienSound)
	alienSoundPlayer.SetVolume(0.5)
	g.alienSoundPlayer = alienSoundPlayer

	return g
}

func (g *GameScene) Update(state *State) error {
	g.player.Update()

	g.updateExhaust()

	g.updateShield()

	g.isPlayerDying()

	g.isPlayerDead(state)

	g.spawnMeteors()

	g.spawnAliens()

	for _, a := range g.aliens {
		a.Update()
	}

	g.letAliensAttack()

	for _, al := range g.alienLasers {
		al.Update()
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, l := range g.lasers {
		l.Update()
	}

	g.speedUpMeteors()

	g.isPlayerCollidingWithMeteor()

	g.isMeteorHitByPlayerLaser()

	// Check if the player is colliding with a alien
	g.isPlayerCollidingWithAlien()

	// Check if alien laser is colliding with the player
	g.isPlayerHitByAlienLaser()

	// Check if player laser is colliding with alien
	g.isAlienHitByPlayerLaser()

	g.cleanupMeteorsAndAliens()

	g.beatSound()

	g.isLevelComplete(state)

	g.removeOffScreenAliens()

	g.removeOffScreenLasers()

	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	// Draw the stars
	for _, s := range g.stars {
		s.Draw(screen)
	}

	g.player.Draw(screen)

	// Draw the exhaust
	if g.exhaust != nil {
		g.exhaust.Draw(screen)
	}

	// Draw the shield
	if g.shield != nil {
		g.shield.Draw(screen)
	}

	// Draw the meteors
	for _, m := range g.meteors {
		m.Draw(screen)
	}

	// Draw the lasers
	for _, l := range g.lasers {
		l.Draw(screen)
	}

	// Draw life indicators
	if len(g.player.lifeIndicators) > 0 {
		for _, li := range g.player.lifeIndicators {
			li.Draw(screen)
		}
	}

	// Draw shield indicators
	if len(g.player.shieldIndicators) > 0 {
		for _, si := range g.player.shieldIndicators {
			si.Draw(screen)
		}
	}

	// Draw the hyperspace indicator
	if g.player.hyperSpaceTimer == nil || g.player.hyperSpaceTimer.IsReady() {
		g.player.hyperSpaceIndicator.Draw(screen)
	}

	// Draw the aliens
	for _, a := range g.aliens {
		a.Draw(screen)
	}

	// Draw the alien lasers
	for _, al := range g.alienLasers {
		al.Draw(screen)
	}

	// Update and draw the score
	textToDraw := fmt.Sprintf("%06d", g.score)
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(ScreenWidth/2, 40)
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.ScoreFont,
		Size:   24,
	}, op)

	// Update and draw the High score
	if g.score > highScore {
		highScore = g.score
	}

	textToDraw = fmt.Sprintf("HIGH SCORE %06d", highScore)
	op = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(ScreenWidth/2, 75)
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.ScoreFont,
		Size:   16,
	}, op)

	// Update and draw the Current level
	textToDraw = fmt.Sprintf("Level %d", g.currentLevel)
	op = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight-40)
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.LevelFont,
		Size:   16,
	}, op)

}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *GameScene) isPlayerCollidingWithAlien() {
	for _, a := range g.aliens {
		if a.alienObj.IsIntersecting(g.player.playerObj) {
			if !g.player.isShielded {
				if !g.explosionPlayer.IsPlaying() {
					_ = g.explosionPlayer.Rewind()
					g.explosionPlayer.Play()
				}
				g.player.isDying = true
			} else {
				// Bounce the alien off the player
				g.bounceAlien(a)
			}
		}
	}
}

func (g *GameScene) isPlayerHitByAlienLaser() {
	for _, al := range g.alienLasers {
		if al.lasterObj.IsIntersecting(g.player.playerObj) {
			if !g.player.isShielded {
				if !g.explosionPlayer.IsPlaying() {
					_ = g.explosionPlayer.Rewind()
					g.explosionPlayer.Play()
				}
				g.player.isDying = true
			} else {
				// Bounce the alien laser off the player
				g.bounceAlienLaser(al)
			}
		}
	}
}

func (g *GameScene) isAlienHitByPlayerLaser() {
	for _, a := range g.aliens {
		for _, l := range g.lasers {
			if a.alienObj.IsIntersecting(l.lasterObj) {
				laserData := l.lasterObj.Data().(*ObjectData)
				delete(g.alienLasers, laserData.index)
				g.space.Remove(l.lasterObj)
				a.sprite = g.explosionSprite
				g.score = g.score + 50
				if !g.explosionPlayer.IsPlaying() {
					_ = g.explosionPlayer.Rewind()
					g.explosionPlayer.Play()
				}
			}
		}
	}
}

func (g *GameScene) letAliensAttack() {
	if len(g.aliens) > 0 {
		if !g.alienSoundPlayer.IsPlaying() {
			_ = g.alienSoundPlayer.Rewind()
			g.alienSoundPlayer.Play()
		}

		// Update the alien attack timer.
		g.alienAttackTimer.Update()

		// Is the timer reached ? If so, reset the timer and attack.
		if g.alienAttackTimer.IsReady() {
			g.alienAttackTimer.Reset()

			for _, a := range g.aliens {
				bounds := a.sprite.Bounds()
				halfW := float64(bounds.Dx()) / 2
				halfH := float64(bounds.Dy()) / 2

				var degreesRadian float64

				// Is the alien intelligent ?
				if !a.isIntelligent {
					// Fire in a random direction.
					degreesRadian = rand.Float64() * (math.Pi * 2)
				} else {
					// Fire with some accuracy.
					degreesRadian = math.Atan2(g.player.position.Y-a.position.Y, g.player.position.X-a.position.X)
					degreesRadian = degreesRadian - math.Pi*-0.5
				}

				r := degreesRadian

				offsetX := float64(a.sprite.Bounds().Dx() - int(halfW))
				offsetY := float64(a.sprite.Bounds().Dy() - int(halfH))

				spawnPos := Vector{
					X: a.position.X + halfW + math.Sin(r) - offsetX,
					Y: a.position.Y + halfH + math.Cos(r) - offsetY,
				}

				laser := NewAlienLaser(spawnPos, r)
				g.alienCount++
				g.alienLasers[g.alienCount] = laser
				if !g.alineLaserPlayer.IsPlaying() {
					_ = g.alineLaserPlayer.Rewind()
					g.alineLaserPlayer.Play()
				}
			}
		}

	}
}

func (g *GameScene) removeOffScreenLasers() {
	for i, l := range g.lasers {
		if l.position.X > ScreenWidth+200 ||
			l.position.Y > ScreenHeight+200 ||
			l.position.X < -200 ||
			l.position.Y < -200 {
			g.space.Remove(l.lasterObj)
			delete(g.lasers, i)
		}
	}

	for i, l := range g.alienLasers {
		if l.position.X > ScreenWidth+200 ||
			l.position.Y > ScreenHeight+200 ||
			l.position.X < -200 ||
			l.position.Y < -200 {
			g.space.Remove(l.lasterObj)
			delete(g.alienLasers, i)
		}
	}
}

func (g *GameScene) spawnAliens() {
	//TODO: spawnAliens()
	g.alienSpawnTimer.Update()
	if len(g.aliens) < maxNumberOfAliens {
		if g.alienSpawnTimer.IsReady() {
			g.alienSpawnTimer.Reset()
			rnd := rand.Intn(100-1) + 1
			if rnd > 50 {
				a := NewAlien(baseAlienVelocity, g)
				g.space.Add(a.alienObj)
				g.alienCount++
				g.aliens[g.alienCount] = a
			}
		}
	}
}

func (g *GameScene) removeOffScreenAliens() {
	for i, a := range g.aliens {
		if a.position.X > ScreenWidth+200 ||
			a.position.Y > ScreenHeight+200 ||
			a.position.X < -200 ||
			a.position.Y < -200 {
			g.space.Remove(a.alienObj)
			delete(g.aliens, i)
		}
	}
}

func (g *GameScene) updateShield() {
	if g.shield != nil {
		g.shield.Update()
	}
}

func (g *GameScene) isLevelComplete(state *State) {
	if g.meteorCount >= g.meteorsForLevel && len(g.meteors) == 0 {
		g.baseVelocity = baseMeteorVelocity
		g.currentLevel++

		if g.currentLevel%5 == 0 {
			if g.player.livesRemaining < 6 {
				g.player.livesRemaining++
				x := float64(20 + len(g.player.lifeIndicators)*50.0)
				y := 20.0
				g.player.lifeIndicators = append(g.player.lifeIndicators, NewLifeIndicator(Vector{X: x, Y: y}))
			}
		}

		g.beatWaitTime = baseBeatWaitTime

		state.SceneManager.GoToScene(&LevelStartsScene{
			game:           g,
			nextLevelTimer: NewTimer(2 * time.Second),
			stars:          GenerateStars(numberOfStars),
		})
	}
}

func (g *GameScene) beatSound() {
	g.beatTimer.Update()
	if g.beatTimer.IsReady() {
		if g.playBeatOne {
			_ = g.beatOnePlayer.Rewind()
			g.beatOnePlayer.Play()
			g.beatTimer.Reset()
		} else {
			_ = g.beatTwoPlayer.Rewind()
			g.beatTwoPlayer.Play()
			g.beatTimer.Reset()
		}
		g.playBeatOne = !g.playBeatOne

		// Speed up the timer
		if g.beatWaitTime > 400 {
			g.beatWaitTime -= 25
			g.beatTimer = NewTimer(time.Duration(g.beatWaitTime) * time.Millisecond)
		}
	}
}

func (g *GameScene) updateExhaust() {
	if g.exhaust != nil {
		g.exhaust.Update()
	}
}

func (g *GameScene) isMeteorHitByPlayerLaser() {
	for _, m := range g.meteors {
		for _, l := range g.lasers {
			if m.meteorObj.IsIntersecting(l.lasterObj) {
				if m.meteorObj.Tags().Has(TagSmall) {
					// Small meteor hit
					m.sprite = g.explosionSmallSprite
					g.score++

					if !g.explosionPlayer.IsPlaying() {
						_ = m.game.explosionPlayer.Rewind()
						m.game.explosionPlayer.Play()
					}
				} else {
					// Large meteor hit
					oldPos := m.position
					m.sprite = g.explosionSprite
					g.score++

					if !g.explosionPlayer.IsPlaying() {
						_ = m.game.explosionPlayer.Rewind()
						m.game.explosionPlayer.Play()
					}

					numToSpawn := rand.Intn(numberOfSmallMeteorsFromLargeMeteor)
					for i := 0; i < numToSpawn; i++ {
						meteor := NewSmallMeteor(baseMeteorVelocity, g, len(m.game.meteors)-1)
						meteor.position = Vector{
							oldPos.X + float64(rand.Intn(100-50)+50),
							oldPos.Y + float64(rand.Intn(100-50)+50),
						}
						meteor.meteorObj.SetPosition(meteor.position.X, meteor.position.Y)
						g.space.Add(meteor.meteorObj)
						g.meteorCount++
						g.meteors[m.game.meteorCount] = meteor
					}
				}
			}
		}
	}
}

func (g *GameScene) isPlayerDying() {
	if g.player.isDying {
		g.player.dyingTimer.Update()
		if g.player.dyingTimer.IsReady() {
			g.player.dyingTimer.Reset()
			g.player.dyingCounter++
			if g.player.dyingCounter == 12 {
				g.player.isDying = false
				g.player.isDead = true
			} else if g.player.dyingCounter < 12 {
				g.player.sprite = g.explosionFrames[g.player.dyingCounter]
			} else {
				// Do nothing.
			}
		}
	}
}

func (g *GameScene) isPlayerDead(state *State) {
	if g.player.isDead {
		g.player.livesRemaining--
		if g.player.livesRemaining == 0 {

			// New High Score ?
			log.Println("Score:", g.score, ", High Score:", highScore)
			if g.score > originalHighScore {
				err := updateHighScore(g.score)
				if err != nil {
					log.Println("Error updating high score:", err)
				}
			}

			state.SceneManager.GoToScene(&GameOverScene{
				game:        g,
				meteors:     make(map[int]*Meteor),
				meteorCount: 5,
				stars:       GenerateStars(numberOfStars),
			})
		} else {
			score := g.score
			livesRemaining := g.player.livesRemaining
			lifeSlice := g.player.lifeIndicators[:len(g.player.lifeIndicators)-1]
			stars := g.stars
			shieldsRemaining := g.player.shieldsRemaining
			shieldIndicatorSlice := g.player.shieldIndicators

			g.Reset()

			g.player.livesRemaining = livesRemaining
			g.score = score
			g.player.lifeIndicators = lifeSlice
			g.stars = stars
			g.player.shieldsRemaining = shieldsRemaining
			g.player.shieldIndicators = shieldIndicatorSlice
		}
	}
}

func (g *GameScene) spawnMeteors() {
	g.meteorSpawnTimer.Update()

	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		if len(g.meteors) < g.meteorsForLevel && g.meteorCount < g.meteorsForLevel {
			m := NewMeteor(g.baseVelocity, g, len(g.meteors)-1)
			g.space.Add(m.meteorObj)
			g.meteorCount++
			g.meteors[g.meteorCount] = m
		}
	}
}

func (g *GameScene) speedUpMeteors() {
	g.velocityTimer.Update()

	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}
}

func (g *GameScene) isPlayerCollidingWithMeteor() {
	for _, m := range g.meteors {
		if m.meteorObj.IsIntersecting(g.player.playerObj) {
			if !g.player.isShielded {
				m.game.player.isDying = true

				if !g.explosionPlayer.IsPlaying() {
					_ = m.game.explosionPlayer.Rewind()
					m.game.explosionPlayer.Play()
				}
				break
			} else {
				// Bounce the meteor off the player
				g.bounceMeteor(m)
			}
		}
	}
}

func (g *GameScene) bounceMeteor(m *Meteor) {
	direction := Vector{
		X: (ScreenWidth/2 - m.position.X) * -1,
		Y: (ScreenHeight/2 - m.position.Y) * -1,
	}
	normalizedDirection := direction.Normalize()
	velocity := g.baseVelocity

	movement := Vector{
		X: normalizedDirection.X * velocity,

		Y: normalizedDirection.Y * velocity,
	}

	m.movement = movement
}

func (g *GameScene) bounceAlien(a *Alien) {
	direction := Vector{
		X: (ScreenWidth/2 - a.position.X) * -1,
		Y: (ScreenHeight/2 - a.position.Y) * -1,
	}
	normalizedDirection := direction.Normalize()
	velocity := g.baseVelocity

	movement := Vector{
		X: normalizedDirection.X * velocity,

		Y: normalizedDirection.Y * velocity,
	}

	a.movement = movement
}

func (g *GameScene) bounceAlienLaser(al *AlienLaser) {
	al.rotation = math.Atan2(g.player.position.Y-al.position.Y, g.player.position.X-al.position.X)
}

func (g *GameScene) cleanupMeteorsAndAliens() {
	g.cleanupTimer.Update()
	if g.cleanupTimer.IsReady() {
		for i, m := range g.meteors {
			if m.sprite == g.explosionSprite || m.sprite == g.explosionSmallSprite {
				delete(g.meteors, i)
				g.space.Remove(m.meteorObj)
			}
		}

		for i, a := range g.aliens {
			if a.sprite == g.explosionSprite {
				delete(g.aliens, i)
				g.space.Remove(a.alienObj)
			}
		}
		g.cleanupTimer.Reset()
	}
}

func (g *GameScene) Reset() {
	g.player = NewPlayer(g)
	g.meteors = make(map[int]*Meteor)
	g.meteorCount = 0
	g.lasers = make(map[int]*Laser)
	g.laserCount = 0
	g.score = 0
	g.meteorSpawnTimer.Reset()
	g.baseVelocity = baseMeteorVelocity
	g.velocityTimer.Reset()
	g.playerIsDead = false
	g.exhaust = nil
	g.space.RemoveAll()
	g.space.Add(g.player.playerObj)
	g.stars = GenerateStars(numberOfStars)
	g.player.shieldsRemaining = numberOfShields
	g.player.isShielded = false
	g.aliens = make(map[int]*Alien)
	g.alienCount = 0
	g.alienLasers = make(map[int]*AlienLaser)
	g.alienLaserCount = 0
}
