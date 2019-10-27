package main

import (
	"fmt"
)

type UserGraph struct {
	User []User `json:"user_graph"`
}

type User struct {
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	Rating   int          `json:"rating"`
	Warning  bool         `json:"warning"`
	Studied  []Course     `json:"studied_course"`
	Passed   []Topic      `json:"passed_topic"`
	Results  []TestResult `json:"has_result"`
	Activity []Activity   `json:"has_activity"`
}

type Course struct {
	Name      string   `json:"name"`
	Topics    []Topic  `json:"has_topic"`
	Connected []Course `json:"related_to"`
}

type Topic struct {
	Name      string   `json:"name"`
	Strength  int   `json:"passed_topic|strength"`
	Course    []Course `json:"from_course"`
	Connected []Topic  `json:"related_to"`
}

type TestResult struct {
	Topic []Topic `json:"for_topic"`
	User  []User  `json:"for_user"`
}

type Activity struct {
	User  []User  `json:"of_user"`
	Topic []Topic `json:"on_topic"`
}

type nodeList struct {
	Nodes []node `json:"nodes"`
	Edges []edge `json:"edges"`
}

type node struct {
	Name string `json:"name"`
}

type edge struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Weight string `json:"weigth,omitempty"`
}

func (g UserGraph) transformToList() nodeList {
	list := nodeList{}
	passed := make(map[string]Topic)
	for _, topic := range g.User[0].Passed {
		passed[topic.Name] = topic
	}
	for _, course := range g.User[0].Studied {
		list.Nodes = append(list.Nodes, node{Name: course.Name})
		for _, topic := range course.Topics {
			list.Nodes = append(list.Nodes, node{Name: topic.Name})
			edge := edge{
				Src: topic.Name, 
				Dest: course.Name,
			}
			if t, ok := passed[topic.Name]; ok {
				edge.Weight = fmt.Sprintf("passed|%d", t.Strength) 
			}
			list.Edges = append(list.Edges, edge)
		}
	}
	return list
} 
