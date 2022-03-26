import MentorsPageQuestElement from "./mentors-page-quest-element";

const MentorsPageQuest = () => {
    return (
        <div className="mentors-page-quest">
            <h2 className="mentors-page-quest-header"></h2>

            <div className="mentors-page-quest-main">
                <MentorsPageQuestElement />
                <MentorsPageQuestElement />
                <MentorsPageQuestElement />
            </div>
        </div>
    );
};

export default MentorsPageQuest;
