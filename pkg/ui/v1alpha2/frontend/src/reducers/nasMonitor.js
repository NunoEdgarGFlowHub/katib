import * as actions from '../actions/nasMonitorActions';

const initialState = {
    filter: '',
    filterType: {
        "Created": true,
        "Running": true,
        "Restarting": true,
        "Succeeded": true,
        "Failed": true,
    },
    jobsList: [
    ],
    filteredJobsList: [
    ],
    loading: false,
    steps: [
        
    ]
};

const nasMonitorReducer = (state = initialState, action) => {
    switch (action.type) {
        case actions.FILTER_JOBS:
            let jobs = state.jobsList.slice();
            let newList = jobs.filter(job => job.name.includes(action.filter));

            let avTypes = Object.assign({}, state.filterType);
            var typeKeys = Object.keys(avTypes);

            var avFilters = typeKeys.filter((key) => {
                return avTypes[key]
            });

            let filteredJobs = newList.filter(job => avFilters.includes(job.status));

            return {
                ...state,
                filteredJobsList: filteredJobs,
                filter: action.filter,
            }
        case actions.CHANGE_TYPE:
            const types = Object.assign({}, state.filterType)
            types[action.filter] = action.checked;
            var keys = Object.keys(types);

            var filters = keys.filter((key) => {
                return types[key]
            });
            const jobsList = state.jobsList.slice();
            const filtered = jobsList.filter(job => filters.includes(job.status));
            
            return {
                ...state,
                filterType: types,
                filteredJobsList: filtered,
            }
        case actions.FETCH_NAS_JOBS_SUCCESS:
            jobs = action.jobs
            avTypes = Object.assign({}, state.filterType);
            typeKeys = Object.keys(avTypes);

            avFilters = typeKeys.filter((key) => {
                return avTypes[key]
            });

            filteredJobs = jobs.filter(job => avFilters.includes(job.status));
            return {
                ...state,
                jobsList: action.jobs,
                filteredJobsList: filteredJobs,
            }
        case actions.FETCH_NAS_JOB_INFO_REQUEST:
            return {
                ...state,
                loading: true,
            }
        case actions.FETCH_NAS_JOB_INFO_FAILURE:
            return {
                ...state,
                loading: false,
            }
        case actions.FETCH_NAS_JOB_INFO_SUCCESS:
            return {
                ...state,
                loading: false,
                steps: action.steps,
            }
        default:
            return state;
    }
};


export default nasMonitorReducer;
