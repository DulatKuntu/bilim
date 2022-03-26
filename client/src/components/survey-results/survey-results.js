import SurveyResultsCard from "./survey-results-comp/survey-result-card";

import "./survey-results-sass/survey-results.sass";

const SurveyResults = () => {
    return (
        <div className="survey-results">
            <h2 className="survey-results-main">Твои результаты:</h2>

            <div className="survey-results-card__container">
                <div className="survey-results-number">1</div>
                <SurveyResultsCard></SurveyResultsCard>
            </div>

            <div className="survey-results-card__container">
                <div className="survey-results-number">2</div>
                <SurveyResultsCard></SurveyResultsCard>
            </div>
            <div className="survey-results-card__container">
                <div className="survey-results-number">3</div>
                <SurveyResultsCard></SurveyResultsCard>
            </div>
        </div>
    );
};

export default SurveyResults;
