import classNames from 'classnames'
import React, { useCallback } from 'react'
import { useHistory } from 'react-router'
import StickyBox from 'react-sticky-box'

import { VersionContextProps } from '@sourcegraph/shared/src/search/util'
import { SettingsCascadeProps } from '@sourcegraph/shared/src/settings/settings'
import { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'

import { CaseSensitivityProps, PatternTypeProps, SearchContextProps } from '../../..'
import { submitSearch, toggleSearchFilter } from '../../../helpers'
import { Filter } from '../../../stream'

import { getDynamicFilterLinks, getRepoFilterLinks, getSearchScopeLinks } from './FilterLink'
import { getQuickLinks } from './QuickLink'
import styles from './SearchSidebar.module.scss'
import { SearchSidebarSection } from './SearchSidebarSection'
import { getSearchTypeLinks } from './SearchTypeLink'

export interface SearchSidebarProps
    extends Omit<PatternTypeProps, 'setPatternType'>,
        Omit<CaseSensitivityProps, 'setCaseSensitivity'>,
        VersionContextProps,
        Pick<SearchContextProps, 'selectedSearchContextSpec'>,
        SettingsCascadeProps,
        TelemetryProps {
    query: string
    filters?: Filter[]
    className?: string
}

export const SearchSidebar: React.FunctionComponent<SearchSidebarProps> = props => {
    const history = useHistory()

    const onFilterClicked = useCallback(
        (value: string) => {
            props.telemetryService.log('DynamicFilterClicked', {
                search_filter: { value },
            })

            const newQuery = toggleSearchFilter(props.query, value)

            submitSearch({ ...props, query: newQuery, source: 'filter', history })
        },
        [history, props]
    )

    return (
        <div className={classNames(styles.searchSidebar, props.className)}>
            <StickyBox className={styles.searchSidebarStickyBox}>
                <SearchSidebarSection className={styles.searchSidebarItem} header="Search types">
                    {getSearchTypeLinks(props)}
                </SearchSidebarSection>
                <SearchSidebarSection className={styles.searchSidebarItem} header="Dynamic filters">
                    {getDynamicFilterLinks(props.filters, onFilterClicked)}
                </SearchSidebarSection>
                <SearchSidebarSection className={styles.searchSidebarItem} header="Repositories" showSearch={true}>
                    {getRepoFilterLinks(props.filters, onFilterClicked)}
                </SearchSidebarSection>
                <SearchSidebarSection className={styles.searchSidebarItem} header="Search snippets">
                    {getSearchScopeLinks(props.settingsCascade, onFilterClicked)}
                </SearchSidebarSection>
                <SearchSidebarSection className={styles.searchSidebarItem} header="Quicklinks">
                    {getQuickLinks(props.settingsCascade)}
                </SearchSidebarSection>
            </StickyBox>
        </div>
    )
}
