import "../mentors-sass/mentors-page-top.sass";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faPhoneAlt } from "@fortawesome/free-solid-svg-icons";

const MentorsPageTop = () => {
    return (
        <div className="mentors-page-top">
            <div className="mentors-page-top-information">
                <div className="mentors-page-top-information__img_container">
                    <div className="mentors-page-top-information__img"></div>
                </div>
            </div>

            <div className="mentors-page-top-information__main">
                <div className="mentors-page-top-information__main_name">
                    Test Test
                </div>
                <div className="mentors-page-top-information__main_username">
                    Test
                </div>
                <div className="mentors-page-top-information__main_interests">
                    <div className="mentors-page-top-information__main_interests_element">
                        test
                    </div>
                    <div className="mentors-page-top-information__main_interests_element">
                        test
                    </div>
                    <div className="mentors-page-top-information__main_interests_element">
                        test
                    </div>
                </div>

                <button className="mentors-page-top-information__main_contacts">
                    Contact me!
                </button>
            </div>
        </div>
    );
};

export default MentorsPageTop;
