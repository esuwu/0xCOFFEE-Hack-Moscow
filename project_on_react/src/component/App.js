import React from "react";
import {HashRouter, Switch, Link} from "react-router-dom";

import QuestionList from "./qustions/question_list"
import "./App.css";

export default class App extends React.Component {
  state = {
    questions: [
      {
        question: "Question first",
        answers: [
          "answer 1",
          "answer 2",
          "answer 3",
        ]
      }, {
        question: "Question second",
        answers: [
          "answer 1",
          "answer 2",
          "answer 3",
        ]
      },
      
    ],
    
  };


  render() {
    console.log("Questions: ", this.state.questions);
    return (
      <div className="component-app">
        <QuestionList value={this.state.questions}/>
      </div>
    );
  }
}

// <Display value={this.state.next || this.state.total || "0"} />
// <ButtonPanel clickHandler={this.handleClick} />
