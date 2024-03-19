package lemin

// func (g *Graph) PathFinder() {
// 	paths := make([][]string, 0)
// 	startVertex := g.findStartVertex()
// 	if startVertex == nil {
// 		fmt.Println("Start room not found")
// 		return
// 	}
//
// 	g.findPaths(startVertex, []string{g.startRoom}, &paths)
// 	for _, path := range paths {
// 		fmt.Printf("Path through room: %v (%d rooms)==> Path:%v\n", path[1], len(path), path)
// 	}
// 	fmt.Println()
//
// 	paths = QuickSort(paths)
//
// 	fmt.Println("------------ After sorting ------------")
// 	for _, path := range paths {
// 		fmt.Printf("Path through room: %v (%d rooms)==> Path:%v\n", path[1], len(path), path)
// 	}
// }
//
// // findStartVertex returns the vertex for the start room.
// func (g *Graph) findStartVertex() *vertex {
// 	for _, v := range g.Vertices {
// 		if v.key == g.startRoom {
// 			return v
// 		}
// 	}
// 	return nil
// }
//
// // findPaths recursively finds all paths from the current vertex.
// func (g *Graph) findPaths(v *vertex, path []string, paths *[][]string) {
// 	if v.key == g.endRoom {
// 		*paths = append(*paths, path)
// 		return
// 	}
// 	for _, conn := range v.Connections {
// 		// Avoid revisiting vertices in the same path to prevent cycles.
// 		if !contains(path, conn.key) {
// 			newPath := append([]string(nil), path...) // Make a copy of the path
// 			newPath = append(newPath, conn.key)
// 			g.findPaths(conn, newPath, paths)
// 		}
// 	}
// }
//
// // contains checks if a slice contains a string.
// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }
