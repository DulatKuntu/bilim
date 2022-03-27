import "../mentors-sass/mentors-page-top.sass";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faPhoneAlt } from "@fortawesome/free-solid-svg-icons";

const MentorsPageTop = ({ name, surname, username, interests, img }) => {
    return (
        <div className="mentors-page-top">
            <div className="mentors-page-top-information">
                <div className="mentors-page-top-information__img_container">
                    <img
                        src={img}
                        className="mentors-page-top-information__img"
                    ></img>
                </div>
            </div>

            <div className="mentors-page-top-information__main">
                <div className="mentors-page-top-information__main_name">
                    {`${name} ${surname}`}
                </div>
                <div className="mentors-page-top-information__main_username">
                    {username}
                </div>
                <div className="mentors-page-top-information__main_interests">
                    {interests}
                </div>
            </div>
        </div>
    );
};

export default MentorsPageTop;
