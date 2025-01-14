// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Ssis environment reference.
 */
export interface SsisEnvironmentReferenceResponse {
    /**
     * Environment folder name.
     */
    environmentFolderName?: string;
    /**
     * Environment name.
     */
    environmentName?: string;
    /**
     * Environment reference id.
     */
    id?: number;
    /**
     * Reference type
     */
    referenceType?: string;
}

/**
 * Ssis environment.
 */
export interface SsisEnvironmentResponse {
    /**
     * Metadata description.
     */
    description?: string;
    /**
     * Folder id which contains environment.
     */
    folderId?: number;
    /**
     * Metadata id.
     */
    id?: number;
    /**
     * Metadata name.
     */
    name?: string;
    /**
     * The type of SSIS object metadata.
     * Expected value is 'Environment'.
     */
    type: "Environment";
    /**
     * Variable in environment
     */
    variables?: outputs.SsisVariableResponse[];
}

/**
 * Ssis folder.
 */
export interface SsisFolderResponse {
    /**
     * Metadata description.
     */
    description?: string;
    /**
     * Metadata id.
     */
    id?: number;
    /**
     * Metadata name.
     */
    name?: string;
    /**
     * The type of SSIS object metadata.
     * Expected value is 'Folder'.
     */
    type: "Folder";
}

/**
 * Ssis Package.
 */
export interface SsisPackageResponse {
    /**
     * Metadata description.
     */
    description?: string;
    /**
     * Folder id which contains package.
     */
    folderId?: number;
    /**
     * Metadata id.
     */
    id?: number;
    /**
     * Metadata name.
     */
    name?: string;
    /**
     * Parameters in package
     */
    parameters?: outputs.SsisParameterResponse[];
    /**
     * Project id which contains package.
     */
    projectId?: number;
    /**
     * Project version which contains package.
     */
    projectVersion?: number;
    /**
     * The type of SSIS object metadata.
     * Expected value is 'Package'.
     */
    type: "Package";
}

/**
 * Ssis parameter.
 */
export interface SsisParameterResponse {
    /**
     * Parameter type.
     */
    dataType?: string;
    /**
     * Default value of parameter.
     */
    defaultValue?: string;
    /**
     * Parameter description.
     */
    description?: string;
    /**
     * Design default value of parameter.
     */
    designDefaultValue?: string;
    /**
     * Parameter id.
     */
    id?: number;
    /**
     * Parameter name.
     */
    name?: string;
    /**
     * Whether parameter is required.
     */
    required?: boolean;
    /**
     * Whether parameter is sensitive.
     */
    sensitive?: boolean;
    /**
     * Default sensitive value of parameter.
     */
    sensitiveDefaultValue?: string;
    /**
     * Parameter value set.
     */
    valueSet?: boolean;
    /**
     * Parameter value type.
     */
    valueType?: string;
    /**
     * Parameter reference variable.
     */
    variable?: string;
}

/**
 * Ssis project.
 */
export interface SsisProjectResponse {
    /**
     * Metadata description.
     */
    description?: string;
    /**
     * Environment reference in project
     */
    environmentRefs?: outputs.SsisEnvironmentReferenceResponse[];
    /**
     * Folder id which contains project.
     */
    folderId?: number;
    /**
     * Metadata id.
     */
    id?: number;
    /**
     * Metadata name.
     */
    name?: string;
    /**
     * Parameters in project
     */
    parameters?: outputs.SsisParameterResponse[];
    /**
     * The type of SSIS object metadata.
     * Expected value is 'Project'.
     */
    type: "Project";
    /**
     * Project version.
     */
    version?: number;
}

/**
 * Ssis variable.
 */
export interface SsisVariableResponse {
    /**
     * Variable type.
     */
    dataType?: string;
    /**
     * Variable description.
     */
    description?: string;
    /**
     * Variable id.
     */
    id?: number;
    /**
     * Variable name.
     */
    name?: string;
    /**
     * Whether variable is sensitive.
     */
    sensitive?: boolean;
    /**
     * Variable sensitive value.
     */
    sensitiveValue?: string;
    /**
     * Variable value.
     */
    value?: string;
}

/**
 * An access key for the storage account.
 */
export interface StorageAccountKeyResponse {
    /**
     * Creation time of the key, in round trip date format.
     */
    creationTime: string;
    /**
     * Name of the key.
     */
    keyName: string;
    /**
     * Permissions for the key -- read-only or full permissions.
     */
    permissions: string;
    /**
     * Base 64-encoded value of the key.
     */
    value: string;
}

