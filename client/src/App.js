import HeaderNav from "./components/header/header-comp/header-nav";
import Header from "./components/header/header";
import Sign from "./components/sign/sign";
import Mentors from "./components/mentors/mentors";
import MentorsPage from "./components/mentors-page/mentors-page";
import Survey from "./components/survey/survey";
import SurveyResults from "./components/survey-results/survey-results";

import "./app.sass";

import { Routes, Route } from "react-router-dom";

function App() {
    return (
        <div className="App">
            <HeaderNav />

            <Routes>
                <Route path="/" element={<Header />} />
                <Route path="/sign" element={<Sign />} />
                <Route path="/sign/survey" element={<Survey />} />
                <Route
                    path="/sign/survey/results"
                    element={<SurveyResults />}
                />
                <Route path="/mentors" element={<Mentors />} />
                <Route path="/mentors/:id" element={<MentorsPage />} />
            </Routes>
        </div>
    );
}

export default App;
