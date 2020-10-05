import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios'

// const api = axios.create({
//   baseURL: `/bubbleSort`
// })

class App extends Component {
  constructor(){
    super();
    this.state ={
      numbers:[1,6,3,5,60,2]
    };
    const{ numbers} = this.state
    let newNumbers = [...numbers]
    console.log(newNumbers)
    axios({
      method: 'post',
      url: '/bubbleSort',
      data: {"numbers":[11,9,8,7,50,-60,70]},
      headers: {'Content-Type':  "application/json" }
      }).then(res=>{
      console.log(res.data)
    })
  }
  render(){
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}
}
export default App;
