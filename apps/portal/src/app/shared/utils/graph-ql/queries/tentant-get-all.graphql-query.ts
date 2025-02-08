import {gql} from "apollo-angular";

export const Tenant_GetAll_Query = gql`
    query GetTenants {
        getTenants {
            description
            id
            name
        }
    }
`;
