package engine

//type dependencyCollector struct {
//	m mask
//}
//
//func (d *dependencyCollector) ChangeScene(_ Scene) {}
//
//func (d *dependencyCollector) Bounds() image.Rectangle {
//	return image.Rect(0, 0, 800, 600)
//}
//
//func (d *dependencyCollector) AddEntity(_ ...any) {}
//
//func (d *dependencyCollector) AddSystem(_ any) {}
//
//func (d *dependencyCollector) Get(_ ...matcher) (entity Entity, ok bool) {
//	return d, true
//}
//
//func (d *dependencyCollector) Append(entities []Entity, _ ...matcher) {
//	entities = append(entities, d)
//}
//
//func (d *dependencyCollector) Each(f func(d Entity), _ ...matcher) {
//	f(d)
//}
//
//func (d *dependencyCollector) Components() int {
//	return 10
//}
//
//func (d *dependencyCollector) Entities() int {
//	return 10
//}
//
//func (d *dependencyCollector) Systems() int {
//	return 10
//}
//
//func (d *dependencyCollector) ID() int {
//	return 0
//}
//
//func (d *dependencyCollector) Rem() {}
//
//func (d *dependencyCollector) setComponent(component any) {
//	d.m.set()
//}
//
//func (d *dependencyCollector) getComponent(component any) (any, bool) {
//	return nil, false
//}
//
//func (d *dependencyCollector) hasComponent(component any) bool {
//	return true
//}
//
//func (d *dependencyCollector) remComponent(component any) {
//
//}
