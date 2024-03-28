import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';

// Mock data
const axiosServices = axios.create({baseURL: process.env.NEXT_PUBLIC_REACT_APP_API_URL || '/'});
var mock = new MockAdapter(axiosServices);

mock.onAny().passThrough();

export default axiosServices;