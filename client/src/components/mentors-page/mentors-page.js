import MentorsPageTop from "./mentors-comp/mentors-page-top";
import MentorsPageQuest from "./mentors-comp/mentors-page-quest";

import "./mentors-sass/mentors-page.sass";

const MentorsPage = () => {
    return (
        <div className="mentors-page">
            <MentorsPageTop />

            <div className="mentors-page-bio">
                <h2 className="mentors-page-bio__header">Bio</h2>

                <div className="mentors-page-bio__text">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Adipisci nemo nihil neque sed, tempore fugiat ad esse. Ab
                    aliquid aperiam quidem blanditiis cumque assumenda ea sequi
                    doloremque sint. Sapiente, tempore!
                </div>
            </div>

            <MentorsPageQuest />
        </div>
    );
};

export default MentorsPage;
