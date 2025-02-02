# Bulk operations on changesets

Bulk operations allow a single action to be performed across many changesets in a batch change.

## Selecting changesets for a bulk operation

1. Click the checkbox next to a changeset in the list view. You can select all changesets you have permission to view.

    <img src="https://sourcegraphstatic.com/docs/images/batch_changes/select_changeset.png" class="screenshot">
1. If you like, select all changesets in the list by using the checkbox in the list header.
    
    <img src="https://sourcegraphstatic.com/docs/images/batch_changes/select_all_changesets_in_view.png" class="screenshot">
    
    If you want to select _all_ changesets that meet the filters and search currently set, click the **(Select XX changesets)** link in the header toolbar.

    <img src="https://sourcegraphstatic.com/docs/images/batch_changes/select_all_changesets.png" class="screenshot">
1. In the top right, select the action to perform on all the changesets.

    <img src="https://sourcegraphstatic.com/docs/images/batch_changes/select_bulk_operation_type.png" class="screenshot">

## Supported types of bulk operations

- Commenting: Post a comment on all selected changesets. This can be particularly useful for pinging people, reminding them to take a look at the changeset, or posting your favorite emoji 🦡.
- Detach: Only available in the archived tab. Detach a selection of changesets from the batch change to remove them from the archived tab.

_More types coming soon._

## Monitoring bulk operations

On the **Bulk operations** tab, you can view all bulk operations that have been run over the batch change. Since bulk operations can involve quite some operations to perform, you can track the progress, and see what operations have been performed in the past.

<img src="https://sourcegraphstatic.com/docs/images/batch_changes/bulk_operations_tab.png" class="screenshot">
