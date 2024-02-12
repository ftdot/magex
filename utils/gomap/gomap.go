package gomap

import (
	"fmt"
	"sort"

	"github.com/ftdot/magex/game"
	"github.com/ftdot/magex/utils/gomap/goi"
	"github.com/ftdot/magex/utils/queue"

	"github.com/hajimehoshi/ebiten/v2"
)

// GOMap stands for GameObject Map.
// GameObject map defines map with the game objects.
// Also, defines all the GameObjectBase functional,
// functions like Start(), Update(), Draw(), etc.
type GOMap struct {
	// Map with objects that implements Awake() function.
	//
	// Has following format:
	//
	// KEY: name of game object (string) = VALUE: Awakable (interface)
	mapAs map[string]goi.Awakable
	// List with the queued startables. List clears on first Scene.Enter() call.
	sliceSQs []goi.Startable
	// Map with objects that implements PhysUpdate() function.
	//
	// Has following format:
	//
	// KEY: name of game object (string) = VALUE: PhysUpdatable (interface)
	mapPUs map[string]goi.PhysUpdatable
	// Map with objects that implements Phys2Update() function.
	//
	// Has following format:
	//
	// KEY: name of game object (string) = VALUE: Phys2Updatable (interface)
	mapP2Us map[string]goi.Phys2Updatable
	// Map with objects that implements Update() function.
	//
	// Has following format:
	//
	// KEY: name of game object (string) = VALUE: Updatable (interface)
	mapUs map[string]goi.Updatable
	// Map with objects that implements Draw() and DrawPriority() functions.
	//
	// Has following format:
	//
	// KEY: name of game object (string) = VALUE: DrawableQueued (interface)
	mapDQs map[string]goi.DrawableQueued
	// Map with objects that implements DrawUI() and DrawUIPriority() functions.
	//
	// Has following format:
	//
	// KEY: name of game object (string) = VALUE: DrawableUIQueued (interface)
	mapDUIQs map[string]goi.DrawableUIQueued

	gameObjects map[string]any // Map with the game objects
	gameBase    *game.GameBase // Can be nil, if gameBase isn't provided
}

// Creates new empty GOMap.
func New() *GOMap {
	return &GOMap{
		mapAs:       map[string]goi.Awakable{},
		sliceSQs:    []goi.Startable{},
		mapPUs:      map[string]goi.PhysUpdatable{},
		mapP2Us:     map[string]goi.Phys2Updatable{},
		mapUs:       map[string]goi.Updatable{},
		mapDQs:      map[string]goi.DrawableQueued{},
		mapDUIQs:    map[string]goi.DrawableUIQueued{},
		gameObjects: map[string]any{},
		gameBase:    nil,
	}
}

////

// Must be called before all the operations.
func (m *GOMap) SetupGame(gameBase *game.GameBase) {
	m.gameBase = gameBase
}

////

// Tries to get the game object by given name.
func (m *GOMap) Get(name string) (i interface{}, ok bool) {
	i, ok = m.gameObjects[name]
	return
}

// Registers a new game object. If Awake() function implemented,
// it will be called. This will be applied to all the components
// of the game object.
func (m *GOMap) Register(name string, gameObject interface{}) error {
	if c, ok := gameObject.(goi.WithComponents); ok {
		for _, c := range c.GetComponents() {
			if err := m.register(c.GetID(), c); err != nil {
				return err
			}
		}
	}

	if err := m.register(name, gameObject); err != nil {
		return err
	}

	m.gameObjects[name] = gameObject

	return nil
}

// Unregisters the game object. If Destroy() function implemented,
// it will be called. This will be applied to all the components
// of the game object.
func (m *GOMap) Unregister(name string) error {

	gameObject, ok := m.gameObjects[name]
	if !ok {
		return fmt.Errorf("game object \"%s\" not found", name)
	}

	if c, ok := gameObject.(goi.WithComponents); ok {
		for _, c := range c.GetComponents() {
			if err := m.unregister(c.GetID(), c); err != nil {
				return err
			}
		}
	}

	if err := m.unregister(name, gameObject); err != nil {
		return err
	}

	delete(m.gameObjects, name)

	return nil
}

// Unregisters all the game objects.
func (m *GOMap) UnregisterAll() error {

	for n := range m.gameObjects {
		if err := m.Unregister(n); err != nil {
			return err
		}
	}

	return nil
}

////

// Registers a new game object. If it supports Awake() function,
// it will be called.
func (m *GOMap) register(id string, gameObject interface{}) error {
	if s, ok := gameObject.(goi.Awakable); ok {
		if err := s.Awake(); err != nil {
			return err
		}
	}

	if sq, ok := gameObject.(goi.Startable); ok {
		m.sliceSQs = append(m.sliceSQs, sq)
	}

	if u, ok := gameObject.(goi.PhysUpdatable); ok {
		m.mapPUs[id] = u
	}

	if u, ok := gameObject.(goi.Phys2Updatable); ok {
		m.mapP2Us[id] = u
	}

	if u, ok := gameObject.(goi.Updatable); ok {
		m.mapUs[id] = u
	}

	if dq, ok := gameObject.(goi.DrawableQueued); ok {
		m.mapDQs[id] = dq
	}

	if duiq, ok := gameObject.(goi.DrawableUIQueued); ok {
		m.mapDUIQs[id] = duiq
	}

	return nil
}

// Unregisters the game object. If Destroy() function implemented,
// it will be called.
func (m *GOMap) unregister(id string, gameObject interface{}) error {
	if d, ok := gameObject.(goi.Destroyable); ok {
		if err := d.Destroy(); err != nil {
			return err
		}
	}

	if _, ok := gameObject.(goi.PhysUpdatable); ok {
		delete(m.mapPUs, id)
	}

	if _, ok := gameObject.(goi.Phys2Updatable); ok {
		delete(m.mapP2Us, id)
	}

	if _, ok := gameObject.(goi.Updatable); ok {
		delete(m.mapUs, id)
	}

	if _, ok := gameObject.(goi.DrawableQueued); ok {
		delete(m.mapDQs, id)
	}

	if _, ok := gameObject.(goi.DrawableUIQueued); ok {
		delete(m.mapDUIQs, id)
	}

	return nil
}

////

// Calls Update(), PhysUpdate() functions.
func (m *GOMap) Update() error {
	for _, u := range m.mapUs {
		if err := u.Update(m.gameBase); err != nil {
			return err
		}
	}

	for _, u := range m.mapPUs {
		if err := u.PhysUpdate(m.gameBase); err != nil {
			return err
		}
	}
	for _, u := range m.mapP2Us {
		if err := u.Phys2Update(m.gameBase); err != nil {
			return err
		}
	}

	return nil
}

// Calls Draw(), DrawUI() functions, previously sorts it by
// priority.
func (m *GOMap) Draw(screen *ebiten.Image) error {

	// sort the drawables by priority

	qs := make([]queue.Queue[goi.DrawableQueued], 0, len(m.mapDQs))
	i := 0
	for _, dq := range m.mapDQs {
		qs = append(qs, queue.Queue[goi.DrawableQueued]{Priority: shiftPriority(i, dq.DrawPriority()), Value: dq})
		i++
	}

	sort.Sort(queue.ByPriority[goi.DrawableQueued](qs))

	// call Draw() function on drawables

	for _, dq := range qs {
		dq.Value.Draw(m.gameBase, screen)
	}

	qs = nil

	// sort the ui drawables by priority

	qss := make([]queue.Queue[goi.DrawableUIQueued], 0, len(m.mapDQs))
	i = 0
	for _, duiq := range m.mapDUIQs {
		qss = append(qss, queue.Queue[goi.DrawableUIQueued]{Priority: shiftPriority(i, duiq.DrawUIPriority()), Value: duiq})
		i++
	}

	sort.Sort(queue.ByPriority[goi.DrawableUIQueued](qss))

	// call DrawUI() function on drawables

	for _, duiq := range qss {
		duiq.Value.DrawUI(m.gameBase, screen)
	}

	return nil
}

// One-time function, that must be called from Scene.Enter()
// only once. Starts all the queued startables, that implements
// StartableQueued interface.
func (m *GOMap) StartQueuedStartables() error {

	// sort startables by priority
	qs := make([]queue.Queue[goi.Startable], 0, len(m.sliceSQs))

	for i, sq := range m.sliceSQs {
		qs = append(qs, queue.Queue[goi.Startable]{Priority: shiftPriority(i, sq.StartPriority()), Value: sq})
	}

	sort.Sort(queue.ByPriority[goi.Startable](qs))

	// start all the queued startables

	for _, q := range qs {
		// run StartQueued() function
		if err := q.Value.Start(m.gameBase); err != nil {
			return err
		}
	}

	return nil
}

////

var priorityShift = 0.00001

// Shifts the priority with priorityShift
func shiftPriority(i int, p float64) float64 {
	return p + (priorityShift * float64(i+1))
}
