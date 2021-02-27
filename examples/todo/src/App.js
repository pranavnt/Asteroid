import { GeistProvider, CssBaseline } from "@geist-ui/react";
import "./styling/App.css";
import Todo from "./components/Todos";

function App() {
	return (
		<div className="App">
			<GeistProvider>
				<CssBaseline />
				<Todo />
			</GeistProvider>
		</div>
	);
}

export default App;
