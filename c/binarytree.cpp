#include <stdio.h>
#include <stdlib.h>

struct node
{
    int value;
    node *left;
    node *right;
};

node* newNode(int value)
{
    node *result = (node *)malloc(sizeof(node));
    result->value = value;
    result->left = NULL;
    result->right = NULL;

    return result;
}
void invert(node *current, node *parent)
{
    if(current == NULL)
        return;
    if(current->left != NULL) {
        invert(current->left, current);
        current->left = NULL;
    }
    if(current->right != NULL) {
        invert(current->right, current);
        current->right = NULL;
    }

    current->left = parent;
}

int size(node *tree)
{
    if(tree == NULL)
        return 0;
    else
        return size(tree->left) + 1 + size(tree->right);         
}

void insert(int value, node **leaf)
{
    if(*leaf == NULL)
        *leaf = newNode(value);
    else if(value <= (*leaf)->value)
        insert(value, &(*leaf)->left);
    else if(value > (*leaf)->value)
        insert(value, &(*leaf)->right);
}

node* search(int value, node *leaf) 
{
    if(leaf != NULL) {
        if(value == leaf->value)
            return leaf;
        else if(value <= leaf->value)
            return search(value, leaf->left);
        else
            return search(value, leaf->right);
    }
    else
        return NULL;
}

node* minValueNode(node *tree)
{
    node* result = tree;
    while(result->left != NULL)
        result = result->left;

    return result;
}

void deleteNode(int value, node **tree)
{
    if(*tree != NULL)
    {
        if(value < (*tree)->value)
            deleteNode(value, &(*tree)->left);
        if(value > (*tree)->value)
            deleteNode(value, &(*tree)->right);
        if(value == (*tree)->value) {
            if((*tree)->left == NULL) {
                node *temp = *tree;
                *tree = (*tree)->right;
                free(temp);
            }
            else if((*tree)->right == NULL) {
                node *temp = *tree;
                *tree = (*tree)->left;
                free(temp);
            }
            else if((*tree)->left != NULL &&  (*tree)->right != NULL) {
                node *temp = minValueNode((*tree)->right);
                (*tree)->value = temp->value;
                deleteNode(temp->value, &(*tree)->right);
            }
        }
    }
}

int maxDepth(node *tree)
{
    if(tree == NULL)
        return 0;
    else {
        int lDepth = maxDepth(tree->left); 
        int rDepth = maxDepth(tree->right); 

        if(lDepth > rDepth)
            return lDepth + 1;
        else
            return rDepth + 1;
    }
}

// Does a path values sum == int sum
bool hasPathSum(node *tree, int sum)
{
    if(tree == NULL)
        return sum == 0;
    else {
        int sub = sum - tree->value;
        return (hasPathSum(tree->left, sub) || 
                hasPathSum(tree->right, sub));
    }
}

void destroy(node *tree)
{
    if(tree != NULL) {
        destroy(tree->left);
        destroy(tree->right);
    }

    free(tree);
}

void printInOrder(node *tree)
{
    if(tree != NULL) {
        printInOrder(tree->left);
        printf("%d ", tree->value);
        printInOrder(tree->right);
    }
}

void printPostOrder(node *tree)
{
    if(tree == NULL)     
        return;

    printPostOrder(tree->left);
    printPostOrder(tree->right);
    printf("%d ", tree->value);
}

// Doesn't work?
void printRootToLeaf(node *tree)
{
    if(tree != NULL) {
        printf("%d ", tree->value);
        printRootToLeaf(tree->left);
        printRootToLeaf(tree->right);
    }
}

node* copyTree(node *tree)
{
    if(tree == NULL)
        return NULL;

    node *result = newNode(tree->value);
    if(tree->left != NULL)
        result->left = copyTree(tree->left);
    if(tree->right != NULL)
        result->right = copyTree(tree->right);
}

// Doesn't work
void printPaths(node *tree, node *path)
{
    if(tree = NULL) {
        printRootToLeaf(path);
        printf("\n");
    }
    else {
        path = newNode(tree->value); 
        if(tree->left != NULL)
            printPaths(tree->left, copyTree(path)->left);
        if(tree->right != NULL)
            // Can copy to left or right, doesn't matter
            printPaths(tree->right, copyTree(path)->left);
    }
}

int main(int argc, const char *argv[]) {
    node *root = newNode(10);

    insert(5, &root);
    insert(20, &root);
    insert(15, &root);
    insert(17, &root);
    insert(25, &root);

    printf("size: %d\n", size(root));

    printf("printInOrder: ");
    printInOrder(root);
    printf("\n");

    //deleteNode(15, &root);
    //deleteNode(10, &root);

    printf("printInOrder: ");
    printInOrder(root);
    printf("\n");

    node *copy = copyTree(root);

    printf("Copy printInOrder: ");
    printInOrder(root);
    printf("\n");

    printf("Copy printRootToLeaf: ");
    printRootToLeaf(copy);
    printf("\n");

    //printPaths(copy, NULL);

    return 0;
}
