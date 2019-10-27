import React from "react";
// import PropTypes from "prop-types";
import "./answer_grid.css";

export default function AnswerGrid(props) {
  console.log("answerGrid: ", props);
  return (
    <div>
      <label className="radio">
        <input name="answer" type="radio" value={props.answerIndex}/>
        <span>{ props.answerElement }</span>
      </label>
    </div>
  )
}