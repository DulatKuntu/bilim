import HeaderNav from "./components/header/header-comp/header-nav";
import Header from "./components/header/header";
import Sign from "./components/sign/sign";

import "./app.sass";

import { Routes, Route } from "react-router-dom";

function App() {
    return (
        <div className="App">
            <HeaderNav />

            <Routes>
                <Route path="/" element={<Header />} />

                <Route path="/sign" element={<Sign />} />
            </Routes>
        </div>
    );
}

export default App;
