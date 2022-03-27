import HeaderNav from "./components/header/header-comp/header-nav";
import Header from "./components/header/header";

import SignMiddle from "./components/sign/sign-middle";
import Sign from "./components/sign/sign";

import Mentors from "./components/mentors/mentors";
import MentorsPage from "./components/mentors-page/mentors-page";
import Survey from "./components/survey/survey";
import SurveyResults from "./components/survey-results/survey-results";
import Profession from "./components/profession/profession";
import UserPage from "./components/user-page/user-page";
import Friend from "./components/friend/friend";

import "./app.sass";

import { useState } from "react";
import { Routes, Route } from "react-router-dom";

function App() {
    const [userRegistered, setUserRegistered] = useState(false);
    const [userId, setUserId] = useState(0);
    const [card, setCard] = useState([]);
    const [user, setUser] = useState("s");

    return (
        <div className="App">
            <HeaderNav
                userRegistered={userRegistered}
                userId={userId}
                user={user}
            />

            <Routes>
                <Route path="/" element={<Header />} />
                <Route
                    path="/sign"
                    element={<SignMiddle setUser={setUser} />}
                />
                <Route
                    path="/sign/user"
                    element={
                        <Sign
                            setUserRegistered={setUserRegistered}
                            setUserId={setUserId}
                        />
                    }
                />
                <Route
                    path="/sign/mentor"
                    element={
                        <Sign
                            setUserRegistered={setUserRegistered}
                            setUserId={setUserId}
                            user={user}
                        />
                    }
                />
                <Route path="/sign/user/survey" element={<Survey />} />
                <Route
                    path="/sign/user/survey/results"
                    element={<SurveyResults />}
                />
                <Route
                    path="/sign/user/survey/results/profession"
                    element={<Profession />}
                />
                <Route
                    path="/mentors"
                    element={<Mentors setCard={setCard} card={card} />}
                />
                <Route
                    path="/mentors/:id"
                    element={<MentorsPage card={card} />}
                />

                <Route path="/user/:id" element={<UserPage />} />
                <Route path="/friend" element={<Friend />} />
            </Routes>
        </div>
    );
}

export default App;
