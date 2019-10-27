var color = "gray";
var len = undefined;

var nodes = [
  { id: 0, label: "Biology", group: 0 },
  { id: 1, label: "", group: 0 },
  { id: 2, label: "", group: 0 },
  { id: 3, label: "", group: 1 },
  { id: 4, label: "Chemistry", group: 1 },
  { id: 5, label: "", group: 1 },
  { id: 6, label: "", group: 2 },
  { id: 7, label: "Geographic", group: 2 },
  { id: 8, label: "", group: 2 },
  { id: 9, label: "", group: 3 },
  { id: 10, label: "English", group: 3 },
  { id: 11, label: "", group: 3 },
  { id: 12, label: "", group: 4 },
  { id: 13, label: "Phys", group: 4 },
  { id: 14, label: "", group: 4 },
  { id: 15, label: "", group: 5 },
  { id: 16, label: "Psychology", group: 5 },
  { id: 17, label: "", group: 5 },
  { id: 18, label: "", group: 6 },
  { id: 19, label: "History", group: 6 },
  { id: 20, label: "", group: 6 },
  { id: 21, label: "", group: 7 },
  { id: 22, label: "Math", group: 7 },
  { id: 23, label: "", group: 7 },
  { id: 24, label: "", group: 8 },
  { id: 25, label: "Art", group: 8 },
  { id: 26, label: "", group: 8 },
  { id: 27, label: "", group: 9 },
  { id: 28, label: "IT", group: 9 },
  { id: 29, label: "", group: 9 }
];
var edges = [
  { from: 1, to: 0 },
  { from: 2, to: 0 },
  { from: 4, to: 3 },
  { from: 5, to: 4 },
  { from: 4, to: 0 },
  { from: 7, to: 6 },
  { from: 8, to: 7 },
  { from: 7, to: 0 },
  { from: 10, to: 9 },
  { from: 11, to: 10 },
  { from: 10, to: 4 },
  { from: 13, to: 12 },
  { from: 14, to: 13 },
  { from: 13, to: 0 },
  { from: 16, to: 15 },
  { from: 17, to: 15 },
  { from: 15, to: 10 },
  { from: 19, to: 18 },
  { from: 20, to: 19 },
  { from: 19, to: 4 },
  { from: 22, to: 21 },
  { from: 23, to: 22 },
  { from: 22, to: 13 },
  { from: 25, to: 24 },
  { from: 26, to: 25 },
  { from: 25, to: 7 },
  { from: 28, to: 27 },
  { from: 29, to: 28 },
  { from: 28, to: 0 }
];

document.addEventListener("DOMContentLoaded", ready);

// create a network

function ready() {
var container = document.getElementById("mynetwork");
var data = {
  nodes: nodes,
  edges: edges
};
var options = {
  nodes: {
    shape: "dot",
    size: 30,
    font: {
      size: 32,
      color: "#ffffff"
    },
    borderWidth: 2
  },
  edges: {
    width: 2
  }
};
network = new vis.Network(container, data, options);
}
