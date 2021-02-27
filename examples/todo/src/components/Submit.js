import React, { Component } from "react";
import { Button, Input } from "@geist-ui/react";

export default class Submit extends Component {
	render() {
		return (
			<div className="container">
				<div className="form" style={{ marginTop: "100px" }}>
					<Input
						type="text"
						placeholder="Enter Todo"
						value={this.props.makeTodo}
						onChange={this.props.handleTodoChange}
					/>
					<Button type="submit" onClick={this.props.createTodo}>
						Make Todo
					</Button>
				</div>
			</div>
		);
	}
}
