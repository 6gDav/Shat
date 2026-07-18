import type { PageServerLoad } from './$types';

export const load: PageServerLoad = ({ getClientAddress }) => {
    try {
        const userIP = getClientAddress();
        return {
            ip: userIP
        };
    } catch (error) {
        console.error("Error occured while trying to get the ip address:", error);
        return {
            ip: ""
        };
    }
};