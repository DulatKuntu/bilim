import "./mentors-sass/mentors.sass";

import MentorsCard from "./mentors-comp/mentors-card";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

import { Link } from "react-router-dom";

const Mentors = () => {
    return (
        <div className="mentors">
            <div className="mentors-top">
                <div className="mentors-top-text">mentors</div>

                <form className="mentors-top-search">
                    <FontAwesomeIcon
                        icon={faSearch}
                        className="mentors-top-search__emoji"
                    />
                    <input
                        type="text"
                        className="mentors-top-search__input"
                        placeholder="Search among 1000+ mentors"
                    />
                </form>
            </div>

            <div className="mentors-container">
                <Link to="/mentors/1">
                    <MentorsCard />
                </Link>
                <MentorsCard />
                <MentorsCard />
            </div>
        </div>
    );
};

export default Mentors;
