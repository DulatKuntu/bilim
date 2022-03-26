import SurveyResultsCard from "./survey-results-comp/survey-result-card";

import "./survey-results-sass/survey-results.sass";

const SurveyResults = () => {
    return (
        <div className="survey-results">
            <h2 className="survey-results-main">Твои результаты:</h2>

            <SurveyResultsCard></SurveyResultsCard>
            <SurveyResultsCard></SurveyResultsCard>
            <SurveyResultsCard></SurveyResultsCard>
        </div>
    );
};

export default SurveyResults;
