import React, { useCallback, useMemo, useState } from 'react'
import { VersionContext } from '../schema/site.schema'
import { VirtualList } from '@sourcegraph/shared/src/components/VirtualList'
import { ConvertVersionContextNode } from './ConvertVersionContextNode'
import { concat, from, Observable, of, Subject } from 'rxjs'
import { useEventObservable } from '@sourcegraph/shared/src/util/useObservable'
import { convertVersionContextToSearchContext } from '../search/backend'
import { catchError, delay, map, mergeMap, reduce, tap } from 'rxjs/operators'
import { asError, isErrorLike } from '@sourcegraph/shared/src/util/errors'

export interface ConvertVersionContextsTabProps {
    availableVersionContexts: VersionContext[] | undefined
}

const initialItemsToShow = 15
const incrementalItemsToShow = 10
const LOADING = 'LOADING' as const

export const ConvertVersionContextsTab: React.FunctionComponent<ConvertVersionContextsTabProps> = ({
    availableVersionContexts,
}) => {
    const itemKey = useCallback((item: VersionContext): string => item.name, [])

    const versionContexts = useMemo(() => {
        if (!availableVersionContexts) {
            return []
        }
        return availableVersionContexts.map(versionContext => ({
            ...versionContext,
            isConvertedUpdates: new Subject<void>(),
        }))
    }, [availableVersionContexts])

    const renderResult = useCallback(
        (item: VersionContext & { isConvertedUpdates: Subject<void> }): JSX.Element => (
            <ConvertVersionContextNode name={item.name} isConvertedUpdates={item.isConvertedUpdates} />
        ),
        []
    )

    const [itemsToShow, setItemsToShow] = useState(initialItemsToShow)
    const onBottomHit = useCallback(() => {
        setItemsToShow(items => Math.min(versionContexts.length || 0, items + incrementalItemsToShow))
    }, [versionContexts])

    const [convertAll, convertAllResult] = useEventObservable(
        useCallback(
            (event: Observable<React.MouseEvent>) =>
                event.pipe(
                    mergeMap(() => {
                        const convertAll = from(versionContexts).pipe(
                            mergeMap(({ name, isConvertedUpdates }) =>
                                convertVersionContextToSearchContext(name).pipe(
                                    tap(() => isConvertedUpdates.next()),
                                    catchError(error => [asError(error)])
                                )
                            ),
                            map(result => (isErrorLike(result) ? 0 : 1)),
                            reduce((accumulator, result) => accumulator + result, 0)
                        )
                        return concat(of(LOADING), convertAll.pipe(delay(500)))
                    })
                ),
            [versionContexts]
        )
    )

    return (
        <div className="convert-version-contexts-tab">
            <div className="convert-version-contexts-tab__header ml-3 mr-3 mb-3">
                <div className="convert-version-contexts-tab__header__title">Available version contexts</div>
                <button
                    type="button"
                    className="btn btn-outline-primary"
                    onClick={convertAll}
                    disabled={convertAllResult === LOADING}
                >
                    {convertAllResult === LOADING ? 'Converting All...' : 'Convert All'}
                </button>
            </div>
            {typeof convertAllResult !== 'undefined' &&
                convertAllResult !== LOADING &&
                (convertAllResult === 0 ? (
                    <div className="alert alert-info">No version contexts to convert.</div>
                ) : (
                    <div className="alert alert-success">
                        Sucessfully converted <strong>{convertAllResult}</strong> version contexts.
                    </div>
                ))}
            <VirtualList<VersionContext & { isConvertedUpdates: Subject<void> }>
                className="mt-2"
                itemsToShow={itemsToShow}
                onShowMoreItems={onBottomHit}
                items={versionContexts}
                itemProps={undefined}
                itemKey={itemKey}
                renderItem={renderResult}
            />
        </div>
    )
}
