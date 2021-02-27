import React, { Component } from "react";
import Submit from "./Submit";
import Task from "./Task";

class Todos extends Component {
	constructor() {
		super();
		this.state = {
			tasks: [],
			makeTodo: "",
		};
	}

	componentDidMount() {}

	//updates state with value in the input prior to submission
	handleTodoChange = (e) => {
		e.preventDefault();

		this.setState({
			tasks: this.state.tasks,
			makeTodo: e.target.value,
		});
		console.log("changed");
	};

	//creates a todo by adding it to the state and displays it
	createTodo = (e) => {
		e.preventDefault();
		if (this.state.makeTodo === "") return false;

		//adds to db
		const taskId = 1;
		this.setState({
			tasks: [
				...this.state.tasks,
				{
					user: "me",
					id: taskId,
					name: this.state.makeTodo,
				},
			],
			makeTodo: "",
		});
		console.log("created");
	};

	//removes todo from state
	deleteTodo = (id) => {
		this.setState({
			tasks: [...this.state.tasks.filter((tasks) => tasks.id !== id)],
		});
	};

	render() {
		//creates all the todos to be rendered
		const todos = this.state.tasks.map((tasks, index) => {
			return <Task tasks={tasks} key={index} deleteTodo={this.deleteTodo} />;
		});

		//send the rendered version of the todos and the submit
		return (
			<div>
				{todos}
				<Submit
					handleTodoChange={this.handleTodoChange}
					makeTodo={this.state.makeTodo}
					createTodo={this.createTodo}
				/>
				{/* <button onClick={() => {firebase.auth().signOut()}}>SIgnout</button> */}
			</div>
		);
	}
}

export default Todos;
