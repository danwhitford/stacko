#lang br/quicklang

(define (read-syntax path port)
  (define src-lines
    (port->lines port))
  (define src-datums
    (format-datums '(handle ~a) src-lines))
  (define module-datum
    `(module stacker-mod "stacker.rkt" ,@src-datums))
  (datum->syntax #f module-datum))

(provide read-syntax)

(define-macro (stacker-module-begin HANDLE-EXPR ...)
  #'(#%module-begin     
     HANDLE-EXPR ...))

(provide (rename-out [stacker-module-begin #%module-begin]))

(define stack empty)

(define (pop-stack!)
  (define arg (first stack))
  (set! stack (rest stack))
  arg)

(define (push-stack! val)
  (set! stack (cons val stack)))

(define (handle [arg #f])  
  (cond
    [(number? arg) (push-stack! arg)]
    [(equal? \t arg) (t)]
    [(equal? \v arg) (v)]
    [(or (equal? + arg) (equal? * arg))
     (define op-result (arg (pop-stack!) (pop-stack!)))
     (push-stack! op-result)]))

(define (v)
  (println "--> top")
  (for-each println stack)
  (println "---"))

(define (t)
  (println (pop-stack!)))

(provide handle)
(provide + * t v)
