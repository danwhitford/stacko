#lang br/quicklang

(define (read-syntax path port)
  (define src-lines
    (port->lines port))
  (define src-datums
    (format-datums '~a src-lines))
  (define module-datum
    `(module funstacker-mod "funstacker.rkt" (handle-args ,@src-datums)))
  (datum->syntax #f module-datum))

(provide read-syntax)

(define-macro (funstacker-module-begin HANDLE-ARGS-EXPR)
  #'(#%module-begin     
     HANDLE-ARGS-EXPR))

(provide (rename-out [funstacker-module-begin #%module-begin]))

(define (handle-args . args)  
  (for/fold ([stack-acc empty])
            ([arg (in-list args)]
             #:unless (void? arg))
   (cond
    [(number? arg) (cons arg stack-acc)]
    [(equal? \t arg) (t stack-acc)]
    [(equal? \v arg) (v stack-acc) stack-acc]
    [(or (equal? + arg) (equal? * arg))
     (define op-result (arg (first stack-acc) (second stack-acc)))
     (cons op-result (drop stack-acc 2))])))

(define (v stack)
  (if (empty? stack)
      (println "--> empty")
      (begin
        (println "--> top")
        (for-each println stack)
        (println "---"))))

(define (t stack)
  (define val (first stack))
  (println val)
  (rest stack))

(provide handle-args)
(provide + * t v)
