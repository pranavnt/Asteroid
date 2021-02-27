import React, { Component } from "react";
import { Button } from "@geist-ui/react";

export default class Task extends Component {
	render() {
		return (
			<div
				key={this.props.index}
				style={{ display: "flex", marginLeft: "45vw", marginRight: "45vw" }}
			>
				<h2 style={{ marginRight: "10px", marginTop: "-5px" }}>
					{this.props.tasks.name}
				</h2>
				<Button
					onClick={() => this.props.deleteTodo(this.props.tasks.id)}
					className="ml-auto"
				>
					Delete
				</Button>
			</div>
		);
	}
}
