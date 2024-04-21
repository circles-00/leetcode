#include <stdio.h>
#include <stdlib.h>

struct ListNode {
  int val;
  struct ListNode *next;
};

int constructNumberFromDigits(struct ListNode *list) {
  int number = 0;
  int multiplier = 1;
  while (list) {
    number += list->val * multiplier;
    multiplier *= 10;
    list = list->next;
  }
  return number;
}

struct ListNode *constructListFromNumber(int number) {
  struct ListNode *list = NULL;
  struct ListNode *current = NULL;
  while (number) {
    struct ListNode *node = malloc(sizeof(struct ListNode));
    node->val = number % 10;
    node->next = NULL;
    if (list == NULL) {
      list = node;
      current = node;
    } else {
      current->next = node;
      current = node;
    }
    number /= 10;
  }
  return list;
}

struct ListNode *addTwoNumbers(struct ListNode *l1, struct ListNode *l2) {
  int number1 = constructNumberFromDigits(l1);
  int number2 = constructNumberFromDigits(l2);

  int sum = number1 + number2;
  if(sum == 0) {
    struct ListNode *node = malloc(sizeof(struct ListNode));
    node->val = 0;
    node->next = NULL;
    return node;
  }
  return constructListFromNumber(sum);
}
